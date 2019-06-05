// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Microsoft.WmiCodeGen.CSharp.Interface
{
    #region Exception
    /// <summary>
    /// 
    /// </summary>
    public class WmiSessionException : Exception
    {
        #region Basic constructors

        /// <summary>
        /// Constructs a WmiException from a custom error message
        /// </summary>
        /// <param name="message">Custom error message</param>
        public WmiSessionException(string message)
            : this(message, null)
        {
        }

        /// <summary>
        /// Constructs a WmiException from a custom error message and an inner exceptioon
        /// </summary>
        /// <param name="message"></param>
        /// <param name="innerException"></param>
        public WmiSessionException(string message, Exception innerException)
            : base(message, innerException)
        {
        }

        #endregion
    }
    #endregion

    /// <summary>
    /// 
    /// </summary>
    public interface IWmiSession : IDisposable
    {
        bool Connect(String ServerName, string Namespace, String UserName, String Password, String Domain);
        bool Connect(String ServerName, string Namespace, WmiCredentials credentials);
    }
}
