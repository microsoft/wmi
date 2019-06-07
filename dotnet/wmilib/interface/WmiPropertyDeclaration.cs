// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Microsoft.WmiCodeGen.DotNet.Interface
{
    /// <summary>
    /// 
    /// </summary>
    public interface IWmiPropertyDeclaration
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
        }

        #endregion Properties
    }

    /// <summary>
    /// 
    /// </summary>
    public class WmiPropertyDeclarationCollection : KeyedCollection<string, IWmiPropertyDeclaration>
    {
        #region Methods

        /// <summary>
        /// Gets the key for item.
        /// </summary>
        /// <param name="property">The property.</param>
        /// <returns></returns>
        protected override string GetKeyForItem(IWmiPropertyDeclaration property)
        {
            return property.Name;
        }

        #endregion Methods

    }
}
