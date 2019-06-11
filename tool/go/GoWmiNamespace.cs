// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Management;
using System.Text;
using Microsoft.WmiCodeGen.Common;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiNamespace : WmiNamespace
    {
        public GOWmiNamespace(string name) : base(name)
        {
        }

        protected override WmiReference GetReference(string reference)
        {
            return new GOWmiReference(reference);
        }

        protected override WmiSource GetWmiSource(ManagementClass wmiClass, WmiModule wModule)
        {
            return new GOWmiSource(wmiClass, wModule);
        }

        public override string GetSourceCode()
        {
            throw new NotImplementedException();
        }
    }
}
