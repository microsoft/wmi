// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

namespace Microsoft.wmi.Common
{
    public abstract class WmiReference : IWmiBase
    {
        public WmiReference(string reference)
            :this(reference, string.Empty)
        {
        }
        public WmiReference(string reference, string alias)
        {
            Reference = reference;
            Alias = alias;
        }

        public string Reference { get; set; }
        public string Alias { get; set; }
    }
}
