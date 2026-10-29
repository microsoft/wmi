// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using Microsoft.wmi.Common;

namespace Microsoft.wmi.GO
{
    public class GoFormat : IFormat
    {
        private string m_rootmodulepath;
        #region Constructor
        public GoFormat(string[] wmiNamespaces, string outDir, bool recurse, string rootmodulepath) : base(wmiNamespaces, outDir, recurse)
        {
            m_rootmodulepath = rootmodulepath;
        }
        #endregion

        protected override WmiNamespace GetWmiNamespace(string wmins)
        {
            return new GOWmiNamespace(wmins, m_rootmodulepath);
        }
    }
}
