// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.
using Microsoft.WmiCodeGen.DotNet.Interface;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Microsoft.WmiCodeGen.DotNet.Lib
{
    /// <summary>
    /// Implementation for Cim_Error
    /// </summary>
    public class CimWmiError : WmiError
    {
        /// <summary>
        /// 
        /// </summary>
        public CimWmiError(
            string Error_Object) : base(Error_Object)
        {
        }
    }
}
