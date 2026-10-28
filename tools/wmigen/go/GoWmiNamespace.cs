// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.IO;
using System.Linq;
using System.Management;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.GO
{
    public class GOWmiNamespace : WmiNamespace
    {
        private string m_rootmodulepath;
        
        public GOWmiNamespace(string name, string rootmodulepath) : base(name)
        {
            m_rootmodulepath = rootmodulepath;

            AddSources(GenerateWmiSources());
        }

        protected override WmiReference GetReference(string reference)
        {
            return new GOWmiReference(reference);
        }

        protected override WmiSource GetWmiSource(ManagementClass wmiClass, WmiModule wModule)
        {
            return new GOWmiSource(wmiClass, wModule, m_rootmodulepath);
        }

        public override string GetSourceCode()
        {
            throw new NotImplementedException();
        }

        public override string GetModuleName(string source)
        {
            return Name.Split(new char[] { '/' }).Last().ToLowerInvariant();
        }

        protected override string GetCommonModuleName()
        {
            return GetModuleName("common");
        }

        private string GetCuratedName(string name)
        {
            // WMI exposes a "default" namespace that wmigen would convert to a module name. 
            // Since in GO the "default" keyword is reserved and cannot be used for module names, we need to modify that name to something go is happy with.
            if (!string.IsNullOrEmpty(name) && name.Equals("root/default", StringComparison.OrdinalIgnoreCase))
            {
                return "root/wmidefault";
            }

            return name;
        }

        public override void GenerateSources(string outputDir)
        {
            string name = GetCuratedName(Name.ToLowerInvariant());
            string path = Path.Combine(Environment.CurrentDirectory, outputDir, name);
            if (!Directory.Exists(path)) { Directory.CreateDirectory(path); }

            foreach (var item in Modules)
            {
                if (!Directory.Exists(path)) Directory.CreateDirectory(path);
                item.Value.GenerateSources(path);
            }

        }
    }
}
