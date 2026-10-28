// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Management;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Microsoft.wmi.Common;
using Microsoft.wmi.GO;

namespace Microsoft.wmi.Tests
{
    [TestClass]
    public class Server2025RegressionTests
    {
        [TestMethod]
        public void LdapReferencesCanResolveClassesInTheirOwnNamespace()
        {
            const string ns = "root/directory/LDAP";
            List<string> previous;
            bool hadPrevious = WmiSource.ClassList.TryGetValue(ns, out previous);
            try
            {
                WmiSource.ClassList[ns] = new List<string> { "__Provider" };
                Assert.IsTrue(EmitterFixture.CreateSource().CheckClass("__Provider", ns));
            }
            finally
            {
                if (hadPrevious)
                    WmiSource.ClassList[ns] = previous;
                else
                    WmiSource.ClassList.Remove(ns);
            }
        }

        [TestMethod]
        public void EnumReferencesReuseTheDeclarationCaseOnWindows()
        {
            var source = EmitterFixture.CreateSource();
            using (var propertySchema = new ManagementClass())
            using (var parameterSchema = new ManagementClass())
            {
                propertySchema.Properties.Add("DownloadMode", CimType.UInt32, false);
                parameterSchema.Properties.Add("downloadMode", CimType.UInt32, false);
                foreach (var property in new[] {
                    propertySchema.Properties["DownloadMode"],
                    parameterSchema.Properties["downloadMode"] })
                {
                    property.Qualifiers.Add("Values", new[] { "HTTP", "LAN" });
                    property.Qualifiers.Add("ValueMap", new[] { "0", "1" });
                }

                Type systemType;
                WmiReference reference;
                var declaration = source.GetType(propertySchema.Properties["DownloadMode"],
                    out systemType, out reference);
                var parameter = source.GetType(parameterSchema.Properties["downloadMode"],
                    out systemType, out reference);

                Assert.AreEqual("Device_DownloadMode", declaration);
                Assert.AreEqual(declaration, parameter);
                Assert.AreSame(source.Parent.Parent.GetEnum(declaration),
                    source.Parent.Parent.GetEnum("Device_downloadMode"));
                Assert.AreEqual(1, source.Parent.Parent.Modules["Common"].Enums.Count);
            }
        }

        [DataTestMethod]
        [DataRow("downloadMode", "HTTP", "2")]
        [DataRow("downloadMode", "HTTPS", "1")]
        [DataRow("DownloadMode", "HTTP", "2")]
        public void ConflictingEnumDefinitionsRemainDistinct(
            string duplicateName, string firstName, string secondValue)
        {
            var source = EmitterFixture.CreateSource();
            using (var declaration = new ManagementClass())
            using (var conflicting = new ManagementClass())
            {
                declaration.Properties.Add("DownloadMode", CimType.UInt32, false);
                conflicting.Properties.Add(duplicateName, CimType.UInt32, false);
                var original = declaration.Properties["DownloadMode"];
                original.Qualifiers.Add("Values", new[] { "HTTP", "LAN" });
                original.Qualifiers.Add("ValueMap", new[] { "0", "1" });
                var duplicate = conflicting.Properties[duplicateName];
                duplicate.Qualifiers.Add("Values", new[] { firstName, "LAN" });
                duplicate.Qualifiers.Add("ValueMap", new[] { "0", secondValue });
                var canonical = source.GetEnum(original);

                var distinct = source.GetEnum(duplicate);

                Assert.AreNotSame(canonical, distinct);
                Assert.IsFalse(string.Equals(canonical.Name, distinct.Name,
                    StringComparison.OrdinalIgnoreCase));
                Assert.AreEqual(0, canonical.EnumValues["HTTP"]);
                Assert.AreEqual(1, canonical.EnumValues["LAN"]);
                Assert.AreEqual(0, distinct.EnumValues[firstName]);
                Assert.AreEqual(int.Parse(secondValue), distinct.EnumValues["LAN"]);
                Assert.AreSame(distinct, source.GetEnum(duplicate));
                Assert.AreSame(canonical, source.GetEnum(original));
                Assert.AreEqual(2, source.Parent.Parent.Modules["Common"].Enums.Count);
            }
        }

        [TestMethod]
        public void SelfReferencesUsePointersInsteadOfRecursiveValueFields()
        {
            var source = EmitterFixture.CreateSource();
            source.Parent.Parent.AddSource(source);
            using (var schema = new ManagementClass())
            {
                schema.Properties.Add("Callback", CimType.Reference, false);
                schema.Properties["Callback"].Qualifiers.Add("CimType", "ref:Test_Device");
                var property = new GOWmiProperty(schema.Properties["Callback"],
                    (GOWmiClass)source.Class);

                StringAssert.Contains(property.GetSourceCode(), "Callback *Test_Device");
                StringAssert.Contains(property.Setter, "SetPropertyCallback(value *Test_Device)");
                StringAssert.Contains(property.Getter, "retValue.(*Test_Device)");
                StringAssert.Contains(property.Getter, "value = (*Test_Device)(valuetmp)");
                StringAssert.Contains(property.Getter, "instance.GetProperty(\"Callback\")");
            }
        }
    }
}
