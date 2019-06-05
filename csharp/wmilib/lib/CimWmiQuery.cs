// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.
using Microsoft.WmiCodeGen.CSharp.Interface;
using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Management;
using System.Text;
using System.Threading.Tasks;

namespace Microsoft.WmiCodeGen.CSharp.Lib
{
    /// <summary>
    /// 
    /// </summary>
    public class WmiQuery : IWmiQuery
    {
        #region Constructor

        public WmiQuery(string className)
            : this(className, string.Empty)
        {

        }
        public WmiQuery(string className, WmiQueryFilter filters)
            : this(className, new List<WmiQueryFilter>() { filters })
        {

        }
        public WmiQuery(string className, List<WmiQueryFilter> filters)
        {
            WmiClassName = className;
            Filters = filters;
        }

        public WmiQuery(string className, WhereOperation op, params string[] filterElements)
            : this(className, GetQueryFilters(op, filterElements))
        {

        }
        public WmiQuery(string className, params string[] filterElements)
            : this(className, GetQueryFilters(WhereOperation.Equals, filterElements))
        {
        }

        private static List<WmiQueryFilter> GetQueryFilters(WhereOperation op,
            params string[] filterElements)
        {
            List<WmiQueryFilter> filters = new List<WmiQueryFilter>();
            for (int i = 1; i < filterElements.Length; i += 2)
            {
                if (filterElements[i - 1] != null)
                    filters.Add(new WmiQueryFilter(filterElements[i - 1], filterElements[i], op));
            }
            return filters;
        }
        #endregion

        #region IWmiQuery Members

        public List<WmiQueryFilter> Filters { get; set; }

        public string QueryString
        {
            get
            {
                StringBuilder sb = new StringBuilder();
                sb.Append(String.Format(CultureInfo.InvariantCulture,
                    "SELECT * FROM {0} ",
                    WmiClassName));


                for (int i = 0; i < Filters.Count; i++)
                {
                    if (i == 0) sb.Append(" WHERE ");
                    sb.AppendFormat(CultureInfo.InvariantCulture, " {0} ", Filters[i].GetFilter());
                    if (i < Filters.Count - 1) sb.Append(" AND ");
                }
                return sb.ToString();
            }
        }

        public string WmiClassName { get; set; }

        #endregion

        public override string ToString()
        {
            return QueryString;
        }
    }
}
