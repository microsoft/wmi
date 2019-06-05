// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Xml.Linq;
using System.Linq;
using System.IO;

namespace Microsoft.WmiCodeGen.CSharp.Interface
{
    /// <summary>
    /// Hyper-V Error codes
    /// </summary>
    public enum WmiErrorCode
    {
        Success = 0,
        ManagementError = 1,
        JobStarted = 4096,
        Failed = 32768,
        AccessDenied = 32769,
        NotSupported = 32770,
        Unknown = 32771,
        TimeOut = 32772,
        InvalidArgs = 32773,
        InUse = 32774,
        InvalidState = 32775,
        IncorrectType = 32776,
        Unavailable = 32777,
        OutOfMemory = 32778,
        FileNotFound = 32779,
        NotReady = 32780,
        InsufficientResources = 32788,
        ObjectNotFound = 32789,
        ObjectExists = 32790,
    }

    public interface IWmiError
    {
        string ErrorSource { get; }
        string Message { get; }
    }

    /// <summary>
    /// Wrapper for Msvm_Error.
    /// </summary>
    public class WmiError : IWmiError
    {
        XDocument error;

        /// <summary>
        /// 
        /// </summary>
        public WmiError(
            string Error_Object)
        {
            if (string.IsNullOrEmpty(Error_Object))
            {
                throw new ArgumentNullException("Error_Object");
            }

            error = XDocument.Parse(Error_Object);
        }

        #region Properties

        /// <summary>
        /// Gets the path to the management object that logged the error.
        /// </summary>
        public string
        ErrorSource
        {
            get
            {
                //string xpath_expression = "//PROPERTY[@NAME = \"ErrorSource\"]";
                XElement node = error.Descendants("PROPERTY")
                    .Where(e => String.Equals((string)(e.Attribute("NAME")), "ErrorSource", StringComparison.OrdinalIgnoreCase))
                    .FirstOrDefault();
                if (node == null) return null;

                // This node will have only one child, the "Value" node
                XElement value = node.Element("VALUE");
                if (value == null) return null;

                return value.Value;
            }
        }

        /// <summary>
        /// Gets the error message.
        /// </summary>
        public string
        Message
        {
            get
            {
                //string xpath_expression = "//PROPERTY[@NAME = \"Message\"]";
                XElement node = error.Descendants("PROPERTY")
                    .Where(e => String.Equals((string)(e.Attribute("NAME")), "Message", StringComparison.OrdinalIgnoreCase))
                    .FirstOrDefault();
                if (node == null) return null;

                // This node will have only one child, the "Value" node
                XElement value = node.Element("VALUE");
                if (value == null) return null;

                return value.Value;
            }
        }

        #endregion
    }

    /// <summary>
    /// VMError collection class.
    /// Reimplement the whole collection class to keep strong type.
    /// </summary>
    public class WmiErrorCollection : Collection<IWmiError>
    {
        public WmiErrorCollection(string[] Errors)
        {
            if (Errors == null)
                return;
            foreach (string error in Errors)
            {
                Add(new WmiError(error));
            }
        }
    }
}
