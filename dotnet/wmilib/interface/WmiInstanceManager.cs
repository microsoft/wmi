// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

namespace Microsoft.WmiCodeGen.DotNet.Interface
{
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Text;
    using System.Threading.Tasks;
    using System.Collections.ObjectModel;

    #region Exception
    /// <summary>
    /// 
    /// </summary>
    public class WmiInstanceManagerException : Exception
    {
        #region Basic constructors

        /// <summary>
        /// Constructs a WmiException from a custom error message
        /// </summary>
        /// <param name="message">Custom error message</param>
        public WmiInstanceManagerException(string message)
            : this(message, null)
        {
        }

        /// <summary>
        /// Constructs a WmiException from a custom error message and an inner exceptioon
        /// </summary>
        /// <param name="message"></param>
        /// <param name="innerException"></param>
        public WmiInstanceManagerException(string message, Exception innerException)
            : base(message, innerException)
        {
        }

        #endregion
    } 
    #endregion

    /// <summary>
    /// 
    /// </summary>
    public interface IWmiInstanceManager // : IDisposable // TODO
    {
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

        string Namespace
        {
            get;
        }

        /// <summary>
        /// 
        /// </summary>
        WmiCredentials Credentials
        {
            get;
        }
        /// <summary>
        /// Enumerates the instances.
        /// </summary>
        /// <param name="className">Name of the class.</param>
        /// <returns></returns>
        WmiInstanceCollection EnumerateInstances(String className);

        /// <summary>
        /// Queries the instances.
        /// </summary>
        /// <param name="query">The query.</param>
        /// <returns></returns>
        WmiInstanceCollection QueryInstances(String query);
        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="query"></param>
        /// <returns></returns>
        WmiInstanceCollection<T> QueryInstances<T>(String query) where T : IWmiInstance;

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="query"></param>
        /// <returns></returns>
        T QueryInstance<T>(String query) where T : IWmiInstance;
        
        /// <summary>
        /// 
        /// </summary>
        /// <param name="query"></param>
        /// <returns></returns>
        WmiInstanceCollection QueryInstances(IWmiQuery query);
        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="query"></param>
        /// <returns></returns>
        WmiInstanceCollection<T> QueryInstances<T>(IWmiQuery query) where T : IWmiInstance;

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="query"></param>
        /// <returns></returns>
        T QueryInstance<T>(IWmiQuery query) where T : IWmiInstance;

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="className"></param>
        /// <param name="propertyValues"></param>
        /// <returns></returns>
        T CreateInstance<T>(String className, WmiKeyValueCollection propertyValues);

        /// <summary>
        /// Creates the instance.
        /// </summary>
        /// <param name="className">Name of the class.</param>
        /// <param name="propertyValues">The property values.</param>
        /// <returns></returns>
        IWmiInstance CreateInstance(String className, WmiKeyValueCollection propertyValues);

        /// <summary>
        /// 
        /// </summary>
        /// <param name="className"></param>
        /// <returns></returns>
        IWmiInstance CreateInstance(String className);

        /// <summary>
        /// Gets the first instance.
        /// </summary>
        /// <param name="className">Name of the class.</param>
        /// <returns></returns>
        IWmiInstance GetFirstInstance(String className);
        
        /// <summary>
        /// Gets the instance.
        /// </summary>
        /// <param name="instance">The instance.</param>
        /// <returns></returns>
        IWmiInstance GetInstance(IWmiInstance instance);


        /// <summary>
        /// Gets the instance.
        /// </summary>
        /// <param name="className">Name of the class.</param>
        /// <param name="vmPropertyValues">The vm property values.</param>
        /// <returns></returns>
        IWmiInstance GetInstance(string className, WmiKeyValueCollection vmPropertyValues);
        
        /// <summary>
        /// Gets the class.
        /// </summary>
        /// <param name="className">Name of the class.</param>
        /// <returns></returns>
        IWmiClass GetClass(String className);
        
        /// <summary>
        /// Enumerates the classes.
        /// </summary>
        /// <returns></returns>
        WmiClassCollection EnumerateClasses();

        /// <summary>
        /// Creates the instance on server and submits it
        /// </summary>
        /// <param name="className">Name of the class.</param>
        /// <param name="propertyValues">The property values.</param>
        /// <returns></returns>
        IWmiInstance CreateInstanceAndSubmit(String className, WmiKeyValueCollection propertyValues);

        /// <summary>
        /// Gets the instance from object path.
        /// </summary>
        /// <param name="path">The path.</param>
        /// <returns></returns>
        IWmiInstance GetInstanceFromObjectPath(string path);

        /// <summary>
        /// Gets the instance from object path.
        /// </summary>
        /// <param name="path">The path.</param>
        /// <returns></returns>
        IWmiInstance[] GetInstanceArrayFromObjectPathArray(string[] pathArray);

        /// <summary>
        /// Gets an event watcher.
        /// </summary>
        /// <param name="wmiNamespace">The WMI namespace to watch.</param>
        /// <param name="query">The event query.</param>
        /// <returns></returns>
        IWmiEventWatcher GetEventWatcher(string query);
    }
}
