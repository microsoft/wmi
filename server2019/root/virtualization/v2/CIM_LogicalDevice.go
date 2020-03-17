// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_LogicalDevice struct
type CIM_LogicalDevice struct {
	CIM_EnabledLogicalElement

	// Additional availability and status of the Device, beyond that specified in the Availability property. The Availability property denotes the primary status and availability of the Device. In some cases, this will not be sufficient to denote the complete status of the Device. In those cases, the AdditionalAvailability property can be used to provide further information. For example, a Device's primary Availability may be "Off line" (value=8), but it may also be in a low power state (AdditonalAvailability value=14), or the Device could be running Diagnostics (AdditionalAvailability value=5, "In Test").
	AdditionalAvailability []LogicalDevice_AdditionalAvailability

	// The primary availability and status of the Device. (Additional status information can be specified using the Additional Availability array property.) For example, the Availability property indicates that the Device is running and has full power (value=3), or is in a warning (4), test (5), degraded (10) or power save state (values 13-15 and 17). Regarding the Power Save states, these are defined as follows: Value 13 ("Power Save - Unknown") indicates that the Device is known to be in a power save mode, but its exact status in this mode is unknown; 14 ("Power Save - Low Power Mode") indicates that the Device is in a power save state but still functioning, and may exhibit degraded performance; 15 ("Power Save - Standby") describes that the Device is not functioning but could be brought to full power 'quickly'; and value 17 ("Power Save - Warning") indicates that the Device is in a warning state, though also in a power save mode.
	Availability LogicalDevice_Availability

	// CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
	CreationClassName string

	// An address or other identifying information to uniquely name the LogicalDevice.
	DeviceID string

	// ErrorCleared is a boolean property indicating that the error reported in LastErrorCode is now cleared.
	ErrorCleared bool

	// ErrorDescription is a free-form string supplying more information about the error recorded in LastErrorCode, and information on any corrective actions that may be taken.
	ErrorDescription string

	// An array of free-form strings providing explanations and details behind the entries in the OtherIdentifyingInfo array. Note, each entry of this array is related to the entry in OtherIdentifyingInfo that is located at the same index.
	IdentifyingDescriptions []string

	// LastErrorCode captures the last error code reported by the LogicalDevice.
	LastErrorCode uint32

	// The MaxQuiesceTime property has been deprecated. When evaluating the use of Quiesce, it was determine that this single property is not adequate for describing when a device will automatically exit a quiescent state. In fact, the most likely scenario for a device to exit a quiescent state was determined to be based on the number of outstanding requests queued rather than on a maximum time. This will be re-evaluated and repositioned later.
	///Maximum time in milliseconds, that a Device can run in a "Quiesced" state. A Device's state is defined in its Availability and AdditionalAvailability properties, where "Quiesced" is conveyed by the value 21. What occurs at the end of the time limit is device-specific. The Device may unquiesce, may offline or take other action. A value of 0 indicates that a Device can remain quiesced indefinitely.
	MaxQuiesceTime uint64

	// OtherIdentifyingInfo captures additional data, beyond DeviceID information, that could be used to identify a LogicalDevice. One example would be to hold the Operating System's user friendly name for the Device in this property.
	OtherIdentifyingInfo []string

	// An enumerated array describing the power management capabilities of the Device. The use of this property has been deprecated. Instead, the PowerCapabilites property in an associated PowerManagementCapabilities class should be used.
	PowerManagementCapabilities []LogicalDevice_PowerManagementCapabilities

	// Boolean indicating that the Device can be power managed. The use of this property has been deprecated. Instead, the existence of an associated PowerManagementCapabilities class (associated using the ElementCapabilities relationhip) indicates that power management is supported.
	PowerManagementSupported bool

	// The number of consecutive hours that this Device has been powered, since its last power cycle.
	PowerOnHours uint64

	// The StatusInfo property indicates whether the Logical Device is in an enabled (value = 3), disabled (value = 4) or some other (1) or unknown (2) state. If this property does not apply to the LogicalDevice, the value, 5 ("Not Applicable"), should be used. StatusInfo has been deprecated in lieu of a more clearly named property with additional enumerated values (EnabledState), that is inherited from ManagedSystemElement.
	///If a Device is ("Enabled")(value=3), it has been powered up, and is configured and operational. The Device may or may not be functionally active, depending on whether its Availability (or AdditionalAvailability) indicate that it is ("Running/Full Power")(value=3) or ("Off line") (value=8). In an enabled but offline mode, a Device may be performing out-of-band requests, such as running Diagnostics. If ("Disabled") StatusInfo value=4), a Device can only be "enabled" or powered off. In a personal computer environment, ("Disabled") means that the Device's driver is not available in the stack. In other environments, a Device can be disabled by removing its configuration file. A disabled device is physically present in a System and consuming resources, but can not be communicated with until a load of a driver, a load of a configuration file or some other "enabling" activity has occurred.
	StatusInfo LogicalDevice_StatusInfo

	// The scoping System's CreationClassName.
	SystemCreationClassName string

	// The scoping System's Name.
	SystemName string

	// The total number of hours that this Device has been powered.
	TotalPowerOnHours uint64
}

// SetAdditionalAvailability sets the value of AdditionalAvailability for the instance
func (instance *CIM_LogicalDevice) SetPropertyAdditionalAvailability(value []LogicalDevice_AdditionalAvailability) (err error) {
	return instance.SetProperty("AdditionalAvailability", value)
}

// GetAdditionalAvailability gets the value of AdditionalAvailability for the instance
func (instance *CIM_LogicalDevice) GetPropertyAdditionalAvailability() (value []LogicalDevice_AdditionalAvailability, err error) {
	retValue, err := instance.GetProperty("AdditionalAvailability")
	if err != nil {
		return
	}
	value, ok := retValue.([]LogicalDevice_AdditionalAvailability)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvailability sets the value of Availability for the instance
func (instance *CIM_LogicalDevice) SetPropertyAvailability(value LogicalDevice_Availability) (err error) {
	return instance.SetProperty("Availability", value)
}

// GetAvailability gets the value of Availability for the instance
func (instance *CIM_LogicalDevice) GetPropertyAvailability() (value LogicalDevice_Availability, err error) {
	retValue, err := instance.GetProperty("Availability")
	if err != nil {
		return
	}
	value, ok := retValue.(LogicalDevice_Availability)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_LogicalDevice) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_LogicalDevice) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceID sets the value of DeviceID for the instance
func (instance *CIM_LogicalDevice) SetPropertyDeviceID(value string) (err error) {
	return instance.SetProperty("DeviceID", value)
}

// GetDeviceID gets the value of DeviceID for the instance
func (instance *CIM_LogicalDevice) GetPropertyDeviceID() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorCleared sets the value of ErrorCleared for the instance
func (instance *CIM_LogicalDevice) SetPropertyErrorCleared(value bool) (err error) {
	return instance.SetProperty("ErrorCleared", value)
}

// GetErrorCleared gets the value of ErrorCleared for the instance
func (instance *CIM_LogicalDevice) GetPropertyErrorCleared() (value bool, err error) {
	retValue, err := instance.GetProperty("ErrorCleared")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorDescription sets the value of ErrorDescription for the instance
func (instance *CIM_LogicalDevice) SetPropertyErrorDescription(value string) (err error) {
	return instance.SetProperty("ErrorDescription", value)
}

// GetErrorDescription gets the value of ErrorDescription for the instance
func (instance *CIM_LogicalDevice) GetPropertyErrorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdentifyingDescriptions sets the value of IdentifyingDescriptions for the instance
func (instance *CIM_LogicalDevice) SetPropertyIdentifyingDescriptions(value []string) (err error) {
	return instance.SetProperty("IdentifyingDescriptions", value)
}

// GetIdentifyingDescriptions gets the value of IdentifyingDescriptions for the instance
func (instance *CIM_LogicalDevice) GetPropertyIdentifyingDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("IdentifyingDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastErrorCode sets the value of LastErrorCode for the instance
func (instance *CIM_LogicalDevice) SetPropertyLastErrorCode(value uint32) (err error) {
	return instance.SetProperty("LastErrorCode", value)
}

// GetLastErrorCode gets the value of LastErrorCode for the instance
func (instance *CIM_LogicalDevice) GetPropertyLastErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxQuiesceTime sets the value of MaxQuiesceTime for the instance
func (instance *CIM_LogicalDevice) SetPropertyMaxQuiesceTime(value uint64) (err error) {
	return instance.SetProperty("MaxQuiesceTime", value)
}

// GetMaxQuiesceTime gets the value of MaxQuiesceTime for the instance
func (instance *CIM_LogicalDevice) GetPropertyMaxQuiesceTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxQuiesceTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherIdentifyingInfo sets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_LogicalDevice) SetPropertyOtherIdentifyingInfo(value []string) (err error) {
	return instance.SetProperty("OtherIdentifyingInfo", value)
}

// GetOtherIdentifyingInfo gets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_LogicalDevice) GetPropertyOtherIdentifyingInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherIdentifyingInfo")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPowerManagementCapabilities sets the value of PowerManagementCapabilities for the instance
func (instance *CIM_LogicalDevice) SetPropertyPowerManagementCapabilities(value []LogicalDevice_PowerManagementCapabilities) (err error) {
	return instance.SetProperty("PowerManagementCapabilities", value)
}

// GetPowerManagementCapabilities gets the value of PowerManagementCapabilities for the instance
func (instance *CIM_LogicalDevice) GetPropertyPowerManagementCapabilities() (value []LogicalDevice_PowerManagementCapabilities, err error) {
	retValue, err := instance.GetProperty("PowerManagementCapabilities")
	if err != nil {
		return
	}
	value, ok := retValue.([]LogicalDevice_PowerManagementCapabilities)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPowerManagementSupported sets the value of PowerManagementSupported for the instance
func (instance *CIM_LogicalDevice) SetPropertyPowerManagementSupported(value bool) (err error) {
	return instance.SetProperty("PowerManagementSupported", value)
}

// GetPowerManagementSupported gets the value of PowerManagementSupported for the instance
func (instance *CIM_LogicalDevice) GetPropertyPowerManagementSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("PowerManagementSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPowerOnHours sets the value of PowerOnHours for the instance
func (instance *CIM_LogicalDevice) SetPropertyPowerOnHours(value uint64) (err error) {
	return instance.SetProperty("PowerOnHours", value)
}

// GetPowerOnHours gets the value of PowerOnHours for the instance
func (instance *CIM_LogicalDevice) GetPropertyPowerOnHours() (value uint64, err error) {
	retValue, err := instance.GetProperty("PowerOnHours")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatusInfo sets the value of StatusInfo for the instance
func (instance *CIM_LogicalDevice) SetPropertyStatusInfo(value LogicalDevice_StatusInfo) (err error) {
	return instance.SetProperty("StatusInfo", value)
}

// GetStatusInfo gets the value of StatusInfo for the instance
func (instance *CIM_LogicalDevice) GetPropertyStatusInfo() (value LogicalDevice_StatusInfo, err error) {
	retValue, err := instance.GetProperty("StatusInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(LogicalDevice_StatusInfo)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_LogicalDevice) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", value)
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_LogicalDevice) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *CIM_LogicalDevice) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_LogicalDevice) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalPowerOnHours sets the value of TotalPowerOnHours for the instance
func (instance *CIM_LogicalDevice) SetPropertyTotalPowerOnHours(value uint64) (err error) {
	return instance.SetProperty("TotalPowerOnHours", value)
}

// GetTotalPowerOnHours gets the value of TotalPowerOnHours for the instance
func (instance *CIM_LogicalDevice) GetPropertyTotalPowerOnHours() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalPowerOnHours")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Sets the power state of the Device. The use of this method has been deprecated. Instead, use the SetPowerState method in the associated PowerManagementService class.

// <param name="PowerState" type="LogicalDevice_PowerState ">The power state to set.</param>
// <param name="Time" type="string ">Time indicates when the power state should be set, either as a regular date-time value or as an interval value (where the interval begins when the method invocation is received.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) SetPowerState( /* IN */ PowerState LogicalDevice_PowerState,
	/* IN */ Time string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPowerState", PowerState, Time)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// Requests a reset of the LogicalDevice. The return value should be 0 if the request was successfully executed, 1 if the request is not supported and some other value if an error occurred. In a subclass, the set of possible return codes could be specified, using a ValueMap qualifier on the method. The strings to which the ValueMap contents are 'translated' may also be specified in the subclass as a Values array qualifier.

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) Reset() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Reset")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// The EnableDevice method has been deprecated in lieu of the more general RequestStateChange method that directly overlaps with the functionality provided by this method.
///Requests that the LogicalDevice be enabled ("Enabled" input parameter = TRUE) or disabled (= FALSE). If successful, the Device's StatusInfo/EnabledState properties should reflect the desired state (enabled/disabled). Note that this method's function overlaps with the RequestedState property. RequestedState was added to the model to maintain a record (i.e., a persisted value) of the last state request. Invoking the EnableDevice method should set the RequestedState property appropriately.
///The return code should be 0 if the request was successfully executed, 1 if the request is not supported and some other value if an error occurred. In a subclass, the set of possible return codes could be specified, using a ValueMap qualifier on the method. The strings to which the ValueMap contents are 'translated' may also be specified in the subclass as a Values array qualifier.

// <param name="Enabled" type="bool ">If TRUE enable the device, if FALSE disable the device.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) EnableDevice( /* IN */ Enabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EnableDevice", Enabled)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// The OnlineDevice method has been deprecated in lieu of the more general RequestStateChange method that directly overlaps with the functionality provided by this method.
///Requests that the LogicalDevice be brought online ("Online" input parameter = TRUE) or taken offline (= FALSE). "Online" indicates that the Device is ready to accept requests, and is operational and fully functioning. In this case, the Device's Availability property would be set to a value of 3 ("Running/Full Power"). "Offline" indicates that a Device is powered up and operational, but not processing functional requests. In an offline state, a Device may be capable of running diagnostics or generating operational alerts. For example, when the "Offline" button is pushed on a Printer, the Device is no longer available to process print jobs, but could be available for diagnostics or maintenance.
///If this method is successful, the Device's Availability and AdditionalAvailability properties should reflect the updated status. If a failure occurs trying to bring the Device online or offline, it should remain in its current state. IE, the request, if unsuccessful, should not leave the Device in an indeterminate state. When bringing a Device back "Online", from an "Offline" mode, the Device should be restored to its last "Online" state, if at all possible. Only a Device that has an EnabledState/StatusInfo of "Enabled" and has been configured can be brought online or taken offline.
///OnlineDevice should return 0 if successful, 1 if the request is not supported at all, 2 if the request is not supported due to the current state of the Device, and some other value if any other error occurred. In a subclass, the set of possible return codes could be specified, using a ValueMap qualifier on the method. The strings to which the ValueMap contents are 'translated' may also be specified in the subclass as a Values array qualifier.
///Note that this method's function overlaps with the RequestedState property. RequestedState was added to the model to maintain a record (i.e., a persisted value) of the last state request. Invoking the OnlineDevice method should set the RequestedState property appropriately.

// <param name="Online" type="bool ">If TRUE, take the device online, if FALSE, take the device OFFLINE.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) OnlineDevice( /* IN */ Online bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("OnlineDevice", Online)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// The QuiesceDevice method has been deprecated in lieu of the more general RequestStateChange method that directly overlaps with the functionality provided by this method.
///Requests that the LogicalDevice cleanly cease all current activity ("Quiesce" input parameter = TRUE) or resume activity (= FALSE). For this method to quiesce a Device, that Device should have an Availability (or Additional Availability) of "Running/Full Power" (value=3) and an EnabledStatus/StatusInfo of "Enabled". For example, if quiesced, a Device may then be offlined for diagnostics, or disabled for power off and hot swap. For the method to "unquiesce" a Device, that Device should have an Availability (or AdditionalAvailability) of "Quiesced" (value=21) and an EnabledStatus/StatusInfo of "Enabled". In this case, the Device would be returned to an "Enabled" and "Running/Full Power" status.
///The method's return code should indicate the success or failure of the quiesce. It should return 0 if successful, 1 if the request is not supported at all, 2 if the request is not supported due to the current state of the Device, and some other value if any other error occurred. In a subclass, the set of possible return codes could be specified, using a ValueMap qualifier on the method. The strings to which the ValueMap contents are 'translated' may also be specified in the subclass as a Values array qualifier.

// <param name="Quiesce" type="bool ">If set to TRUE then cleanly cease all activity, if FALSE resume activity.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) QuiesceDevice( /* IN */ Quiesce bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("QuiesceDevice", Quiesce)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// Requests that the Device capture its current configuration, setup and/or state information in a backing store. The goal would be to use this information at a later time (via the RestoreProperties method), to return a Device to its present "condition". This method may not be supported by all Devices. The method should return 0 if successful, 1 if the request is not supported, and some other value if any other error occurred. In a subclass, the set of possible return codes could be specified, using a ValueMap qualifier on the method. The strings to which the ValueMap contents are 'translated' may also be specified in the subclass as a Values array qualifier.

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) SaveProperties() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SaveProperties")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// Requests that the Device re-establish its configuration, setup and/or state information from a backing store. The intent is to capture this information at an earlier time (via the SaveProperties method), and use it to return a Device to this earlier "condition". This method may not be supported by all Devices. The method should return 0 if successful, 1 if the request is not supported, and some other value if any other error occurred. In a subclass, the set of possible return codes could be specified, using a ValueMap qualifier on the method. The strings to which the ValueMap contents are 'translated' may also be specified in the subclass as a Values array qualifier.

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) RestoreProperties() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RestoreProperties")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
