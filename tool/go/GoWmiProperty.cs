// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Management;
using System.Text;
using Microsoft.WmiCodeGen.Common;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiProperty : WmiProperty
    {
        public GOWmiProperty(PropertyData pData, GOWmiClass wClass) : base(pData, wClass)
        {
            HeaderComment = GOWmiSource.GetHeaderCommentText(pData.Qualifiers);
            Setter = GetPropertySetter(pData);
            Getter = GetPropertyGetter(pData);
            Name = pData.Name;
            //Name = FixPropertyName(pData.Name);
            //FixPropertyName(pData.Name, wmiClass);
        }

        public override string GetSourceCode()
        {
            StringBuilder sb = new StringBuilder();
            sb.AppendLine();
            sb.AppendLine(HeaderComment);
            sb.AppendFormat(CultureInfo.InvariantCulture,
                "{1} {2}{0}", Type, FixPropertyName(GOWmiMethod.FixName(Name)), IsArray ? "[]" : "");
            return sb.ToString();
        }

        public override string ToString()
        {
            return string.Format(CultureInfo.InvariantCulture, "WmiProperty [Name:{0}] [Type:{1}]", Name, Type);
        }
        private string GetPropertySetter(PropertyData pData)
        {
            return String.Format(CultureInfo.InvariantCulture,
@"
// Set{1} sets the value of {1} for the instance
func (instance {3}) Set{1}(value {2}{0}) error {{ 
    return SetProperty(""{1}"", value)
}}", Type, pData.Name, IsArray ? "[]" : "", Parent.Name);
        }

        private string GetPropertyGetter(PropertyData pData)
        {
            return String.Format(CultureInfo.InvariantCulture,
@"
// Get{1} gets the value of {1} for the instance
func (instance {3}) Get{1}(value {2}{0}) error {{ 
    return GetProperty(""{1}"", value)
}}", Type, pData.Name, IsArray ? "[]" : "", Parent.Name);
        }

    }
}
