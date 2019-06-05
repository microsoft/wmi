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
    class WmiQualifier : IWmiQualifier
    {
        #region Data

        private CimQualifier m_Qualifier;

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
                try
                {
                    return m_Qualifier.Name;
                }
                catch (CimException ex)
                {
                    throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
                }
            }
        }

        /// <summary>
        /// Gets the value.
        /// </summary>
        /// <value>
        /// The value.
        /// </value>
        public object Value
        {
            get
            {
                try
                {
                    return m_Qualifier.Value;
                }
                catch (CimException ex)
                {
                    throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
                }
            }
        }

        #endregion Properties

        #region Construction

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiQualifier" /> class.
        /// </summary>
        /// <param name="cimQualifier">The cim qualifier.</param>
        public WmiQualifier(CimQualifier cimQualifier)
        {
            m_Qualifier = cimQualifier;
        }

        #endregion Construction
    }
}
