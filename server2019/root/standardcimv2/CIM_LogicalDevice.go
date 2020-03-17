// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_LogicalDevice struct
type CIM_LogicalDevice struct {
	CIM_EnabledLogicalElement

	//
	AdditionalAvailability []uint16

	//
	Availability uint16

	//
	CreationClassName string

	//
	DeviceID string

	//
	ErrorCleared bool

	//
	ErrorDescription string

	//
	IdentifyingDescriptions []string

	//
	LastErrorCode uint32

	//
	MaxQuiesceTime uint64

	//
	OtherIdentifyingInfo []string

	//
	PowerManagementCapabilities []uint16

	//
	PowerManagementSupported bool

	//
	PowerOnHours uint64

	//
	StatusInfo uint16

	//
	SystemCreationClassName string

	//
	SystemName string

	//
	TotalPowerOnHours uint64
}

// SetAdditionalAvailability sets the value of AdditionalAvailability for the instance
func (instance *CIM_LogicalDevice) SetPropertyAdditionalAvailability(value []uint16) (err error) {
	return instance.SetProperty("AdditionalAvailability", value)
}

// GetAdditionalAvailability gets the value of AdditionalAvailability for the instance
func (instance *CIM_LogicalDevice) GetPropertyAdditionalAvailability() (value []uint16, err error) {
	retValue, err := instance.GetProperty("AdditionalAvailability")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAvailability sets the value of Availability for the instance
func (instance *CIM_LogicalDevice) SetPropertyAvailability(value uint16) (err error) {
	return instance.SetProperty("Availability", value)
}

// GetAvailability gets the value of Availability for the instance
func (instance *CIM_LogicalDevice) GetPropertyAvailability() (value uint16, err error) {
	retValue, err := instance.GetProperty("Availability")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
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
func (instance *CIM_LogicalDevice) SetPropertyPowerManagementCapabilities(value []uint16) (err error) {
	return instance.SetProperty("PowerManagementCapabilities", value)
}

// GetPowerManagementCapabilities gets the value of PowerManagementCapabilities for the instance
func (instance *CIM_LogicalDevice) GetPropertyPowerManagementCapabilities() (value []uint16, err error) {
	retValue, err := instance.GetProperty("PowerManagementCapabilities")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
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
func (instance *CIM_LogicalDevice) SetPropertyStatusInfo(value uint16) (err error) {
	return instance.SetProperty("StatusInfo", value)
}

// GetStatusInfo gets the value of StatusInfo for the instance
func (instance *CIM_LogicalDevice) GetPropertyStatusInfo() (value uint16, err error) {
	retValue, err := instance.GetProperty("StatusInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
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

//

// <param name="PowerState" type="uint16 "></param>
// <param name="Time" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) SetPowerState( /* IN */ PowerState uint16,
	/* IN */ Time string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPowerState", PowerState, Time)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) Reset() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Reset")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Enabled" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) EnableDevice( /* IN */ Enabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EnableDevice", Enabled)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Online" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) OnlineDevice( /* IN */ Online bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("OnlineDevice", Online)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Quiesce" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) QuiesceDevice( /* IN */ Quiesce bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("QuiesceDevice", Quiesce)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) SaveProperties() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SaveProperties")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_LogicalDevice) RestoreProperties() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RestoreProperties")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
