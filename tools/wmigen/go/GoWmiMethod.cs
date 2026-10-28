// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Management;
using System.Runtime.InteropServices;
using System.Text;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.GO
{
    public class GOWmiMethod : WmiMethod
    {
        public new static string FixName(string name)
        {
            switch (name)
            {
                case "break":
                case "default":
                case "func":
                case "interface":
                case "select":
                case "case":
                case "defer":
                case "go":
                case "map":
                case "struct":
                case "chan":
                case "else":
                case "goto":
                case "package":
                case "switch":
                case "const":
                case "fallthrough":
                case "if":
                case "range":
                case "type":
                case "continue":
                case "for":
                case "import":
                case "return":
                case "var":
                    return char.ToUpperInvariant(name[0]) + name.Substring(1);
                default:
                    return WmiMethod.FixName(name);
            }
        }

        private String GoReturnType { get; set; }
        public GOWmiMethod() { }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="mData"></param>
        /// <param name="wClass"></param>
        public GOWmiMethod(MethodData mData, GOWmiClass wClass)
            : base(mData, wClass)
        {
            // ReturnType = getReturnType(ReturnSystemType);
            GoReturnType = GOWmiSource.ConvertCimTypeToGoType(ReturnCimType);
            HeaderComments.Add(GOWmiSource.GetHeaderCommentText(mData.Qualifiers));
            if (mData.InParameters != null)
                HeaderComments.Add(GetWmiMethodComment(mData.InParameters.Properties));
            if (mData.OutParameters != null)
                HeaderComments.Add(GetWmiMethodComment(mData.OutParameters.Properties));
            BodyText = GetMethodBodyText(mData);
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
                sb.Remove(sb.Length - 2, 2);
            }
            if (GoReturnType == "none") {
                sb.AppendFormat(") (err error) {{");
            }else {
                sb.AppendFormat(") (result {0}, err error) {{", GoReturnType);
            }
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
                WmiReference reference;
                Type tmpType;
                foreach (var pData in collection)
                {
                    sbComment.AppendFormat(CultureInfo.InvariantCulture,
                            "\n// <param name=\"{0}\" type=\"{1} {3}\">{2}</param>",
                            pData.Name,
                            Parent.Parent.GetType(pData, out tmpType, out reference),
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
            if (pType == ParamType.Custom && Type == "UserAction") {
                this.Parent.Parent.AddReference("github.com/microsoft/wmi/pkg/wmiinstance");
            }
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
            StringBuilder sbInParams = new StringBuilder();
            if (mData.InParameters != null)
            {
                /*
                sbInParams.AppendFormat("\n arguments  := []MethodParameter{{}}\n");
                sbInParams.AppendFormat("\n wParam  := MethodParameter{{}}\n");
                */

                foreach (var item in Params.FindAll(p => p.ParamterType == ParamType.Input))
                {
                    /*
                    if (item.Optional)
                        sbInParams.AppendFormat(CultureInfo.InvariantCulture, "if ({0} != null) {{\n", WmiMethod.FixName(item.Name));

                    sbInParams.AppendFormat(CultureInfo.InvariantCulture, 
                        "wParam = MethodParameter{{ Name: \"{0}\", Value : {0} }}\n", 
                        WmiMethod.FixName(item.Name));
                    sbInParams.AppendFormat(CultureInfo.InvariantCulture, 
                        "arguments = append(arguments, wParam);\n");
                    if (item.Optional)
                        sbInParams.AppendFormat(CultureInfo.InvariantCulture, "}}\n");
                        */
                    //sbInParams.AppendFormat(",\"{0}\", {0}", WmiMethod.FixName(item.Name));
                    sbInParams.AppendFormat(", {0}", FixName(item.Name));
                }

                //sb.Append(sbInParams.ToString());
            }

            string inParamText = sbInParams.Length > 0 ? sbInParams.ToString() : "";

            if (GoReturnType.Equals("none", StringComparison.OrdinalIgnoreCase))
            {
                sb.Append("_, err = ");
            }
            else {
                sb.Append("retVal, err := ");
            }

            //string returnType = GetMethodReturnType(mData);
            if (IFormat.HasOutParams(mData))
            {
                if (IFormat.HasJobOutputParams(mData))
                {
                    sb.AppendFormat(CultureInfo.InvariantCulture,
                        "instance.InvokeMethodAsync(\"{0}\", Action, PercentComplete, Timeout {1})\n",
                        mData.Name, inParamText);
                }
                else
                {
                    sb.AppendFormat(CultureInfo.InvariantCulture,
                        "instance.InvokeMethod(\"{0}\" {1})\n",
                        mData.Name, inParamText);
                }
                sb.AppendLine(GOWmiSource.GetErrorHandler());

                /*
                Type tmpType;
                foreach (var pData in mData.OutParameters.Properties)
                {
                    if (!pData.Name.Equals("ReturnValue", StringComparison.OrdinalIgnoreCase))
                    {
                        string typeString = Parent.Parent.GetType(pData, out tmpType);
                        string outParamGetValueMethod = pData.IsArray ? "GetValueArray" : "GetValue";

                        sb.AppendFormat(CultureInfo.InvariantCulture,
                            "\terr = methodResult.OutParameters.{2}(\"{0}\", {0});\n",
                            pData.Name, typeString, outParamGetValueMethod);
                        sb.AppendLine(GOWmiSource.GetErrorHandler());
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
                                    "\t\t\t{0}, err = Job.GetAllRelated(\"{1}\").GetInstanceArray();\n",
                                    pData.Name, typeString);
                            }
                            else
                            {
                                sb.AppendFormat(CultureInfo.InvariantCulture,
                                    "\t\t{0}, err = Job.GetRelated(\"{1}\");\n",
                                    pData.Name, typeString);
                            }
                            sb.AppendLine("\t\t" + GOWmiSource.GetErrorHandler());
                            sb.AppendLine("\t\t\t}");
                            sb.AppendLine("\t\t}");
                            sb.AppendLine("\t}");
                        }
                    }
                }
                */
                if (GoReturnType.Equals("bool", StringComparison.OrdinalIgnoreCase))
                {
                    sb.AppendFormat(CultureInfo.InvariantCulture,
                             "retValue := retVal[0].(int32)\n", ReturnType);
                    sb.AppendFormat(CultureInfo.InvariantCulture,"result = bool(retValue != 0)\n");
                }
                else if (!GoReturnType.Equals("none", StringComparison.OrdinalIgnoreCase))
                {
                    sb.AppendFormat(CultureInfo.InvariantCulture,
                             "retValue := retVal[0].(int32)\n", ReturnType);
                    sb.AppendFormat(CultureInfo.InvariantCulture,"result = {0}(retValue)\n", GoReturnType);
                }
            }
            else
            {
                sb.AppendFormat(CultureInfo.InvariantCulture,
                             "instance.InvokeMethodWithReturn(\"{0}\" {1});\n",
                             mData.Name, inParamText);
                sb.AppendLine(GOWmiSource.GetErrorHandler());
                
                if (GoReturnType.Equals("none", StringComparison.OrdinalIgnoreCase))
                {
                    // Do nothing.
                }
                else
                {
                    if (GoReturnType == "bool"){
                        sb.AppendFormat(CultureInfo.InvariantCulture,"result = (retVal > 0)\n");
                    } else {
                        sb.AppendFormat(CultureInfo.InvariantCulture,"result = {0}(retVal)\n", GoReturnType);
                    }
                }
            }
            sb.AppendFormat(CultureInfo.InvariantCulture,
                             "return\n");
            return sb.ToString();
        }
    }
}
