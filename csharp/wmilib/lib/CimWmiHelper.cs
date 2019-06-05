// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using Microsoft.Management.Infrastructure;
using System.Collections.ObjectModel;
using Microsoft.Management.Infrastructure.Serialization;
using System.Globalization;
using Microsoft.WmiCodeGen.CSharp.Interface;

namespace Microsoft.WmiCodeGen.CSharp.Lib
{
    public class WmiHelper
    {
        #region PathToInstance Helpers

        /// <summary>
        /// Gets the cim instance from path.
        /// Eg: __PATH : \\MADHANM-T2\root\virtualization\v2:Msvm_ComputerSystem.CreationClassName="Msvm_ComputerSystem",Name="26B0C45F-B797-4BEE-B2F5-F51A13894B7B"
        /// </summary>
        /// <param name="path">The path.</param>
        /// <returns></returns>
        public static IWmiInstance GetInstanceFromPath(string path)
        {
            int indexOfKVP = path.IndexOf(".", StringComparison.OrdinalIgnoreCase) + 1;

            string serverName = null;
            string nameSpace = null;
            string className = null;

            string serverAndNameSpace = path.Substring(0, indexOfKVP - 1).Replace("/", "\\");

            int indexOfServer = serverAndNameSpace.IndexOf("\\\\", StringComparison.OrdinalIgnoreCase) + 2;
            int indexOfNameSpace = serverAndNameSpace.IndexOf("\\", indexOfServer, StringComparison.OrdinalIgnoreCase) + 1;
            int indexOfClass = serverAndNameSpace.IndexOf(":", StringComparison.OrdinalIgnoreCase) + 1;

            serverName = serverAndNameSpace.Substring(indexOfServer, indexOfNameSpace - indexOfServer - 1);
            nameSpace = serverAndNameSpace.Substring(indexOfNameSpace, indexOfClass - indexOfNameSpace - 1);
            className = serverAndNameSpace.Substring(indexOfClass);

            //Get KVP Values
            WmiKeyValueCollection kvpCollection = new WmiKeyValueCollection();
            string kvpString = path.Substring(indexOfKVP);
            string delimiter = ",";
            string[] kvpValues = kvpString.Split(delimiter.ToCharArray());
            foreach (string kvp in kvpValues)
            {
                string kvpDelimiter = "=";
                string[] kvpKeyValuePair = kvp.Split(kvpDelimiter.ToCharArray());

                string propertyValue = kvpKeyValuePair[1].Replace("\\\\", "\\");
                propertyValue = propertyValue.TrimStart("\"".ToCharArray());
                propertyValue = propertyValue.TrimEnd("\"".ToCharArray());
                kvpCollection.Add(kvpKeyValuePair[0], propertyValue);

            }

            //Get instance using KVP values
            IWmiInstanceManager objectManager = WmiInstanceFactory.GetInstanceManager(serverName, nameSpace);

            return objectManager.GetInstance(className, kvpCollection);
        }


        /// <summary>
        /// Gets the instance array from path array.
        /// </summary>
        /// <param name="paths">The paths.</param>
        /// <returns></returns>
        public static IWmiInstance[] GetInstanceArrayFromPathArray(string[] paths)
        {
            if (paths == null)
            {
                return null;
            }

            if (paths.Length == 0)
            {
                return new IWmiInstance[] { };
            }

            IWmiInstance[] instanceArray = new IWmiInstance[paths.Length];
            for (int index = 0; index < paths.Length; index++)
            {
                instanceArray[index] = GetInstanceFromPath(paths[index]);
            }
            return instanceArray;
        }

        #endregion PathToInstance Helpers

        #region InstanceToPath Helpers
        /// <summary>
        /// Gets the path value.
        /// </summary>
        /// <param name="Value">The value.</param>
        /// <returns></returns>
        public static string GetPathFromInstance(IWmiInstance Value)
        {
            string path = String.Empty;
            CimInstance instance = Value.GetInstance() as CimInstance;

            path += @"\\" + instance.CimSystemProperties.ServerName + @"\" + instance.CimSystemProperties.Namespace.Replace(@"/", @"\") + ":" + instance.CimSystemProperties.ClassName;

            int index = 0;
            foreach (CimProperty property in instance.CimInstanceProperties)
            {
                if ((property.Flags & CimFlags.Key) != 0)
                {
                    if (index != 0)
                    {
                        path += ",";
                    }
                    else
                    {
                        path += ".";
                    }

                    path += property.Name + "=" + "\"" + property.Value.ToString().Replace("\\", "\\\\") + "\"";
                    index++;
                }
            }
            return path;
        }

        #endregion InstanceToPath Helpers

        #region GetCimInstance Helpers

        /// <summary>
        /// Gets the reference value array.
        /// </summary>
        /// <param name="inputValue">The input value.</param>
        /// <returns></returns>
        public static CimInstance[] GetCimInstanceArray(IWmiInstance[] WmiInstancearray)
        {
            if (WmiInstancearray == null)
            {
                return null;
            }

            if (WmiInstancearray.Length == 0)
            {
                return new CimInstance[] { };
            }

            CimInstance[] instanceArray = new CimInstance[WmiInstancearray.Length];
            for (int index = 0; index < WmiInstancearray.Length; index++)
            {
                instanceArray[index] = WmiInstancearray[index].GetInstance() as CimInstance;
            }
            return instanceArray;
        }

        /// <summary>
        /// Gets the reference value.
        /// </summary>
        /// <param name="inputValue">The input value.</param>
        /// <returns></returns>
        public static CimInstance GetCimInstance(IWmiInstance inputValue)
        {
            CimInstance instance = (inputValue as IWmiInstance).GetInstance() as CimInstance;
            return instance;
        }

        #endregion GetCimInstance Helpers

        #region EmbeddedInstance Helpers

        /// <summary>
        /// Gets the embedded instance.
        /// </summary>
        /// <param name="inputValue">The input value.</param>
        /// <returns></returns>
        public static string GetEmbeddedInstance(IWmiInstance inputValue)
        {
            if (inputValue == null)
            {
                return null;
            }

            CimInstance instance = inputValue.GetInstance() as CimInstance;
            return WmiHelper.GetSerializedXMLFromCimInstance(instance);
        }


        /// <summary>
        /// Gets the CIMXML from cim instance.
        /// </summary>
        /// <param name="instance">The instance.</param>
        /// <returns></returns>
        public static String GetSerializedXMLFromCimInstance(CimInstance instance)
        {
            Byte[] buffer = null;
            using (CimSerializer cimSerializer = CimSerializer.Create())
            {
                buffer = cimSerializer.Serialize(instance, InstanceSerializationOptions.None);
            }
            return Encoding.Unicode.GetString(buffer, 0, buffer.Length);
        }

        #endregion EmbeddedInstance Helpers

        #region PropertyFlag Helpers
        /// <summary>
        /// Gets the vm property flags from cim property.
        /// </summary>
        /// <param name="property">The property.</param>
        /// <returns></returns>
        public static WmiPropertyFlags GetWmiPropertyFlagsFromCimProperty(CimProperty property)
        {
            return GetWmiPropertyFlags(property.Flags);
        }

        /// <summary>
        /// Gets the vm property flags.
        /// </summary>
        /// <param name="cimFlags">The cim flags.</param>
        /// <returns></returns>
        private static WmiPropertyFlags GetWmiPropertyFlags(CimFlags cimFlags)
        {
            return (WmiPropertyFlags)cimFlags;
        }
        #endregion PropertyFlag Helpers

        #region ExceptionHandlingHelpers
        /// <summary>
        /// Gets WmiManagementStatus from CimException
        /// </summary>
        /// <param name="property">The CimException.</param>
        /// <returns></returns>
        public static WmiManagementStatus GetWmiManagementStatus(CimException exception)
        {
            NativeErrorCode errorCode = exception.NativeErrorCode;

            if (errorCode == NativeErrorCode.Failed)
            {
                uint extendedErrorCode = (uint)exception.ErrorData.CimInstanceProperties["error_Code"].Value;

                return (WmiManagementStatus)(0xFFFFFFFF00000000 + extendedErrorCode);
            }

            switch (errorCode)
            {
                case NativeErrorCode.Ok:
                    return WmiManagementStatus.NoError;
                case NativeErrorCode.Failed:
                    return WmiManagementStatus.HrFailed;
                case NativeErrorCode.AccessDenied:
                    return WmiManagementStatus.HrAccessDenied;
                case NativeErrorCode.InvalidNamespace:
                    return WmiManagementStatus.HrInvalidNamespace;
                case NativeErrorCode.InvalidParameter:
                    return WmiManagementStatus.HrInvalidParameter;
                case NativeErrorCode.InvalidClass:
                    return WmiManagementStatus.HrInvalidClass;
                case NativeErrorCode.NotFound:
                    return WmiManagementStatus.HrNotFound;
                case NativeErrorCode.NotSupported:
                    return WmiManagementStatus.HrNotSupported;
                case NativeErrorCode.ClassHasChildren:
                    return WmiManagementStatus.HrClassHasChildren;
                case NativeErrorCode.ClassHasInstances:
                    return WmiManagementStatus.HrClassHasInstances;
                case NativeErrorCode.InvalidSuperClass:
                    return WmiManagementStatus.HrInvalidSuperclass;
                case NativeErrorCode.AlreadyExists:
                    return WmiManagementStatus.HrAlreadyExists;
                case NativeErrorCode.NoSuchProperty:
                    return WmiManagementStatus.HrInvalidProperty;
                case NativeErrorCode.TypeMismatch:
                    return WmiManagementStatus.HrTypeMismatch;
                case NativeErrorCode.QueryLanguageNotSupported:
                    return WmiManagementStatus.QueryLanguageNotSupported;
                case NativeErrorCode.InvalidQuery:
                    return WmiManagementStatus.HrInvalidQuery;
                case NativeErrorCode.MethodNotAvailable:
                    return WmiManagementStatus.HrInvalidMethod;
                case NativeErrorCode.MethodNotFound:
                    return WmiManagementStatus.HrMethodNotImplemented;
                case NativeErrorCode.NamespaceNotEmpty:
                    return WmiManagementStatus.NamespaceNotEmpty;
                case NativeErrorCode.InvalidEnumerationContext:
                    return WmiManagementStatus.InvalidEnumerationContext;
                case NativeErrorCode.InvalidOperationTimeout:
                    return WmiManagementStatus.InvalidOperationTimeout;
                case NativeErrorCode.PullHasBeenAbandoned:
                    return WmiManagementStatus.PullHasBeenAbandoned;
                case NativeErrorCode.PullCannotBeAbandoned:
                    return WmiManagementStatus.PullCannotBeAbandoned;
                case NativeErrorCode.FilteredEnumerationNotSupported:
                    return WmiManagementStatus.FilteredEnumerationNotSupported;
                case NativeErrorCode.ContinuationOnErrorNotSupported:
                    return WmiManagementStatus.ContinuationOnErrorNotSupported;
                case NativeErrorCode.ServerLimitsExceeded:
                    return WmiManagementStatus.HrServerTooBusy;
                case NativeErrorCode.ServerIsShuttingDown:
                    return WmiManagementStatus.ServerIsShuttingDown;
                default:
                    return WmiManagementStatus.UnrecognizedError;
            }
        }

        #endregion ExceptionHandlingHelpers

        #region Converters for Property

        /// <summary>
        /// Converts CimType to DotNet Type
        /// </summary>
        /// <param name="property">The property.</param>
        /// <returns></returns>
        public static Type GetDotNetTypeFromCimProperty(CimProperty property)
        {
            return GetDotNetTypeFromCimData(property.CimType, property.Value);
        }

        /// <summary>
        /// Gets the dot net value from cim property.
        /// </summary>
        /// <param name="property">The property.</param>
        /// <returns></returns>
        public static object GetDotNetValueFromCimProperty(CimProperty property)
        {
            if (property == null)
                return null;

            return GetDotNetValueFromCimData(property.CimType, property.Value);
        }

        #endregion Converters for Property

        #region Converters for Methods


        /// <summary>
        /// Converts the cim method resultto VM method result.
        /// </summary>
        /// <param name="cimResult">The cim result.</param>
        /// <returns></returns>
        public static WmiMethodResult ConvertCimMethodResulttoWmiMethodResult(CimMethodResult cimResult)
        {
            WmiMethodResult methodResult = new WmiMethodResult();
            if (cimResult.ReturnValue != null && cimResult.ReturnValue.Value != null)
            {
                methodResult.ReturnValue.Value = cimResult.ReturnValue.Value;
            }

            foreach (CimMethodParameter cimParam in cimResult.OutParameters)
            {
                WmiMethodParameter methodParams = new WmiMethodParameter();
                methodParams.Name = cimParam.Name;
                methodParams.Value = WmiHelper.GetDotNetValueFromCimMethodParameter(cimParam);
                methodParams.Type = WmiHelper.GetDotNetTypeFromCimMethodParameter(cimParam);
                methodResult.OutParameters.Add(methodParams);
            }
            return methodResult;
        }

        /// <summary>
        /// Gets the dot net type from cim method parameter.
        /// </summary>
        /// <param name="methodParam">The method param.</param>
        /// <returns></returns>
        public static Type GetDotNetTypeFromCimMethodParameter(CimMethodParameter methodParam)
        {
            return GetDotNetTypeFromCimData(methodParam.CimType, methodParam.Value);
        }

        /// <summary>
        /// Gets the dot net value from cim method parameter.
        /// </summary>
        /// <param name="methodParam">The method param.</param>
        /// <returns></returns>
        public static object GetDotNetValueFromCimMethodParameter(CimMethodParameter methodParam)
        {
            return GetDotNetValueFromCimData(methodParam.CimType, methodParam.Value);
        }

        #endregion Converters for Methods

        #region Generic Converters


        /// <summary>
        /// Gets the cim value from dot net value.
        /// </summary>
        /// <param name="Value">The value.</param>
        /// <returns></returns>
        public static object GetCimValueFromDotNetValue(object Value, CimType type)
        {
            if (Value == null)
                return null;

            object outputValue = null;
            if (Value is IWmiInstance)
            {
                outputValue = (Value as IWmiInstance).GetInstancePath();
            }
            else if (Value is IWmiInstance[])
            {
                outputValue = (Value as IWmiInstance[]).GetInstancePaths();
            }
            else
            {
                outputValue = Value;
            }
            return outputValue;
        }

        /// <summary>
        /// Converts CimType to DotNet Type
        /// </summary>
        /// <param name="cimType">CimType Value</param>
        /// <returns></returns>
        private static Type GetDotNetTypeFromCimData(CimType cimType, object cimValue)
        {
            Type systemType = null;
            switch (cimType)
            {
                case CimType.Unknown:
                    break;

                case CimType.DateTime:
                case CimType.DateTimeArray:
                    object newValue = GetDotNetValueFromCimData(cimType, cimValue);
                    if (newValue != null)
                    {
                        systemType = newValue.GetType();
                    }
                    break;

                case CimType.Reference:
                case CimType.Instance:
                    systemType = typeof(IWmiInstance);
                    break;

                case CimType.InstanceArray:
                case CimType.ReferenceArray:
                    systemType = typeof(IWmiInstance[]);
                    break;

                default:
                    systemType = CimConverter.GetDotNetType(cimType);
                    break;
            }
            return systemType;
        }

        /// <summary>
        /// Gets the dot net value from cim data.
        /// </summary>
        /// <param name="cimType">Type of the cim.</param>
        /// <param name="cimValue">The cim value.</param>
        /// <returns></returns>
        private static object GetDotNetValueFromCimData(CimType cimType, object cimValue)
        {
            if (cimValue == null)
                return null;

            object propertyValue = null;

            switch (cimType)
            {
                case CimType.Instance:
                case CimType.Reference:
                    propertyValue = new WmiInstance(cimValue as CimInstance);
                    break;

                case CimType.InstanceArray:
                case CimType.ReferenceArray:
                    WmiInstanceCollection instanceCollection = new WmiInstanceCollection();
                    foreach (CimInstance instance in cimValue as Array)
                    {
                        instanceCollection.Add(new WmiInstance(instance));
                    }
                    propertyValue = instanceCollection.ToArray();
                    break;

                case CimType.DateTime:
                case CimType.DateTimeArray:
                //TODO(JAYANTM): Do we need any special handling for DateTime?
                //TODO(JAYANTM): Do we need any special handling for other Arrays?
                default:
                    propertyValue = cimValue;
                    break;
            }
            return propertyValue;
        }

        // Converts a given datetime in DMTF format to System.DateTime object.
        public static System.DateTime ToDateTime(string dmtfDate)
        {
            System.DateTime initializer = System.DateTime.MinValue;
            int year = initializer.Year;
            int month = initializer.Month;
            int day = initializer.Day;
            int hour = initializer.Hour;
            int minute = initializer.Minute;
            int second = initializer.Second;
            long ticks = 0;
            string dmtf = dmtfDate;
            System.DateTime datetime = System.DateTime.MinValue;
            string tempString = string.Empty;
            if ((dmtf == null))
            {
                throw new ArgumentNullException("dmtfDate");
            }
            if ((dmtf.Length == 0))
            {
                throw new System.ArgumentOutOfRangeException("dmtfDate");
            }
            if ((dmtf.Length != 25))
            {
                throw new System.ArgumentOutOfRangeException("dmtfDate");
            }
            try
            {
                tempString = dmtf.Substring(0, 4);
                if (("****" != tempString))
                {
                    year = int.Parse(tempString, CultureInfo.InvariantCulture);
                }
                tempString = dmtf.Substring(4, 2);
                if (("**" != tempString))
                {
                    month = int.Parse(tempString, CultureInfo.InvariantCulture);
                }
                tempString = dmtf.Substring(6, 2);
                if (("**" != tempString))
                {
                    day = int.Parse(tempString, CultureInfo.InvariantCulture);
                }
                tempString = dmtf.Substring(8, 2);
                if (("**" != tempString))
                {
                    hour = int.Parse(tempString, CultureInfo.InvariantCulture);
                }
                tempString = dmtf.Substring(10, 2);
                if (("**" != tempString))
                {
                    minute = int.Parse(tempString, CultureInfo.InvariantCulture);
                }
                tempString = dmtf.Substring(12, 2);
                if (("**" != tempString))
                {
                    second = int.Parse(tempString, CultureInfo.InvariantCulture);
                }
                tempString = dmtf.Substring(15, 6);
                if (("******" != tempString))
                {
                    ticks = (long.Parse(tempString, CultureInfo.InvariantCulture) * ((long)((System.TimeSpan.TicksPerMillisecond / 1000))));
                }
                if (((((((((year < 0)
                            || (month < 0))
                            || (day < 0))
                            || (hour < 0))
                            || (minute < 0))
                            || (minute < 0))
                            || (second < 0))
                            || (ticks < 0)))
                {
                    throw new System.ArgumentOutOfRangeException("dmtfDate");
                }
            }
            catch (System.Exception e)
            {
                throw new System.ArgumentOutOfRangeException(null, e.Message);
            }
            datetime = new System.DateTime(year, month, day, hour, minute, second, 0);
            datetime = datetime.AddTicks(ticks);
            System.TimeSpan tickOffset = System.TimeZone.CurrentTimeZone.GetUtcOffset(datetime);
            int UTCOffset = 0;
            int OffsetToBeAdjusted = 0;
            long OffsetMins = ((long)((tickOffset.Ticks / System.TimeSpan.TicksPerMinute)));
            tempString = dmtf.Substring(22, 3);
            if ((tempString != "******"))
            {
                tempString = dmtf.Substring(21, 4);
                try
                {
                    UTCOffset = int.Parse(tempString, CultureInfo.InvariantCulture);
                }
                catch (System.Exception e)
                {
                    throw new System.ArgumentOutOfRangeException(null, e.Message);
                }
                OffsetToBeAdjusted = ((int)((OffsetMins - UTCOffset)));
                datetime = datetime.AddMinutes(((double)(OffsetToBeAdjusted)));
            }
            return datetime;
        }

        // Converts a given System.DateTime object to DMTF datetime format.
        public static string ToDmtfDateTime(System.DateTime date)
        {
            string utcString = string.Empty;
            System.TimeSpan tickOffset = System.TimeZone.CurrentTimeZone.GetUtcOffset(date);
            long OffsetMins = ((long)((tickOffset.Ticks / System.TimeSpan.TicksPerMinute)));
            if ((System.Math.Abs(OffsetMins) > 999))
            {
                date = date.ToUniversalTime();
                utcString = "+000";
            }
            else
            {
                if ((tickOffset.Ticks >= 0))
                {
                    utcString = string.Concat("+",
                        ((long)((tickOffset.Ticks / System.TimeSpan.TicksPerMinute))).ToString(CultureInfo.InvariantCulture).PadLeft(3, '0'));
                }
                else
                {
                    string strTemp = ((long)(OffsetMins)).ToString(CultureInfo.InvariantCulture);
                    utcString = string.Concat("-", strTemp.Substring(1, (strTemp.Length - 1)).PadLeft(3, '0'));
                }
            }
            string dmtfDateTime = ((int)(date.Year)).ToString(CultureInfo.InvariantCulture).PadLeft(4, '0');
            dmtfDateTime = string.Concat(dmtfDateTime, ((int)(date.Month)).ToString(CultureInfo.InvariantCulture).PadLeft(2, '0'));
            dmtfDateTime = string.Concat(dmtfDateTime, ((int)(date.Day)).ToString(CultureInfo.InvariantCulture).PadLeft(2, '0'));
            dmtfDateTime = string.Concat(dmtfDateTime, ((int)(date.Hour)).ToString(CultureInfo.InvariantCulture).PadLeft(2, '0'));
            dmtfDateTime = string.Concat(dmtfDateTime, ((int)(date.Minute)).ToString(CultureInfo.InvariantCulture).PadLeft(2, '0'));
            dmtfDateTime = string.Concat(dmtfDateTime, ((int)(date.Second)).ToString(CultureInfo.InvariantCulture).PadLeft(2, '0'));
            dmtfDateTime = string.Concat(dmtfDateTime, ".");
            System.DateTime dtTemp = new System.DateTime(date.Year, date.Month, date.Day, date.Hour, date.Minute, date.Second, 0);
            long microsec = ((long)((((date.Ticks - dtTemp.Ticks)
                        * 1000)
                        / System.TimeSpan.TicksPerMillisecond)));
            string strMicrosec = ((long)(microsec)).ToString(CultureInfo.InvariantCulture);
            if ((strMicrosec.Length > 6))
            {
                strMicrosec = strMicrosec.Substring(0, 6);
            }
            dmtfDateTime = string.Concat(dmtfDateTime, strMicrosec.PadLeft(6, '0'));
            dmtfDateTime = string.Concat(dmtfDateTime, utcString);
            return dmtfDateTime;
        }

        #endregion Generic Converters

        #region Deprecated Methods

        //
        //TODO(JAYANTM): Remove these methods
        //

        /*
         
        /// <summary>
        /// Determines whether the specified value is reference.
        /// </summary>
        /// <param name="Value">The value.</param>
        /// <returns>
        ///   <c>true</c> if the specified value is reference; otherwise, <c>false</c>.
        /// </returns>
        public static bool IsReference(object Value)
        {
            if (Value != null)
            {
                if (Value is String[] || Value is string[])
                {
                    String[] stringArray = Value as String[];
                    if (stringArray[0] != null && IsObjectPath(stringArray[0]))
                    {
                        return true;
                    }
                }
                else if (Value is String || Value is string[])
                {
                    if (IsObjectPath(Value as String))
                    {
                        return true;
                    }
                }
                else
                {
                    return false;
                }
            }
            return false;
        }

        /// <summary>
        /// Determines whether [is object path] [the specified input].
        /// </summary>
        /// <param name="input">The input.</param>
        /// <returns>
        ///   <c>true</c> if [is object path] [the specified input]; otherwise, <c>false</c>.
        /// </returns>
        public static bool IsObjectPath(String input)
        {
            try
            {
                if (!input.StartsWith(@"\\"))
                {
                    return false;
                }

                //TODO: Add path parsing logic here from IWbemServivces
                if (input.Contains(@"root\virtualization"))
                {
                    return true;
                }

                return false;
            }
            catch (Exception)
            {
                return false;
            }

        }
        
        
        /// <summary>
        /// Determines whether [is embedded instance or reference] [the specified value].
        /// </summary>
        /// <param name="Value">The value.</param>
        /// <returns>
        ///   <c>true</c> if [is embedded instance or reference] [the specified value]; otherwise, <c>false</c>.
        /// </returns>
        public static bool IsEmbeddedInstanceOrReference(object Value)
        {
            //TODO(JAYANTM): Find a better way to identify embedded instance and REF
            if (Value != null)
            {
                if (Value is IWmiInstance || Value is IWmiInstance[])
                {
                    return true;
                }
            }
            return false;
        } 
        
                /// <summary>
        /// Gets the path value array.
        /// </summary>
        /// <param name="Value">The value.</param>
        /// <returns></returns>
        public static string[] GetPathValueArray(IWmiInstance[] Value)
        {
            string[] pathArray = new string[Value.Length];
            for (int index = 0; index < Value.Length; index++)
            {
                string path = GetPathFromInstance(Value[index]);
                pathArray[index] = path;
            }
            return pathArray;
        } 
         
         
         * */

        #endregion DeprecatedMethods

        #region Private Methods
        /// <summary>
        /// Refreshes the cim method output.
        /// </summary>
        /// <param name="methodResult">The method result.</param>


        /// <summary>
        /// Adds the method parameters.
        /// </summary>
        /// <param name="cimClass">The cim class.</param>
        /// <param name="methodName">Name of the method.</param>
        /// <param name="inputParams">The input params.</param>
        /// <param name="methodParameters">The method parameters.</param>
        public static void AddMethodParameters(CimClass cimClass, String methodName,
            WmiMethodParameterCollection inputParams, CimMethodParametersCollection methodParameters)
        {
            try
            {
                foreach (CimMethodParameterDeclaration methodParameter in cimClass.CimClassMethods[methodName].Parameters)
                {
                    bool bInQualifier = false;

                    CimQualifier inQualifier = methodParameter.Qualifiers["In"];

                    if (inQualifier != null)
                    {
                        bInQualifier = true;
                    }

                    if (bInQualifier && inputParams.Contains(methodParameter.Name))
                    {
                        if (inputParams[methodParameter.Name].Value == null)
                        {
                            continue;
                        }
                        AddMethodParam(methodParameters, methodParameter, inputParams[methodParameter.Name].Value);
                    }
                }
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        /// <summary>
        /// Adds the method param.
        /// </summary>
        /// <param name="methodParameters">The method parameters.</param>
        /// <param name="methodDeclaration">The method declaration.</param>
        /// <param name="Value">The value.</param>
        public static void AddMethodParam(CimMethodParametersCollection methodParameters,
            CimMethodParameterDeclaration methodDeclaration, object Value)
        {
            try
            {
                if (Value == null)
                    return;

                switch (methodDeclaration.CimType)
                {
                    case CimType.Reference:
                        methodParameters.Add(CimMethodParameter.Create(methodDeclaration.Name,
                            WmiHelper.GetCimInstance(Value as IWmiInstance), methodDeclaration.CimType, CimFlags.In));
                        break;

                    case CimType.ReferenceArray:
                        methodParameters.Add(CimMethodParameter.Create(methodDeclaration.Name,
                            WmiHelper.GetCimInstanceArray(Value as IWmiInstance[]), methodDeclaration.CimType, CimFlags.In));
                        break;

                    case CimType.DateTime:
                        if ((DateTime)Value != default(DateTime))
                        {
                            methodParameters.Add(CimMethodParameter.Create(methodDeclaration.Name,
                                Value, methodDeclaration.CimType, CimFlags.In));
                        }
                        break;
                    default:
                        methodParameters.Add(CimMethodParameter.Create(methodDeclaration.Name,
                            Value, methodDeclaration.CimType, CimFlags.In));
                        break;
                }
            }
            catch (CimException ex)
            {
                throw new WmiException(WmiHelper.GetWmiManagementStatus(ex), ex.Message, ex);
            }
        }

        public static void RefreshCimMethodOutput(CimMethodResult methodResult)
        {
            foreach (CimMethodParameter methodParam in methodResult.OutParameters)
            {
                //Refresh out parameter if it's an instance type
                if (methodParam.CimType == CimType.Reference)
                {
                    using (CimInstance instanceOut = methodParam.Value as CimInstance)
                    {
                        if (instanceOut != null)
                        {
                            //
                            // BUGBUG: Resolvers are missing for some of the association classes. 
                            // Adding workaround till we get this fixed in product
                            //
                            if (instanceOut.CimSystemProperties.ClassName.Equals("Msvm_ActiveConnection"))
                            {
                                continue;
                            }
                            CimSession cimSession = CimSessionManager.GetSession(instanceOut.GetCimSessionComputerName());

                            using (CimInstance newinstance = cimSession.GetInstance(
                                instanceOut.CimSystemProperties.Namespace,
                                instanceOut))
                            {
                                methodParam.Value = new CimInstance(newinstance);
                            }
                        }
                    }
                }
                else if (methodParam.CimType == CimType.ReferenceArray)
                {
                    CimInstance[] instanceArray = methodParam.Value as CimInstance[];
                    if (instanceArray != null)
                    {
                        for (int index = 0; index < instanceArray.Length; index++)
                        {
                            using (CimInstance oldInstance = instanceArray[index])
                            {
                                CimSession cimSession = CimSessionManager.GetSession(oldInstance.GetCimSessionComputerName());
                                using (CimInstance instance = cimSession.GetInstance(
                                    oldInstance.CimSystemProperties.Namespace,
                                    instanceArray[index]))
                                {
                                    instanceArray[index] = new CimInstance(instance);
                                }
                            }
                        }

                        methodParam.Value = instanceArray;
                    }
                }
            }
        }
        #endregion

    }
}
