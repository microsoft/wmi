// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.ObjectModel;

namespace Microsoft.wmi.DotNet.Interface
{
    /// <summary>
    /// 
    /// </summary>
    public interface IWmiClass : IDisposable
    {
        #region Properties

        /// <summary>
        /// Gets the name of the super class.
        /// </summary>
        /// <value>
        /// The name of the super class.
        /// </value>
        string SuperClassName
        {
            get;
        }

        /// <summary>
        /// Gets the name of the class.
        /// </summary>
        /// <value>
        /// The name of the class.
        /// </value>
        string ClassName
        {
            get;
        }

        /// <summary>
        /// Gets the name of the server.
        /// </summary>
        /// <value>
        /// The name of the server.
        /// </value>
        string ServerName
        {
            get;
        }

        /// <summary>
        /// Gets the namespace.
        /// </summary>
        /// <value>
        /// The namespace.
        /// </value>
        string Namespace
        {
            get;
        }

        /// <summary>
        /// Gets the super class.
        /// </summary>
        /// <value>
        /// The super class.
        /// </value>
        IWmiClass SuperClass
        {
            get;
        }

        /// <summary>
        /// Gets the properties.
        /// </summary>
        /// <value>
        /// The properties.
        /// </value>
        WmiPropertyDeclarationCollection Properties
        {
            get;

        }

        /// <summary>
        /// Gets the qualifiers.
        /// </summary>
        /// <value>
        /// The qualifiers.
        /// </value>
        WmiQualifierCollection Qualifiers
        {
            get;
        }

        /// <summary>
        /// Gets the methods.
        /// </summary>
        /// <value>
        /// The methods.
        /// </value>
        WmiMethodDeclarationCollection Methods
        {
            get;
        }

        #endregion Properties

        #region Methods

        /// <summary>
        /// Gets all properites.
        /// </summary>
        /// <returns></returns>
        WmiPropertyDeclarationCollection GetAllProperties();

        /// <summary>
        /// Gets all qualifiers.
        /// </summary>
        /// <returns></returns>
        WmiQualifierCollection GetAllQualifiers();

        /// <summary>
        /// Gets all methods.
        /// </summary>
        /// <returns></returns>
        WmiMethodDeclarationCollection GetAllMethods();

        /// <summary>
        /// Gets the super class.
        /// </summary>
        /// <returns></returns>
        IWmiClass GetSuperClass();

        /// <summary>
        /// Gets the class.
        /// </summary>
        /// <returns></returns>
        object GetClass();

        /// <summary>
        /// Gets the method parameters.
        /// </summary>
        /// <param name="methodName">Name of the method.</param>
        /// <returns></returns>
        WmiMethodParameterCollection GetMethodParameters(string methodName);

        WmiMethodResult InvokeMethod(
            String methodName,
            WmiMethodParameterCollection inputParams,
            WmiOperationOptions inputOptions);


        #endregion Methods
    }
}