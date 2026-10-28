// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Diagnostics;
using System.IO;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Microsoft.wmi.GO;

namespace Microsoft.wmi.Tests
{
    [TestClass]
    public class CliTests
    {
        [DataTestMethod]
        [DataRow(false)]
        [DataRow(true)]
        public void HelpShortCircuitsGenerationEvenAfterGenerationOptions(bool generationOptions)
        {
            var directory = Path.Combine(Path.GetTempPath(), "wmigen-cli-" + Guid.NewGuid().ToString("N"));
            var arguments = generationOptions
                ? "-generatesource -ns root/wmigen_test_must_not_query -outdir \"" + directory + "\" -help"
                : "-help";
            try
            {
                using (var process = new Process())
                {
                    process.StartInfo = new ProcessStartInfo
                    {
                        FileName = typeof(GOWmiSource).Assembly.Location,
                        Arguments = arguments,
                        UseShellExecute = false,
                        RedirectStandardOutput = true,
                        RedirectStandardError = true,
                        CreateNoWindow = true
                    };
                    Assert.IsTrue(process.Start());
                    var stdout = process.StandardOutput.ReadToEndAsync();
                    var stderr = process.StandardError.ReadToEndAsync();
                    bool exited = process.WaitForExit(15000);
                    if (!exited)
                    {
                        process.Kill();
                        process.WaitForExit();
                    }
                    var output = stdout.GetAwaiter().GetResult();
                    var error = stderr.GetAwaiter().GetResult();
                    Assert.IsTrue(exited, "Help did not finish within 15 seconds.");
                    Assert.AreEqual(0, process.ExitCode, error);
                    StringAssert.Contains(output, "-generatesource");
                    StringAssert.Contains(output, "-rootmodulepath");
                    Assert.IsFalse(Directory.Exists(directory), "Help must not start generation.");
                }
            }
            finally
            {
                if (Directory.Exists(directory))
                    Directory.Delete(directory, true);
            }
        }
    }
}
