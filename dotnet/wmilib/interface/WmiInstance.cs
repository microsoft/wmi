// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.


namespace Microsoft.wmi.DotNet.Interface
{
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Text;
    using System.Threading.Tasks;
    using System.Collections.ObjectModel;
    using System.Globalization;

    /// <summary>
    /// WMI Namespaces
    /// </summary>
    public class WMINamespaces
    {
        /// V2 Namespace
        /// </summary>
        public const string V2 = @"root\virtualization\v2";
        public const string V1 = @"root\virtualization";
        public const string CIMV2 = @"root\cimv2";
        public const string ClusterNS = @"root\mscluster";
        /// <summary>
        /// Default Namespace
        /// </summary>
        public static string DefaultVMNamespace = WMINamespaces.V2;
    }
 
    public class WmiInstancePath
    {
        public WmiInstancePath(string path)
        {
            this.InstancePath = path;
        }
        string InstancePath;

        public override string ToString()
        {
            return InstancePath;
        }
    }
    public interface IWmiInstance : IDisposable, ICloneable
    {
#if false
        #region Constructor
        IWmiInstance(string instancePath);
        IWmiInstance(IWmiInstance instance);
        IWmiInstance(String HostName, String WMINameSpace, String className);
        IWmiInstance(String HostName, String WMINameSpace, IWmiQuery query);
        IWmiInstance(String HostName, String WMINameSpace, IWmiQuery query, string userName, string password);
        IWmiInstance(String HostName, String WMINameSpace, String className, string userName, string password);

        #endregion
        
#endif
        /// <summary>
        /// Gets the instance.
        /// </summary>
        /// <returns></returns>
        object GetInstance();

        /// <summary>
        /// Gets the property value.
        /// </summary>
        /// <param name="propertyName">Name of the property.</param>
        /// <returns></returns>
        object GetProperty(String propertyName);


        /// <summary>
        /// Gets the and cast.
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="property">The property.</param>
        /// <returns></returns>
        object GetProperty<T>(string property);

        /// <summary>
        /// Sets the property value.
        /// </summary>
        /// <param name="propertyName">Name of the property.</param>
        /// <param name="propertyValue">The property value.</param>
        void SetProperty<T>(String propertyName, T propertyValue);

        /// <summary>
        /// 
        /// </summary>
        /// <param name="propertyName"></param>
        void ResetProperty(string propertyName);
        /// <summary>
        /// Gets the properties.
        /// </summary>
        /// <value>
        /// The properties.
        /// </value>
        WmiPropertyCollection InstanceProperties
        {
            get;
        }

        /// <summary>
        /// Gets the qualifiers.
        /// </summary>
        /// <value>
        /// The qualifiers.
        /// </value>
        WmiQualifierCollection InstanceQualifiers
        {
            get;
        }

        /// <summary>
        /// Gets or sets the <see cref="System.Object" /> with the specified property name.
        /// </summary>
        /// <value>
        /// The <see cref="System.Object" />.
        /// </value>
        /// <param name="propertyName">Name of the property.</param>
        /// <returns></returns>
        object this[string propertyName]
        { 
            get;
            set;
        }

        /// <summary>
        /// Gets the class.
        /// </summary>
        /// <value>
        /// The class.
        /// </value>
        IWmiClass Class
        {
            get;
        }

        /// <summary>
        /// Gets the embedded instance.
        /// </summary>
        /// <value>
        /// The embedded instance.
        /// </value>
        string EmbeddedInstance
        {
            get;
        }

        IWmiInstanceManager InstanceManager { get; }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="destination"></param>
        /// <returns></returns>
        bool Equals(object destination);

        /// <summary>
        /// Reloads Instance from the server
        /// </summary>
        void RefreshInstance();

        /// <summary>
        /// Creates a new Instance
        /// </summary>
        void CommitInstance();
        /// <summary>
        /// 
        /// </summary>
        /// <param name="options"></param>
        void CommitInstance(WmiOperationOptions options);
        /// <summary>
        /// Puts the local changes to the Server
        /// & Refresh local instace
        /// </summary>
        void ModifyInstance();

        /// <summary>
        /// Deletes this Instance in the server
        /// </summary>
        void DeleteInstance();

        /// <summary>
        /// 
        /// </summary>
        /// <returns></returns>
        string GetInstancePath();

        /// <summary>
        /// 
        /// </summary>
        /// <param name="methodName"></param>
        /// <param name="inputParams"></param>
        /// <returns></returns>
        WmiMethodResult InvokeMethod(string methodName, WmiMethodParameterCollection inputParams);

        /// <summary>
        /// Call the Method with inputParams
        /// </summary>
        /// <returns></returns>
        WmiMethodResult InvokeMethod(string methodName, 
            WmiMethodParameterCollection inputParams,
            WmiOperationOptions inputOptions);

        WmiInstanceCollection GetRelated(String resultClassName);

        /// <summary>
        /// 
        /// </summary>
        /// <param name="resultClassName"></param>
        /// <param name="associationClassName"></param>
        /// <param name="resultRole"></param>
        /// <param name="sourceRole"></param>
        /// <returns></returns>
        WmiInstanceCollection GetRelated(String resultClassName,
                                                   String associationClassName,
                                                   String resultRole,
                                                   String sourceRole);
        /// <summary>
        /// 
        /// </summary>
        /// <param name="resultClassName"></param>
        /// <param name="associationClassName"></param>
        /// <param name="resultRole"></param>
        /// <param name="sourceRole"></param>
        /// <returns></returns>
        WmiInstanceCollection GetAssociatedInstances(String resultClassName,
                                                   String associationClassName,
                                                   String resultRole,
                                                   String sourceRole);
        /// <summary>
        /// 
        /// </summary>
        /// <param name="associationClassName"></param>
        /// <param name="sourceRole"></param>
        /// <returns></returns>
        WmiInstanceCollection EnumerateReferencingInstances(
                                            String associationClassName,
                                            String sourceRole
            );

    }

    #region Common

    public class WmiManager
    {
        private static ClientAPIOptions m_DefaultAPI = ClientAPIOptions.ManagementInfrastructureAPI;

        public static ClientAPIOptions DefaultAPI
        {
            get
            {
                return m_DefaultAPI;
            }            
        }

        public static readonly string DefaultNamespace = WMINamespaces.V2;

        public static void SetClientAPI(ClientAPIOptions clientAPI)
        {
            m_DefaultAPI = clientAPI;
        }
    }

    public enum ClientAPIOptions
    {
        SystemManagementAPI = 1,
        ManagementInfrastructureAPI = 2
    }

    public static class WmiInstanceExtensions
    {
        /// <summary>
        /// Gets the embedded instances.
        /// </summary>
        /// <param name="instanceArray">The instance array.</param>
        /// <returns></returns>
        public static string[] GetEmbeddedInstances(this IWmiInstance[] instanceArray)
        {
            if (instanceArray == null) 
            {
                return null;
            }

            if (instanceArray.Length == 0)
            {
                return new string[] { };
            }

            string[] embeddedInstanceArray = new string[instanceArray.Length];

            for (int index = 0; index < instanceArray.Length; index++)
            {
                embeddedInstanceArray[index] = instanceArray[index].EmbeddedInstance;
            }
            return embeddedInstanceArray;
        }

        /// <summary>
        /// Gets the paths.
        /// </summary>
        /// <param name="instanceArray">The instance array.</param>
        /// <returns></returns>
        public static string[] GetInstancePaths(this IWmiInstance[] instanceArray)
        {
            if (instanceArray == null)
            {
                return null;
            }

            if (instanceArray.Length == 0)
            {
                return new string[] { };
            }

            string[] pathArray = new string[instanceArray.Length];

            for (int index = 0; index < instanceArray.Length; index++)
            {
                pathArray[index] = instanceArray[index].GetInstancePath();
            }
            return pathArray;
        }

        public static void Dispose(this IWmiInstance[] instanceArray)
        {
            if (instanceArray != null)
            {
                foreach (var item in instanceArray)
                {
                    using (item) { }
                }
            }
        }

        /// <summary>
        /// Retrieve a specified property and cast it into the returned type.
        /// Will throw exceptions if casting fails or if property is null.
        /// </summary>
        /// <typeparam name="T">returned type</typeparam>
        /// <param name="obj">implicit ManagementObject</param>
        /// <param name="property">name of property to retrieve</param>
        /// <returns>property's value in specified type</returns>
        public static T GetAndCast<T>(this IWmiInstance obj, string property)
        {
            return (T)obj.GetProperty(property);
        }

        /// <summary>
        /// Retrieve a specified property and convert it to a DateTime
        /// </summary>
        /// <param name="obj">implicit ManagementObject</param>
        /// <param name="property">name of property to retrieve</param>
        /// <returns>DateTime value of specified property or default DateTime if does not exist</returns>
        public static DateTime GetDateTime(this IWmiInstance obj, string property)
        {
            object temp = obj.GetProperty(property);
            if (temp != null)
            {
                if (temp is DateTime)
                {
                    return (DateTime)temp;
                }
                else
                {
                    DateTime.Parse(temp.ToString(), CultureInfo.InvariantCulture);
                }
            }
            return new DateTime();
        }
    }


    #endregion Common

}
