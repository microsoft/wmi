// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.



using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Management.Infrastructure;
using Microsoft.Management.Infrastructure.Options;
using System.Security;
using System.Collections.Concurrent;
using Microsoft.wmi.DotNet.Interface;

namespace Microsoft.wmi.DotNet.Lib
{
    /// <summary>
    /// 
    /// </summary>
    class SessionInfo
    {
        public String ServerName;
        public String UserName;
        public String Password;
        public String Domain;
        public CimSession cimSession;

        /// <summary>
        /// Initializes a new instance of the <see cref="SessionInfo" /> class.
        /// </summary>
        /// <param name="ServerName">Name of the server.</param>
        /// <param name="UserName">Name of the user.</param>
        /// <param name="Password">The password.</param>
        /// <param name="Domain">The domain.</param>
        public SessionInfo(String ServerName, String UserName, String Password, String Domain, CimSession cimSession)
        {
            this.ServerName = ServerName;
            this.UserName = UserName;
            this.Password = Password;
            this.Domain = Domain;
            this.cimSession = cimSession;
        }
    }

    /// <summary>
    /// 
    /// </summary>
    internal class CimSessionManager
    {
        /// <summary>
        ///  This class cannot be instantiated
        /// </summary>
        private CimSessionManager()
        {

        }
        static ConcurrentDictionary<string, SessionInfo> m_sessions = new ConcurrentDictionary<string, SessionInfo>();
        /// <summary>
        /// 
        /// </summary>
        /// <param name="ServerName"></param>
        /// <param name="credentials"></param>
        /// <returns></returns>
        public static String GetSessionKey(String ServerName, WmiCredentials credentials)
        {
            return GetSessionKey(ServerName,
                credentials != null ? credentials.UserName : "",
                credentials != null ? credentials.Password : "",
                credentials != null ? credentials.Domain : ""
                );
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="ServerName"></param>
        /// <param name="UserName"></param>
        /// <param name="Password"></param>
        /// <param name="Domain"></param>
        /// <returns></returns>
        public static String GetSessionKey(String ServerName, String UserName, String Password, String Domain)
        {
            String sessionKey = String.Empty;
            string serverNameToConnect = ServerName;
            if (string.IsNullOrEmpty(ServerName) ||
                ServerName.ToLowerInvariant().Contains(Environment.MachineName.ToLowerInvariant()) ||
                ServerName.Equals("localhost", StringComparison.OrdinalIgnoreCase) ||
                ServerName.Equals(".", StringComparison.OrdinalIgnoreCase))
            {
                serverNameToConnect = Environment.MachineName;
            }

            if (serverNameToConnect != null)
            {
                // Workaround for IPv6 address
                sessionKey += serverNameToConnect.Replace("[", "").Replace("]", "");
            }
            if (UserName != null)
            {
                sessionKey += UserName;
            }
            if (Password != null)
            {
                sessionKey += Password;
            }
            if (Domain != null)
            {
                sessionKey += Domain;
            }
            return sessionKey.ToLowerInvariant();
        }

        /// <summary>
        /// Gets the session.
        /// </summary>
        /// <param name="ServerName">Name of the server.</param>
        /// <returns></returns>
        public static CimSession GetSession(String ServerName)
        {
            return GetSession(ServerName, null);
        }

        /// <summary>
        /// Gets the session. Even if partial data like, ServerName is only passed,
        /// try to retrieve the existing session for this server
        /// </summary>
        /// <param name="ServerName">Name of the server.</param>
        /// <param name="UserName">Name of the user.</param>
        /// <param name="Password">The password.</param>
        /// <param name="Domain">The domain.</param>
        /// <returns></returns>
        public static CimSession GetSession(String ServerName, String UserName, String Password, String Domain)
        {
            string sessionKey = GetSessionKey(ServerName, UserName, Password, Domain);
            foreach (var item in m_sessions.Keys)
            {
                if (item.Contains(sessionKey))
                {
                    return m_sessions[item].cimSession;
                }
            }

            return null;
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="ServerName"></param>
        /// <param name="credentials"></param>
        /// <returns></returns>
        public static CimSession GetSession(String ServerName, WmiCredentials credentials)
        {
            return GetSession(ServerName,
                credentials != null ? credentials.UserName : "",
                credentials != null ? credentials.Password : "",
                credentials != null ? credentials.Domain : ""
                );
        }
        /// <summary>
        /// Adds the session.
        /// </summary>
        /// <param name="ServerName">Name of the server.</param>
        /// <param name="UserName">Name of the user.</param>
        /// <param name="Password">The password.</param>
        /// <param name="Domain">The domain.</param>
        /// <param name="cimSession">The cim session.</param>
        public static void AddSession(String ServerName, String UserName, String Password, String Domain, CimSession cimSession)
        {
            m_sessions[GetSessionKey(ServerName, UserName, Password, Domain)] = new SessionInfo(ServerName, UserName, Password, Domain, cimSession);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="ServerName"></param>
        /// <param name="credentials"></param>
        /// <param name="cimSession"></param>
        public static void AddSession(String ServerName, WmiCredentials credentials, CimSession cimSession)
        {
            AddSession(ServerName,
                credentials != null ? credentials.UserName : "",
                credentials != null ? credentials.Password : "",
                credentials != null ? credentials.Domain : "",
                cimSession
                );
        }
#if false
        /// <summary>
        /// Adds the session.
        /// </summary>
        /// <param name="ServerName">Name of the server.</param>
        /// <param name="cimSession">The cim session.</param>
        public static void AddSession(String ServerName, CimSession cimSession)
        {
            AddSession(ServerName, null, null, null, cimSession);
        } 
#endif

    }

    /// <summary>
    /// 
    /// </summary>
    public class WmiSession : IWmiSession
    {
        #region Data

        private CimSession m_CimSession;

        #endregion Data

        #region Construction

        /// <summary>
        /// Gets the session.
        /// </summary>
        /// <value>
        /// The session.
        /// </value>
        public CimSession Session
        {
            get
            {
                return m_CimSession;
            }
        }

        #endregion Construction

        #region Methods
        /// <summary>
        /// 
        /// </summary>
        /// <param name="ServerName"></param>
        /// <param name="credentials"></param>
        public void CreateSession(String ServerName, WmiCredentials credentials)
        {
            m_CimSession = CimSessionManager.GetSession(ServerName, credentials);
            if (m_CimSession == null)
            {
                if (credentials != null)
                {
                    using (SecureString securePassword = credentials.SecurePassword)
                    {
                        CimCredential cimcredentials = new CimCredential(credentials.AuthMechanism,
                            credentials.Domain, credentials.UserName, securePassword);
                        using (CimSessionOptions options = new CimSessionOptions())
                        {
                            options.AddDestinationCredentials(cimcredentials);
                            m_CimSession = CimSession.Create(ServerName, options);
                        }
                    }
                }
                else
                {
                    m_CimSession = CimSession.Create(ServerName);
                }

                if (m_CimSession != null)
                {
                    CimSessionManager.AddSession(ServerName, credentials, m_CimSession);
                }
                else
                {
                    throw new WmiException("Failed to create CimSession");
                }
            }
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="ServerName"></param>
        /// <param name="UserName"></param>
        /// <param name="Password"></param>
        /// <param name="Domain"></param>
        public void CreateSession(String ServerName, String UserName, String Password, String Domain)
        {
            CreateSession(ServerName, new WmiCredentials(UserName, Password, Domain));
        }


        /// <summary>
        /// Connects the specified server name.
        /// </summary>
        /// <param name="ServerName">Name of the server.</param>
        /// <returns></returns>
        public bool Connect(String ServerName)
        {
            return Connect(ServerName, null, null);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="ServerName"></param>
        /// <param name="Namespace"></param>
        /// <param name="credentials"></param>
        /// <returns></returns>
        public bool Connect(String ServerName, String Namespace,
            WmiCredentials credentials)
        {
            //
            // There is a bug (BLUE: *TODO*) in MI.Net where if you pass the machine
            // name and that is the same machine you are executing on it will still
            // proxy the calls through WinRM. This workaround will make it use the local
            // DCOM calls instead.
            //
            // Note: Ordered by most likely comparison for early evaluation exit.
            //
            string serverNameToConnect = ServerName;
            if (string.IsNullOrEmpty(ServerName) ||
                ServerName.ToLowerInvariant().Contains(Environment.MachineName.ToLowerInvariant()) ||
                ServerName.Equals("localhost", StringComparison.OrdinalIgnoreCase) ||
                ServerName.Equals(".", StringComparison.OrdinalIgnoreCase))
            {
                serverNameToConnect = Environment.MachineName;
            }

            m_CimSession = CimSessionManager.GetSession(serverNameToConnect, credentials);
            if (m_CimSession == null)
            {
                CreateSession(serverNameToConnect, credentials);
            }
            return true;

        }

        /// <summary>
        /// Connects the specified server name.
        /// </summary>
        /// <param name="ServerName">Name of the server.</param>
        /// <param name="UserName">Name of the user.</param>
        /// <param name="Password">The password.</param>
        /// <param name="Domain">The domain.</param>
        /// <returns></returns>
        public bool Connect(String ServerName, String Namespace,
            String UserName, String Password, String Domain)
        {
            return Connect(ServerName, Namespace, new WmiCredentials(UserName, Password, Domain));
        }

        #endregion Methods

        #region IDisposable Members
        private bool m_Disposed;
        protected virtual void Dispose(bool disposing)
        {
            if (disposing && !m_Disposed)
            {
                using (m_CimSession) { }
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
}
