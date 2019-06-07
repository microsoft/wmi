// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Management.Infrastructure;
using System.Collections.ObjectModel;
using System.Globalization;
using Microsoft.Management.Infrastructure.Options;
using System.Threading;
using System.Net;
using Microsoft.WmiCodeGen.DotNet.Interface;

namespace Microsoft.WmiCodeGen.DotNet.Lib
{
    /// <summary>
    /// Flags to enable WMI validation types
    /// </summary>
    [Flags]
    public enum WmiManagementVerify
    {
        None = 0x0,
        PropertyType = 0x1,
        Put = 0x2,
        Method = 0x4,
        All = 0xFFFFF
    }

    /// <summary>
    /// Delegate used with ManagementObject.InvokeMethod called with ManagementVerify.Method enabled
    /// </summary>
    /// <param name="managementObject"></param>
    /// <param name="methodName">Method invoked</param>
    /// <param name="inParameters">Parameters passed to method cannot be used with args</param>
    /// <param name="args">Arguments passed to method cannot be used with inParameters</param>
    public delegate void WmiManagementMethodVerify(
        WmiInstance managementObject,
        String methodName,
        WmiMethodParameterCollection inParameters,
        WmiMethodParameterCollection args);


    public class WmiInstance : IWmiInstance
    {
        #region Data

        private CimInstance m_Instance;
        private CimSession m_CimSession;
        private string m_hostName;
        private bool m_isInstanceCreated = true;

        #endregion Data

        #region Properties

        /// <summary>
        /// Indexer to set property value
        /// </summary>
        /// <param name="propertyName">Property Name</param>
        /// <returns></returns>
        public object this[string propertyName]
        {
            get
            {
                return GetProperty(propertyName);
            }
            set
            {
                SetProperty(propertyName, value);
            }
        }

        /// <summary>
        /// Get all properties of m_Instance
        /// </summary>
        public WmiPropertyCollection InstanceProperties
        {
            get
            {
                return GetAllProperites();
            }
        }

        /// <summary>
        /// Get all qualifiers of m_Instance
        /// </summary>
        public WmiQualifierCollection InstanceQualifiers
        {
            get
            {
                return GetAllQualifiers();
            }
        }

        /// <summary>
        /// Gets the class.
        /// </summary>
        /// <value>
        /// The class.
        /// </value>
        public IWmiClass Class
        {
            get
            {
                return new WmiClass(m_Instance.CimClass);
            }
        }

        /// <summary>
        /// Gets the embedded instance.
        /// </summary>
        /// <value>
        /// The embedded instance.
        /// </value>
        public string EmbeddedInstance
        {
            get
            {
                return this.GetEmbeddedInstance();
            }
        }


        /// <summary>
        /// 
        /// </summary>
        public IWmiInstanceManager InstanceManager
        {
            get
            {
                return WmiInstanceFactory.GetInstanceManager(GetServerName(), GetNamespace());
            }
        }

        /// <summary>
        /// 
        /// </summary>
        public bool AutoCommit { get; set; }

        /// <summary>
        /// 
        /// </summary>
        public bool isEmbedded { get; set; }

        #endregion Properties

        #region Construction

        /// <summary>
        /// Constructs CimWmiInstance which abstracts CimInstance 
        /// </summary>
        /// <param name="instance">instance to be abstracted</param>
        public WmiInstance(CimInstance instance)
        {
            if (instance == null) throw new ArgumentNullException("instance");

            if (instance.CimInstanceProperties.Count > 0)
                m_Instance = new CimInstance(instance); // Create a copy here
            else
                // Workaround. Do not create a copy for empty instance. Ends up in AV
                m_Instance = instance;

            m_CimSession = CimSessionManager.GetSession(GetServerName());
        }

        /// <summary>
        /// Constructor to clone the Wmi Instance
        /// </summary>
        /// <param name="WmiInstanceToClone"></param>
        public WmiInstance(IWmiInstance WmiInstanceToClone)
            : this((CimInstance)WmiInstanceToClone.GetInstance())
        {
        }

        /// <summary>
        /// Create the instance using the Instance __PATH
        /// </summary>
        /// <param name="instancePath"></param>
        public WmiInstance(WmiInstancePath instancePath)
            : this(WmiHelper.GetInstanceFromPath(instancePath.ToString()))
        {
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="className"></param>
        public WmiInstance(String HostName, String WMINameSpace, String className)
            : this(HostName, WMINameSpace, className, null)
        {
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="query"></param>
        public WmiInstance(String HostName, String WMINameSpace, IWmiQuery query)
            : this(HostName, WMINameSpace, query, null)
        {

        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="query"></param>
        /// <param name="userName"></param>
        /// <param name="password"></param>
        public WmiInstance(String HostName, String WMINameSpace, IWmiQuery query, string userName, string password)
            : this(HostName, WMINameSpace, query, new WmiCredentials(userName, password))
        {
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="query"></param>
        /// <param name="userName"></param>
        /// <param name="password"></param>
        public WmiInstance(String HostName, String WMINameSpace, IWmiQuery query, WmiCredentials credentials)
            : this(GetInstance(HostName, WMINameSpace, query, credentials))
        {
        }

        /// <summary>
        /// Creates an Empty instance of WmiInstance specified by className
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="className"></param>
        /// <param name="userName"></param>
        /// <param name="password"></param>
        public WmiInstance(String HostName, String WMINameSpace, String className, string userName, string password)
            : this(HostName, WMINameSpace, className, new WmiCredentials(userName, password))
        {

        }

        /// <summary>
        /// Creates an Empty instance of WmiInstance specicied by class name
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="className"></param>
        /// <param name="credentials"></param>
        public WmiInstance(String HostName, String WMINameSpace, String className, WmiCredentials credentials)
            : this(GetInstance(HostName, WMINameSpace, className, credentials))
        {
            if (string.IsNullOrEmpty(HostName) ||
                HostName.ToLowerInvariant().Contains(Environment.MachineName.ToLowerInvariant()) ||
                HostName.Equals("localhost", StringComparison.OrdinalIgnoreCase) ||
                HostName.Equals(".", StringComparison.OrdinalIgnoreCase))
            {
                m_hostName = Environment.MachineName;
            }
            else
            {
                m_hostName = HostName;
            }

            m_CimSession = CimSessionManager.GetSession(GetServerName());
            m_isInstanceCreated = false;
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="query"></param>
        /// <param name="userName"></param>
        /// <param name="password"></param>
        /// <returns></returns>
        private static IWmiInstance GetInstance(String HostName, String WMINameSpace, IWmiQuery query, string userName, string password)
        {
            return GetInstance(HostName, WMINameSpace, query, new WmiCredentials(userName, password));
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="query"></param>
        /// <param name="credentials"></param>
        /// <returns></returns>
        private static IWmiInstance GetInstance(String HostName, String WMINameSpace, IWmiQuery query, WmiCredentials credentials)
        {
            using (WmiInstanceCollection collection = GetInstances(HostName, WMINameSpace, query, credentials))
            {
                return collection.First().Clone() as IWmiInstance;
            }
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="query"></param>
        /// <param name="credentials"></param>
        /// <returns></returns>
        private static WmiInstanceCollection GetInstances(String HostName, String WMINameSpace,
            IWmiQuery query,
            WmiCredentials credentials)
        {
            if (HostName == null)
            {
                throw new ArgumentNullException("HostName");
            }
            if (WMINameSpace == null)
            {
                throw new ArgumentNullException("WMINameSpace");
            }

            // Gets the machine name of the local host.
            if (string.Equals(HostName, "localhost", StringComparison.CurrentCultureIgnoreCase))
            {
                //HostName = HostEnvironment.MachineName;
            }

            // Check if the name is FQDN, if so, get rid of the domain part.
            // If it is an IP Address, do not do this.
            IPAddress tempIPAddress;
            if (!IPAddress.TryParse(HostName, out tempIPAddress))
            {
                int dotPosition = HostName.IndexOf('.');
                if (dotPosition != -1)
                {
                    HostName = HostName.Substring(0, dotPosition);
                }
            }

            IWmiInstanceManager wmiInstanceManager = WmiInstanceFactory.GetInstanceManager(HostName, WMINameSpace, credentials);
            if (wmiInstanceManager == null) throw new Exception("WMI InstanceManager not initialized");

            WmiInstanceCollection collection = wmiInstanceManager.QueryInstances(query);
            if (collection.Count < 1)
            {
                throw new WmiException(string.Format(CultureInfo.InvariantCulture, "Query[{0}] failed with No instances", query));
            }
            return collection;
        }
        /// <summary>
        /// Gets the Existing Instance, filtering using the query provided
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="query"></param>
        /// <param name="userName"></param>
        /// <param name="password"></param>
        /// <returns></returns>
        private static WmiInstanceCollection GetInstances(String HostName, String WMINameSpace,
            IWmiQuery query, string userName, string password)
        {
            return GetInstances(HostName, WMINameSpace, query, new WmiCredentials(userName, password));
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="className"></param>
        /// <param name="credentials"></param>
        /// <returns></returns>
        private static IWmiInstance GetInstance(String HostName, String WMINameSpace, String className, WmiCredentials credentials)
        {
            if (HostName == null)
            {
                throw new ArgumentNullException("HostName");
            }
            if (WMINameSpace == null)
            {
                throw new ArgumentNullException("WMINameSpace");
            }

            // Gets the machine name of the local host.
            if (string.Equals(HostName, "localhost", StringComparison.CurrentCultureIgnoreCase))
            {
                HostName = Environment.MachineName;
            }

            // Check if the name is FQDN, if so, get rid of the domain part.
            // Check if the name is FQDN, if so, get rid of the domain part.
            // If it is an IP Address, do not do this.
            IPAddress tempIPAddress;
            if (!IPAddress.TryParse(HostName, out tempIPAddress))
            {
                int dotPosition = HostName.IndexOf('.');
                if (dotPosition != -1)
                {
                    HostName = HostName.Substring(0, dotPosition);
                }
            }

            IWmiInstanceManager wmiInstanceManager = WmiInstanceFactory.GetInstanceManager(HostName, WMINameSpace, credentials);
            if (wmiInstanceManager == null) throw new Exception("WMI InstanceManager not initialized");

            return wmiInstanceManager.CreateInstance(className);
        }

        /// <summary>
        /// Creates a new Instance
        /// </summary>
        /// <param name="HostName"></param>
        /// <param name="WMINameSpace"></param>
        /// <param name="className"></param>
        /// <param name="userName"></param>
        /// <param name="password"></param>
        /// <returns></returns>
        private static IWmiInstance GetInstance(String HostName, String WMINameSpace,
            String className, string userName, string password)
        {

            return GetInstance(HostName, WMINameSpace, className, new WmiCredentials(userName, password));
        }
        #endregion Construction

        #region Methods

        /// <summary>
        /// Returns the abstracted CimInstance
        /// </summary>
        /// <returns></returns>
        public object GetInstance()
        {
            return m_Instance;
        }

        /// <summary>
        /// Gets the property value for the specified property name
        /// </summary>
        /// <param name="propertyName">Property Name</param>
        /// <returns></returns>
        public object GetProperty(String propertyName)
        {
            try
            {
                return WmiHelper.GetDotNetValueFromCimProperty(m_Instance.CimInstanceProperties[propertyName]);
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="propertyName"></param>
        public void ResetProperty(string propertyName)
        {
            SetProperty<object>(propertyName, null);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="propertyName"></param>
        /// <returns></returns>
        public object GetProperty<T>(string propertyName)
        {
            object value = GetProperty(propertyName);
            if (value != null)
            {
                if (value is Array)
                {
                    List<object> items = new List<object>();
                    foreach (var item in (value as Array))
                        items.Add(item);
                    return GetObjectArray<T>(items.ToArray());
                }
                else
                {
                    if (typeof(T).IsEnum)
                    {
                        return (T)Enum.Parse(typeof(T), value.ToString());
                    }
                    if (typeof(T).BaseType == typeof(WmiInstance))
                    {
                        return (T)Activator.CreateInstance(typeof(T), value);
                    }
                    return Convert.ChangeType(value, typeof(T));
                }
            }

            return null;
        }

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="value"></param>
        /// <returns></returns>
        private T[] GetObjectArray<T>(object[] value)
        {
            T[] newValue = new T[value.Length];

            for (int i = 0; i < value.Length; i++)
            {
                if (typeof(T).IsEnum)
                {
                    newValue[i] = (T)Enum.Parse(typeof(T), value[i].ToString());
                }
                else if (typeof(T) is IWmiInstance)
                {
                    newValue[i] = (T)Activator.CreateInstance(typeof(T), value);
                }
                else
                {
                    newValue[i] = (T)value[i];
                }
            }
            return newValue;
        }

        /// <summary>
        /// Sets the property value
        /// </summary>
        /// <param name="propertyName">Property Name</param>
        /// <param name="propertyValue">Property Value</param>
        public void SetProperty<T>(String propertyName, T propertyValue)
        {
            try
            {
                if (m_Instance.CimInstanceProperties[propertyName] != null)
                {
                    m_Instance.CimInstanceProperties[propertyName].Value =
                        WmiHelper.GetCimValueFromDotNetValue(
                            propertyValue,
                            m_Instance.CimInstanceProperties[propertyName].CimType
                        );

                }
                else
                {
                    CimType cType = typeof(T).IsEnum ? CimType.UInt16 : CimConverter.GetCimType(typeof(T));
                    // Handle Array of Enum - TODO
                    CimProperty cProperty = CimProperty.Create(
                            propertyName,
                            WmiHelper.GetCimValueFromDotNetValue(propertyValue, cType),
                            cType,
                            CimFlags.Property);
                    m_Instance.CimInstanceProperties.Add(cProperty);
                }

                ModifyInstance();

                // Should we verify here based on the flags set?
                if (ShouldVerify(WmiManagementVerify.Put))
                {
                    // TBD
                }
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets class name of the abstracted CimInstance
        /// </summary>
        /// <returns></returns>
        public String GetClassName()
        {
            try
            {
                return m_Instance.CimSystemProperties.ClassName;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets namespace of the abstracted CimInstance
        /// </summary>
        /// <returns></returns>
        public String GetNamespace()
        {
            try
            {
                string wmiNamespace = m_Instance.CimSystemProperties.Namespace;

                wmiNamespace = wmiNamespace.Replace(@"/", @"\");

                return wmiNamespace;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets server name of the abstracted Server Name
        /// </summary>
        /// <returns></returns>
        public String GetServerName()
        {
            try
            {
                if (string.IsNullOrEmpty(m_Instance.CimSystemProperties.ServerName))
                {
                    return m_hostName;
                }
                return m_Instance.CimSystemProperties.ServerName;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets all properties of the abstracted instance
        /// </summary>
        /// <returns></returns>
        public WmiPropertyCollection GetAllProperites()
        {
            try
            {
                WmiPropertyCollection properties = new WmiPropertyCollection();
                foreach (CimProperty property in m_Instance.CimInstanceProperties)
                {
                    properties.Add(new WmiProperty(property));
                }
                return properties;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }


        /// <summary>
        /// Gets all qualifiers of the abstracted instance
        /// </summary>
        /// <returns></returns>
        public WmiQualifierCollection GetAllQualifiers()
        {
            try
            {
                WmiQualifierCollection qualifierCollection = new WmiQualifierCollection();
                foreach (CimQualifier qualifier in m_Instance.CimClass.CimClassQualifiers)
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
        /// Determines whether the specified <see cref="System.Object" /> is equal to this instance.
        /// </summary>
        /// <param name="destination">The <see cref="System.Object" /> to compare with this instance.</param>
        /// <returns>
        ///   <c>true</c> if the specified <see cref="System.Object" /> is equal to this instance; otherwise, <c>false</c>.
        /// </returns>
        public override bool Equals(object destination)
        {
            IWmiInstance destinationObject = destination as IWmiInstance;

            CimInstance sourceInstance = this.GetInstance() as CimInstance;
            CimInstance destinationInstance = destinationObject.GetInstance() as CimInstance;

            return CimEquals(sourceInstance, destinationInstance);
        }

        /// <summary>
        /// Values the equals.
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="propertyName">Name of the property.</param>
        /// <param name="cimType">Type of the cim.</param>
        /// <param name="other">The other.</param>
        /// <returns></returns>
        public bool ValueEquals<T>(String propertyName, Type cimType, T other)
        {
            //TODO:KR:CSS: Equals is not that simple. PropertyData performs a complicated operation around it. Need to import that once we verify this works OR not.
            IWmiProperty property = this.InstanceProperties[propertyName];
            T value = (T)property.Value;
            return (value.Equals(other));
        }

        /// <summary>
        /// Determines whether the specified CimInstance source and destination are equal. Unlike ManagementObject, CimInstance does
        /// not override Object.Equals for object equality. We use our own wrapper.
        /// </summary>
        /// <param name="source">Source instance</param>
        /// <param name="destination">Destination instance</param>
        /// <returns>
        ///     <c>true</c> if source is equal to destination; otherwise, <c>false</c>.
        /// </returns>
        private bool CimEquals(CimInstance sourceInstance, CimInstance destinationInstance)
        {
            try
            {
                if (!sourceInstance.CimSystemProperties.ClassName.Equals(destinationInstance.CimSystemProperties.ClassName))
                {
                    return false;
                }

                foreach (CimProperty sourceProperty in sourceInstance.CimInstanceProperties)
                {
                    CimProperty currentSourceProperty = sourceProperty;

                    // Only check properties used as keys
                    if ((currentSourceProperty.Flags & CimFlags.Key) == 0)
                    {
                        continue;
                    }

                    if (currentSourceProperty.Value == null)
                    {
                        return false;
                    }

                    bool bMatchingDestinationPropertyFound = false;
                    foreach (CimProperty destinationProperty in destinationInstance.CimInstanceProperties)
                    {
                        if (destinationProperty.Name.Equals(currentSourceProperty.Name))
                        {
                            if (destinationProperty.Value == null)
                            {
                                return false;
                            }

                            // Apparently some CimInstances in the properties are used as keys too. We compare them with our wrapper.
                            // When comparing string properties we use a case-insensitive comparison.
                            if ((destinationProperty.Value is CimInstance) &&
                                (currentSourceProperty.Value is CimInstance) &&
                                (CimEquals((CimInstance)destinationProperty.Value, (CimInstance)currentSourceProperty.Value)) ||
                                (destinationProperty.Value is string) &&
                                (currentSourceProperty.Value is string) &&
                                (((string)destinationProperty.Value).Equals((string)currentSourceProperty.Value, StringComparison.OrdinalIgnoreCase)) ||
                                (destinationProperty.Value.Equals(currentSourceProperty.Value)))
                            {
                                bMatchingDestinationPropertyFound = true;
                                break;
                            }
                        }

                    }

                    if (!bMatchingDestinationPropertyFound)
                    {
                        return false;
                    }
                }

                return true;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Returns a hash code for this instance.
        /// </summary>
        /// <returns>
        /// A hash code for this instance, suitable for use in hashing algorithms and data structures like a hash table. 
        /// </returns>
        public override int GetHashCode()
        {
            return base.GetHashCode();
        }

        /// <summary>
        /// Gets the CIMXML from instance.
        /// </summary>
        /// <param name="instance">The instance.</param>
        /// <returns></returns>
        public string GetEmbeddedInstance()
        {
            try
            {
                return WmiHelper.GetEmbeddedInstance(this);
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets the object path from instance.
        /// </summary>
        /// <param name="instance">The instance.</param>
        /// <returns></returns>
        public string GetInstancePath()
        {
            try
            {
                return WmiHelper.GetPathFromInstance(this);
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Undo all the local changes and get the latest instance from Server
        /// </summary>
        public void RefreshInstance()
        {
            using (CimInstance oldInstance = m_Instance)
                m_Instance = m_CimSession.GetInstance(GetNamespace(), oldInstance);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <returns></returns>
        public override string ToString()
        {
            return string.Format(CultureInfo.InvariantCulture, "Path[{0}]\nInstance[{1}]\nHash[{2}]",
                GetInstancePath(), GetEmbeddedInstance(), GetHashCode());
        }

        /// <summary>
        /// Creates & Commits a new Instances to the Server
        /// </summary>
        public void CommitInstance()
        {
            using (CimInstance oldInstance = m_Instance)
            {
                using (CimInstance newInstance = m_CimSession.CreateInstance(GetNamespace(), oldInstance))
                {
                    m_Instance = m_CimSession.ModifyInstance(newInstance);
                }
            }
        }

        /// <summary>
        /// Creates & Commits a new Instances to the Server
        /// </summary>
        public void CommitInstance(WmiOperationOptions options)
        {
            using (CimOperationOptions cimOptions = new CimOperationOptions())
            {
                foreach (var item in options.Options)
                {
                    cimOptions.SetCustomOption(item.Name, item.Value,
                        CimConverter.GetCimType(item.OperationType), item.Comply);
                }

                using (CimInstance oldInstance = m_Instance)
                {
                    using (CimInstance newInstance = m_CimSession.CreateInstance(GetNamespace(), oldInstance, cimOptions))
                    {
                        m_Instance = m_CimSession.ModifyInstance(newInstance);
                    }
                }
            }
        }
        /// <summary>
        /// Puts the local changes to the Server and 
        /// updates the local copy
        /// </summary>
        public void ModifyInstance()
        {
            if (m_isInstanceCreated && ShouldModify)
            {
                using (CimInstance oldInstance = m_Instance)
                    m_Instance = m_CimSession.ModifyInstance(oldInstance);
            }
        }

        /// <summary>
        /// 
        /// </summary>
        public void DeleteInstance()
        {
            m_CimSession.DeleteInstance(m_Instance);
        }

        public WmiMethodResult InvokeMethod(String methodName,
            WmiMethodParameterCollection inputParams)
        {
            return InvokeMethod(methodName, inputParams, null as WmiOperationOptions);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="methodName"></param>
        /// <param name="inputParams"></param>
        /// <param name="inputOptions"></param>
        /// <returns></returns>
        public WmiMethodResult InvokeMethod(String methodName,
                                            WmiMethodParameterCollection inputParams,
                                            WmiOperationOptions inputOptions)
        {
            return InvokeMethod(methodName, inputParams, inputOptions, UserAction.Async, 0, 0);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="methodName"></param>
        /// <param name="inputParams"></param>
        /// <param name="inputInvokeOptions"></param>
        /// <param name="action"></param>
        /// <param name="percentComplete"></param>
        /// <param name="timeout"></param>
        /// <returns></returns>
        public WmiMethodResult InvokeMethod(String methodName,
                                            WmiMethodParameterCollection inputParams,
                                            WmiOperationOptions inputInvokeOptions,
                                            UserAction action,
                                            uint percentComplete,
                                            int timeout)
        {
            try
            {
                CimOperationOptions operationOptions = null;
                if (inputInvokeOptions != null)
                {
                    operationOptions = new CimOperationOptions();
                    if (inputInvokeOptions.TimeOut != null)
                    {
                        operationOptions.Timeout = inputInvokeOptions.TimeOut;
                    }
                }


                //Add Method Parameters
                CimMethodParametersCollection methodParameters = null;
                if (inputParams != null)
                {
                    methodParameters = new CimMethodParametersCollection();
                    if (!m_isInstanceCreated)
                    {
                        // For Empty instances, use this workaround
                        using (IWmiInstance tmpInstance = InstanceManager.CreateInstance(GetClassName(), null))
                        {
                            WmiHelper.AddMethodParameters(tmpInstance.Class.GetClass() as CimClass,
                                methodName, inputParams, methodParameters);
                        }
                    }
                    else
                    {
                        WmiHelper.AddMethodParameters(m_Instance.CimClass, methodName, inputParams, methodParameters);
                    }
                }

                using (operationOptions)
                {
                    using (methodParameters)
                    {
                        //Invoke Method
                        using (CimMethodResult methodResult = m_CimSession.InvokeMethod(m_Instance, methodName, methodParameters))
                        {
                            //Refresh Output
                            WmiHelper.RefreshCimMethodOutput(methodResult);

                            //Convert Result back to Dot net type
                            WmiMethodResult convertedResult = WmiHelper.ConvertCimMethodResulttoWmiMethodResult(methodResult);
                            if (convertedResult.OutParameters.Contains("Job"))
                            {
                                WmiInstance jobInstance = convertedResult.OutParameters["Job"].Value as WmiInstance;
                                if (jobInstance != null)
                                    WmiJob.WaitForAction(jobInstance, action, percentComplete, timeout);
                            }
                            return convertedResult;
                        }
                    }
                }
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="methodName"></param>
        /// <param name="inputParams"></param>
        /// <param name="inputInvokeOptions"></param>
        /// <param name="action"></param>
        /// <param name="taskPercentage"></param>
        /// <param name="job"></param>
        /// <returns></returns>
        public UInt32 InvokeMethod(string methodName,
            WmiMethodParameterCollection inputParams, WmiOperationOptions inputInvokeOptions,
            out Object returnValue
            )
        {
            using (WmiMethodResult result = InvokeMethod(methodName, inputParams, inputInvokeOptions))
            {
                returnValue = result.ReturnValue.Value;
                if (result.ReturnValue.Type == typeof(bool))
                {
                    return Convert.ToUInt32((bool)result.ReturnValue.Value);
                }
                else if (result.ReturnValue.Type == typeof(uint) || result.ReturnValue.Type == typeof(int))
                {
                    return (uint)result.ReturnValue.Value;
                }
                return 0;
            }
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="resultClassName"></param>
        /// <returns></returns>
        public WmiInstanceCollection GetAssociatedInstances(String resultClassName)
        {
            return GetAssociatedInstances(resultClassName, null, null, null);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="resultClassName"></param>
        /// <param name="associationClassName"></param>
        /// <param name="resultRole"></param>
        /// <param name="sourceRole"></param>
        /// <returns></returns>
        public WmiInstanceCollection GetAssociatedInstances(
                                                   String resultClassName,
                                                   String associationClassName,
                                                   String resultRole,
                                                   String sourceRole
                                                   )
        {
            try
            {
                WmiInstanceCollection associatedInstanceList = new WmiInstanceCollection();

                IEnumerable<CimInstance> associatedInstances = m_CimSession.EnumerateAssociatedInstances(
                                                                    GetNamespace(),
                                                                    m_Instance,
                                                                    associationClassName,
                                                                    resultClassName,
                                                                    sourceRole,
                                                                    resultRole);

                foreach (CimInstance assocInstance in associatedInstances)
                {
                    using (assocInstance)
                    {
                        using (WmiInstance instance = new WmiInstance(assocInstance))
                        {
                            Type cType = WmiInstanceManager.GetType(resultClassName);
                            if (cType != null)
                                associatedInstanceList.Add(
                                    (WmiInstance)Activator.CreateInstance(cType, instance));
                            else
                                associatedInstanceList.Add(instance.Clone() as WmiInstance);
                        }
                    }
                }
                return associatedInstanceList;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }
        /// <summary>
        /// 
        /// </summary>
        /// <returns></returns>
        public WmiInstanceCollection GetRelated()
        {
            return GetAssociatedInstances(null, null, null, null);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="resultClassName"></param>
        /// <returns></returns>
        public T GetRelated<T>(String resultClassName) where T : IWmiInstance
        {
            using (WmiInstanceCollection col = GetRelated(resultClassName))
            {
                if (col.Count > 0)
                {
                    return (T)Activator.CreateInstance(typeof(T), col.GetFirstInstance());
                }
            }
            return default(T);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="resultClassName"></param>
        /// <returns></returns>
        public WmiInstanceCollection<T> GetAllRelated<T>(String resultClassName) where T : IWmiInstance
        {
            using (WmiInstanceCollection col = GetRelated(resultClassName))
            {
                WmiInstanceCollection<T> resultCol = new WmiInstanceCollection<T>();
                foreach (var item in col)
                {
                    resultCol.Add((T)Activator.CreateInstance(typeof(T), item));
                }
                return resultCol;
            }
        }

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="resultClassName"></param>
        /// <returns></returns>
        public WmiInstanceCollection<T> GetAllRelated<T>(IWmiQuery query) where T : IWmiInstance
        {
            using (WmiInstanceCollection col = GetInstances(GetServerName(), GetNamespace(), query, null, null))
            {
                WmiInstanceCollection<T> resultCol = new WmiInstanceCollection<T>();
                foreach (var item in col)
                {
                    resultCol.Add((T)Activator.CreateInstance(typeof(T), item));
                }
                return resultCol;
            }
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="resultClassName"></param>
        /// <returns></returns>
        public WmiInstanceCollection GetRelated(String resultClassName)
        {
            return GetAssociatedInstances(resultClassName, null, null, null);
        }

        public WmiInstanceCollection<T> GetAllRelated<T>(
                                                   String resultClassName,
                                                   String associationClassName,
                                                   String resultRole,
                                                   String sourceRole
                                                   ) where T : IWmiInstance
        {
            using (WmiInstanceCollection col = GetRelated(resultClassName, associationClassName, resultRole, sourceRole))
            {
                WmiInstanceCollection<T> resultCol = new WmiInstanceCollection<T>();
                foreach (var item in col)
                {
                    resultCol.Add((T)Activator.CreateInstance(typeof(T), item));
                }
                return resultCol;
            }
        }

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="resultClassName"></param>
        /// <param name="associationClassName"></param>
        /// <param name="resultRole"></param>
        /// <param name="sourceRole"></param>
        /// <returns></returns>
        public T GetRelated<T>(
                                                   String resultClassName,
                                                   String associationClassName,
                                                   String resultRole,
                                                   String sourceRole
                                                   )
        {
            using (WmiInstanceCollection col = GetRelated(resultClassName, associationClassName, resultRole, sourceRole))
            {
                if (col.Count > 0)
                    return (T)Activator.CreateInstance(typeof(T), col.GetFirstInstance());
            }
            return default(T);
        }

        public WmiInstanceCollection GetRelated(
                                                   String resultClassName,
                                                   String associationClassName,
                                                   String resultRole,
                                                   String sourceRole
                                                   )
        {
            return GetAssociatedInstances(resultClassName, associationClassName, resultRole, sourceRole);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="associationClassName"></param>
        /// <param name="sourceRole"></param>
        /// <returns></returns>
        public WmiInstanceCollection EnumerateReferencingInstances(
                                            String associationClassName,
                                            String sourceRole
                                           )
        {
            try
            {
                WmiInstanceCollection referencingInstances = new WmiInstanceCollection();

                IEnumerable<CimInstance> associatedInstances = m_CimSession.EnumerateReferencingInstances(
                                                                                    GetNamespace(),
                                                                                    m_Instance,
                                                                                    associationClassName,
                                                                                    sourceRole);

                foreach (CimInstance assocInstance in associatedInstances)
                {
                    referencingInstances.Add(new WmiInstance(assocInstance));
                }
                return referencingInstances;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// 
        /// </summary>
        /// <returns></returns>
        public WmiInstanceCollection GetRelationships()
        {
            return EnumerateReferencingInstances(null, null);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="associationClassName"></param>
        /// <returns></returns>
        public WmiInstanceCollection GetRelationships(String associationClassName)
        {
            return EnumerateReferencingInstances(associationClassName, null);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="associationClassName"></param>
        /// <param name="sourceRole"></param>
        /// <returns></returns>
        public WmiInstanceCollection GetRelationships(String associationClassName, string sourceRole)
        {
            return EnumerateReferencingInstances(associationClassName, sourceRole);
        }

        #endregion Methods

        #region Verify
        protected List<String> m_ModifiedProperties = new List<String>();
        /// <summary>
        /// Enabled verification types
        /// </summary>
        private WmiManagementVerify gm_Verify = WmiManagementVerify.None;

        /// <summary>
        /// List of modified property names
        /// </summary>
        public List<String> ModifiedProperties { get { return this.m_ModifiedProperties; } }

        /// <summary>
        /// Default enabled verification types
        /// </summary>
        public WmiManagementVerify DefaultVerify { get { return (WmiManagementVerify.Put | WmiManagementVerify.Method); } }

        private bool gm_Modify = true;
        /// <summary>
        /// 
        /// </summary>
        protected bool ShouldModify { get { return gm_Modify; } set { gm_Modify = value; } }

        /// <summary>
        /// Enabled verification types
        /// </summary>
        public WmiManagementVerify VerifyInstance { get { return gm_Verify; } set { gm_Verify = value; } }

        public Object InvokeMethod(string methodName, WmiMethodParameterCollection args,
            WmiManagementMethodVerify verifyDelegate)
        {
            return InvokeMethod(methodName, args, null, verifyDelegate);
        }
        /// <summary>
        /// Invoke WMI method
        /// http://msdn.microsoft.com/en-us/library/f9ck6sf2.aspx
        /// </summary>
        /// <param name="verifyDelegate">Called at completion of method to verify results</param>
        /// <returns>Object</returns>
        public Object InvokeMethod(string methodName, WmiMethodParameterCollection args,
            WmiOperationOptions options, WmiManagementMethodVerify verifyDelegate)
        {
            WmiMethodResult returnValue = InvokeMethod(methodName, args);

            if (verifyDelegate != null && ShouldVerify(WmiManagementVerify.Method))
            {
                verifyDelegate(this, methodName, args, null);
            }

            return returnValue;
        }

        /// <summary>
        /// Determines if WMI verification type is enabled
        /// </summary>
        /// <param name="verify">Type of verification</param>
        /// <returns>True if verification type is enabled</returns>
        public bool ShouldVerify(WmiManagementVerify verify)
        {
            return ((gm_Verify & verify) == verify);
        }

        /// <summary>
        /// Adds propertyName to modified properties list
        /// </summary>
        /// <param name="propertyName">Name of modified property</param>
        internal void MarkPropertyModified(String propertyName)
        {
            if (!this.m_ModifiedProperties.Contains(propertyName.ToLowerInvariant()))
            {
                this.m_ModifiedProperties.Add(propertyName.ToLowerInvariant());
            }
        }

        #endregion

        #region ICloneable Members
        /// <summary>
        /// 
        /// </summary>
        /// <returns></returns>
        public object Clone()
        {
            return new WmiInstance(m_Instance);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <returns></returns>
        public T Clone<T>()
        {
            return (T)Activator.CreateInstance(typeof(T), m_Instance);
        }

        #endregion

        #region IDisposable Members
        private bool m_Disposed;
        protected virtual void Dispose(bool disposing)
        {
            if (disposing && !m_Disposed)
            {
                using (m_Instance) { }
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

    /// <summary>
    /// Verification exception thrown when VmManagement verification fails
    /// </summary>
    public class WmiManagementVerificationException : Exception
    {
        #region Construction

        public WmiManagementVerificationException(String message)
            : base(message)
        {
        }

        public WmiManagementVerificationException(String message, Exception innerException)
            : base(message, innerException)
        {
        }

        #endregion Construction
    }

    /// <summary>
    /// Verification exception thrown when VmManagement verification fails
    /// </summary>
    public class WmiManagementVerificationArgumentNullException : WmiManagementVerificationException
    {
        #region Data

        private String m_ParamName;

        #endregion Data

        #region Construction

        public WmiManagementVerificationArgumentNullException(String paramName)
            : this(paramName, String.Empty, null)
        {
        }

        public WmiManagementVerificationArgumentNullException(String paramName, String message)
            : this(paramName, message, null)
        {
        }

        public WmiManagementVerificationArgumentNullException(String paramName, Exception innerException)
            : this(paramName, String.Empty, innerException)
        {
        }

        public WmiManagementVerificationArgumentNullException(String paramName, String message, Exception innerException)
            : base(message, innerException)
        {
            this.m_ParamName = paramName;
        }

        #endregion Construction

        #region Public Properties

        public String ParamName { get { return this.m_ParamName; } }

        #endregion Public Properties
    }
}
