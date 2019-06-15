// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.IO;
using System.Linq;
using System.Management;
using Microsoft.WmiCodeGen.Common;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiNamespace : WmiNamespace
    {
        public GOWmiNamespace(string name) : base(name)
        {
        }

        protected override WmiReference GetReference(string reference)
        {
            return new GOWmiReference(reference);
        }

        protected override WmiSource GetWmiSource(ManagementClass wmiClass, WmiModule wModule)
        {
            return new GOWmiSource(wmiClass, wModule);
        }

        public override string GetSourceCode()
        {
            throw new NotImplementedException();
        }

        public override string GetModuleName(string source)
        {
            return Name.Split(new char[] { '/' }).Last();
        }

        protected override string GetCommonModuleName()
        {
            return GetModuleName("Common");
        }

        public override void GenerateSources(string outputDir)
        {
            string path = Path.Combine(Environment.CurrentDirectory, outputDir, Name);
            if (!Directory.Exists(path)) { Directory.CreateDirectory(path); }

            foreach (var item in Modules)
            {
                if (!Directory.Exists(path)) Directory.CreateDirectory(path);
                item.Value.GenerateSources(path);
            }

        }
    }
}
