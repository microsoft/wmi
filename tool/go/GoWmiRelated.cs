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
            
            string methodName = multipleElement ? "GetAllRelated" : "GetRelated";
            string propertyName = Name.Contains('_') ? Name.Split('_')[1] : Name;
            sb.AppendFormat(CultureInfo.InvariantCulture, "func  GetRelated{0}(value {2}{1}) (error) {\n", 
                propertyName, Name, multipleElement ? "[]" : "");
            sb.AppendFormat(CultureInfo.InvariantCulture, "\t return {1}(\"{0}\", value); \n", Name, methodName);
            sb.AppendLine("}");
            return sb.ToString();
        }

        public override string GetSourceCode()
        {
            throw new NotImplementedException();
        }
    }
}
