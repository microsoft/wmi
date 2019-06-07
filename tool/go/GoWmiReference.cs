// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System.Globalization;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiReference : WmiReference
    {
        public GOWmiReference(string reference) : base(reference)
        {
        }

        public override string GetSourceCode()
        {
            return string.Format(CultureInfo.InvariantCulture, "using {0};", Reference);
        }

        

    }
}
