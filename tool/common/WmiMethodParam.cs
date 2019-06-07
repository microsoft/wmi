// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Management;

namespace Microsoft.WmiCodeGen.Common
{
    public enum ParamType
    {
        Input,
        Output,
        Reference,
        Custom,
    }

    public abstract class WmiMethodParam : IWmiBase
    {
        public ParamType ParamterType { get; set; }
        /// <summary>
        /// 
        /// </summary>
        public WmiMethod Parent { get; private set; }
        /// <summary>
        /// 
        /// </summary>
        public bool Optional { get; set; }

        /// <summary>
        /// 
        /// </summary>
        //public bool Custom { get; set; }
        public WmiMethodParam(WmiMethod wMethod)
        {
            Parent = wMethod;
        }
        public WmiMethodParam(PropertyData pData, ParamType pType, WmiMethod wMethod, bool optional) //, bool custom)
        {
            Parent = wMethod;
            ParamterType = pType;
            Name = pData.Name;
            Optional = optional;
            Type tmpType;
            Type = Parent.Parent.Parent.GetType(pData, out tmpType);
            SType = tmpType;
            //Optional = optional;
            //Custom = custom;
            IsArray = pData.IsArray;
        }

        public WmiMethodParam(string comment, ParamType pType, string Type, string Name, bool custom)
        {
            ParamterType = pType;
            Comment = comment;
        }

        public string Comment { get; set; }
        public string Type { get; set; }
        public Type SType { get; set; }
        //public bool IsInput { get; set; }
        public bool IsArray { get; set; }

        public override string ToString()
        {
            return string.Format(CultureInfo.InvariantCulture,
                "WmiMethodParam [Name:{0}] [Return Type:{1}]", Name, Type);
        }
    }
}
