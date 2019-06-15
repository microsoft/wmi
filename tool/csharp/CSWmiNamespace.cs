// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Management;
using System.Text;
using Microsoft.WmiCodeGen.Common;

namespace Microsoft.WmiCodeGen.CSharp
{
    public class CSWmiNamespace : WmiNamespace
    {
        public CSWmiNamespace(string name) : base(name)
        {
        }

        protected override WmiReference GetReference(string reference)
        {
            return new CSWmiReference(reference);
        }

        protected override WmiSource GetWmiSource(ManagementClass wmiClass, WmiModule wModule)
        {
            return new CSWmiSource(wmiClass, wModule);
        }

        public override string GetSourceCode()
        {
            throw new NotImplementedException();
        }

        public override string GetModuleName(string source)
        {
            return CSWmiSource.GetModuleName(source);
        }

        public override void GenerateSources(string outputDir)
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
    }
}
