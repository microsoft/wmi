// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using Microsoft.Management.Infrastructure;
using Microsoft.Management.Infrastructure.Generic;
using Microsoft.Management.Infrastructure.Options;
using Microsoft.wmi.DotNet.Interface;

namespace Microsoft.wmi.DotNet.Lib
{
    public delegate void WmiObserverHandler(CimSubscriptionResult r);
    /// <summary>
    /// 
    /// </summary>
    class WmiObserver : IObserver<CimSubscriptionResult>
    {
        private WmiObserverHandler m_Handler = null;

        public WmiObserverHandler Handler
        {
            set
            {
                m_Handler = value;
            }
        }

        public void OnCompleted()
        {
        }

        public void OnError(Exception error)
        {
        }

        public void OnNext(CimSubscriptionResult result)
        {
            m_Handler(result);
        }

    }

    /// <summary>
    /// 
    /// </summary>
    class WmiEventWatcher : IWmiEventWatcher
    {
        /// <summary>
        /// Default timespan for WMI event query
        /// </summary>
        public TimeSpan m_Timeout = new TimeSpan(0, 0, 1);

        /// <summary>
        /// WMI namespace
        /// </summary>
        private String m_Namespace = null;

        /// <summary>
        /// WMI event query
        /// </summary>
        private String m_Query = null;

        /// <summary>
        /// WMI scope
        /// </summary>
        private CimSession m_CimSession = null;

        /// <summary>
        /// Operation options
        /// </summary>
        private CimOperationOptions m_Options = null;

        /// <summary>
        /// Object used to dispose of WMI results
        /// </summary>
        private IDisposable m_Dispose = null;

        /// <summary>
        /// WMI asynchronous results object
        /// </summary>
        private CimAsyncMultipleResults<CimSubscriptionResult> m_AsyncResults = null;

        /// <summary>
        /// WMI event observer
        /// </summary>
        private WmiObserver m_Observer = null;

        /// <summary>
        /// WMI event handler
        /// </summary>
        private WmiObserverHandler m_Handler = null;

        /// <summary>
        /// WMI event arrrived event
        /// </summary>
        public event WmiEventArrivedHandler EventArrived
        {
            add
            {
                this.m_Handler = (e) => 
                {
                    using (e)
                    {
                        value(e.MachineId, new WmiInstance(e.Instance)); 
                    }
                };
            }
            remove
            {
                throw new NotImplementedException("Cannot remove handler from EventArrived");
            }
        }

        #region Properties
        /// <summary>
        /// Timeout for WMI event query
        /// </summary>
        public TimeSpan Timeout
        {
            get
            {
                return m_Timeout;
            }
            set
            {
                m_Timeout = value;
            }
        }

        #endregion Properties
        /// <summary>
        /// 
        /// </summary>
        /// <param name="Session"></param>
        /// <param name="Namespace"></param>
        /// <param name="query"></param>
        public WmiEventWatcher(CimSession Session, String Namespace, String query)
        {
            if (Session == null)
            {
                throw new ArgumentNullException("Session");
            }

            if (query == null)
            {
                throw new ArgumentNullException("query");
            }

            m_CimSession = Session;
            m_Namespace = Namespace;
            m_Query = query;
        }

        /// <summary>
        /// Start watching for WMI events
        /// </summary>
        public void Start()
        {
            m_Options = new CimOperationOptions();
            m_Options.Timeout = this.m_Timeout;
            this.m_AsyncResults = m_CimSession.SubscribeAsync(this.m_Namespace, "WQL", this.m_Query, this.m_Options);
            this.m_Observer = new WmiObserver();
            this.m_Observer.Handler = this.m_Handler;

            this.m_Dispose = this.m_AsyncResults.Subscribe(this.m_Observer);
        }

        #region IDisposable Members
        private bool m_Disposed = false;
        void IDisposable.Dispose()
        {
            Dispose(true);
        }

        /// <summary>
        ///Clean up resources used by watcher
        /// </summary>
        protected void Dispose(bool disposing)
        {
            if (disposing && !m_Disposed)
            {
                using (this.m_Dispose) { }
                using (m_Options) { }

                m_Disposed = true;
            }
        }
        #endregion
    }
}
