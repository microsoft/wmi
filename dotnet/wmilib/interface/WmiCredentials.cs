// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using Microsoft.Management.Infrastructure.Options;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Security;
using System.Text;
using System.Threading.Tasks;

namespace Microsoft.WmiCodeGen.DotNet.Interface
{
    /// <summary>
    /// 
    /// </summary>
    public class WmiCredentials
    {
        /// <summary>
        /// 
        /// </summary>
        public string UserName { get; set; }
        /// <summary>
        /// 
        /// </summary>
        public string Password { get; set; }
        /// <summary>
        /// 
        /// </summary>
        public string Domain { get; set; }

        /// <summary>
        /// 
        /// </summary>
        public SecureString SecurePassword
        {
            get
            {
                if (string.IsNullOrEmpty(Password)) return null;

                SecureString securePassword = new SecureString();
                foreach (char passwordChar in Password.ToCharArray())
                {
                    securePassword.AppendChar(passwordChar);
                }
                return securePassword;
            }
        }
        /// <summary>
        /// 
        /// </summary>
        public PasswordAuthenticationMechanism AuthMechanism { get; set; }

        #region Constructor
        public WmiCredentials(string userName, string password, string domain, PasswordAuthenticationMechanism auth)
        {
            this.UserName = userName;
            this.Password = password;
            this.Domain = domain;
            this.AuthMechanism = auth;
        }

        /// <summary>
        /// 
        /// </summary>
        public WmiCredentials() : this(null, null, null, PasswordAuthenticationMechanism.Default) { }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="userName"></param>
        /// <param name="password"></param>
        public WmiCredentials(string userName, string password) : this(userName, password, null, PasswordAuthenticationMechanism.Default) { }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="userName"></param>
        /// <param name="password"></param>
        /// <param name="domain"></param>
        public WmiCredentials(string userName, string password, string domain) : this(userName, password, domain, PasswordAuthenticationMechanism.Default) { }

        #endregion
    }
}
