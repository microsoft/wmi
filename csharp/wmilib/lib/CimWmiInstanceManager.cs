// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Management.Infrastructure;
using Microsoft.Management.Infrastructure.Options;
using System.Collections.ObjectModel;
using System.Reflection;
using System.Collections.Concurrent;
using Microsoft.WmiCodeGen.CSharp.Interface;

namespace Microsoft.WmiCodeGen.CSharp.Lib
{
    /// <summary>
    /// 
    /// </summary>
    public class WmiInstanceManager : IWmiInstanceManager
    {
        #region Data
        private WmiCredentials m_Credentials;
        public WmiSession Session { get; set; }
        private String m_Namespace;
        private String m_ServerName;
        #endregion Data

        #region Properties

        /// <summary>
        /// 
        /// </summary>
        public WmiCredentials Credentials
        {
            get
            {
                return m_Credentials;
            }
        }

        public String ServerName
        {
            get
            {
                return m_ServerName;
            }
        }

        public String Namespace
        {
            get
            {
                return m_Namespace;
            }
        }

        #endregion Properties

        #region Construction

        /// <summary>
        /// Constructs object manager that surfaces Cim objects
        /// </summary>
        /// <param name="serverName">Host name</param>
        /// <param name="Namespace">Namespace </param>
        public WmiInstanceManager(String serverName, String Namespace)
            : this(serverName, Namespace, null)
        {
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiInstanceManager" /> class.
        /// </summary>
        /// <param name="serverName">Name of the server.</param>
        /// <param name="namespaceName">The namespace.</param>
        /// <param name="userName">Name of the user.</param>
        /// <param name="password">The password.</param>
        /// <param name="domain">The domain.</param>
        public WmiInstanceManager(String serverName, String namespaceName, String userName, String password, String domain)
            : this(serverName, namespaceName, new WmiCredentials(userName, password, domain, PasswordAuthenticationMechanism.Default))
        {
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="serverName"></param>
        /// <param name="Namespace"></param>
        /// <param name="credentials"></param>
        public WmiInstanceManager(string serverName, string namespaceName, WmiCredentials credentials)
        {
            this.m_Namespace = namespaceName;
            this.m_ServerName = serverName;
            Session = new WmiSession();
            m_Credentials = credentials;
            Session.Connect(serverName, namespaceName, m_Credentials);
        }
        #endregion Construction

        #region Methods

        /// <summary>
        /// Gets first instance of a specific class
        /// </summary>
        /// <param name="className"> Class Name</param>
        /// <returns></returns>
        public IWmiInstance GetFirstInstance(String className)
        {
            if (className == null)
            {
                throw new ArgumentNullException("className", "className cannot be null");
            }

            try
            {
                using (WmiInstanceCollection enumList = EnumerateInstances(className))
                {
                    if (enumList != null && enumList.Count > 0)
                    {
                        IEnumerator<IWmiInstance> enumerator = enumList.GetEnumerator();
                        enumerator.MoveNext();
                        return enumerator.Current.Clone() as WmiInstance;
                    }
                }
                return null;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }


        /// <summary>
        /// Enumerate instances of a specific class
        /// </summary>
        /// <param name="className">Class Name</param>
        /// <returns></returns>
        public WmiInstanceCollection EnumerateInstances(String className)
        {
            if (className == null)
            {
                throw new ArgumentNullException("className", "className cannot be null");
            }

            try
            {
                WmiInstanceCollection enumInstanceList = new WmiInstanceCollection();
                IEnumerable<CimInstance> enumeratedInstances = Session.Session.EnumerateInstances(m_Namespace, className);
                foreach (CimInstance cimInstance in enumeratedInstances)
                {
                    using (cimInstance)
                    {
                        using (IWmiInstance newInstance = new WmiInstance(cimInstance))
                        {
                            try
                            {
                                Type cType = GetType(className);
                                if (cType != null)
                                    enumInstanceList.Add((IWmiInstance)Activator.CreateInstance(cType, newInstance));
                                else
                                    enumInstanceList.Add(newInstance.Clone() as IWmiInstance);
                            }
                            catch (Exception)
                            {
                                enumInstanceList.Add(newInstance.Clone() as IWmiInstance);
                            }
                        }
                    }
                }
                return enumInstanceList;
            }
            catch (CimException ex)
            {
                Console.WriteLine(ex.Message);
                Console.WriteLine(ex.StackTrace);
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="query"></param>
        /// <returns></returns>
        public T QueryInstance<T>(IWmiQuery query) where T : IWmiInstance
        {
            using (WmiInstanceCollection<T> col = QueryInstances<T>(query))
            {
                return (T)Activator.CreateInstance(typeof(T), col.First());
            }
        }
        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="query"></param>
        /// <returns></returns>
        public WmiInstanceCollection<T> QueryInstances<T>(IWmiQuery query) where T : IWmiInstance
        {
            if (query == null) throw new ArgumentNullException("query");
            return QueryInstances<T>(query.QueryString);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="query"></param>
        /// <returns></returns>
        public WmiInstanceCollection QueryInstances(IWmiQuery query)
        {
            if (query == null) throw new ArgumentNullException("query");
            WmiInstanceCollection queryInstanceList = new WmiInstanceCollection();
            Type cType = GetType(query.WmiClassName);
            WmiInstanceCollection collection = QueryInstances(query.QueryString);
            if (cType == null) return collection;

            using (collection)
            {
                foreach (var item in collection)
                {
                    queryInstanceList.Add((IWmiInstance)Activator.CreateInstance(cType, item));
                }
            }
            return queryInstanceList;
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="className"></param>
        /// <param name="filters"></param>
        /// <returns></returns>
        public WmiInstanceCollection QueryInstances(string className, params string[] filters)
        {
            return QueryInstances(new WmiQuery(className, filters));
        }

        public IWmiInstance QueryInstance(string className)
        {
            return QueryInstance(className, string.Empty);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="className"></param>
        /// <param name="selectPropertyName"></param>
        /// <param name="selectPropertyValue"></param>
        /// <returns></returns>
        public IWmiInstance QueryInstance(string className, params string[] filters)
        {
            using (WmiInstanceCollection collection = QueryInstances(className, filters))
            {
                return collection.First().Clone() as IWmiInstance;
            }
        }

        /// <summary>
        /// Query instances using specified select query
        /// </summary>
        /// <param name="selectQuery">select query</param>
        /// <returns></returns>
        public WmiInstanceCollection<T> QueryInstances<T>(String selectQuery) where T : IWmiInstance
        {
            WmiInstanceCollection<T> queryInstanceList = new WmiInstanceCollection<T>();
            using (WmiInstanceCollection collection = QueryInstances(selectQuery))
            {
                foreach (var item in collection)
                {
                    queryInstanceList.Add((T)Activator.CreateInstance(typeof(T), item));
                }
            }
            return queryInstanceList;
        }
        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="selectQuery"></param>
        /// <returns></returns>
        public T QueryInstance<T>(String selectQuery) where T : IWmiInstance
        {
            using (WmiInstanceCollection<T> collection = QueryInstances<T>(selectQuery))
            {
                return (T) Activator.CreateInstance(typeof(T), collection.First());
            }
        }

        public WmiInstanceCollection QueryInstances(String selectQuery)
        {
            if (selectQuery == null)
            {
                throw new ArgumentNullException("selectQuery");
            }

            try
            {
                WmiInstanceCollection queryInstanceList = new WmiInstanceCollection();

                IEnumerable<CimInstance> queryInstances = Session.Session.QueryInstances(m_Namespace, "WQL", selectQuery);
                foreach (CimInstance cimInstance in queryInstances)
                {
                    queryInstanceList.Add(new WmiInstance(cimInstance));
                }

                return queryInstanceList;
            }
            catch (CimException ex)
            {
                Console.WriteLine(selectQuery);
                Console.WriteLine(ex.Message);
                Console.WriteLine(ex.StackTrace);
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="className"></param>
        /// <param name="propertyValues"></param>
        /// <returns></returns>
        public T CreateInstance<T>(String className, WmiKeyValueCollection propertyValues)
        {
            using (IWmiInstance instance = CreateInstance(className, propertyValues))
            {
                return (T)Activator.CreateInstance(typeof(T), instance);
            }
        }
        /// <summary>
        /// Create the Class Instance, will all properties.
        /// Key properties has to be filled to commit this instance
        /// </summary>
        /// <param name="className">Class name whose instance needs to be created</param>
        /// <param name="propertyValues">Property Values</param>
        /// <returns></returns>
        public IWmiInstance CreateInstance(String className, WmiKeyValueCollection propertyValues)
        {
            if (string.IsNullOrEmpty(className))
            {
                throw new ArgumentNullException("className", "className cannot be null");
            }

            try
            {
                using (CimClass cimClass = Session.Session.GetClass(m_Namespace, className))
                {
                    // Create an instalce with all properties, with null
                    using (CimInstance instance = new CimInstance(cimClass))
                    {
                        if (propertyValues != null)
                        {
                            var cimProperties = from p in cimClass.CimClassProperties select p;
                            foreach (CimPropertyDeclaration property in cimProperties)
                            {
                                if (propertyValues.Contains(property.Name))
                                {
                                    instance.CimInstanceProperties[property.Name].Value = propertyValues[property.Name].Value;
                                    //instance.CimInstanceProperties.Add(CimProperty.Create(property.Name, propertyValue, property.CimType, CimFlags.None));
                                }
                            }
                        }

                        using (IWmiInstance newInstance = new WmiInstance(instance))
                        {

                            try
                            {
                                Type cType = GetType(className);
                                if (cType != null)
                                    return (IWmiInstance)Activator.CreateInstance(cType, newInstance);
                                else
                                    return newInstance.Clone() as IWmiInstance;
                            }
                            catch (Exception)
                            {
                                return newInstance.Clone() as IWmiInstance;
                            }
                        }
                    }
                }
            }
            catch (CimException ex)
            {
                Console.WriteLine(ex.Message);
                Console.WriteLine(ex.StackTrace);
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Creates an Empty instance. Properties has to be Added to this instance
        /// & then Commited.
        /// </summary>
        /// <param name="className"></param>
        /// <returns></returns>
        public IWmiInstance CreateInstance(String className)
        {
            if (string.IsNullOrEmpty(className))
            {
                throw new ArgumentNullException("className", "className cannot be null");
            }

            try
            {
                // Create the instance with no properties.
                // Properties needs to be added
                // Do not dispose this instance.
                CimInstance instance = new CimInstance(className, m_Namespace);
                {
                    // Workaround - WmiIntance doesn't create a clone for this instance.
                    IWmiInstance newInstance = new WmiInstance(instance);
                    {
                        try
                        {
                            // Try to create an object of type className
                            Type cType = GetType(className);
                            if (cType != null)
                                return (IWmiInstance)Activator.CreateInstance(cType, newInstance);
                            else
                                return newInstance; //.Clone() as IWmiInstance;
                        }
                        catch (Exception)
                        {
                            return newInstance; //.Clone() as IWmiInstance;
                        }
                    }
                }
            }
            catch (CimException ex)
            {
                Console.WriteLine(ex.Message);
                Console.WriteLine(ex.StackTrace);
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Creates the instance on server.
        /// </summary>
        /// <param name="className">Name of the class.</param>
        /// <param name="propertyValues">The property values.</param>
        /// <returns></returns>
        public IWmiInstance CreateInstanceAndSubmit(String className, WmiKeyValueCollection propertyValues)
        {
            try
            {
                IWmiInstance instance = CreateInstance(className, propertyValues);
                CimInstance inputInstance = instance.GetInstance() as CimInstance;
                CimInstance outputInstance = Session.Session.CreateInstance(m_Namespace, inputInstance);
                return new WmiInstance(outputInstance);
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets the instance.
        /// </summary>
        /// <param name="instance">The instance.</param>
        /// <returns></returns>
        public IWmiInstance GetInstance(IWmiInstance instance)
        {
            try
            {
                CimInstance outInstance = Session.Session.GetInstance(m_Namespace, (CimInstance)instance.GetInstance());
                return new WmiInstance(outInstance);
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets the instance of the class, and fills the property values
        /// </summary>
        /// <param name="className">Name of the class.</param>
        /// <param name="PropertyValues">The vm property values.</param>
        /// <returns></returns>
        public IWmiInstance GetInstance(string className, WmiKeyValueCollection PropertyValues)
        {
            try
            {
                IWmiInstance getInstance = CreateInstance(className, PropertyValues);
                return GetInstance(getInstance);
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Enumerates the classes.
        /// </summary>
        /// <returns></returns>
        public WmiClassCollection EnumerateClasses()
        {
            try
            {
                WmiClassCollection enumInstanceList = new WmiClassCollection();
                IEnumerable<CimClass> enumeratedInstances = Session.Session.EnumerateClasses(m_Namespace);
                foreach (CimClass cimClass in enumeratedInstances)
                {
                    enumInstanceList.Add(new WmiClass(cimClass));
                }
                return enumInstanceList;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets the class.
        /// </summary>
        /// <param name="className">Name of the class.</param>
        /// <returns></returns>
        public IWmiClass GetClass(String className)
        {
            try
            {
                CimClass cimClass = Session.Session.GetClass(m_Namespace, className);

                if (cimClass != null)
                {
                    return new WmiClass(cimClass);
                }
                return null;
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets the instance from object path.
        /// </summary>
        /// <param name="path">The path.</param>
        /// <returns></returns>
        public IWmiInstance GetInstanceFromObjectPath(string path)
        {
            try
            {
                return WmiHelper.GetInstanceFromPath(path);
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets the instance array from object path array.
        /// </summary>
        /// <param name="path">The path.</param>
        /// <returns></returns>
        public IWmiInstance[] GetInstanceArrayFromObjectPathArray(string[] pathArray)
        {
            try
            {
                return WmiHelper.GetInstanceArrayFromPathArray(pathArray);
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Gets an event watcher.
        /// </summary>
        /// <param name="wmiNamespace">The WMI namespace to watch.</param>
        /// <param name="query">The event query.</param>
        /// <returns></returns>
        public IWmiEventWatcher GetEventWatcher(string query)
        {
            try
            {
                return new WmiEventWatcher(this.Session.Session, this.m_Namespace, query);
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        #region Unsupportedmethods

        /*public  IWmiInstance GetInstance(String className, Dictionary<string, object> propertyValues)
        {
            if (className == null)
            {
                throw new ArgumentNullException("className");
            }

            CimInstance getInstance = GetInstanceCore(className, propertyValues);

            CimInstance outInstance = m_VmCimWmiSession.Session.GetInstance(m_Namespace, getInstance);

            return new VmCimWmiInstance(outInstance);
        }

        internal CimInstance GetInstanceCore(string cimClassName, Dictionary<string, object> propertyValues)
        {
            CimInstance getInstance = new CimInstance(cimClassName);

            foreach (var property in propertyValues)
            {
                getInstance.CimInstanceProperties.Add(CimProperty.Create(property.Key, property.Value, CimFlags.Key));
            }
            return getInstance;
        }*/

        #endregion Unsupportedmethods

        #region Static Methods
        private static ConcurrentDictionary<string, Type> TypeMap = new ConcurrentDictionary<string, Type>();
        /// <summary>
        /// 
        /// </summary>
        /// <param name="className"></param>
        /// <returns></returns>
        public static Type GetType(string className)
        {
            if (string.IsNullOrEmpty(className)) return null;

            Type type = Type.GetType(className);

            if (TypeMap.TryGetValue(className, out type)) return type;

            if (type == null)
            {
                foreach (Assembly item in AppDomain.CurrentDomain.GetAssemblies())
                {
                    try
                    {
                        type = item.ExportedTypes.ToList().Find(eType => eType.Name == className);
                        if (type != null)
                        {
                            TypeMap[className] = type;
                            break;
                        }
                    }
                    catch (Exception)
                    {
                        // Ignore Exception, since we are making a best effort to find the Type
                    }
                }
            }
            return type;
        }
        #endregion
        #endregion Methods
    }
}
