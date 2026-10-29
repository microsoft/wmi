// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.IO;
using System.Management;
using System.Text;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.GO
{
    public class GOWmiSource : WmiSource
    {
        private string m_rootmodulepath;

        public GOWmiSource(string sourceName, WmiModule wModule, string rootmodulepath) : base(sourceName, wModule)
        {
            m_rootmodulepath = rootmodulepath;
            HeaderComment = string.Format(CultureInfo.InvariantCulture,
@"
// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.
//
// Author:
//      Auto Generated on {2} using wmigen
//      Source {3}
//////////////////////////////////////////////",
                    Name,
                    "This class doesn't exist. This is just a placeholder",
                    DateTime.Now.ToShortDateString(),
                    Parent.Parent.CSNamespaceName
                    );
            AddReference("github.com/microsoft/wmi/pkg/base/query");
            AddReference("github.com/microsoft/wmi/pkg/wmiinstance");

            Class = new GOWmiClass(this, m_rootmodulepath);
        }

        public GOWmiSource(ManagementClass wmiClass, WmiModule wModule, string rootmodulepath) : base (wmiClass, wModule)
        {
            m_rootmodulepath = rootmodulepath;

            CopyrightText = string.Format(CultureInfo.InvariantCulture,
@"// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.
");

            ///
            HeaderComment = string.Format(CultureInfo.InvariantCulture,
@"// 
// Author:
//      Auto Generated on {2} using wmigen
//      Source {3}
//////////////////////////////////////////////",
                    Name,
                    IFormat.GetDescriptionText(wmiClass.Qualifiers),
                    DateTime.Now.ToShortDateString(),
                    Parent.Parent.CSNamespaceName
                    );

            AddReference("github.com/microsoft/wmi/pkg/base/query");
            AddReference("github.com/microsoft/wmi/pkg/wmiinstance");

            Class = new GOWmiClass(wmiClass, this, m_rootmodulepath);
        }
        

        public override string GetSourceCode() 
        {
            StringBuilder sb = new StringBuilder();

            sb.AppendLine(CopyrightText);
            sb.AppendLine(HeaderComment);
            sb.AppendFormat(CultureInfo.InvariantCulture,
                "package {0}\n", GetCuratedName(Parent.Name));

            // Package imports
            sb.AppendLine("import (");
            foreach (var item in References)
            {
                sb.AppendLine(item.GetSourceCode());
            }
            sb.AppendLine(")");

            // Types
            sb.AppendLine(Class.GetSourceCode());

            return sb.ToString();
        }
        protected override WmiReference GetWmiReference(string reference)
        {
            string csReference = reference.StartsWith("root", StringComparison.OrdinalIgnoreCase)
                ? string.Format(CultureInfo.InvariantCulture,
                "{0}/{1}", m_rootmodulepath, reference) : reference;
            return new GOWmiReference(csReference);
        }

        public GOWmiReference GetGoReference(string reference)
        {
            Logger.Debug("GoWmiSource - GetGoReference {0}", reference);
            WmiReference wReferenceFound = References.Find(a => a.Reference == reference);
            if (wReferenceFound == null)
            {
                throw new Exception("Reference not found");
            }
            return wReferenceFound as GOWmiReference;
        }

        private string GetCuratedName(string name)
        {
            // WMI exposes a "default" namespace that wmigen would convert to a module name. 
            // Since in GO the "default" keyword is reserved and cannot be used for module names, we need to modify that name to something go is happy with.
            if (!string.IsNullOrEmpty(name) && name.Equals("default", StringComparison.OrdinalIgnoreCase))
            {
                return "wmidefault";
            }

            return name;
        }

        public override void GenerateSource(string outdir)
        {
            string path = Path.Combine(outdir, Name.TrimStart(new char[] {'_'}) + ".go");

            File.WriteAllText(path, GetSourceCode());
            Logger.Info("Source {0}", path);
        }

        public void GenerateMakefile(string outputDir)
        {
            throw new NotImplementedException();
        }

        public static string GetHeaderCommentText(QualifierDataCollection qCollection)
        {
            object description = IFormat.GetQualifierValue(qCollection, "description");
            return String.Format(CultureInfo.InvariantCulture,
@"
// {0}",
               description != null ? description.ToString().Replace("\n", "\n///") : string.Empty);
        }

        protected override WmiEnum GetWmiEnum(string enumName, WmiSource wSource)
        {
            return new GOWmiEnum(enumName, wSource as GOWmiSource);
        }

        protected WmiSource GetWmiSource(string sourceName, WmiModule wModule)
        {
            return new GOWmiSource(sourceName, wModule, m_rootmodulepath);
        }

        public static string GetErrorHandler() {
            return "if err != nil { return }";
        }
        //
        // Get the go type of the property
        //
        public override string GetType(PropertyData pData, out Type type, out WmiReference goreference)
        {
            type = ConvertCimTypeToSystemType(pData.Type);
            goreference = null;
            string typeName = type.Name;
            string moduleName = GetCuratedName(Parent.Name);
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
                                        goreference = AddReference(reference);
                                        Parent.Parent.AddReference(reference);
                                    }
                                    else
                                    {
                                        // No reference found anywhere. Create a Dummy Source
                                        Parent.Parent.AddSource(GetWmiSource(typeName, Parent.Parent.AddModule(moduleName)));
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
                    type = typeof(Enum);
                }
                else typeName = type.Name;
            }
            else
            {
                typeName = ConvertCimTypeToGoType(pData.Type);
            }
            return FixTypeName(typeName);
            //return pData.IsArray ? typeName + "[]" : typeName;
        }

        private string FixTypeName(string name)
        {
            if (name.Equals("String")) return "string";
            if (name.Equals("unint32", StringComparison.OrdinalIgnoreCase)) return "uint32";
            if (name.Equals("Object", StringComparison.OrdinalIgnoreCase)) 
            {
                return "interface{}";
            }
            return name;
        }

        public static string ConvertCimTypeToGoType(CimType type)
        {
            switch (type)
            {
                case CimType.Boolean:
                    return "bool";
                case CimType.Char16:
                    return "byte";
                case CimType.DateTime:
                    return "string";
                case CimType.Object:
                    return "interface{}";
                case CimType.Real32:
                    return "float32";
                case CimType.Real64:
                    return "float64";
                case CimType.Reference:
                    return "interface{}";
                case CimType.SInt16:
                    return "int16";
                case CimType.SInt32:
                    return "int32";
                case CimType.SInt64:
                    return "int64";
                case CimType.SInt8:
                    return "int8";
                case CimType.String:
                    return "string";
                case CimType.UInt16:
                    return "uint16";
                case CimType.UInt32:
                    return "uint32";
                case CimType.UInt64:
                    return "uint64";
                case CimType.UInt8:
                    return "uint8";
                case CimType.None:
                    return "none";
                default:
                    throw new NotImplementedException();
            }
        }
    }
}
