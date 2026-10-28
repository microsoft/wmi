// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

namespace Microsoft.wmi.Common
{
    public abstract class WmiRelated : IWmiBase
    {
        public WmiClass Parent { get; private set; }
        public WmiRelated(string name, WmiClass parent)
        {
            Parent = parent;
            Name = name;
        }
    }
}
