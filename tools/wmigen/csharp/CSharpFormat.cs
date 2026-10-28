// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using Microsoft.wmi.Common;

namespace Microsoft.wmi.CSharp
{
    public class CSharpFormat : IFormat
    {
        #region Constructor
        public CSharpFormat(string[] wmiNamespaces, string outDir, bool recurse) : base(wmiNamespaces, outDir, recurse)
        {
        }
        #endregion

        protected override WmiNamespace GetWmiNamespace(string wmins)
        {
            return new CSWmiNamespace(wmins);
        }
    }
}
