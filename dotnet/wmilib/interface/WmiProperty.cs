// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

namespace Microsoft.WmiCodeGen.DotNet.Interface
{
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Text;

    public interface IWmiProperty
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
        /// Gets or sets the value.
        /// </summary>
        /// <value>
        /// The value.
        /// </value>
        object Value
        {
            get;
            set;
        }

        /// <summary>
        /// Gets the type.
        /// </summary>
        /// <value>
        /// The type.
        /// </value>
        Type Type
        {
            get;
        }

        WmiPropertyFlags Flags
        {
            get;
        }

        #endregion Properties
    }

    [Flags]
    public enum WmiPropertyFlags
    {
        None = 0,
        Class = 1,
        Method = 2,
        Property = 4,
        Parameter = 8,
        Association = 16,
        Indication = 32,
        Reference = 64,
        Any = 127,
        EnableOverride = 128,
        DisableOverride = 256,
        Restricted = 512,
        ToSubclass = 1024,
        Translatable = 2048,
        Key = 4096,
        In = 8192,
        Out = 16384,
        Required = 32768,
        Static = 65536,
        Abstract = 131072,
        Terminal = 262144,
        Expensive = 524288,
        Stream = 1048576,
        ReadOnly = 2097152,
        NotModified = 33554432,
        NullValue = 536870912,
        Borrow = 1073741824,
        //Adopt = 2147483648,
    }
}
