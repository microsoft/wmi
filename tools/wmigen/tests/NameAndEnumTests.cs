// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

using Microsoft.VisualStudio.TestTools.UnitTesting;
using Microsoft.wmi.Common;
using Microsoft.wmi.GO;

namespace Microsoft.wmi.Tests
{
    [TestClass]
    public class NameAndEnumTests
    {
        [DataTestMethod]
        [DataRow("Not Applicable", "Not_Applicable")]
        [DataRow("Read/Write", "Read_Write")]
        [DataRow("C++", "Cplusplus")]
        [DataRow("32-bit", "_32_bit")]
        [DataRow("true", "_true")]
        [DataRow("interface", "_interface")]
        [DataRow("string", "_string")]
        public void EnumLabelsBecomeIdentifiers(string label, string expected)
        {
            Assert.AreEqual(expected, WmiEnum.FixEnumName(label));
        }

        [TestMethod]
        public void CollidingEnumLabelsRetainTheirValuesAndGetUniqueDeclarations()
        {
            var values = new GOWmiEnum("Device_State", EmitterFixture.CreateSource());
            values.Add(WmiEnum.FixEnumName("Read/Write"), 2);
            values.Add(WmiEnum.FixEnumName("Read Write"), 7);
            // An existing numeric suffix must not be overwritten by a collision.
            values.Add("Read_Write1", 11);
            values.Add(WmiEnum.FixEnumName("Read/Write"), 15);

            Assert.AreEqual(4, values.EnumValues.Count);
            var code = values.GetSourceCode();
            StringAssert.Contains(code, "type Device_State int");
            StringAssert.Contains(code, "Device_State_Read_Write Device_State = 2\n");
            StringAssert.Contains(code, "Device_State_Read_Write1 Device_State = 7\n");
            StringAssert.Contains(code, "Device_State_Read_Write11 Device_State = 11\n");
            StringAssert.Contains(code, "Device_State_Read_Write2 Device_State = 15\n");
        }

        [DataTestMethod]
        [DataRow("__namespace", "__Namespace")]
        [DataRow("__NAMESPACE", "__Namespace")]
        [DataRow("DN_WIth_Binary", "DN_With_Binary")]
        [DataRow("Win32_OperatingSystem", "Win32_OperatingSystem")]
        public void ClassNameRepairsPreserveGeneratedTypeIdentity(string name, string expected)
        {
            Assert.AreEqual(expected, WmiClass.FixClassName(name));
        }

        [DataTestMethod]
        [DataRow("namespace", "Namespace")]
        [DataRow("volatile", "Volatile")]
        [DataRow("Result", "Result")]
        public void ExistingMethodNameEscapesArePreserved(string name, string expected)
        {
            Assert.AreEqual(expected, WmiMethod.FixName(name));
        }
    }
}
