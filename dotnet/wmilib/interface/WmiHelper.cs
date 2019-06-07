// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Text;
using System.Threading;
using System.Globalization;

namespace Microsoft.WmiCodeGen.DotNet.Interface.Helper
{
    /// <summary>
    /// Wrapper for WMI Helper
    /// </summary>
    public class WmiHelper
    {
        private string m_targetMachine;
        private IWmiInstanceManager m_objectManager;
        private IWmiInstanceManager m_cimInstanceManager;
        private string m_VirtualizationNamespace = WMINamespaces.DefaultVMNamespace;

        public const String QuickVmStaticSnapshotPrefix = "QuickVmSnapshot";

        private void ValidateParameters(string targetMachine, string vmNamespace)
        {
            if (string.IsNullOrEmpty(targetMachine) || string.IsNullOrEmpty(vmNamespace))
            {
                throw new Exception("Either TargetMachine or VMNamespace is Null or Empty string...");
            }
        }

        private void Initialize()
        {
            ValidateParameters(m_targetMachine, m_VirtualizationNamespace);
            m_objectManager = WmiInstanceFactory.GetInstanceManager(m_targetMachine, m_VirtualizationNamespace);
            m_cimInstanceManager = WmiInstanceFactory.GetInstanceManager(m_targetMachine, "root\\cimv2");
        }

        private void Initialize(string userName, string passwd, string domain)
        {
            ValidateParameters(m_targetMachine, m_VirtualizationNamespace);

            m_objectManager = WmiInstanceFactory.GetInstanceManager(m_targetMachine, m_VirtualizationNamespace, userName, passwd, domain);
            m_cimInstanceManager = WmiInstanceFactory.GetInstanceManager(m_targetMachine, "root\\cimv2");
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiHelper"/> class.
        /// </summary>
        /// <param name="targetMachine">The target machine.</param>
        /// <param name="vmNamespace">The vm namespace.</param>
        public WmiHelper(string targetMachine, string vmNamespace)
        {
            m_targetMachine = targetMachine;
            m_VirtualizationNamespace = vmNamespace;
            Initialize();
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiHelper"/> class.
        /// </summary>
        /// <param name="targetMachine">The target machine.</param>
        public WmiHelper(string targetMachine)
        {
            m_targetMachine = targetMachine;
            Initialize();
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiHelper"/> class.
        /// </summary>
        /// <param name="targetMachine">The target machine.</param>
        /// <param name="vmNamespace">The vm namespace.</param>
        /// <param name="userName">Name of the user.</param>
        /// <param name="passwd">The passwd.</param>
        /// <param name="domain">The domain.</param>
        public WmiHelper(string targetMachine, string vmNamespace, string userName, string passwd, string domain)
        {
            m_targetMachine = targetMachine;
            m_VirtualizationNamespace = vmNamespace;
            Initialize(userName, passwd, domain);
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiHelper"/> class.
        /// </summary>
        /// <param name="targetMachine">The target machine.</param>
        /// <param name="userName">Name of the user.</param>
        /// <param name="passwd">The passwd.</param>
        /// <param name="domain">The domain.</param>
        public WmiHelper(string targetMachine, string userName, string passwd, string domain)
        {
            m_targetMachine = targetMachine;
            Initialize(userName, passwd, domain);
        }

        /// <summary>
        /// Returns ManagementScope object for the CIMv2 namespace
        /// </summary>
        /// <value>The cim scope.</value>
        public IWmiInstanceManager CimScope
        {
            get
            {
                return m_cimInstanceManager;
            }
        }

        /// <summary>
        /// Gets the target machine.
        /// </summary>
        /// <value>The target machine.</value>
        public string TargetMachine
        {
            get
            {
                return m_targetMachine;
            }
        }

        
        /// <summary>
        /// Retrieves a property from a single-instance WMI object (e.g. VMMS, etc).
        /// </summary>
        /// <param name="objectName">Name of the object instance to retrieve property from</param>
        /// <param name="propertyName">Name of the property to retrieve value from</param>
        /// <returns>Property value</returns>
        public string GetObjectProperty(string objectName, string propertyName)
        {
            /*
            IWmiInstanceSearcher searcher = new IWmiInstanceSearcher(
                m_cimInstanceManager,
                new ObjectQuery(string.Format("SELECT * FROM {0}", objectName)));
            */
            WmiInstanceCollection collection = m_cimInstanceManager.QueryInstances(
                string.Format(CultureInfo.InvariantCulture,
                "SELECT * FROM {0}", objectName));
            if (collection.Count == 1)
            {
                foreach (IWmiInstance obj in collection)
                {
                    return obj[propertyName].ToString();
                }
                return string.Empty;
            }
            else
            {
                return string.Empty;
            }
        }

        /// <summary>
        /// Provides a clean way to grab a full instance of a WMI object, optionally providing
        /// for selection by a WHERE clause if it's a multi-instance object.
        /// </summary>
        /// <param name="objectName">Object class to search</param>
        /// <param name="scope">Management scope to search</param>
        /// <param name="criteria">Optional (set to string.Empty if not desired)</param>
        /// <returns>IWmiInstance</returns>
        public IWmiInstance GetObjectInstance(string objectName, IWmiInstanceManager cimInstanceManager, string criteria)
        {
            return this.GetObjectInstances(objectName, cimInstanceManager, criteria).GetFirstInstance();
        }

        /// <summary>
        /// Provides a clean way to grab a full instance of WMI objects, optionally providing
        /// for selection by a WHERE clause if it's a multi-instance object.
        /// </summary>
        /// <param name="objectName">Object class to search</param>
        /// <param name="scope">Management scope to search</param>
        /// <param name="criteria">Optional (set to string.Empty if not desired)</param>
        /// <returns>IWmiInstance</returns>
        public WmiInstanceCollection GetObjectInstances(string objectName, IWmiInstanceManager cimInstanceManager, string criteria)
        {
            if (cimInstanceManager == null)
            {
                throw new ArgumentNullException("cimInstanceManager");
            }

            String query = string.Format(CultureInfo.InvariantCulture, "SELECT * FROM {0}", objectName);
            if (criteria != string.Empty)
            {
                query += " WHERE " + criteria;
            }
            WmiInstanceCollection collection = cimInstanceManager.QueryInstances(query);
            return collection;
        }
    }
}
