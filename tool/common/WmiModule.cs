// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System.Collections.Generic;

namespace Microsoft.WmiCodeGen.Common
{
    public class WmiModule : IWmiBase
    {
        public WmiNamespace Parent { get; private set; }
        public WmiModule(string name, WmiNamespace wNamespace)
        {
            Parent = wNamespace;
            Name = name;
            Logger.Debug("WmiModule {0} Namespace {1}", name, wNamespace.Name);
        }

        Dictionary<string, WmiSource> m_source = new Dictionary<string, WmiSource>();
        Dictionary<string, WmiEnum> m_enum = new Dictionary<string, WmiEnum>();

        public void Add(WmiSource source)
        {
            m_source[source.Name] = source;
        }

        public bool Contains(string sourceName)
        {
            return m_source.ContainsKey(sourceName);
        }

        public WmiSource Get(string sourceName)
        {
            WmiSource source = null;
            m_source.TryGetValue(sourceName, out source);
            return source;
        }
        public Dictionary<string, WmiSource> Sources { get { return m_source; } }
        public Dictionary<string, WmiEnum> Enums { get { return m_enum; } }

        public void GenerateSources(string outdir)
        {
            foreach (var item in Sources)
            {
                Logger.Debug(item.ToString());
                item.Value.GenerateSource(outdir);
            }

            foreach (var item in Enums)
            {
                item.Value.GenerateSource(outdir);
            }
        }

        public static string GetModuleName(string className)
        {
            if (className.StartsWith("__")) return "System";
            string[] classPrefixes = className.Split(new char[] { '_' });

            if (classPrefixes.Length > 1)
            {
                return classPrefixes[0];
            }
            return "Common";
        }

        public override string GetSourceCode()
        {
            throw new System.NotImplementedException();
        }
    }
}
