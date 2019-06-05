// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Globalization;
using System.Management;
using System.Text;

namespace Microsoft.WmiCodeGen
{

    public abstract class WmiMethod : IWmiBase
    {
        public WmiMethod() { }
        public WmiClass Parent { get; set; }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="mData"></param>
        /// <param name="wClass"></param>
        public WmiMethod(MethodData mData, WmiClass wClass)
        {
            Parent = wClass;
            Name = mData.Name;
            Type tmpType;
            GetMethodReturnType(mData, out tmpType);
            ReturnSystemType = tmpType;
            Params.AddRange(GetWmiMethodParams(mData));
            BodyText = GetMethodBodyText(mData);
            ReturnType = (ReturnTypeIsArray ? "[]" : "");
        }
        public string ReturnType { get; set; }

        public bool ReturnTypeIsArray { get; set; }
        public Type ReturnSystemType { get; set; }
        List<string> m_comment = new List<string>();
        public List<string> HeaderComments { get { return m_comment; } }
        public string Declaration { get; set; }
        public string BodyText { get; set; }

        List<WmiMethodParam> m_params = new List<WmiMethodParam>();

        public List<WmiMethodParam> Params { get { return m_params; } }

        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.AppendFormat(CultureInfo.InvariantCulture,
                "WmiMethod [Name:{0}] [Return Type:{1}]", Name, ReturnSystemType);
            foreach (var item in Params)
            {
                sb.AppendFormat(CultureInfo.InvariantCulture, "[{0}]\n", item.ToString());
            }
            return sb.ToString();
        }

        private PropertyData GetWmiMethodParamByID(MethodData mData, int index, out ParamType pType)
        {
            pType = ParamType.Input;
            PropertyData iParam = null;
            if (mData.InParameters != null)
                foreach (PropertyData item in mData.InParameters.Properties)
                {
                    try
                    {
                        if ((int)(item.Qualifiers["Id"].Value) == index ||
                                    (int)(item.Qualifiers["ID"].Value) == index)
                        {
                            pType = ParamType.Input;
                            iParam = item;
                            break;
                        }
                    }
                    catch (Exception)
                    {
                    }
                }

            PropertyData oParam = null;
            if (mData.OutParameters != null)
                foreach (PropertyData item in mData.OutParameters.Properties)
                {
                    try
                    {
                        if ((int)(item.Qualifiers["Id"].Value) == index ||
                                    (int)(item.Qualifiers["ID"].Value) == index)
                        {
                            pType = ParamType.Output;
                            oParam = item;
                        }
                    }
                    catch (Exception)
                    {
                    }
                }

            if (iParam != null && oParam != null)
            {
                pType = ParamType.Reference;
                return iParam;
            }
            else
            {
                return iParam != null ? iParam : oParam;
            }
        }

        protected abstract WmiMethodParam GetWmiMethodParam(PropertyData pData, ParamType pType, WmiMethod wMethod, bool optional);
        protected abstract WmiMethodParam GetWmiMethodParam(string comment, ParamType pType, string Type, string Name, bool custom);

        private List<WmiMethodParam> GetWmiMethodParams(MethodData mData)
        {
            List<WmiMethodParam> wMethodParams = new List<WmiMethodParam>();

            int paramCount = 0;
            bool outputParamSwitchFlag = false;
            bool optionalParamFlag = false;
            // input 
            if (mData.InParameters != null)
                paramCount += mData.InParameters.Properties.Count;

            if (mData.OutParameters != null)
                paramCount += mData.OutParameters.Properties.Count;

            {
                ParamType pType = ParamType.Input;
                for (int i = 0; i < paramCount; i++)
                {
                    PropertyData pData = GetWmiMethodParamByID(mData, i, out pType);
                    if (pData != null && pData.Name != "ReturnValue")
                    {
                        if (pType == ParamType.Output) outputParamSwitchFlag = true;
                        try
                        {
                            if (
                                (bool)(pData.Qualifiers["optional"].Value) == true ||
                                (bool)(pData.Qualifiers["Optional"].Value) == true
                            )
                            {
                                wMethodParams.Add(GetWmiMethodParam(pData, pType, this, true));
                                continue;
                            }
                        }
                        catch (Exception)
                        {
                        }

                        if (pType == ParamType.Input && outputParamSwitchFlag)
                            optionalParamFlag = true;

                        wMethodParams.Add(GetWmiMethodParam(pData, pType, this, optionalParamFlag));
                    }
                }
            }

            // output. TODO : optional output
            if (mData.OutParameters != null)
            {
#if false
                int outputParamCount = mData.OutParameters.Properties.Count;
                for (int i = 0; i <= outputParamCount; i++)
                {
                    PropertyData outParam = GetWmiMethodParamByID(mData.OutParameters.Properties, i);
                    if (outParam != null && outParam.Name != "ReturnValue")
                    {
                        WmiMethodParam wParam = wMethodParams.Find(p => p.Name == outParam.Name);
                        if (wParam != null)
                        {
                            //wParam.GetSourceCode().Replace("IN */  ", "INOUT */ out ");
                            wParam.GetSourceCode = string.Format(CultureInfo.InvariantCulture,
                                "{0} {1} {2} {3}", " /* IN/OUT */", "ref", wParam.Type, wParam.Name);
                        }
                        else
                        {
                            wMethodParams.Add(new WmiMethodParam(outParam, false, this, false));
                        }
                    }
                } 
#endif

                if (CSharpFormat.HasJobOutputParams(mData))
                {
                    wMethodParams.Add(GetWmiMethodParam("/*Custom IN*/", ParamType.Custom, "UserAction", "Action", true));
                    wMethodParams.Add(GetWmiMethodParam("/*Custon IN*/", ParamType.Custom, "uint", "PercentComplete", true));
                    wMethodParams.Add(GetWmiMethodParam("/*Custon IN*/", ParamType.Custom, "int", "Timeout", true));
                }
            }

#if false
            // Sort the params
            List<WmiMethodParam> sortedwParams = new List<WmiMethodParam>();
            sortedwParams.AddRange(wMethodParams.FindAll(p => p.IsInput && !p.Optional));
            sortedwParams.AddRange(wMethodParams.FindAll(p => !p.IsInput));
            sortedwParams.AddRange(wMethodParams.FindAll(p => p.IsInput && p.Optional));
            sortedwParams.AddRange(wMethodParams.FindAll(p => p.IsInput && p.Custom));

            return sortedwParams; 
#endif

            return wMethodParams;
        }

#if false
        private List<WmiMethodParam> GetWmiMethodParams(
            PropertyDataCollection collection, bool input)
        {
            int expectedParamId = 0;
            bool optionalParamFlag = false;
            List<WmiMethodParam> wMethodParams = new List<WmiMethodParam>();
            if (collection != null)
            {
                foreach (PropertyData pData in collection)
                {
                    if (pData.Name != "ReturnValue")
                    {
                        // Find the optional Input Parameter using the ID qualifier Tag
                        if (input)
                        {
                            QualifierData qualifier = pData.Qualifiers["ID"];
                            if (qualifier != null)
                            {
                                if ((int)qualifier.Value != expectedParamId++)
                                {
                                    optionalParamFlag = true;
                                }
                            }
                        }

                        wMethodParams.Add(new WmiMethodParam(pData, input, this, optionalParamFlag));
                    }
                }
            }
            return wMethodParams;
        } 
#endif

        /// <summary>
        /// 
        /// </summary>
        /// <param name="mData"></param>
        /// <returns></returns>
        protected abstract string GetMethodBodyText(MethodData mData);
        private void GetMethodReturnType(MethodData mData, out Type returnType)
        {
            string returnValue = "void";
            ReturnTypeIsArray = false;
            returnType = typeof(void);
            if (mData.OutParameters != null)
            {
                foreach (var outParam in mData.OutParameters.Properties)
                {
                    if (outParam.Name == "ReturnValue")
                    {
                        returnValue = Parent.Parent.GetType(outParam, out returnType);
                        ReturnTypeIsArray = outParam.IsArray;
                        break;
                    }
                }
            }
        }

        private List<Type> convertibleTypesList = new List<Type>() {
            typeof(bool),
            typeof(UInt64),
            typeof(UInt16),
            typeof(UInt32),
            typeof(Int16),
            typeof(Int32),
            typeof(Int64),
            typeof(Double),
            typeof(Single),
            typeof(String),
            typeof(SByte),
            typeof(Char),
            typeof(Byte),
            typeof(DateTime),
            typeof(Decimal),
        };
        protected bool IsSystemConvertible(Type type)
        {
            if (convertibleTypesList.Contains(type)) return true;
            return false;
        }

        public static string FixName(string name)
        {
            if (name.Equals("volatile")) return "Volatile";
            if (name.Equals("namespace")) return "Namespace";

            return name;
        }
    }

}
