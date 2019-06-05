// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Microsoft.WmiCodeGen.CSharp.Interface
{
    /// <summary>
    /// 
    /// </summary>
    public interface IWmiMethodDeclaration
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
        /// Gets method parameters.
        /// </summary>
        /// <value>
        /// Method parameters.
        /// </value>
        WmiMethodParameterCollection Parameters
        {
            get;
        }

        /// <summary>
        /// Gets method qualifiers.
        /// </summary>
        /// <value>
        /// Method qualifiers.
        /// </value>
        WmiQualifierCollection Qualifiers
        {
            get;
        }

        #endregion Properties
    }

    /// <summary>
    /// 
    /// </summary>
    public class WmiMethodDeclarationCollection : KeyedCollection<string, IWmiMethodDeclaration>
    {
        #region Methods

        /// <summary>
        /// Gets the key for item.
        /// </summary>
        /// <param name="method">The method.</param>
        /// <returns></returns>
        protected override string GetKeyForItem(IWmiMethodDeclaration method)
        {
            return method.Name;
        }

        #endregion Methods

    }
}
