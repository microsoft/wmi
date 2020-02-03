// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

namespace Microsoft.wmi.Common
{
    public abstract class WmiConstructor : WmiMethod
    {
        public WmiConstructor(WmiClass wClass, string derivation)
        {
            Parent = wClass;
            Derivation = derivation;
        }

        public string Derivation { get; set; }
    }
}
