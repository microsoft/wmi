// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Management;
using System.Text;
using Microsoft.WmiCodeGen.Common;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiMethod : WmiMethod
    {
        public GOWmiMethod() { }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="mData"></param>
        /// <param name="wClass"></param>
        public GOWmiMethod(MethodData mData, GOWmiClass wClass)
            : base(mData, wClass)
        {
            HeaderComments.Add(GOWmiSource.GetHeaderCommentText(mData.Qualifiers));
            if (mData.InParameters != null)
                HeaderComments.Add(GetWmiMethodComment(mData.InParameters.Properties));
            if (mData.OutParameters != null)
                HeaderComments.Add(GetWmiMethodComment(mData.OutParameters.Properties));
        }

        public override string GetSourceCode()
        {
            StringBuilder sb = new StringBuilder();
            foreach (var item in HeaderComments)
            {
                sb.AppendLine(item);
            }
            //sb.AppendLine(Declaration);
            sb.AppendFormat(CultureInfo.InvariantCulture, "func (instance *{1}) {0}(", FixName(Name), Parent.Name);
            if (Params.Count > 0)
            {
                foreach (var item in Params)
                {
                    sb.AppendFormat(CultureInfo.InvariantCulture, "{0},\n", item.GetSourceCode().Replace("\n", "\n\t"));
                }
                sb.Remove(sb.Length - 2, 1);
            }
            sb.AppendFormat(") ({0}, error) {{", ReturnType);
            sb.AppendLine(BodyText.Replace("\n", "\n\t"));
            sb.AppendLine("}");


            return sb.ToString();
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="collection"></param>
        /// <returns></returns>
        private string GetWmiMethodComment(PropertyDataCollection collection)
        {
            StringBuilder sbComment = new StringBuilder();
            if (collection != null)
            {
                Type tmpType;
                foreach (var pData in collection)
                {
                    sbComment.AppendFormat(CultureInfo.InvariantCulture,
                            "\n// <param name=\"{0}\" type=\"{1} {3}\">{2}</param>",
                            pData.Name,
                            Parent.Parent.GetType(pData, out tmpType),
                            IFormat.GetDescription(pData.Qualifiers).Replace("\n", "\n///"),
                            pData.IsArray ? "[]" : "");
                }
            }
            return sbComment.ToString();
        }
        protected override WmiMethodParam GetWmiMethodParam(PropertyData pData, ParamType pType, WmiMethod wMethod, bool optional)
        {
            return new GOWmiMethodParam(pData, pType, wMethod as GOWmiMethod, optional);
        }

        protected override WmiMethodParam GetWmiMethodParam(string comment, ParamType pType, string Type, string Name, bool custom)
        {
            return new GOWmiMethodParam(comment, pType, Type, Name, custom);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="mData"></param>
        /// <returns></returns>
        protected override string GetMethodBodyText(MethodData mData)
        {
            StringBuilder sb = new StringBuilder();
            if (mData.InParameters != null)
            {
                StringBuilder sbInParams = new StringBuilder();
                sbInParams.AppendFormat("\n arguments  := []MethodParameter");

                foreach (var item in Params.FindAll(p => p.ParamterType == ParamType.Input))
                {
                    if (item.Optional)
                        sbInParams.AppendFormat(CultureInfo.InvariantCulture, "if ({0} != null) ", WmiMethod.FixName(item.Name));

                    sbInParams.AppendFormat(CultureInfo.InvariantCulture, "wParam := MethodParameter{{ Name: \"{0}\", Value : {0} }}", WmiMethod.FixName(item.Name));
                    sbInParams.AppendFormat(CultureInfo.InvariantCulture, "arguments = append(arguments, .wParam);\n");
                }

                sb.Append(sbInParams.ToString());
            }

            //string returnType = GetMethodReturnType(mData);
            if (IFormat.HasOutParams(mData))
            {
                if (IFormat.HasJobOutputParams(mData))
                {
                    sb.AppendFormat(CultureInfo.InvariantCulture,
                        "\nmethodResult, err := InvokeMethod(\"{0}\", {1}, null, Action, PercentComplete, Timeout)\n",
                        mData.Name, mData.InParameters != null ? "arguments" : "null");
                }
                else
                {
                    sb.AppendFormat(CultureInfo.InvariantCulture,
                        "\nmethodResult, err := InvokeMethod(\"{0}\", {1})\n",
                        mData.Name, mData.InParameters != null ? "arguments" : "null");
                }

                Type tmpType;

                foreach (var pData in mData.OutParameters.Properties)
                {
                    if (!pData.Name.Equals("ReturnValue", StringComparison.OrdinalIgnoreCase))
                    {
                        string typeString = Parent.Parent.GetType(pData, out tmpType);
                        string outParamGetValueMethod = pData.IsArray ? "GetValueArray" : "GetValue";

                        sb.AppendFormat(CultureInfo.InvariantCulture,
                            "\t{0} = methodResult.OutParameters.{2}(\"{0}\");\n",
                            pData.Name, typeString, outParamGetValueMethod);
                        sb.AppendLine();
                        // Assuming Job param comes first, then the resulting object param

                        // Sometimes, the resulting object might be null. Try to 
                        // Get it via Job, in case of a Wmi Class 
                        // Assumption is that the out param type is a Wmi class 
                        if (IFormat.HasJobOutputParams(mData) && tmpType.IsClass
                            && pData.Name != "Job" && pData.Name.Contains("Resulting"))
                        {
                            sb.AppendLine("\tif (Action != UserAction.Async || Action != UserAction.Cancel) && PercentComplete == 100 {");
                            sb.AppendFormat(CultureInfo.InvariantCulture,
                                "\t\tif {0} == nil {{\n", pData.Name);
                            sb.AppendLine("\t\t\tif Job != null {");
                            if (pData.IsArray)
                            {
                                sb.AppendFormat(CultureInfo.InvariantCulture,
                                    "\t\t\t{0}, err := Job.GetAllRelated(\"{1}\").GetInstanceArray();\n",
                                    pData.Name, typeString);
                            }
                            else
                            {
                                sb.AppendFormat(CultureInfo.InvariantCulture,
                                    "\t\t{0}, err := Job.GetRelated(\"{1}\");\n",
                                    pData.Name, typeString);
                            }
                            sb.AppendLine("\t\t\t}");
                            sb.AppendLine("\t\t}");
                            sb.AppendLine("\t}");
                        }
                    }
                }

                if (!ReturnType.Equals("void", StringComparison.OrdinalIgnoreCase))
                {
                    sb.AppendFormat(CultureInfo.InvariantCulture,
                             "\treturn methodResult.ReturnValue.Value, err;\n", ReturnType);
                }
            }
            else
            {
                sb.AppendFormat(CultureInfo.InvariantCulture,
                             "result, err := InvokeMethodWithReturn(\"{0}\", {1}, null, returnValue interface{{}});\n",
                             mData.Name, mData.InParameters != null ? "arguments" : "null");

                if (ReturnType.Equals("void", StringComparison.OrdinalIgnoreCase))
                {
                    // Do nothing.
                }
                else
                {
                    sb.AppendFormat(CultureInfo.InvariantCulture,
                             "return result, err;\n");
                }
            }


            return sb.ToString();
        }
    }
}
