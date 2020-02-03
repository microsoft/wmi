// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Microsoft.wmi.DotNet.Interface
{
    public enum WhereOperation
    {
        Equals = 0,
        LessThan,
        GreaterThan,
        LessThanEquals,
        GreaterThenEquals,
        NotEqual,
        Like,
    }
    public interface IWmiQuery
    {
        string WmiClassName { get; set; }
        string QueryString { get; }
        List<WmiQueryFilter> Filters { get; set; }
    }

    public class WmiQueryFilter
    {
        public string Name { get; set; }
        public string Value { get; set; }

        public WhereOperation Operation { get; set; }

        public WmiQueryFilter(string name, string value)
            : this(name, value, WhereOperation.Equals)
        {
        }

        public WmiQueryFilter(string name, string value, WhereOperation op)
        {
            Name = name; Value = value; Operation = op;
        }

        public string GetFilter()
        {
            string Operator = "=";
            switch (Operation)
            {
                case WhereOperation.LessThan:
                    Operator = "<";
                    break;
                case WhereOperation.GreaterThan:
                    Operator = ">";
                    break;
                case WhereOperation.LessThanEquals:
                    Operator = "<=";
                    break;
                case WhereOperation.GreaterThenEquals:
                    Operator = ">=";
                    break;
                case WhereOperation.NotEqual:
                    Operator = "!=";
                    break;
                case WhereOperation.Like:
                    Operator = "LIKE";
                    return String.Format(CultureInfo.InvariantCulture, " {0} {2} '%{1}%' ", Name, Value, Operator);
                case WhereOperation.Equals:
                    Operator = "=";
                    break;
                default:
                    break;
            }
            return String.Format(CultureInfo.InvariantCulture, " {0}{2}'{1}' ", Name, Value, Operator);
        }
    }
}
