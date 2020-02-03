// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

namespace Microsoft.wmi.Common
{
    public abstract class IWmiBase
    {
        public string Name { get; set; }
        public string CopyrightText { get; set; }
        public string HeaderComment { get; set; }
        abstract public string GetSourceCode();
    }
}
