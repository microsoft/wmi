// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Linq;
using System.Text;
using Microsoft.WmiCodeGen.Common;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiRelated : WmiRelated
    {
        public GOWmiRelated(string name, GOWmiClass parent) : base(name, parent)
        {
        }

        public string GetSourceCode(bool multipleElement)
        {
            StringBuilder sb = new StringBuilder();
            string specifier = "virtual";
            if (Parent.Parent.Parent.Parent.HasRelated(Parent.Derivation, Name))
            {
                specifier = "override";
            }
            string returnValue = multipleElement ? string.Format(CultureInfo.InvariantCulture, "WmiInstanceCollection<{0}>", Name) : Name;
            string methodName = multipleElement ? "GetAllRelated" : "GetRelated";
            string propertyName = Name.Contains('_') ? Name.Split('_')[1] : Name;
            sb.AppendLine("\n#region " + methodName);
            sb.AppendFormat(CultureInfo.InvariantCulture, "public {2} {1} Related{0}\n", 
                propertyName, returnValue, specifier);
            sb.AppendLine("{");
            sb.AppendFormat(CultureInfo.InvariantCulture, "\tget {{  return {1}<{0}>(\"{0}\"); }}\n", Name, methodName);
            sb.AppendLine("}");
            sb.AppendLine("#endregion " + methodName);
            return sb.ToString();
        }

        public override string GetSourceCode()
        {
            throw new NotImplementedException();
        }
    }
}
