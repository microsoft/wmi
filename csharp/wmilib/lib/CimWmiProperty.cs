// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using Microsoft.Management.Infrastructure;
using Microsoft.WmiCodeGen.CSharp.Interface;

namespace Microsoft.WmiCodeGen.CSharp.Lib
{
    public class WmiProperty : IWmiProperty
    {
        #region Data

        private CimProperty m_Property;

        #endregion Data

        #region Properties

        /// <summary>
        /// Gets the name.
        /// </summary>
        /// <value>
        /// The name.
        /// </value>
        public string Name
        {
            get
            {
                return m_Property.Name;
            }
        }

        /// <summary>
        /// Gets or sets the value.
        /// </summary>
        /// <value>
        /// The value.
        /// </value>
        public object Value
        {
            get
            {
                return WmiHelper.GetDotNetValueFromCimProperty(m_Property);
            }
            set
            {
                m_Property.Value = WmiHelper.GetCimValueFromDotNetValue(Value, m_Property.CimType);
            }
        }

        /// <summary>
        /// Gets the type.
        /// </summary>
        /// <value>
        /// The type.
        /// </value>
        public Type Type
        {
            get
            {
                return WmiHelper.GetDotNetTypeFromCimProperty(m_Property);
            }
        }

        /// <summary>
        /// Gets the flags.
        /// </summary>
        /// <value>
        /// The flags.
        /// </value>
        public WmiPropertyFlags Flags
        {
            get
            {
                return WmiHelper.GetWmiPropertyFlagsFromCimProperty(m_Property);
            }
        }

        #endregion Properties

        #region Construction

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiProperty" /> class.
        /// </summary>
        /// <param name="cimProperty">The cim property.</param>
        public WmiProperty(CimProperty cimProperty)
        {
            m_Property = cimProperty;
        }

        #endregion Construction

    }
}
