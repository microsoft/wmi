// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Management;

namespace Microsoft.WmiCodeGen.CSharp
{
    public class CSWmiMethodParam : WmiMethodParam
    {
        /// <summary>
        /// 
        /// </summary>
        //public bool Custom { get; set; }
        public CSWmiMethodParam(WmiMethod wMethod, string sourceCode) : base(wMethod)
        {
            SourceCode = sourceCode;
        }

        public CSWmiMethodParam(PropertyData pData, ParamType pType, CSWmiMethod wMethod, bool optional) : base(pData, pType, wMethod, optional)
        {
            Comment = String.Format(CultureInfo.InvariantCulture, 
                " /* {0} */",
                ParamterType == ParamType.Input ? (optional ? "OPTIONAL IN" : "IN") : 
                    (ParamterType == ParamType.Reference ? "IN/OUT" : "OUT")
                );
            SourceCode = string.Format(CultureInfo.InvariantCulture,
                //"{/*Comment*/} {4} {1} {2} {4} {3}", 
                "{0} {1} {2} {3} {4} {5} {6} ", 
                Comment, 
                optional ? "[Optional]" : "",
                ParamterType == ParamType.Output ? "out" : (ParamterType == ParamType.Reference ? "ref" : ""), 
                Type, 
                optional ? (SType.IsClass ? "" : "?") : "",
                IsArray ? "[]" : "",
                WmiMethod.FixName(Name)
            );
        }

        public CSWmiMethodParam(string Comment, ParamType pType, string Type, string Name, bool custom)
            : base(Comment, pType, Type, Name, custom)
        {
            SourceCode = string.Format(CultureInfo.InvariantCulture,
                "{0} [Optional, DefaultValue(default({2}))] {1} {2} {3}", 
                Comment,
                ParamterType == ParamType.Output ? "out" : "", 
                Type, 
                WmiMethod.FixName(Name));
        }

        public string SourceCode { get; set; }

        public override string GetSourceCode()
        {
            return SourceCode;
        }
    }
}
