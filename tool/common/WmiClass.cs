// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Management;
using System.Text;

namespace Microsoft.WmiCodeGen.Common
{
    public abstract class WmiClass : IWmiBase
    {
        public WmiSource Parent { get; private set; }

        public WmiClass(WmiSource wSource)
        {
            Name = wSource.Name;
            Derivation = "WmiInstance";
            Parent = wSource;
            Methods = new List<WmiMethod>();
            Properties = new List<WmiProperty>();
            Constructors = GetWmiConstructors();
            Related = new List<WmiRelated>();
            Logger.Debug("WmiClass {0} Namespace {1}", Name, wSource.Parent.Parent.Name);
        }

        protected abstract WmiSource GetWmiSource(string sourceName, WmiModule wModule);
        private void AddReference(string className)
        {
            if (!Parent.Parent.Parent.HasSource(className) &&
                    !Parent.CheckClass(className, Parent.Parent.Parent.Name))
            {
                Logger.Debug("AddReference: Class not found in the current Namespace. Start searching from root namespace");
                string reference = Parent.GetReference(className, "root");
                if (!string.IsNullOrEmpty(reference))
                {
                    Logger.Debug("WmiClass AddReference {0}", reference);
                    Parent.AddReference(reference);
                    Parent.Parent.Parent.AddReference(reference);
                }
                else
                {
                    Logger.Debug("No reference found anywhere. Create a Dummy Source - {0}", className);
                    Parent.Parent.Parent.AddSource(GetWmiSource(className,
                        Parent.Parent.Parent.AddModule(WmiModule.GetModuleName(className))));
                }
            }
        }

        public WmiClass(ManagementClass wmiClass, WmiSource wSource)
        {
            Parent = wSource;
            string className = FixClassName(wmiClass.ClassPath.ClassName);
            Name = className;
            Abstract = IFormat.IsAbstract(wmiClass.Qualifiers);

            //sb.Append(GetDescriptionText(wmiClass.Qualifiers));

            string baseClass = "WmiInstance";

            if (wmiClass.Derivation.Count > 0)
            {
                baseClass = WmiClass.FixClassName(wmiClass.Derivation[0]);
                AddReference(baseClass);
            }

            Methods = GetWmiMethods(wmiClass);
            Properties = GetWmiProperties(wmiClass);
            Constructors = GetWmiConstructors();
            Related = GetRelated(wmiClass);
            Derivation = baseClass;

#if false
            if (Name.Equals("CIM_ConcreteJob") && Parent.Parent.Parent.Name.ToLowerInvariant().Contains("virtualization"))
            {
                // Add reference to Thread
                Parent.AddReference("System.Threading");
            } 
#endif
        }
        public string Derivation { get; set; }

        public bool Abstract { get; set; }

        //Dictionary<string, WmiMethod> m_methods = new Dictionary<string, WmiMethod>();
        //Dictionary<string, WmiProperty> m_properties = new Dictionary<string, WmiProperty>();
        public List<WmiConstructor> Constructors { get; set; }
        public List<WmiMethod> Methods { get; set; }
        public List<WmiProperty> Properties { get; set; }

        public List<WmiRelated> Related { get; set; }

        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.AppendFormat("WmiClass [{0}][\n", Name);
            foreach (var item in Properties)
            {
                sb.AppendFormat("[Properties : {0}]\n", item.ToString());
            }
            foreach (var item in Methods)
            {
                sb.AppendFormat("[Methods : {0}]\n", item.ToString());
            }
            sb.AppendFormat("]\n", Name);
            return sb.ToString();
        }

        private static Dictionary<string, string> classNameFix = new Dictionary<string, string>()
        {
            {"__namespace", "__Namespace"},
            {"__NAMESPACE", "__Namespace"},
            {"__ntlmuser9x", "__NTLMUser9X"},
            {"DN_WIth_Binary", "DN_With_Binary"}
        };
        // Workaround to maintain the case of the class name
        public static string FixClassName(string value)
        {
            if (classNameFix.ContainsKey(value))
            {
                Logger.Debug("FixClassName - {0}", value);
                return classNameFix[value];
            }
            else
            {
                return value;
            }
        }

        protected abstract WmiConstructor GetWmiConstructor(List<string> paramText, string derivation, string body);
        protected abstract List<WmiConstructor> GetWmiConstructors();
        protected abstract WmiMethod GetWmiMethod(MethodData mData, WmiClass wClass);

        /// <summary>
        /// 
        /// </summary>
        /// <param name="wmiClass"></param>
        /// <returns></returns>
        private List<WmiMethod> GetWmiMethods(ManagementClass wmiClass)
        {
            List<WmiMethod> wMethods = new List<WmiMethod>();

            foreach (MethodData mData in wmiClass.Methods)
            {
                if (!mData.Origin.Equals(wmiClass.ClassPath.ClassName, StringComparison.OrdinalIgnoreCase))
                {
                    continue;
                }
                wMethods.Add(GetWmiMethod(mData, this));
            }

            return wMethods;
        }

        protected abstract WmiProperty GetWmiProperty(PropertyData pData, WmiClass wClass);

        private List<WmiProperty> GetWmiProperties(ManagementClass wmiClass)
        {
            string className = wmiClass.ClassPath.ClassName;
            List<WmiProperty> wProperties = new List<WmiProperty>();
            foreach (var pData in wmiClass.Properties)
            {
                if (!pData.Origin.Equals(className, StringComparison.OrdinalIgnoreCase))
                {
                    continue;
                }
                //FixPropertyName(pData.Name, wmiClass);

                wProperties.Add(GetWmiProperty(pData, this));
            }
            return wProperties;
        }

        protected abstract WmiRelated GetWmiRelated(string name, WmiClass parent);
        /// <summary>
        /// 
        /// </summary>
        /// <param name="wmiClass"></param>
        /// <returns></returns>
        private List<WmiRelated> GetRelated(ManagementClass wmiClass)
        {
            string className = wmiClass.ClassPath.ClassName;
            List<WmiRelated> related = new List<WmiRelated>();
            bool needRelated = false;
            string[] relatedClass = new string[] { "msvm" };
            foreach (var item in relatedClass)
            {
                if (className.ToLowerInvariant().StartsWith(item))
                {
                    needRelated = true;
                    break;
                }
            }

            if (!needRelated) return related;

            Logger.Debug("GetRelated ClassName :{0}, Namespace: {1}",
                className, wmiClass.Scope.Path.Path);
            using (ManagementObjectSearcher mos =
                new ManagementObjectSearcher(wmiClass.Scope.Path.Path,
                    "select * from " + wmiClass.ClassPath.ClassName))
            {
                try
                {
                    using (ManagementObjectCollection moCollection = mos.Get())
                    {
                        foreach (ManagementObject item in moCollection)
                        {
                            using (ManagementObjectCollection moRelatedCollection = item.GetRelated())
                            {
                                //Logger.Debug("GetRelated - Found {0}", moRelatedCollection.Count);
                                foreach (var robjects in moRelatedCollection)
                                {
                                    string rClass = robjects["__CLASS"] as string;
                                    Logger.Debug("Related Class [{0}]", rClass);
                                    related.Add(GetWmiRelated(rClass, this));
                                    AddReference(rClass);
                                }
                            }
                            break;
                        }
                    }
                }
                catch (Exception)
                {
                }
            }
            return related;
        }

        public bool HasRelated(string related)
        {
            if (Related != null)
            {
                WmiRelated wrelated = Related.Find(r => r.Name.Equals(related));
                if (wrelated != null) return true;
                if (Derivation != "WmiInstance")
                {
                    return Parent.Parent.Parent.HasRelated(Derivation, related);
                }
            }
            return false;
        }

    }
}
