// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ComputerSystem struct
type CIM_ComputerSystem struct {
	*CIM_System

	// Enumeration indicating the purpose(s) to which the ComputerSystem is dedicated, if any, and what functionality is provided. For example, one could specify that the System is dedicated to "Print" (value=11) or acts as a "Hub" (value=8).
	///Also, one could indicate that this is a general purpose system by indicating 'Not Dedicated' (value=0) but that it also hosts 'Print' (value=11) or mobile phone 'Mobile User Device' (value=17) services.
	///A clarification is needed with respect to the value 17 ("Mobile User Device"). An example of a dedicated user device is a mobile phone or a barcode scanner in a store that communicates via radio frequency. These systems are quite limited in functionality and programmability, and are not considered 'general purpose' computing platforms. Alternately, an example of a mobile system that is 'general purpose' (i.e., is NOT dedicated) is a hand-held computer. Although limited in its programmability, new software can be downloaded and its functionality expanded by the user.
	///A value of "Management" indicates this instance is dedicated to hosting system management software.
	///A value of "Management Controller" indicates this instance represents specialized hardware dedicated to systems management (i.e., a Baseboard Management Controller (BMC) or service processor).
	///The management scope of a "Management Controller" is typically a single managed system in which it is contained.
	///A value of "Chassis Manager" indicates this instance represents a system dedicated to management of a blade chassis and its contained devices. This value would be used to represent a Shelf Controller. A "Chassis Manager" is an aggregation point for management and may rely on subordinate management controllers for the management of constituent parts. A value of "Host-based RAID Controller" indicates this instance represents a RAID storage controller contained within a host computer. A value of "Storage Device Enclosure" indicates this instance represents an enclosure that contains storage devices. A "Virtual Tape Library" is the emulation of a tape library by a Virtual Library System. A "Virtual Library System" uses disk storage to emulate tape libraries.A "FC Switch" indicates this instance is dedicated to switching layer 2 fibre channel frames. An "Ethernet Switch" indicates this instance is dedicated to switching layer 2 ethernet frames.
	Dedicated []ComputerSystem_Dedicated

	// A string describing how or why the system is dedicated when the Dedicated array includes the value 2, "Other".
	OtherDedicatedDescriptions []string

	// An enumerated array describing the power management capabilities of the ComputerSystem. The use of this property has been deprecated. Instead, the Power Capabilites property in an associated PowerManagement Capabilities class should be used.
	PowerManagementCapabilities []ComputerSystem_PowerManagementCapabilities

	// If enabled (value = 4), the ComputerSystem can be reset via hardware (e.g. the power and reset buttons). If disabled (value = 3), hardware reset is not allowed. In addition to Enabled and Disabled, other Values for the property are also defined - "Not Implemented" (5), "Other" (1) and "Unknown" (2).
	ResetCapability ComputerSystem_ResetCapability
}

func NewCIM_ComputerSystemEx1(instance *cim.WmiInstance) (newInstance *CIM_ComputerSystem, err error) {
	tmp, err := NewCIM_SystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ComputerSystem{
		CIM_System: tmp,
	}
	return
}

func NewCIM_ComputerSystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ComputerSystem, err error) {
	tmp, err := NewCIM_SystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ComputerSystem{
		CIM_System: tmp,
	}
	return
}

// SetDedicated sets the value of Dedicated for the instance
func (instance *CIM_ComputerSystem) SetPropertyDedicated(value []ComputerSystem_Dedicated) (err error) {
	return instance.SetProperty("Dedicated", (value))
}

// GetDedicated gets the value of Dedicated for the instance
func (instance *CIM_ComputerSystem) GetPropertyDedicated() (value []ComputerSystem_Dedicated, err error) {
	retValue, err := instance.GetProperty("Dedicated")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ComputerSystem_Dedicated(valuetmp))
	}

	return
}

// SetOtherDedicatedDescriptions sets the value of OtherDedicatedDescriptions for the instance
func (instance *CIM_ComputerSystem) SetPropertyOtherDedicatedDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherDedicatedDescriptions", (value))
}

// GetOtherDedicatedDescriptions gets the value of OtherDedicatedDescriptions for the instance
func (instance *CIM_ComputerSystem) GetPropertyOtherDedicatedDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherDedicatedDescriptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetPowerManagementCapabilities sets the value of PowerManagementCapabilities for the instance
func (instance *CIM_ComputerSystem) SetPropertyPowerManagementCapabilities(value []ComputerSystem_PowerManagementCapabilities) (err error) {
	return instance.SetProperty("PowerManagementCapabilities", (value))
}

// GetPowerManagementCapabilities gets the value of PowerManagementCapabilities for the instance
func (instance *CIM_ComputerSystem) GetPropertyPowerManagementCapabilities() (value []ComputerSystem_PowerManagementCapabilities, err error) {
	retValue, err := instance.GetProperty("PowerManagementCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ComputerSystem_PowerManagementCapabilities(valuetmp))
	}

	return
}

// SetResetCapability sets the value of ResetCapability for the instance
func (instance *CIM_ComputerSystem) SetPropertyResetCapability(value ComputerSystem_ResetCapability) (err error) {
	return instance.SetProperty("ResetCapability", (value))
}

// GetResetCapability gets the value of ResetCapability for the instance
func (instance *CIM_ComputerSystem) GetPropertyResetCapability() (value ComputerSystem_ResetCapability, err error) {
	retValue, err := instance.GetProperty("ResetCapability")
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

	value = ComputerSystem_ResetCapability(valuetmp)

	return
}

// Sets the power state of the computer. The use of this method has been deprecated. Instead, use the SetPowerState method in the associated PowerManagementService class.

// <param name="PowerState" type="ComputerSystem_PowerState ">The Desired state for the COmputerSystem.</param>
// <param name="Time" type="string ">Time indicates when the power state should be set, either as a regular date-time value or as an interval value (where the interval begins when the method invocation is received.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ComputerSystem) SetPowerState( /* IN */ PowerState ComputerSystem_PowerState,
	/* IN */ Time string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPowerState", PowerState, Time)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
