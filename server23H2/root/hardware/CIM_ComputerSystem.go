// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ComputerSystem struct
type CIM_ComputerSystem struct {
	*CIM_System

	//
	Dedicated []uint16

	//
	IdentifyingDescriptions []string

	//
	OtherDedicatedDescriptions []string

	//
	OtherIdentifyingInfo []string

	//
	PowerManagementCapabilities []uint16

	//
	ResetCapability uint16
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
func (instance *CIM_ComputerSystem) SetPropertyDedicated(value []uint16) (err error) {
	return instance.SetProperty("Dedicated", (value))
}

// GetDedicated gets the value of Dedicated for the instance
func (instance *CIM_ComputerSystem) GetPropertyDedicated() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Dedicated")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetIdentifyingDescriptions sets the value of IdentifyingDescriptions for the instance
func (instance *CIM_ComputerSystem) SetPropertyIdentifyingDescriptions(value []string) (err error) {
	return instance.SetProperty("IdentifyingDescriptions", (value))
}

// GetIdentifyingDescriptions gets the value of IdentifyingDescriptions for the instance
func (instance *CIM_ComputerSystem) GetPropertyIdentifyingDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("IdentifyingDescriptions")
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

// SetOtherIdentifyingInfo sets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_ComputerSystem) SetPropertyOtherIdentifyingInfo(value []string) (err error) {
	return instance.SetProperty("OtherIdentifyingInfo", (value))
}

// GetOtherIdentifyingInfo gets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_ComputerSystem) GetPropertyOtherIdentifyingInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherIdentifyingInfo")
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
func (instance *CIM_ComputerSystem) SetPropertyPowerManagementCapabilities(value []uint16) (err error) {
	return instance.SetProperty("PowerManagementCapabilities", (value))
}

// GetPowerManagementCapabilities gets the value of PowerManagementCapabilities for the instance
func (instance *CIM_ComputerSystem) GetPropertyPowerManagementCapabilities() (value []uint16, err error) {
	retValue, err := instance.GetProperty("PowerManagementCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetResetCapability sets the value of ResetCapability for the instance
func (instance *CIM_ComputerSystem) SetPropertyResetCapability(value uint16) (err error) {
	return instance.SetProperty("ResetCapability", (value))
}

// GetResetCapability gets the value of ResetCapability for the instance
func (instance *CIM_ComputerSystem) GetPropertyResetCapability() (value uint16, err error) {
	retValue, err := instance.GetProperty("ResetCapability")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

//

// <param name="PowerState" type="uint32 "></param>
// <param name="Time" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ComputerSystem) SetPowerState( /* IN */ PowerState uint32,
	/* IN */ Time string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPowerState", PowerState, Time)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
