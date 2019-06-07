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
        public string Setter { get; set; }
        public string Getter { get; set; }

        public override string GetSourceCode()
        {
            StringBuilder sb = new StringBuilder();
            sb.AppendLine();
            sb.AppendLine("#region Property " + Name);
            sb.AppendLine(HeaderComment);
            sb.AppendFormat(CultureInfo.InvariantCulture,
                "public virtual {0} {2} {1}", Type, FixPropertyName(GOWmiMethod.FixName(Name)), IsArray ? "[]" : "");
            sb.AppendLine("{");
            sb.AppendLine(Setter.Replace("\n", "\n\t"));
            sb.AppendLine(Getter.Replace("\n", "\n\t"));
            sb.AppendLine("}");
            sb.AppendLine("#endregion Property " + Name);
            return sb.ToString();
        }

        public override string ToString()
        {
            return string.Format(CultureInfo.InvariantCulture, "WmiProperty [Name:{0}] [Type:{1}]", Name, Type);
        }
        private string GetPropertySetter(PropertyData pData)
        {
            return String.Format(CultureInfo.InvariantCulture,
                "\nset {{ SetProperty<{0}{2}>(\"{1}\", value); }}\n", Type, pData.Name, IsArray ? "[]" : "");
              // "\nset {{ this[\"{0}\"] = value; }}", pData.Name);
        }

        private string GetPropertyGetter(PropertyData pData)
        {
            return String.Format(CultureInfo.InvariantCulture,
                "\nget {{ return ({0}{2})GetProperty<{0}>(\"{1}\"); }}\n", Type, pData.Name, IsArray ? "[]" : "");
        }

    }
}
