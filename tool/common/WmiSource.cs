// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Management;
using System.Text;

namespace Microsoft.WmiCodeGen.Common
{
    public abstract class WmiSource : IWmiBase
    {
        public WmiModule Parent { get; private set; }
        public WmiSource(string sourceName, WmiModule wModule)
        {
            Name = sourceName;
            Parent = wModule;
        }

        public WmiSource(ManagementClass wmiClass, WmiModule wModule)
        {
            Parent = wModule;
            Name = WmiClass.FixClassName(wmiClass.ClassPath.ClassName);
        }

        List<WmiReference> m_references = new List<WmiReference>();


        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.AppendFormat("WmiSource [{0}] [Class : {1}]", Name, Class.ToString());
            return sb.ToString();
        }

        public string Namespace { get { return Parent.Parent.CSNamespaceName; } }
        public string WmiNamespace { get { return Parent.Parent.Name; } }
        public List<WmiReference> References { get { return m_references; } }
        public WmiClass Class { get; set; }

        public void GenerateSource(string outdir)
        {
            string path = Path.Combine(outdir, Name + ".cs");

            File.WriteAllText(path, GetSourceCode());
            Logger.Info("Source {0}", path);
        }

        public static Dictionary<string, List<string>> ClassList = new Dictionary<string, List<string>>();
        /// <summary>
        /// 
        /// </summary>
        /// <param name="cName"></param>
        /// <param name="wNamespace"></param>
        /// <returns></returns>
        public bool CheckClass(string cName, string wNamespace)
        {
            // Do not search user WMI Namespace
            if (wNamespace.Equals("root/rsop/user", StringComparison.OrdinalIgnoreCase)) return false;
            // It takes time to check this. There will not be any dependency in this namespace.so Skip it
            if (wNamespace.Equals("root/directory/LDAP", StringComparison.OrdinalIgnoreCase)) return false;

            Logger.Debug("CheckClass - Class [{0}], NS [{1}]", cName, wNamespace);
            if (!ClassList.ContainsKey(wNamespace))
            {
                ClassList[wNamespace] = new List<string>();
                using (ManagementObjectSearcher searcher = new ManagementObjectSearcher(
                                                            new ManagementScope(wNamespace),
                                                            new WqlObjectQuery("select * from meta_class"), null))
                {
                    try
                    {
                        using (ManagementObjectCollection moc = searcher.Get())
                        {
                            foreach (ManagementClass mgmtClass in moc)
                            {
                                string className = WmiClass.FixClassName(mgmtClass["__CLASS"].ToString());
                                ClassList[wNamespace].Add(className);
                                //if (className.Equals(cName, StringComparison.OrdinalIgnoreCase)) return true;
                            }
                        }
                    }
                    catch (ManagementException) { }
                    catch (Exception e)
                    {
                        Logger.Exception(e, "");
                    }
                }
            }

            foreach (var item in ClassList[wNamespace])
            {
                if (item.Equals(cName))
                {
                    return true;
                }
            }
            
            return false;

#if false
            using (ManagementObjectSearcher searcher = new ManagementObjectSearcher(
                                    new ManagementScope(wNamespace),
                                    new WqlObjectQuery(string.Format(CultureInfo.InvariantCulture,
                                            @"select * from meta_class where __This ISA ""{0}""", cName)), null))
            {
                try
                {
                    using (ManagementObjectCollection moc = searcher.Get())
                    {
                        foreach (ManagementClass mgmtClass in moc)
                        {
                            string className = WmiClass.FixClassName(mgmtClass["__CLASS"].ToString());
                            if (className.Equals(cName, StringComparison.OrdinalIgnoreCase)) return true;
                        }
                    }
                }
                catch (ManagementException e)
                {
                }
                catch (Exception e)
                {
                    Logger.Exception(e, "Class {0}, Namespace {1}", cName, wNamespace);
                }

                return false;
            } 
#endif
        }

        public bool HasReference(string reference)
        {
            string csReference = reference.StartsWith("root", StringComparison.OrdinalIgnoreCase)
                ? string.Format(CultureInfo.InvariantCulture,
                "Microsoft.Test.Wmi.{0}", reference.Replace("/", ".")) : reference;
            if (m_references.Find(a => a.Reference == csReference) != null) return true;
            return false;
        }

        protected abstract WmiReference GetWmiReference(string reference);
        public void AddReference(string reference)
        {
            Logger.Debug("WmiSource - AddReference {0}", reference);
            string csReference = reference.StartsWith("root", StringComparison.OrdinalIgnoreCase)
                ? string.Format(CultureInfo.InvariantCulture,
                "Microsoft.Test.Wmi.{0}", reference.Replace("/", ".")) : reference;
            if (!string.IsNullOrEmpty(csReference))
            {
                WmiReference wReference = m_references.Find(a => a.Reference == csReference);
                if (wReference == null)
                {
                    Logger.Debug("AddReference {0}", csReference);
                    m_references.Add(GetWmiReference(csReference));
                }
            }
        }

        /// <summary>
        /// Pre-Populated reference path
        /// </summary>
        private Dictionary<string, string> namespaceReference = new Dictionary<string, string>()
        {
            {"CIM_RegisteredProfile", "root/Interop"},
            {"CIM_SettingsDefineState", "root/virtualization/v2"},
            {"CIM_ConcreteJob", "root/StandardCimv2"},
        };

        public string GetReference(string className, string wmiNamespace)
        {
            if (namespaceReference.ContainsKey(className)) return namespaceReference[className];

            Logger.Debug("GetReference - CLASS:{0}, Namespace:{1}", className, wmiNamespace);
            List<string> ExcludeFailbackNamespace = new List<string> { "ms_409", "ms_509", "ms_501" };
            if (wmiNamespace.Equals(Name, StringComparison.OrdinalIgnoreCase)) return string.Empty;
            if (CheckClass(className, wmiNamespace))
            {
                return wmiNamespace;
            }

            using (ManagementClass cls = new ManagementClass
                (new ManagementScope(wmiNamespace), new ManagementPath("__namespace"), null))
            {
                try
                {
                    foreach (var instance in cls.GetInstances())
                    {
                        string instanceName = instance["Name"].ToString();
                        // skip the failback namespaces
                        if (!ExcludeFailbackNamespace.Contains(instanceName.ToLowerInvariant()))
                        {
                            string cNamespace = string.Format("{0}/{1}", wmiNamespace, instanceName);
                            string wRef = GetReference(className, cNamespace);
                            if (!string.IsNullOrEmpty(wRef))
                            {
                                Logger.Debug("Found Reference {0} for Class {1}", wRef, className);
                                return wRef;
                            }
                        }
                    }
                }
                catch (ManagementException) { }
                catch (Exception e)
                {
                    Logger.Exception(e, "");
                }
            }

            return string.Empty;
        }

        public static Type ConvertCimTypeToSystemType(CimType type)
        {
            switch (type)
            {
                case CimType.Boolean:
                    return typeof(bool);
                case CimType.Char16:
                    return typeof(Char);
                case CimType.DateTime:
                    return typeof(DateTime);
                case CimType.Object:
                    return typeof(Object);
                case CimType.Real32:
                    return typeof(float);
                case CimType.Real64:
                    return typeof(Double);
                case CimType.Reference:
                    return typeof(Object);
                case CimType.SInt16:
                    return typeof(Int16);
                case CimType.SInt32:
                    return typeof(Int32);
                case CimType.SInt64:
                    return typeof(Int64);
                case CimType.SInt8:
                    return typeof(Byte);
                case CimType.String:
                    return typeof(String);
                case CimType.UInt16:
                    return typeof(UInt16);
                case CimType.UInt32:
                    return typeof(UInt32);
                case CimType.UInt64:
                    return typeof(UInt64);
                case CimType.UInt8:
                    return typeof(Byte);
                case CimType.None:
                default:
                    throw new NotImplementedException();
            }
        }

        private static string FixTypeName(string name)
        {
            if (name.Equals("unint32", StringComparison.OrdinalIgnoreCase)) return typeof(UInt32).Name;
            return name;
        }

        protected abstract WmiSource GetWmiSource(string sourceName, WmiModule wModule);

        public string GetType(PropertyData pData, out Type type)
        {
            type = ConvertCimTypeToSystemType(pData.Type);
            string typeName = type.Name;
            if (pData.Type == CimType.Object || pData.Type == CimType.Reference)
            {
                object typeValue;
                if (IFormat.TryGetQualifierValue(pData.Qualifiers, "CimType", out typeValue))
                {
                    string typeValueString = typeValue.ToString();
                    if (typeValueString.Contains(":"))
                    {
                        string[] typeValues = typeValueString.Split(new char[] { ':' });
                        if (typeValues.Length > 1)
                        {
                            // Add reference to this type
                            typeName = WmiClass.FixClassName(typeValues[1]);

                            if (!typeName.Equals("unint32", StringComparison.OrdinalIgnoreCase))
                            {
                                if (!Parent.Parent.HasSource(typeName) &&
                                    !CheckClass(typeName, Parent.Parent.Name))
                                {
                                    // Class not found in the current Namespace. Start searching from root namespace
                                    Logger.Debug("Class not found in the current Namespace." +
                                    " Start searching from root namespace - {0}",
                                        Namespace);
                                    string reference = GetReference(typeName, "root");
                                    if (!string.IsNullOrEmpty(reference))
                                    {
                                        AddReference(reference);
                                        Parent.Parent.AddReference(reference);
                                    }
                                    else
                                    {
                                        // No reference found anywhere. Create a Dummy Source
                                        Parent.Parent.AddSource(GetWmiSource(typeName, Parent.Parent.AddModule(WmiModule.GetModuleName(typeName))));
                                    }
                                }
                            }
                            else
                            {
                                type = typeof(UInt32);
                            }
                        }
                    }
                }
            }
            else if (pData.Type != CimType.Boolean && IFormat.HasQualifier(pData.Qualifiers, "values"))
            {
                WmiEnum wEnum = GetEnum(pData);
                if (wEnum != null)
                {
                    typeName = wEnum.Name;
                }
                else typeName = type.Name;
            }
            return FixTypeName(typeName);
            //return pData.IsArray ? typeName + "[]" : typeName;
        }

        protected abstract WmiEnum GetWmiEnum(string enumName, WmiSource wSource);
        public WmiEnum GetEnum(PropertyData pData)
        {
            WmiEnum wEnum = null;
            Object[] values = IFormat.GetQualifierValue(pData.Qualifiers, "values") as Object[];
            Object[] valueMaps = IFormat.GetQualifierValue(pData.Qualifiers, "valueMap") as Object[];
            if (values != null && values.Length > 0)
            {
                string enumName = WmiEnum.GetEnumName(pData, this);
                //string classPrefix = WmiModule.GetModuleName(wSource.Name);
                //string enumFile = enumName + ".cs";

                wEnum = this.Parent.Parent.GetEnum(enumName);
                if (wEnum != null) return wEnum;

                // Create new Enum
                wEnum = GetWmiEnum(enumName, this);
                this.Parent.Parent.AddEnum(wEnum);
                //wSource.Parent.Enums[enumName] = wEnum;

                int value = 0;

                for (int i = 0; i < values.Length; i++)
                {
                    //byte[] bytes = Encoding.Unicode.GetBytes(values[i].ToString());
                    //string temp = Encoding.Default.GetString(bytes);
                    string enumValue = (Convert.ToString(values[i], CultureInfo.InvariantCulture));

                    try
                    {
                        value = valueMaps != null && valueMaps.Length > 0 ? Convert.ToInt32(valueMaps[i], CultureInfo.InvariantCulture) : i;
                    }
                    catch (Exception)
                    {
                        value++;
                    }

                    Logger.Debug("Enum Name {0} => Value[{1}]", enumValue, value);
                    wEnum.Add(WmiEnum.FixEnumName(enumValue), value);
                }

#if false
                // Check if the existing enum and the current enum are exactly the same
                // If the Values are different, modify the enumName
                if (!existingEnum.Equals(wEnum))
                {
                    // Find a new Name based on the context of this enum
                    wEnum.Name = pData.Name + this.GetSuffix() + "Values";
                    Logger.Debug("Enum : {0}", wEnum.Name);
                    this.Parent.Parent.AddEnum(wEnum);
                }
                else
                {
                    Logger.Debug("Enum [{0}] already Exists", wEnum.Name);
                } 
#endif

                return wEnum;
            }
            return wEnum;
            // If there are no values defined in enum, use the regular type
            //return ConvertCimTypeToSystemType(pData.Type).Name;
        }

        public string GetPrefix()
        {
            return WmiModule.GetModuleName(Name);
        }

        public string GetSuffix()
        {
            string[] classPrefixes = Name.Split(new char[] { '_' });
            return classPrefixes[classPrefixes.Length - 1];
        }
    }
}
