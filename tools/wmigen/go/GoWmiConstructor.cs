// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Management;
using System.Text;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.GO
{
    public class GOWmiConstructor : WmiConstructor
    {
        public GOWmiConstructor(GOWmiClass wClass, string derivation) : base(wClass, derivation)
        {
        }

        public override string GetSourceCode()
        {
            StringBuilder sb = new StringBuilder();
            foreach (var item in HeaderComments)
            {
                sb.AppendLine(item);
            }
            sb.AppendLine(Declaration);
            sb.AppendFormat(CultureInfo.InvariantCulture, "func New{0}Ex{1}", Name,Params.Count);

            sb.Append("(");
            if (Params.Count > 0)
            {
                foreach (var item in Params)
                {
                    sb.Append(item.GetSourceCode().Replace("\n", "\n\t"));
                    sb.AppendLine(",");
                }
                sb.Remove(sb.Length - 3, 3);
            }
            sb.Append(") "); 
            sb.AppendFormat(CultureInfo.InvariantCulture, "(newInstance *{0}, err error) {{", Name);
            sb.AppendLine(BodyText.Replace("\n", "\n\t"));
            string tmpDerivation = Derivation;
            if (Derivation == "cim.WmiInstance") {
                tmpDerivation = "WmiInstance";
            } else {
            }
            sb.AppendLine(GOWmiSource.GetErrorHandler());
            sb.AppendFormat(CultureInfo.InvariantCulture, "newInstance = &{0} {{\n", Name);
            sb.AppendFormat(CultureInfo.InvariantCulture, "{0}: tmp,\n", tmpDerivation);
            sb.AppendLine("}");
            sb.AppendLine("return");
            sb.AppendLine("}");
            return sb.ToString();
        }

        protected override string GetMethodBodyText(MethodData mData)
        {
            throw new NotImplementedException();
        }

        protected override WmiMethodParam GetWmiMethodParam(PropertyData pData, ParamType pType, WmiMethod wMethod, bool optional)
        {
            throw new NotImplementedException();
        }

        protected override WmiMethodParam GetWmiMethodParam(string comment, ParamType pType, string Type, string Name, bool custom)
        {
            throw new NotImplementedException();
        }
    }
}
