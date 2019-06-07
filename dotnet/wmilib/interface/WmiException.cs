// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;

namespace Microsoft.WmiCodeGen.DotNet.Interface
{
    /// <summary>
    /// Hyper-V exception
    /// </summary>
    public class WmiException : Exception /* derived from ManagementException for compatibility reasons */
    {
        #region Basic constructors

        /// <summary>
        /// Constructs a WmiException from a custom error message
        /// </summary>
        /// <param name="message">Custom error message</param>
        public WmiException(string message)
            : this(message, null)
        {
        }

        /// <summary>
        /// Constructs a WmiException from a custom error message and an inner exceptioon
        /// </summary>
        /// <param name="message"></param>
        /// <param name="innerException"></param>
        public WmiException(string message, Exception innerException)
            : base(message, innerException)
        {
        }

        #endregion

        #region Constructors from a WmiStatus code

        public WmiException(ushort statusCode)
            : this(statusCode, null, null)
        {
        }

        public WmiException(ushort statusCode, Exception innerException)
            : this(statusCode, null, innerException)
        {
        }

        public WmiException(ushort statusCode, string detailedMessage)
            : this(statusCode, detailedMessage, null)
        {
        }

        public WmiException(ushort statusCode, string detailedMessage, Exception innerException)
            : this(FormatErrorMessage(statusCode, detailedMessage), innerException)
        {
            StatusCode = statusCode;
            DetailedMessage = detailedMessage;
        }

        public WmiException(WmiManagementStatus errorCode)
            : this(errorCode, null, null)
        {
        }

        public WmiException(WmiManagementStatus errorCode, Exception innerException)
            : this(errorCode, null, innerException)
        {
        }

        public WmiException(WmiManagementStatus errorCode, string detailedMessage)
            : this(errorCode, detailedMessage, null)
        {
        }

        public WmiException(WmiManagementStatus errorCode, string detailedMessage, Exception innerException)
            : this(FormatErrorMessage(errorCode, detailedMessage), innerException)
        {
            StatusCode = (ushort)WmiErrorCode.ManagementError; // FIXME
            ErrorCode = errorCode;
            DetailedMessage = detailedMessage;
        }

        #endregion

        #region Public fields

        /// <summary>
        /// Status code returned by a Hyper-V function
        /// </summary>
        public ushort StatusCode
        {
            get;
            private set;
        }

        /// <summary>
        /// Management Status code returned by a Hyper-V function
        /// </summary>
        public WmiManagementStatus ErrorCode
        {
            get;
            private set;
        }

        /// <summary>
        /// Detailed error message from Hyper-V
        /// </summary>
        public string DetailedMessage
        {
            get;
            private set;
        }

        #endregion

        #region Helper functions

        /// <summary>
        /// Formats an error message from a Hyper-V status code
        /// </summary>
        /// <param name="statusCode">Hyper-V status code</param>
        /// <param name="detailedMessage">Optional detailed message</param>
        /// <returns></returns>
        private static string FormatErrorMessage(ushort statusCode, string detailedMessage)
        {
            string pattern = detailedMessage == null
                ? "{0}: Wmi operation failed (error code: {1})"
                : "{0}: Wmi operation failed (error code: {1}, details: \"{2}\")";

            return string.Format(CultureInfo.InvariantCulture,
                pattern,
                statusCode,
                (ushort)statusCode,
                detailedMessage);
        }

        /// <summary>
        /// Formats an error message from a Hyper-V status code
        /// </summary>
        /// <param name="statusCode">Hyper-V status code</param>
        /// <param name="detailedMessage">Optional detailed message</param>
        /// <returns></returns>
        private static string FormatErrorMessage(WmiManagementStatus statusCode, string detailedMessage)
        {
            string pattern = detailedMessage == null
                ? "{0}: Wmi operation failed (Management error code: {1})"
                : "{0}: Wmi operation failed (Management error code: {1}, details: \"{2}\")";

            return string.Format(CultureInfo.InvariantCulture,
                pattern,
                statusCode,
                (ushort)statusCode,
                detailedMessage);
        }

        #endregion
    }

    public enum WmiManagementStatus
    {
        // These are all HRESULTS, not CIM status codes. 
        HrFailed = -2147217407,
        HrNotFound = -2147217406,
        HrAccessDenied = -2147217405,
        HrProviderFailure = -2147217404,
        HrTypeMismatch = -2147217403,
        HrOutOfMemory = -2147217402,
        HrInvalidContext = -2147217401,
        HrInvalidParameter = -2147217400,
        HrNotAvailable = -2147217399,
        HrCriticalError = -2147217398,
        HrInvalidStream = -2147217397,
        HrNotSupported = -2147217396,
        HrInvalidSuperclass = -2147217395,
        HrInvalidNamespace = -2147217394,
        HrInvalidObject = -2147217393,
        HrInvalidClass = -2147217392,
        HrProviderNotFound = -2147217391,
        HrInvalidProviderRegistration = -2147217390,
        HrProviderLoadFailure = -2147217389,
        HrInitializationFailure = -2147217388,
        HrTransportFailure = -2147217387,
        HrInvalidOperation = -2147217386,
        HrInvalidQuery = -2147217385,
        HrInvalidQueryType = -2147217384,
        HrAlreadyExists = -2147217383,
        HrOverrideNotAllowed = -2147217382,
        HrPropagatedQualifier = -2147217381,
        HrPropagatedProperty = -2147217380,
        HrUnexpected = -2147217379,
        HrIllegalOperation = -2147217378,
        HrCannotBeKey = -2147217377,
        HrIncompleteClass = -2147217376,
        HrInvalidSyntax = -2147217375,
        HrNondecoratedObject = -2147217374,
        HrReadOnly = -2147217373,
        HrProviderNotCapable = -2147217372,
        HrClassHasChildren = -2147217371,
        HrClassHasInstances = -2147217370,
        HrQueryNotImplemented = -2147217369,
        HrIllegalNull = -2147217368,
        HrInvalidQualifierType = -2147217367,
        HrInvalidPropertyType = -2147217366,
        HrValueOutOfRange = -2147217365,
        HrCannotBeSingleton = -2147217364,
        HrInvalidCimType = -2147217363,
        HrInvalidMethod = -2147217362,
        HrInvalidMethodParameters = -2147217361,
        HrSystemProperty = -2147217360,
        HrInvalidProperty = -2147217359,
        HrCallCanceled = -2147217358,
        HrShuttingDown = -2147217357,
        HrPropagatedMethod = -2147217356,
        HrUnsupportedParameter = -2147217355,
        HrMissingParameterID = -2147217354,
        HrInvalidParameterID = -2147217353,
        HrNonconsecutiveParameterIDs = -2147217352,
        HrParameterIDOnRetval = -2147217351,
        HrInvalidObjectPath = -2147217350,
        HrOutOfDiskSpace = -2147217349,
        HrBufferTooSmall = -2147217348,
        HrUnsupportedPutExtension = -2147217347,
        HrUnknownObjectType = -2147217346,
        HrUnknownPacketType = -2147217345,
        HrMarshalVersionMismatch = -2147217344,
        HrMarshalInvalidSignature = -2147217343,
        HrInvalidQualifier = -2147217342,
        HrInvalidDuplicateParameter = -2147217341,
        HrTooMuchData = -2147217340,
        HrServerTooBusy = -2147217339,
        HrInvalidFlavor = -2147217338,
        HrCircularReference = -2147217337,
        HrUnsupportedClassUpdate = -2147217336,
        HrCannotChangeKeyInheritance = -2147217335,
        HrCannotChangeIndexInheritance = -2147217328,
        HrTooManyProperties = -2147217327,
        HrUpdateTypeMismatch = -2147217326,
        HrUpdateOverrideNotAllowed = -2147217325,
        HrUpdatePropagatedMethod = -2147217324,
        HrMethodNotImplemented = -2147217323,
        HrMethodDisabled = -2147217322,
        HrRefresherBusy = -2147217321,
        HrUnparsableQuery = -2147217320,
        HrNotEventClass = -2147217319,
        HrMissingGroupWithin = -2147217318,
        HrMissingAggregationList = -2147217317,
        HrPropertyNotAnObject = -2147217316,
        HrAggregatingByObject = -2147217315,
        HrUninterpretableProviderQuery = -2147217313,
        HrBackupRestoreWinmgmtRunning = -2147217312,
        HrQueueOverflow = -2147217311,
        HrPrivilegeNotHeld = -2147217310,
        HrInvalidOperator = -2147217309,
        HrLocalCredentials = -2147217308,
        HrCannotBeAbstract = -2147217307,
        HrAmendedObject = -2147217306,
        HrClientTooSlow = -2147217305,
        HrRegistrationTooBroad = -2147213311,
        HrRegistrationTooPrecise = -2147213310,


        // Not sure where these came from 
        UnrecognizedError = -1,
        NoError = 0,
        False = 1,
        ResetToDefault = 262146,
        Different = 262147,
        Timedout = 262148,
        NoMoreData = 262149,
        OperationCanceled = 262150,
        Pending = 262151,
        DuplicateObjects = 262152,
        PartialResults = 262160,

        QueryLanguageNotSupported = 14,
        NamespaceNotEmpty = 20,
        InvalidEnumerationContext = 21,
        InvalidOperationTimeout = 22,
        PullHasBeenAbandoned = 23,
        PullCannotBeAbandoned = 24,
        FilteredEnumerationNotSupported = 25,
        ContinuationOnErrorNotSupported = 26,
        ServerLimitsExceeded = 27,
        ServerIsShuttingDown = 28,

        // From vm\mgmtv2\msvm\Msvm_ConcreteJob.mof
        //   See CIM_ConcreteJob.GetError
        Failed = 32768,
        AccessDenied = 32769,
        NotSupported = 32770,
        StatusIsUnknown = 32771,
        Timeout = 32772,
        InvalidParameter = 32773,
        SystemIsInUsed = 32774,
        InvalidStateForThisOperation = 32775,
        IncorrectDataType = 32776,
        SystemIsNotAvailable = 32777,
        OutOfMemory = 32778,
        NotFound = 32779 // WINBLUE: 129827 - Defined in Msvm_ComputerSystem.mof, but not Msvm_VirtualSystemManagementService.mof
    }
}
