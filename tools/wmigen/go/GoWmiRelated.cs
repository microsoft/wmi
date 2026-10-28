// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Linq;
using System.Text;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.GO
{
    public class GOWmiRelated : WmiRelated
    {
        public GOWmiRelated(string name, GOWmiClass parent) : base(name, parent)
        {
            parent.Parent.AddReference("github.com/microsoft/wmi/pkg/wmiinstance");
        }

        public string GetSourceCode(bool multipleElement)
        {
            StringBuilder sb = new StringBuilder();
            
            string methodName = multipleElement ? "GetAllRelated" : "GetRelated";
            string propertyName = Name.Contains('_') ? Name.Split('_')[1] : Name;
            //sb.AppendFormat(CultureInfo.InvariantCulture, "func  (instance* {3}) GetRelated{0}() (value {2}*{1}, err error) {{\n", 
             //   propertyName, Name, multipleElement ? "[]" : "", Parent.Name);
            sb.AppendFormat(CultureInfo.InvariantCulture, "func  (instance* {3}) GetRelated{0}() (value {2}*{1}, err error) {{\n", 
                propertyName, "cim.WmiInstance", multipleElement ? "[]" : "", Parent.Name);
            sb.AppendFormat(CultureInfo.InvariantCulture, "\t return instance.{1}(\"{0}\"); \n", Name, methodName);
            sb.AppendLine("}");
            return sb.ToString();
        }

        public override string GetSourceCode()
        {
            throw new NotImplementedException();
        }
    }
}
