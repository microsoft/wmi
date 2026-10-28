// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.IO;
using System.Linq;
using System.Management;
using System.Text.RegularExpressions;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Microsoft.wmi.Common;
using Microsoft.wmi.CSharp;

namespace Microsoft.wmi.Tests
{
    [TestClass]
    public class CSharpEmissionTests
    {
        [TestMethod]
        public void SourcesEmitCSharpImportsNamespacesAndDelegatingConstructors()
        {
            var source = EmitterFixture.CreateCSharpSource();
            source.AddReference("root/standardcimv2");
            source.AddReference("root/standardcimv2");
            Assert.AreEqual(1, source.References.Count(
                r => r.Reference == "Microsoft.Test.Wmi.root.standardcimv2"));

            var code = source.GetSourceCode();
            StringAssert.Contains(code, "using Microsoft.Test.Wmi.root.standardcimv2;");
            StringAssert.Contains(code, "namespace Microsoft.Test.Wmi.root.cimv2");
            StringAssert.Contains(code, "class Test_Device : WmiInstance");
            StringAssert.Contains(code,
                "this(hostName, \"root/cimv2\", \"Test_Device\", credentials)");
            StringAssert.Contains(code, "base(hostName, wmiNamespace, className, credentials)");
        }

        [DataTestMethod]
        [DataRow("namespace", CimType.String, false, "String Namespace", "String")]
        [DataRow("Counters", CimType.UInt32, true, "UInt32 [] Counters", "UInt32[]")]
        public void PropertiesPreserveWireNamesWhileEmittingTypesAndAccessors(
            string name, CimType type, bool array, string declaration, string setterType)
        {
            var source = EmitterFixture.CreateCSharpSource();
            using (var schema = new ManagementClass())
            {
                schema.Properties.Add(name, type, array);
                var property = new CSWmiProperty(schema.Properties[name], (CSWmiClass)source.Class);
                StringAssert.Contains(Regex.Replace(property.GetSourceCode(), @"\s+", " "),
                    "public virtual " + declaration);
                StringAssert.Contains(property.Setter,
                    "SetProperty<" + setterType + ">(\"" + name + "\", value)");
                StringAssert.Contains(property.Getter, "(\"" + name + "\")");
            }
        }

        [DataTestMethod]
        [DataRow("Count", CimType.UInt32, ParamType.Input, true, false,
            "/* OPTIONAL IN */ [Optional] UInt32 ? Count")]
        [DataRow("namespace", CimType.String, ParamType.Output, false, false,
            "/* OUT */ out String Namespace")]
        [DataRow("Values", CimType.UInt16, ParamType.Reference, false, true,
            "/* IN/OUT */ ref UInt16 [] Values")]
        public void MethodParametersEmitNullableOptionalValuesAndRefOutDirections(
            string name, CimType type, ParamType direction, bool optional, bool array, string expected)
        {
            var source = EmitterFixture.CreateCSharpSource();
            var method = new CSWmiMethod { Parent = source.Class };
            using (var schema = new ManagementClass())
            {
                schema.Properties.Add(name, type, array);
                var parameter = new CSWmiMethodParam(schema.Properties[name], direction, method, optional);
                Assert.AreEqual(expected, Regex.Replace(parameter.GetSourceCode(), @"\s+", " ").Trim());
            }
        }

        [TestMethod]
        public void EnumFilesRetainTheirNamespaceAndWireValues()
        {
            var source = EmitterFixture.CreateCSharpSource();
            var values = new CSWmiEnum("Device_State", source);
            values.Add("Ready", 4);
            values.Add("Busy", 16);
            var directory = Path.Combine(Path.GetTempPath(), "wmigen-tests-" + Guid.NewGuid().ToString("N"));
            var path = Path.Combine(directory, "Device_State.cs");
            Directory.CreateDirectory(directory);
            try
            {
                values.GenerateSource(directory);
                var code = File.ReadAllText(path);
                StringAssert.Contains(code, "namespace Microsoft.Test.Wmi.root.cimv2");
                StringAssert.Contains(code, "public enum Device_State");
                StringAssert.Contains(code, "Ready = 4,");
                StringAssert.Contains(code, "Busy = 16,");
            }
            finally
            {
                File.Delete(path);
                Directory.Delete(directory);
            }
        }
    }
}
