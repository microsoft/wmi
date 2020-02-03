// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System.Globalization;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.CSharp
{
    public class CSWmiReference : WmiReference
    {
        public CSWmiReference(string reference) : base(reference)
        {
        }

        public override string GetSourceCode()
        {
            return string.Format(CultureInfo.InvariantCulture, "using {0};", Reference);
        }

        

    }
}
