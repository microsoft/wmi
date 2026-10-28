// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Linq;
using System.Management;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Microsoft.wmi.Common;
using Microsoft.wmi.GO;

namespace Microsoft.wmi.Tests
{
    [TestClass]
    public class GoEmissionTests
    {
        [DataTestMethod]
        [DataRow(CimType.Boolean, "bool")]
        [DataRow(CimType.SInt8, "int8")]
        [DataRow(CimType.SInt16, "int16")]
        [DataRow(CimType.SInt32, "int32")]
        [DataRow(CimType.SInt64, "int64")]
        [DataRow(CimType.UInt8, "uint8")]
        [DataRow(CimType.UInt16, "uint16")]
        [DataRow(CimType.UInt32, "uint32")]
        [DataRow(CimType.UInt64, "uint64")]
        [DataRow(CimType.Real32, "float32")]
        [DataRow(CimType.Real64, "float64")]
        [DataRow(CimType.DateTime, "string")]
        [DataRow(CimType.Reference, "interface{}")]
        [DataRow(CimType.Object, "interface{}")]
        public void CimTypesPreserveGoWidthsAndWireRepresentations(CimType type, string expected)
        {
            Assert.AreEqual(expected, GOWmiSource.ConvertCimTypeToGoType(type));
        }

        [TestMethod]
        public void UnsupportedCimTypeFailsInsteadOfEmittingAnInvalidType()
        {
            Assert.ThrowsException<NotImplementedException>(
                () => GOWmiSource.ConvertCimTypeToGoType((CimType)int.MaxValue));
        }

        [DataTestMethod]
        [DataRow("type", "_type")]
        [DataRow("namespace", "Namespace")]
        [DataRow("DisplayName", "DisplayName")]
        public void PropertyIdentifiersAreEscapedWithoutChangingWmiLookupNames(string name, string field)
        {
            var source = EmitterFixture.CreateSource();
            // An unbound, unsaved schema: no Get/Put, queries or instance operations.
            using (var schema = new ManagementClass())
            {
                schema.Properties.Add(name, CimType.String, false);
                var property = new GOWmiProperty(schema.Properties[name], (GOWmiClass)source.Class);
                Assert.AreEqual(field + " string",
                    property.GetSourceCode().Split('\n').Last().Trim());
                StringAssert.Contains(property.Getter, "instance.GetProperty(\"" + name + "\")");
                StringAssert.Contains(property.Setter, "instance.SetProperty(\"" + name + "\"");
                StringAssert.Contains(property.Getter, "GetProperty" + name + "()");
            }
        }

        [TestMethod]
        public void ArrayPropertyEmitsSliceAndConvertsEachElement()
        {
            var source = EmitterFixture.CreateSource();
            using (var schema = new ManagementClass())
            {
                schema.Properties.Add("Codes", CimType.UInt32, true);
                var property = new GOWmiProperty(schema.Properties["Codes"], (GOWmiClass)source.Class);
                source.Class.Properties.Add(property);

                var code = source.Class.GetSourceCode();
                StringAssert.Contains(code, "Codes []uint32");
                StringAssert.Contains(code, "SetPropertyCodes(value []uint32)");
                StringAssert.Contains(code, "GetPropertyCodes()(value []uint32, err error)");
                StringAssert.Contains(code, "range retValue.([]interface{})");
                StringAssert.Contains(code, "interfaceValue.(uint32)");
                StringAssert.Contains(code, "value = append(value, uint32(valuetmp))");
            }
        }

        [DataTestMethod]
        [DataRow("namespace", ParamType.Output, false, false, "/* OUT */ Namespace string")]
        [DataRow("volatile", ParamType.Input, true, false, "/* OPTIONAL IN */ Volatile string")]
        [DataRow("Values", ParamType.Reference, false, true, "/* IN/OUT */ Values []string")]
        public void MethodParametersEmitEscapedNamesDirectionsAndSlices(
            string name, ParamType direction, bool optional, bool array, string expected)
        {
            var source = EmitterFixture.CreateSource();
            var method = new GOWmiMethod { Parent = source.Class };
            using (var schema = new ManagementClass())
            {
                schema.Properties.Add(name, CimType.String, array);
                var parameter = new GOWmiMethodParam(schema.Properties[name],
                    direction, method, optional);
                Assert.AreEqual(expected, parameter.GetSourceCode().Trim());
            }
        }

        [TestMethod]
        public void NamespaceImportsUseConfiguredRootAndAreDeduplicated()
        {
            var source = EmitterFixture.CreateSource();
            source.AddReference("root/standardcimv2");
            source.AddReference("root/standardcimv2");

            Assert.AreEqual(1, source.References.Count(
                r => r.Reference == "example.com/wmi/server/root/standardcimv2"));
            var code = source.GetSourceCode();
            StringAssert.Contains(code, "standardcimv2 \"example.com/wmi/server/root/standardcimv2\"");
            StringAssert.Contains(code, "cim \"github.com/microsoft/wmi/pkg/wmiinstance\"");
            StringAssert.Contains(code, "type Test_Device struct");
            StringAssert.Contains(code, "*cim.WmiInstance");
        }

        [DataTestMethod]
        [DataRow("default", "wmidefault")]
        [DataRow("DEFAULT", "wmidefault")]
        [DataRow("cimv2", "cimv2")]
        public void SourcePackageAvoidsTheDefaultKeyword(string module, string expected)
        {
            StringAssert.Contains(EmitterFixture.CreateSource(module).GetSourceCode(),
                "package " + expected + "\n");
        }
    }
}
