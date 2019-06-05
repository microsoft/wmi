// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Management.Infrastructure;
using System.Collections.ObjectModel;
using Microsoft.Management.Infrastructure.Options;
using Microsoft.WmiCodeGen.CSharp.Interface;

namespace Microsoft.WmiCodeGen.CSharp.Lib
{
    public class WmiClass : IWmiClass
    {
        #region Data

        private CimClass m_Class;
        private CimSession m_CimSession;

        #endregion Data

        #region Properties

        /// <summary>
        /// Gets the name of the super class.
        /// </summary>
        /// <value>
        /// The name of the super class.
        /// </value>
        public string SuperClassName
        {
            get
            {
                return GetSuperClassName();
            }
        }

        /// <summary>
        /// Gets the name of the class.
        /// </summary>
        /// <value>
        /// The name of the class.
        /// </value>
        public string ClassName
        {
            get
            {
                return GetClassName();
            }
        }

        /// <summary>
        /// Gets the name of the server.
        /// </summary>
        /// <value>
        /// The name of the server.
        /// </value>
        public string ServerName
        {
            get
            {
                return GetServerName();
            }
        }

        /// <summary>
        /// Gets the namespace.
        /// </summary>
        /// <value>
        /// The namespace.
        /// </value>
        public string Namespace
        {
            get
            {
                return GetNameSpace();
            }
        }

        /// <summary>
        /// Gets the super class.
        /// </summary>
        /// <value>
        /// The super class.
        /// </value>
        public IWmiClass SuperClass
        {
            get
            {
                return GetSuperClass();
            }
        }

        /// <summary>
        /// Gets the properties.
        /// </summary>
        /// <value>
        /// The properties.
        /// </value>
        public WmiPropertyDeclarationCollection Properties
        {
            get
            {
                return GetAllProperties();
            }
        }

        /// <summary>
        /// Gets the qualifiers.
        /// </summary>
        /// <value>
        /// The qualifiers.
        /// </value>
        public WmiQualifierCollection Qualifiers
        {
            get
            {
                return GetAllQualifiers();
            }
        }

        /// <summary>
        /// Gets the methods.
        /// </summary>
        /// <value>
        /// The methods.
        /// </value>
        public WmiMethodDeclarationCollection Methods
        {
            get
            {
                return GetAllMethods();
            }
        }

        #endregion Properties

        #region Construction

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiClass" /> class.
        /// </summary>
        /// <param name="cimClass">The cim class.</param>
        public WmiClass(CimClass cimClass)
        {
            m_Class = cimClass;
            m_CimSession = CimSessionManager.GetSession(GetServerName());

        }

        #endregion Construction

        #region Methods

        /// <summary>
        /// Gets the name of the super class.
        /// </summary>
        /// <returns></returns>
        public string GetSuperClassName()
        {
            try
            {
                return m_Class.CimSuperClassName;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets the class.
        /// </summary>
        /// <returns></returns>
        public object GetClass()
        {
            return m_Class;
        }

        /// <summary>
        /// Gets the name of the class.
        /// </summary>
        /// <returns></returns>
        public string GetClassName()
        {
            try
            {
                return m_Class.CimSystemProperties.ClassName;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets the name of the server.
        /// </summary>
        /// <returns></returns>
        public string GetServerName()
        {
            try
            {
                return m_Class.CimSystemProperties.ServerName;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets the name space.
        /// </summary>
        /// <returns></returns>
        public string GetNameSpace()
        {
            try
            {
                return m_Class.CimSystemProperties.Namespace;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets all properites.
        /// </summary>
        /// <returns></returns>
        public WmiPropertyDeclarationCollection GetAllProperties()
        {
            try
            {
                WmiPropertyDeclarationCollection vmPropertyDeclarationColl = new WmiPropertyDeclarationCollection();
                foreach (CimPropertyDeclaration propertyDeclaration in m_Class.CimClassProperties)
                {

                    vmPropertyDeclarationColl.Add(new WmiCimPropertyDeclaration(propertyDeclaration));
                }
                return vmPropertyDeclarationColl;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets all qualifiers.
        /// </summary>
        /// <returns></returns>
        public WmiQualifierCollection GetAllQualifiers()
        {
            try
            {
                WmiQualifierCollection qualifierCollection = new WmiQualifierCollection();
                foreach (CimQualifier qualifier in m_Class.CimClassQualifiers)
                {
                    qualifierCollection.Add(new WmiQualifier(qualifier));
                }
                return qualifierCollection;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets all methods.
        /// </summary>
        /// <returns></returns>
        public WmiMethodDeclarationCollection GetAllMethods()
        {
            try
            {
                WmiMethodDeclarationCollection methodCollection = new WmiMethodDeclarationCollection();
                foreach (CimMethodDeclaration method in m_Class.CimClassMethods)
                {
                    methodCollection.Add(new WmiCimMethodDeclaration(method));
                }
                return methodCollection;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets the super class.
        /// </summary>
        /// <returns></returns>
        public IWmiClass GetSuperClass()
        {
            try
            {
                return new WmiClass(m_Class.CimSuperClass);
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets the method parameters.
        /// </summary>
        /// <param name="classInstance">The class instance.</param>
        /// <param name="methodName">Name of the method.</param>
        /// <returns></returns>
        public WmiMethodParameterCollection GetMethodParameters(string methodName)
        {
            return this.Methods[methodName].Parameters;
        }

        public WmiMethodResult InvokeMethod(
            String methodName,
            WmiMethodParameterCollection inputParams,
            WmiOperationOptions inputInvokeOptions)
        {
            if (methodName == null)
            {
                throw new ArgumentNullException("methodName", "methodName cannot be null");
            }

            try
            {
                using (CimOperationOptions operationOptions = new CimOperationOptions())
                {
                    if (inputInvokeOptions != null)
                    {
                        if (inputInvokeOptions != null)
                        {
                            operationOptions.Timeout = inputInvokeOptions.TimeOut;
                        }
                    }

                    //Add Method Parameters
                    using (CimMethodParametersCollection methodParameters = new CimMethodParametersCollection())
                    {
                        if (inputParams != null)
                        {
                            WmiHelper.AddMethodParameters(m_Class, methodName, inputParams, methodParameters);
                        }

                        //Invoke Method
                        using (CimMethodResult methodResult = m_CimSession.InvokeMethod(
                                                Namespace, ClassName, methodName, methodParameters))
                        {
                            //Refresh Output
                            WmiHelper.RefreshCimMethodOutput(methodResult);

                            //Convert Result back to Dot net type
                            return WmiHelper.ConvertCimMethodResulttoWmiMethodResult(methodResult);
                        }
                    }
                }
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        #endregion Methods

        #region IDisposable Members
        private bool m_Disposed;
        protected virtual void Dispose(bool disposing)
        {
            if (disposing && !m_Disposed)
            {
                using (m_Class) { }
                m_Disposed = true;
            }
        }
        public void Dispose()
        {
            Dispose(true);
            GC.SuppressFinalize(this);
        }

        #endregion

    }

    public class WmiCimPropertyDeclaration : IWmiPropertyDeclaration
    {
        #region Data

        private CimPropertyDeclaration m_PropertyDeclaration;

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
                return m_PropertyDeclaration.Name;
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
                return m_PropertyDeclaration.Value;
            }
        }

        #endregion Properties

        #region Construction

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiCimPropertyDeclaration" /> class.
        /// </summary>
        /// <param name="propertyDeclaration">The property declaration.</param>
        public WmiCimPropertyDeclaration(CimPropertyDeclaration propertyDeclaration)
        {
            m_PropertyDeclaration = propertyDeclaration;
        }

        #endregion Construction


    }

    public class WmiCimMethodDeclaration : IWmiMethodDeclaration
    {
        #region Data

        private CimMethodDeclaration m_MethodDeclaration;

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
                return m_MethodDeclaration.Name;
            }
        }

        /// <summary>
        /// Gets method parameters.
        /// </summary>
        /// <value>
        /// Method parameters.
        /// </value>
        public WmiMethodParameterCollection Parameters
        {
            get
            {
                WmiMethodParameterCollection methodParams = new WmiMethodParameterCollection();
                foreach (CimMethodParameterDeclaration cimMethodParamDeclaration in m_MethodDeclaration.Parameters)
                {
                    WmiMethodParameter methodParam = new WmiMethodParameter();
                    methodParam.Name = cimMethodParamDeclaration.Name;
                    methodParam.Value = null;

                    switch (cimMethodParamDeclaration.CimType)
                    {
                        case CimType.Instance:
                        case CimType.Reference:
                            methodParam.Type = typeof(IWmiInstance);
                            break;
                        case CimType.InstanceArray:
                        case CimType.ReferenceArray:
                            methodParam.Type = typeof(IWmiInstance[]);
                            break;
                        case CimType.DateTime:
                            methodParam.Type = typeof(DateTime);
                            break;
                        case CimType.DateTimeArray:
                            methodParam.Type = typeof(DateTime[]);
                            break;
                        case CimType.Boolean:
                            methodParam.Type = typeof(bool);
                            break;
                        case CimType.BooleanArray:
                            methodParam.Type = typeof(bool[]);
                            break;
                        case CimType.String:
                            methodParam.Type = typeof(String);
                            break;
                        case CimType.StringArray:
                            methodParam.Type = typeof(String[]);
                            break;
                        case CimType.UInt8:
                            methodParam.Type = typeof(Byte);
                            break;
                        case CimType.UInt8Array:
                            methodParam.Type = typeof(Byte[]);
                            break;
                        case CimType.SInt8:
                            methodParam.Type = typeof(SByte);
                            break;
                        case CimType.SInt8Array:
                            methodParam.Type = typeof(SByte[]);
                            break;
                        case CimType.UInt16:
                            methodParam.Type = typeof(UInt16);
                            break;
                        case CimType.UInt16Array:
                            methodParam.Type = typeof(UInt16[]);
                            break;
                        case CimType.SInt16:
                            methodParam.Type = typeof(Int16);
                            break;
                        case CimType.SInt16Array:
                            methodParam.Type = typeof(Int16[]);
                            break;
                        case CimType.UInt32:
                            methodParam.Type = typeof(UInt32);
                            break;
                        case CimType.UInt32Array:
                            methodParam.Type = typeof(UInt32[]);
                            break;
                        case CimType.SInt32:
                            methodParam.Type = typeof(Int32);
                            break;
                        case CimType.SInt32Array:
                            methodParam.Type = typeof(Int32[]);
                            break;
                        case CimType.UInt64:
                            methodParam.Type = typeof(UInt64);
                            break;
                        case CimType.UInt64Array:
                            methodParam.Type = typeof(UInt64[]);
                            break;
                        case CimType.SInt64:
                            methodParam.Type = typeof(Int64);
                            break;
                        case CimType.SInt64Array:
                            methodParam.Type = typeof(Int64[]);
                            break;
                        default:
                            methodParam.Type = null;
                            break;
                    }
                    methodParams.Add(methodParam);
                }
                return methodParams;
            }
        }

        /// <summary>
        /// Gets method qualifiers.
        /// </summary>
        /// <value>
        /// Method qualifiers.
        /// </value>
        public WmiQualifierCollection Qualifiers
        {
            get
            {
                WmiQualifierCollection qualifierCollection = new WmiQualifierCollection();
                foreach (CimQualifier qualifier in m_MethodDeclaration.Qualifiers)
                {
                    qualifierCollection.Add(new WmiQualifier(qualifier));
                }
                return qualifierCollection;
            }
        }

        #endregion Properties

        #region Construction

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiCimMethodDeclaration" /> class.
        /// </summary>
        /// <param name="methodDeclaration">The method declaration.</param>
        public WmiCimMethodDeclaration(CimMethodDeclaration methodDeclaration)
        {
            m_MethodDeclaration = methodDeclaration;
        }

        #endregion Construction
    }
}
