// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

namespace Microsoft.WmiCodeGen.DotNet.Interface
{
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Text;

    public interface IWmiQualifier
    {
        
        #region Properties

        /// <summary>
        /// Gets the name.
        /// </summary>
        /// <value>
        /// The name.
        /// </value>
        string Name
        {
            get;
        }

        /// <summary>
        /// Gets the value.
        /// </summary>
        /// <value>
        /// The value.
        /// </value>
        object Value
        {
            get;
        }

        #endregion Properties

    }
}
