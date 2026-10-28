// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.Intel_ME
//////////////////////////////////////////////
package intel_me

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ME_System struct
type ME_System struct {
	*CIM_ComputerSystem

	// Capabilities represents the supported features on this system
	Capabilities uint32

	// Represents the status of the chipset CryptoFuse, when disabled, Intel(R) ME will not perform any encryption operations
	CryptoFuseEnabled bool

	// EnabledCapabilities represents the enabled features on this system
	EnabledCapabilities uint32

	// Intel(R) ME firmware version in format: Major.Minor.Revision.build
	FWVersion string

	// ManageabilityMode represents the manageability mode. Every mode supports a different set of capabilities
	ManageabilityMode System_ManageabilityMode

	// PlatformType represents the category the system belongs to
	PlatformType System_PlatformType

	// Represents the OS Control status of Unique Platform ID Feature
	UniquePlatformIDFeatureOSControlState bool

	// Represents the state of Unique Platform ID Feature
	UniquePlatformIDFeatureState bool

	// Represents the support status of Unique Platform ID Feature
	UniquePlatformIDFeatureSupported bool

	// UserSegmentType represents the market segment the system belongs to
	UserSegmentType System_UserSegmentType
}

func NewME_SystemEx1(instance *cim.WmiInstance) (newInstance *ME_System, err error) {
	tmp, err := NewCIM_ComputerSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ME_System{
		CIM_ComputerSystem: tmp,
	}
	return
}

func NewME_SystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ME_System, err error) {
	tmp, err := NewCIM_ComputerSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ME_System{
		CIM_ComputerSystem: tmp,
	}
	return
}

// SetCapabilities sets the value of Capabilities for the instance
func (instance *ME_System) SetPropertyCapabilities(value uint32) (err error) {
	return instance.SetProperty("Capabilities", (value))
}

// GetCapabilities gets the value of Capabilities for the instance
func (instance *ME_System) GetPropertyCapabilities() (value uint32, err error) {
	retValue, err := instance.GetProperty("Capabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetCryptoFuseEnabled sets the value of CryptoFuseEnabled for the instance
func (instance *ME_System) SetPropertyCryptoFuseEnabled(value bool) (err error) {
	return instance.SetProperty("CryptoFuseEnabled", (value))
}

// GetCryptoFuseEnabled gets the value of CryptoFuseEnabled for the instance
func (instance *ME_System) GetPropertyCryptoFuseEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("CryptoFuseEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetEnabledCapabilities sets the value of EnabledCapabilities for the instance
func (instance *ME_System) SetPropertyEnabledCapabilities(value uint32) (err error) {
	return instance.SetProperty("EnabledCapabilities", (value))
}

// GetEnabledCapabilities gets the value of EnabledCapabilities for the instance
func (instance *ME_System) GetPropertyEnabledCapabilities() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnabledCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetFWVersion sets the value of FWVersion for the instance
func (instance *ME_System) SetPropertyFWVersion(value string) (err error) {
	return instance.SetProperty("FWVersion", (value))
}

// GetFWVersion gets the value of FWVersion for the instance
func (instance *ME_System) GetPropertyFWVersion() (value string, err error) {
	retValue, err := instance.GetProperty("FWVersion")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetManageabilityMode sets the value of ManageabilityMode for the instance
func (instance *ME_System) SetPropertyManageabilityMode(value System_ManageabilityMode) (err error) {
	return instance.SetProperty("ManageabilityMode", (value))
}

// GetManageabilityMode gets the value of ManageabilityMode for the instance
func (instance *ME_System) GetPropertyManageabilityMode() (value System_ManageabilityMode, err error) {
	retValue, err := instance.GetProperty("ManageabilityMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = System_ManageabilityMode(valuetmp)

	return
}

// SetPlatformType sets the value of PlatformType for the instance
func (instance *ME_System) SetPropertyPlatformType(value System_PlatformType) (err error) {
	return instance.SetProperty("PlatformType", (value))
}

// GetPlatformType gets the value of PlatformType for the instance
func (instance *ME_System) GetPropertyPlatformType() (value System_PlatformType, err error) {
	retValue, err := instance.GetProperty("PlatformType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = System_PlatformType(valuetmp)

	return
}

// SetUniquePlatformIDFeatureOSControlState sets the value of UniquePlatformIDFeatureOSControlState for the instance
func (instance *ME_System) SetPropertyUniquePlatformIDFeatureOSControlState(value bool) (err error) {
	return instance.SetProperty("UniquePlatformIDFeatureOSControlState", (value))
}

// GetUniquePlatformIDFeatureOSControlState gets the value of UniquePlatformIDFeatureOSControlState for the instance
func (instance *ME_System) GetPropertyUniquePlatformIDFeatureOSControlState() (value bool, err error) {
	retValue, err := instance.GetProperty("UniquePlatformIDFeatureOSControlState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUniquePlatformIDFeatureState sets the value of UniquePlatformIDFeatureState for the instance
func (instance *ME_System) SetPropertyUniquePlatformIDFeatureState(value bool) (err error) {
	return instance.SetProperty("UniquePlatformIDFeatureState", (value))
}

// GetUniquePlatformIDFeatureState gets the value of UniquePlatformIDFeatureState for the instance
func (instance *ME_System) GetPropertyUniquePlatformIDFeatureState() (value bool, err error) {
	retValue, err := instance.GetProperty("UniquePlatformIDFeatureState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUniquePlatformIDFeatureSupported sets the value of UniquePlatformIDFeatureSupported for the instance
func (instance *ME_System) SetPropertyUniquePlatformIDFeatureSupported(value bool) (err error) {
	return instance.SetProperty("UniquePlatformIDFeatureSupported", (value))
}

// GetUniquePlatformIDFeatureSupported gets the value of UniquePlatformIDFeatureSupported for the instance
func (instance *ME_System) GetPropertyUniquePlatformIDFeatureSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("UniquePlatformIDFeatureSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUserSegmentType sets the value of UserSegmentType for the instance
func (instance *ME_System) SetPropertyUserSegmentType(value System_UserSegmentType) (err error) {
	return instance.SetProperty("UserSegmentType", (value))
}

// GetUserSegmentType gets the value of UserSegmentType for the instance
func (instance *ME_System) GetPropertyUserSegmentType() (value System_UserSegmentType, err error) {
	retValue, err := instance.GetProperty("UserSegmentType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = System_UserSegmentType(valuetmp)

	return
}

// Receives local  FW Update counters (In Intel(R) AMT version 7.0 and later this function is not supported and always returns NOT SUPPORTED error)

// <param name="Counter" type="uint32 ">FW Update override counter. When 255 refer to always</param>
// <param name="OverrideQualifier" type="System_OverrideQualifier ">FW Update override qualifier</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *ME_System) getFwUpdateOverrideParams( /* OUT */ Counter uint32,
	/* OUT */ OverrideQualifier System_OverrideQualifier) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("getFwUpdateOverrideParams")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Returns the cause for the last Intel(R) ME Reset operation

// <param name="ReasonCode" type="System_ReasonCode ">A description of the last Intel(R) ME reset reason</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *ME_System) getLastMEResetReason( /* OUT */ ReasonCode System_ReasonCode) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("getLastMEResetReason")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Specifies the list of ACPI states Intel(R) ME is operational in

// <param name="PowerPolicy" type="string ">A description of the power policy</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *ME_System) getCurrentPowerPolicy( /* OUT */ PowerPolicy string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("getCurrentPowerPolicy")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Indicates whether the Firmware Update capability is currently enabled

// <param name="enabled" type="bool ">Indicates whether the Firmware Update capability is currently enabled</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *ME_System) IsFirmwareUpdateEnabled( /* OUT */ enabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsFirmwareUpdateEnabled")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Specifies the list of supported features on this system

// <param name="Capabilities" type="string []">An array of strings defining system capabilities. Should match the bitvalues of the system Capabilities</param>
// <param name="EnabledCapabilities" type="string []">An array of strings defining system capabilities. Should match the bitvalues of the system Enabled Capabilities</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *ME_System) getCapabilities( /* OUT */ Capabilities []string,
	/* OUT */ EnabledCapabilities []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("getCapabilities")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Get Intel Unique Platform ID feature State

// <param name="ReturnValue" type="uint32 "></param>
// <param name="state" type="bool ">Unique Platform ID feature State</param>
func (instance *ME_System) getUniquePlatformIDFeatureState( /* OUT */ state bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("getUniquePlatformIDFeatureState")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Set Intel Unique Platform ID feature State

// <param name="state" type="bool ">Unique Platform ID feature State</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *ME_System) setUniquePlatformIDFeatureState( /* IN */ state bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("setUniquePlatformIDFeatureState", state)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// Get Intel Unique Platform ID

// <param name="CSMEPlatformID" type="string "></param>
// <param name="OEMPlatformID" type="string ">OEM Platform ID </param>
// <param name="OEMPlatformIDType" type="System_OEMPlatformIDType ">Platform ID Type</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *ME_System) getUniquePlatformID( /* OUT */ OEMPlatformIDType System_OEMPlatformIDType,
	/* OUT */ OEMPlatformID string,
	/* OUT */ CSMEPlatformID string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("getUniquePlatformID")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Intel Unique Platform ID feature supported on platform

// <param name="ReturnValue" type="uint32 "></param>
// <param name="supported" type="bool ">Intel Unique Platform ID feature supported on platform</param>
func (instance *ME_System) getUniquePlatformIDFeatureSupported( /* OUT */ supported bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("getUniquePlatformIDFeatureSupported")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Intel Unique Platform ID feature OS Control State

// <param name="ReturnValue" type="uint32 "></param>
// <param name="state" type="bool ">Unique Platform ID OS Control State</param>
func (instance *ME_System) getUniquePlatformIDFeatureOSControlState( /* OUT */ state bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("getUniquePlatformIDFeatureOSControlState")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
