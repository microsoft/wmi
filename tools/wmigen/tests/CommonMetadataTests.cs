// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.IO;
using System.Management;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.Tests
{
    [TestClass]
    [DoNotParallelize]
    public class CommonMetadataTests
    {
        [TestMethod]
        public void QualifierLookupIsCaseInsensitiveAndDistinguishesMissingValues()
        {
            using (var schema = new ManagementClass())
            {
                schema.Qualifiers.Add("Description", "Device metadata");
                object value;
                Assert.IsTrue(IFormat.TryGetQualifierValue(schema.Qualifiers, "DESCRIPTION", out value));
                Assert.AreEqual("Device metadata", value);
                Assert.IsTrue(IFormat.HasQualifier(schema.Qualifiers, "description"));
                Assert.IsFalse(IFormat.TryGetQualifierValue(schema.Qualifiers, "Missing", out value));
                Assert.IsNull(value);
            }
        }

        [DataTestMethod]
        [DataRow(null, "", "")]
        [DataRow("", "", "")]
        [DataRow("A \"quoted\" device", "A \"quoted\" device", "\n[Description(@\"A 'quoted' device\")]")]
        public void DescriptionAttributesOmitEmptyMetadataAndEscapeQuotes(
            string input, string description, string attribute)
        {
            using (var schema = new ManagementClass())
            {
                if (input != null)
                    schema.Qualifiers.Add("Description", input);
                Assert.AreEqual(description, IFormat.GetDescription(schema.Qualifiers));
                Assert.AreEqual(attribute, IFormat.GetDescriptionText(schema.Qualifiers));
            }
        }

        [DataTestMethod]
        [DataRow(null, false, false)]
        [DataRow("Result", true, false)]
        [DataRow("Job", true, true)]
        public void MethodSignaturesDistinguishReturnValuesOutputsAndAsyncJobs(
            string output, bool hasOutput, bool hasJob)
        {
            using (var schema = new ManagementClass())
            using (var signature = new ManagementClass())
            {
                signature.Properties.Add("ReturnValue", CimType.UInt32, false);
                if (output != null)
                {
                    signature.Properties.Add(output, CimType.String, false);
                    signature.Properties[output].Qualifiers.Add("ID", 0);
                }
                schema.Methods.Add("Run", null, signature);
                var method = schema.Methods["Run"];
                Assert.AreEqual(hasOutput, IFormat.HasOutParams(method));
                Assert.AreEqual(hasJob, IFormat.HasJobOutputParams(method));
            }
        }

        [TestMethod]
        public void EnumValueMapsPreserveNonSequentialAndNegativeWireValues()
        {
            var source = EmitterFixture.CreateSource();
            using (var schema = new ManagementClass())
            {
                schema.Properties.Add("State", CimType.SInt32, false);
                var property = schema.Properties["State"];
                property.Qualifiers.Add("Values", new[] { "Unknown", "Ready", "Busy" });
                property.Qualifiers.Add("ValueMap", new[] { "-1", "4", "16" });
                var values = source.GetEnum(property);
                Assert.AreEqual(-1, values.EnumValues["Unknown"]);
                Assert.AreEqual(4, values.EnumValues["Ready"]);
                Assert.AreEqual(16, values.EnumValues["Busy"]);
                Assert.AreSame(values, source.GetEnum(property));
            }
        }

        [TestMethod]
        public void ExceptionLoggingPropagatesTheOriginalFailure()
        {
            InvalidOperationException expected;
            try
            {
                throw new InvalidOperationException("Generation failed");
            }
            catch (InvalidOperationException error)
            {
                expected = error;
            }

            var originalStack = expected.StackTrace;
            var originalOutput = Console.Out;
            using (var output = new StringWriter())
            {
                try
                {
                    Console.SetOut(output);
                    var actual = Assert.ThrowsException<InvalidOperationException>(
                        () => Logger.Exception(expected, "Namespace {0}", "root/cimv2"));
                    Assert.AreSame(expected, actual);
                    StringAssert.Contains(output.ToString(), "Generation failed");
                    StringAssert.Contains(output.ToString(), originalStack);
                    StringAssert.Contains(output.ToString(), "Namespace root/cimv2");
                }
                finally
                {
                    Console.SetOut(originalOutput);
                }
            }
        }
    }
}
