// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Management;
using Microsoft.WmiCodeGen.Common;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiMethodParam : WmiMethodParam
    {
        /// <summary>
        /// 
        /// </summary>
        //public bool Custom { get; set; }
        public GOWmiMethodParam(WmiMethod wMethod, string sourceCode) : base(wMethod)
        {
            SourceCode = sourceCode;
        }

        public GOWmiMethodParam(PropertyData pData, ParamType pType, GOWmiMethod wMethod, bool optional) : base(pData, pType, wMethod, optional)
        {
            Comment = String.Format(CultureInfo.InvariantCulture, 
                " /* {0} */",
                ParamterType == ParamType.Input ? (optional ? "OPTIONAL IN" : "IN") : 
                    (ParamterType == ParamType.Reference ? "IN/OUT" : "OUT")
                );
            SourceCode = string.Format(CultureInfo.InvariantCulture,
                //"{/*Comment*/} {4} {1} {2} {4} {3}", 
                "{0} {1} {2}{3}", 
                Comment, 
                WmiMethod.FixName(Name),
                IsArray ? "[]" : "",
                Type
            );
        }

        public GOWmiMethodParam(string Comment, ParamType pType, string Type, string Name, bool custom)
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
