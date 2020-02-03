// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Microsoft.wmi.DotNet.Interface
{
    /// <summary>
    /// 
    /// </summary>
    public class WmiMethodParameter : IDisposable
    {
        #region IDisposable Members
        private bool m_Disposed;
        protected virtual void Dispose(bool disposing)
        {
            if (disposing && !m_Disposed)
            {
                if (Value != null)
                {
                    if (Value is IWmiInstance)
                    {
                        using (Value as IWmiInstance) { }
                    }
                    else if (Value is IWmiInstance[])
                    {
                        (Value as IWmiInstance[]).Dispose();
                    }
                }
                m_Disposed = true;
            }
        }
        public void Dispose()
        {
            Dispose(true);
            GC.SuppressFinalize(this);
        }

        #endregion

        public string Name;
        public object Value;
        public Type Type;
        //public int Flags; //Reserved

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiMethodParameter" /> class.
        /// </summary>
        /// <param name="Name">The name.</param>
        /// <param name="Value">The value.</param>
        public WmiMethodParameter(string Name, object Value)
        {
            this.Name = Name;
            this.Value = Value;
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="WmiMethodParameter" /> class.
        /// </summary>
        public WmiMethodParameter()
        {
        }
    }

}
