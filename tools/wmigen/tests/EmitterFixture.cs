// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Management;
using Microsoft.wmi.Common;
using Microsoft.wmi.GO;

namespace Microsoft.wmi.Tests
{
    internal static class EmitterFixture
    {
        internal static GOWmiSource CreateSource(string moduleName = "cimv2")
        {
            var ns = new InMemoryNamespace(moduleName);
            return new GOWmiSource("Test_Device", ns.AddModule(moduleName),
                "example.com/wmi/server");
        }

        // Do not construct GOWmiNamespace: its constructor discovers live WMI
        // classes. This fixture uses the real module/source/class emitters but
        // makes any accidental metadata discovery fail instead.
        private sealed class InMemoryNamespace : WmiNamespace
        {
            private readonly string moduleName;

            internal InMemoryNamespace(string moduleName) : base("root/" + moduleName)
            {
                this.moduleName = moduleName;
            }

            public override string GetModuleName(string source) { return moduleName; }

            protected override WmiSource GetWmiSource(ManagementClass wmiClass, WmiModule module)
            {
                throw new InvalidOperationException("Tests must not discover WMI classes.");
            }

            protected override WmiReference GetReference(string reference)
            {
                return new GOWmiReference(reference);
            }

            public override string GetSourceCode() { throw new NotSupportedException(); }
            public override void GenerateSources(string outputDir) { throw new NotSupportedException(); }
        }
    }
}
