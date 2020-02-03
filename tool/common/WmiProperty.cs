// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Management;

namespace Microsoft.wmi.Common
{
    public abstract class WmiProperty : IWmiBase
    {
        public WmiClass Parent { get; private set; }
        public WmiProperty(PropertyData pData, WmiClass wClass)
        {
            Parent = wClass;
            Type tmpType;
            Type = wClass.Parent.GetType(pData, out tmpType);
            SType = tmpType;
            IsArray = pData.IsArray;
            Name = pData.Name;
            //FixPropertyName(pData.Name, wmiClass);
        }
        public string Declaration { get; set; }
        public string Type { get; set; }
        public Type SType { get; set; }
        public bool IsArray { get; set; }

        public string Setter { get; set; }
        public string Getter { get; set; }
        public override string ToString()
        {
            return string.Format(CultureInfo.InvariantCulture, "WmiProperty [Name:{0}] [Type:{1}]", Name, Type);
        }
        
        protected string FixPropertyName(string name)
        {
            string pName = name;

            while (true)
            {
                WmiSource wSource = Parent.Parent.Parent.Parent.GetSource(Parent.Derivation);
                while (wSource != null)
                {
                    if (wSource.Class.Methods.Find(m => m.Name == pName) != null)
                    {
                        pName = pName + Type; break;
                        // Found a method with the same name
                    }
                    if (wSource.Class.Properties.Find(m => m.Name == pName) != null)
                    {
                        //
                        pName = pName + Type; break;
                    }
                    wSource = Parent.Parent.Parent.Parent.GetSource(wSource.Class.Derivation);
                }
                if (wSource == null) break;
            }
            return pName;
        }
    }
}
