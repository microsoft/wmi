// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Microsoft.WmiCodeGen.DotNet.Interface
{
    /// <summary>
    /// 
    /// </summary>
    public class OperationOption
    {
        public string Name { get; set; }
        public object Value { get; set; }
        public Type OperationType { get; set; }
        public bool Comply { get; set; }

        #region Constructor
        public OperationOption(string optionName, object value, Type opType, bool mustComply)
        {
            Name = optionName;
            Value = value;
            OperationType = opType;
            Comply = mustComply;
        }
        #endregion

    }

    /// <summary>
    /// 
    /// </summary>
    public class WmiOperationOptions
    {
        public List<OperationOption> Options { get; set; }
        public TimeSpan TimeOut { get; set; }

        #region Constructor
        public WmiOperationOptions()
        {
            Options = new List<OperationOption>();
            TimeOut = default(TimeSpan);
        }
        #endregion
        public void SetCustomOption(string optionName, object optionValue, Type type, bool mustComply)
        {
            OperationOption oo = new OperationOption(optionName, optionValue, type, mustComply);
            Options.Add(oo);
        }
    }
}
