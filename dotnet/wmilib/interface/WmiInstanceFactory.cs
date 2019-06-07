// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

namespace Microsoft.WmiCodeGen.DotNet.Interface
{
    using System;
    using System.Collections.Generic;
    using System.IO;
    using System.Linq;
    using System.Text;
    using System.Threading.Tasks;
    using System.Reflection;
    using System.Globalization;

    public class WmiInstanceFactory
    {
        private static string LegacyWmiAssemblyName = "Microsoft.WmiCodeGen.DotNet.Interface.Legacy";

        //private static string WmiInstanceManagerName = "Microsoft.WmiCodeGen.DotNet.Interface.WmiInstanceManager";

        private static string CimWmiAssemblyName = "Microsoft.WmiCodeGen.DotNet.Interface.Cim";

        private static string WmiInstanceManagerName = "Microsoft.WmiCodeGen.DotNet.Interface.WmiInstanceManager";

        private static string CssCimWmiAssemblyName = "Microsoft.WmiCodeGen.DotNet.Interface.Cim.CSS";

        private static Dictionary<string, IWmiInstanceManager> WmiInstanceManagerCollection = new Dictionary<string, IWmiInstanceManager>();

        /// <summary>
        /// In case, only hostname is passed, we can still try to retrieve the correct instance manager previously created.
        /// </summary>
        /// <param name="key"></param>
        /// <param name="instanceManager"></param>
        /// <returns></returns>
        private static bool LookupWmiInstanceManagerCollection(string key, out IWmiInstanceManager instanceManager)
        {
            instanceManager = null;

            lock (WmiInstanceManagerCollection)
            {
                foreach (var item in WmiInstanceManagerCollection.Keys)
                {
                    if (item.Contains(key))
                    {
                        instanceManager = WmiInstanceManagerCollection[item];
                        return true;
                    }
                }
            }

            return false;
        }

        /// <summary>
        /// If we get a bigger key, remove the old instance, and keep the new one
        /// </summary>
        /// <param name="key"></param>
        /// <param name="instanceManager"></param>
        private static void AddWmiInstanceManagerCollection(string key, IWmiInstanceManager instanceManager)
        {
            lock (WmiInstanceManagerCollection)
            {
                bool addedFlag = false;
                foreach (var item in WmiInstanceManagerCollection.Keys)
                {
                    if (!key.Equals(item) && key.Contains(item))
                    {
                        WmiInstanceManagerCollection.Remove(item); // TODO : Clean up WmiSession
                        WmiInstanceManagerCollection[key] = instanceManager;
                        addedFlag = true;
                        break;
                    }
                }

                if (!addedFlag)
                    WmiInstanceManagerCollection[key] = instanceManager;
            }
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="parameters"></param>
        /// <returns></returns>
        private static string GetWmiInstanceManagerKey(params string[] parameters)
        {
            String key = String.Empty;
            if (parameters != null)
            {
                foreach (var item in parameters)
                {
                    if (!string.IsNullOrEmpty(item))
                    {
                        key += item.ToLowerInvariant().Replace('\\', '_').Replace('/', '_') + "#";
                    }
                }
            }
            return key;
        }

        public static IWmiInstanceManager GetInstanceManager(String SystemName)
        {
            return WmiInstanceFactory.GetWmiInstanceManager(
                WmiInstanceFactory.GetWmiAssemblyName(WmiManager.DefaultAPI),
                WmiInstanceFactory.GetWmiInstanceManagerName(WmiManager.DefaultAPI),
                SystemName);
        }

        public static IWmiInstanceManager GetInstanceManager(String SystemName, String Namespace)
        {
            return WmiInstanceFactory.GetWmiInstanceManager(
                WmiInstanceFactory.GetWmiAssemblyName(WmiManager.DefaultAPI),
                WmiInstanceFactory.GetWmiInstanceManagerName(WmiManager.DefaultAPI),
                SystemName,
                Namespace);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="SystemName"></param>
        /// <param name="credentials"></param>
        /// <returns></returns>
        public static IWmiInstanceManager GetInstanceManager(String SystemName, String Namespace, WmiCredentials credentials)
        {
            return WmiInstanceFactory.GetWmiInstanceManager(
                WmiInstanceFactory.GetWmiAssemblyName(WmiManager.DefaultAPI),
                WmiInstanceFactory.GetWmiInstanceManagerName(WmiManager.DefaultAPI),
                SystemName,
                Namespace,
                credentials);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="SystemName"></param>
        /// <param name="UserName"></param>
        /// <param name="Password"></param>
        /// <param name="Domain"></param>
        /// <returns></returns>
        public static IWmiInstanceManager GetInstanceManager(String SystemName, String UserName,
            String Password, String Domain)
        {
            return WmiInstanceFactory.GetWmiInstanceManager(
                WmiInstanceFactory.GetWmiAssemblyName(WmiManager.DefaultAPI),
                WmiInstanceFactory.GetWmiInstanceManagerName(WmiManager.DefaultAPI),
                SystemName,
                UserName,
                Password,
                Domain);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="SystemName"></param>
        /// <param name="Namespace"></param>
        /// <param name="UserName"></param>
        /// <param name="Password"></param>
        /// <param name="Domain"></param>
        /// <returns></returns>
        public static IWmiInstanceManager GetInstanceManager(String SystemName, String Namespace,
            String UserName, String Password, String Domain)
        {
            return WmiInstanceFactory.GetWmiInstanceManager(
                WmiInstanceFactory.GetWmiAssemblyName(WmiManager.DefaultAPI),
                WmiInstanceFactory.GetWmiInstanceManagerName(WmiManager.DefaultAPI),
                SystemName,
                Namespace,
                UserName,
                Password,
                Domain);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="assemblyName"></param>
        /// <param name="objectManagerName"></param>
        /// <param name="parameters"></param>
        /// <returns></returns>
        private static IWmiInstanceManager GetWmiInstanceManager(AssemblyName assemblyName,
            string objectManagerName, params string[] parameters)
        {
            // Below section reduces unnecessary loading of DLL
            IWmiInstanceManager vwom = null;
            string key = WmiInstanceFactory.GetWmiInstanceManagerKey(parameters);

            if (WmiInstanceFactory.LookupWmiInstanceManagerCollection(key, out vwom))
            {
                return vwom;
            }

            Assembly cimWmi = Assembly.Load(assemblyName);

            TypeInfo cimInstanceManager = cimWmi.GetType(objectManagerName).GetTypeInfo();

            foreach (ConstructorInfo constructor in cimInstanceManager.DeclaredConstructors)
            {
                ParameterInfo[] paramtersInfo = constructor.GetParameters();

                // Find the correct constructor
                if (paramtersInfo.Length == parameters.Length)
                {
                    vwom = constructor.Invoke(parameters) as IWmiInstanceManager;
                    AddWmiInstanceManagerCollection(key, vwom);
                    return vwom;
                }
            }

            throw new ArgumentException(string.Format(CultureInfo.InvariantCulture, "Failed to find {0} constructor with {1} arguments",
                objectManagerName, parameters.Length));
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="assemblyName"></param>
        /// <param name="objectManagerName"></param>
        /// <param name="serverName"></param>
        /// <param name="Namespace"></param>
        /// <param name="credentials"></param>
        /// <returns></returns>
        private static IWmiInstanceManager GetWmiInstanceManager(AssemblyName assemblyName,
            string objectManagerName, string serverName, string Namespace, WmiCredentials credentials)
        {
            // Below section reduces unnecessary loading of DLL
            IWmiInstanceManager vwom = null;

            string key = WmiInstanceFactory.GetWmiInstanceManagerKey(
                serverName, Namespace,
                credentials != null ? credentials.UserName : "",
                credentials != null ? credentials.Password : "",
                credentials != null ? credentials.Domain : ""
                );

            object[] parameters = new object[] { serverName, Namespace, credentials };
            if (WmiInstanceFactory.LookupWmiInstanceManagerCollection(key, out vwom))
            {
                return vwom;
            }

            Assembly cimWmi = Assembly.Load(assemblyName);

            TypeInfo cimInstanceManager = cimWmi.GetType(objectManagerName).GetTypeInfo();

            foreach (ConstructorInfo constructor in cimInstanceManager.DeclaredConstructors)
            {
                ParameterInfo[] paramtersInfo = constructor.GetParameters();

                // Find the correct constructor
                if (paramtersInfo.Length == parameters.Length)
                {
                    vwom = constructor.Invoke(parameters) as IWmiInstanceManager;
                    AddWmiInstanceManagerCollection(key, vwom);
                    return vwom;
                }
            }

            throw new ArgumentException(string.Format(CultureInfo.InvariantCulture, 
                "Failed to find {0} constructor with {1} arguments",
                objectManagerName, parameters.Length));
        }

        private static AssemblyName GetWmiAssemblyName(ClientAPIOptions api)
        {
            AssemblyName assemblyName = null;

            switch (api)
            {
                case ClientAPIOptions.SystemManagementAPI:
                    assemblyName = new AssemblyName(WmiInstanceFactory.LegacyWmiAssemblyName);
                    break;

                case ClientAPIOptions.ManagementInfrastructureAPI:
                default:
                    {
                        // Detect if FullSKU
                        if (File.Exists(Environment.ExpandEnvironmentVariables(@"%SystemRoot%\system32\kernel32.dll")))
                        {
                            assemblyName = new AssemblyName(WmiInstanceFactory.CimWmiAssemblyName);
                        }
                        else
                        {
                            assemblyName = new AssemblyName(WmiInstanceFactory.CssCimWmiAssemblyName);
                        }
                        break;
                    }
            }

            return assemblyName;
        }

        private static string GetWmiInstanceManagerName(ClientAPIOptions api)
        {
            switch (api)
            {
                case ClientAPIOptions.SystemManagementAPI:
                    return WmiInstanceFactory.WmiInstanceManagerName;

                case ClientAPIOptions.ManagementInfrastructureAPI:
                default:
                    return WmiInstanceFactory.WmiInstanceManagerName;
            }
        }
    }
}
