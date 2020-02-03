// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;

namespace Microsoft.wmi.DotNet.Interface
{
    public interface IWmiEventWatcher : IDisposable
    {
        #region Properties

        /// <summary>
        /// Timeout for WMI event query
        /// </summary>
        TimeSpan Timeout
        {
            get;
            set;
        }

        #endregion Properties

        /// <summary>
        /// WMI event arrrived event
        /// </summary>
        event WmiEventArrivedHandler EventArrived;

        /// <summary>
        /// Start watching for WMI events
        /// </summary>
        void Start();
    }

    /// <summary>
    /// Delegate for expected WMI events
    /// </summary>
    /// <param name="sender">Sender of event</param>
    /// <param name="e">Event arguments</param>
    public delegate void WmiEventArrivedHandler(object sender, IWmiInstance e);

    /// <summary>
    /// Delegate for unexpected WMI events
    /// </summary>
    /// <param name="sender">Sender of event</param>
    /// <param name="e">Event arguments</param>
    public delegate void WmiUnexpectedEventArrivedHandler(object sender, IWmiInstance e);
}
