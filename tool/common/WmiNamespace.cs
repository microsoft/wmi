// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.IO;
using System.Management;

namespace Microsoft.WmiCodeGen.Common
{
    public abstract  class WmiNamespace : IWmiBase
    {
        Dictionary<string /* module Name */,
            WmiModule> namespaceList = new Dictionary<string, WmiModule>();
        public List<WmiReference> sourceReferences = new List<WmiReference>();
        public WmiNamespace(string name)
        {
            Name = name;
            CSNamespaceName = name.Replace("\\", ".").Replace("/", ".");
            AddSources(GenerateWmiSources());
            Logger.Debug("WmiNamespace {0}", name);
        }

        public Dictionary<string, WmiModule> Modules { get { return namespaceList; } }

        void AddModule(WmiModule module)
        {
            Modules[module.Name] = module;
        }

        public WmiModule AddModule(string moduleName)
        {
            if (!namespaceList.ContainsKey(moduleName))
                namespaceList[moduleName] = new WmiModule(moduleName, this);

            return namespaceList[moduleName];
        }

        WmiModule GetModule(string moduleName)
        {
            WmiModule module = null;
            namespaceList.TryGetValue(moduleName, out module);
            return module;
        }

        public void AddSource(WmiSource wSource)
        {
            string moduleName = WmiModule.GetModuleName(wSource.Name);
            WmiModule module = GetModule(moduleName);
            if (module == null)
            {
                module = AddModule(moduleName);
            }
            module.Add(wSource);
        }

        public void AddSources(Dictionary<string, WmiSource> wSources)
        {
            foreach (var item in wSources)
            {
                AddSource(item.Value);
            }
        }

        protected abstract WmiReference GetReference(string reference);
        public void AddReference(string reference)
        {
            if (sourceReferences.Find(r => r.Reference == reference) == null)
                sourceReferences.Add(GetReference(reference));
        }

        public void AddEnum(WmiEnum wEnum)
        {
            string moduleName = "Common";
            WmiModule module = GetModule(moduleName);
            if (module == null)
            {
                module = AddModule(moduleName);
            }
            module.Enums[wEnum.Name] = wEnum;
        }

        public WmiEnum GetEnum(string enumName)
        {
            string moduleName = "Common";
            WmiModule module = GetModule(moduleName);
            if (module != null)
            {
                if (module.Enums.ContainsKey(enumName))
                {
                    return module.Enums[enumName];
                }
            }

            return null;
        }

        public bool HasRelated(string className, string related)
        {
            if (HasSource(className))
            {
                return GetSource(className).Class.HasRelated(related);
            }
            return false;
        }
#if false
        public WmiSource GetSource(string sourceName)
        {
            string moduleName = WmiModule.GetModuleName(sourceName);
            WmiModule module = GetModule(moduleName);
            if (module != null)
            {
                if (module.Sources.ContainsKey(sourceName))
                {
                    return module.Sources[sourceName];
                }
            }

            return null;
        } 
#endif
        public string CSNamespaceName { get; private set; }

        public void GenerateSources(string outputDir)
        {
            string path = Path.Combine(Environment.CurrentDirectory, outputDir, Name, CSNamespaceName);
            if (!Directory.Exists(path)) { Directory.CreateDirectory(path); }

            foreach (var item in Modules)
            {
                string mpath = Path.Combine(path, item.Value.Name);
                if (!Directory.Exists(mpath)) Directory.CreateDirectory(mpath);
                item.Value.GenerateSources(mpath);
            }

        }
        private Dictionary<string, WmiSource> GenerateWmiSources()
        {
            Dictionary<string, WmiSource> wSource = new Dictionary<string, WmiSource>();
            using (ManagementObjectSearcher searcher = new ManagementObjectSearcher(
                                                            new ManagementScope(Name),
                                                            new WqlObjectQuery("select * from meta_class"), null))
            {
                try
                {
                    using (ManagementObjectCollection moc = searcher.Get())
                    {
                        foreach (ManagementClass mgmtClass in moc)
                        {
                            string className = WmiClass.FixClassName(mgmtClass["__CLASS"].ToString());
                            //if (!WmiSource.ClassList.ContainsKey(Name)) WmiSource.ClassList[Name] = new List<string>();
                            //WmiSource.ClassList[Name].Add(className);
                            if (!wSource.ContainsKey(className))
                                GetSource(className);
                        }
                    }
                }
                catch (Exception e)
                {
                    Logger.Exception(e, "");
                }
            }
            return wSource;
        }

        public bool HasSource(string sourceName)
        {
            if (sourceName.Equals("WmiInstance")) return false;

            string moduleName = WmiModule.GetModuleName(sourceName);
            WmiModule module = GetModule(moduleName);
            if (module != null)
            {
                if (module.Sources.ContainsKey(sourceName))
                {
                    return true;
                }
            }
            return false;
        }

        protected abstract WmiSource GetWmiSource(ManagementClass wmiClass, WmiModule wModule);
        public WmiSource GetSource(string className)
        {
            if (className.Equals("WmiInstance")) return null;

            string moduleName = WmiModule.GetModuleName(className);
            WmiModule module = GetModule(moduleName);
            if (module != null)
            {
                if (module.Sources.ContainsKey(className))
                {
                    return module.Sources[className];
                }
            }

            using (ManagementClass mClass = new ManagementClass(Name, className, null))
            {
                WmiSource wSource = GetWmiSource(mClass, AddModule(moduleName));
                AddSource(wSource);
                return wSource;
            }

        }

    }
}
