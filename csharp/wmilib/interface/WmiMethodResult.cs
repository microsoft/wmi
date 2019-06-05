// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Microsoft.WmiCodeGen.CSharp.Interface
{
    /// <summary>
    /// 
    /// </summary>
    public class WmiMethodResult : IDisposable
    {
        #region IDisposable Members
        private bool m_Disposed;
        /// <summary>
        /// 
        /// </summary>
        /// <param name="disposing"></param>
        protected virtual void Dispose(bool disposing)
        {
            if (disposing && !m_Disposed)
            {
                using (OutParameters) { }
                using (ReturnValue) { }

                m_Disposed = true;
            }
        }
        public void Dispose()
        {
            Dispose(true);
            GC.SuppressFinalize(this);
        }

        #endregion

        /// <summary>
        /// Gets the output.
        /// </summary>
        /// <value>
        /// The output.
        /// </value>
        public object Output
        {
            get
            {
                return this.ReturnValue.Value;
            }
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiMethodResult" /> class.
        /// </summary>
        public WmiMethodResult()
        {
            ReturnValue = new WmiMethodParameter("ReturnValue", (UInt32) 0);
            OutParameters = new WmiMethodParameterCollection();
        }

        public WmiMethodParameterCollection OutParameters;
        public WmiMethodParameter ReturnValue;
    }

}
