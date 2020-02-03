// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System.Globalization;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.GO
{
    public class GOWmiReference : WmiReference
    {
        public GOWmiReference(string reference) : base(reference)
        {
        }

        public override string GetSourceCode()
        {
            return string.Format(CultureInfo.InvariantCulture, "\"{0}\"", Reference);
        }

    }
}
