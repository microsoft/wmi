// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Management;
using System.Text;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiConstructor : WmiConstructor
    {
        public GOWmiConstructor(GOWmiClass wClass, string derivation) : base(wClass, derivation)
        {
        }

        public override string GetSourceCode()
        {
            StringBuilder sb = new StringBuilder();
            sb.AppendLine("\n#region Constructor " + Name);
            foreach (var item in HeaderComments)
            {
                sb.AppendLine(item);
            }
            sb.AppendLine(Declaration);
            sb.AppendFormat(CultureInfo.InvariantCulture, "public {0} {1}", ReturnType, Name);

            sb.AppendLine("(");
            if (Params.Count > 0)
            {
                foreach (var item in Params)
                {
                    sb.Append(item.GetSourceCode().Replace("\n", "\n\t"));
                    sb.AppendLine(",");
                }
                sb.Remove(sb.Length - 3, 3);
            }
            sb.Append(")");
            if (!string.IsNullOrEmpty(Derivation))
            {
                sb.AppendFormat(CultureInfo.InvariantCulture, "\t: {0} \n", Derivation);
            }
            sb.AppendLine("{");
            sb.AppendLine(BodyText.Replace("\n", "\n\t"));
            sb.AppendLine("}");
            sb.AppendLine("#endregion Constructor " + Name);
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
