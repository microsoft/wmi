// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Linq;
using System.Management;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.GO
{
    public class GOWmiMethodParam : WmiMethodParam
    {
        /// <summary>
        /// 
        /// </summary>
        //public bool Custom { get; set; }
        public GOWmiMethodParam(WmiMethod wMethod, string sourceCode) : base(wMethod)
        {
            Name = sourceCode.Split(new char []{' '}).First();
            SourceCode = sourceCode;
            ReferencedType = (Reference != null ? Reference.Alias + "." : "") + Type;
        }


        public GOWmiMethodParam(PropertyData pData, ParamType pType, GOWmiMethod wMethod, bool optional) : base(pData, pType, wMethod, optional)
        {
            ReferencedType = (Reference != null ? Reference.Alias + "." : "") + Type;
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
                ReferencedType
            );
        }

        public GOWmiMethodParam(string Comment, ParamType pType, string paramType, string Name, bool custom)
            : base(Comment, pType, paramType, Name, custom)
        {
            if (pType == ParamType.Custom && paramType == "UserAction")
            {
                ReferencedType = "cim.UserAction";
            } else 
            {
                ReferencedType = (Reference != null ? Reference.Alias + "." : "") + Type;
            }

            SourceCode = string.Format(CultureInfo.InvariantCulture,
                "{0} {1} {3} {2}", 
                Comment,
                ParamterType == ParamType.Output ? "out" : "", 
                ReferencedType, 
                WmiMethod.FixName(Name));
        }

        public string ReferencedType { get; set; }
        public string SourceCode { get; set; }

        public override string GetSourceCode()
        {
            return SourceCode;
        }
    }
}
