// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System.Globalization;
using System.Linq;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.GO
{
    public class GOWmiReference : WmiReference
    {
        public GOWmiReference(string reference) : base(reference) {
            if (reference.Contains("root")) {
                Alias = reference.Split(new char [] {'/'} ).Last().ToLowerInvariant();
            }
            if (reference == "github.com/microsoft/wmi/pkg/wmiinstance") {
                Alias = "cim";
            }
        }

        public override string GetSourceCode()
        {
            return string.Format(CultureInfo.InvariantCulture, "{1} \"{0}\"", 
                Reference,
                string.IsNullOrEmpty(Alias) ? "" : Alias
                ).ToLowerInvariant();
        }
    }
}
