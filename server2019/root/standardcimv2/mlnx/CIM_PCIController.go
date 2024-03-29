// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_PCIController struct
type CIM_PCIController struct {
	*CIM_Controller

	//
	CacheLineSize uint8

	//
	Capabilities []PCIController_Capabilities

	//
	CapabilityDescriptions []string

	//
	ClassCode PCIController_ClassCode

	//
	CommandRegister uint16

	//
	DeviceSelectTiming PCIController_DeviceSelectTiming

	//
	ExpansionROMBaseAddress uint32

	//
	InterruptPin PCIController_InterruptPin

	//
	LatencyTimer uint8

	//
	SelfTestEnabled bool
}

func NewCIM_PCIControllerEx1(instance *cim.WmiInstance) (newInstance *CIM_PCIController, err error) {
	tmp, err := NewCIM_ControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PCIController{
		CIM_Controller: tmp,
	}
	return
}

func NewCIM_PCIControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PCIController, err error) {
	tmp, err := NewCIM_ControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PCIController{
		CIM_Controller: tmp,
	}
	return
}

// SetCacheLineSize sets the value of CacheLineSize for the instance
func (instance *CIM_PCIController) SetPropertyCacheLineSize(value uint8) (err error) {
	return instance.SetProperty("CacheLineSize", (value))
}

// GetCacheLineSize gets the value of CacheLineSize for the instance
func (instance *CIM_PCIController) GetPropertyCacheLineSize() (value uint8, err error) {
	retValue, err := instance.GetProperty("CacheLineSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetCapabilities sets the value of Capabilities for the instance
func (instance *CIM_PCIController) SetPropertyCapabilities(value []PCIController_Capabilities) (err error) {
	return instance.SetProperty("Capabilities", (value))
}

// GetCapabilities gets the value of Capabilities for the instance
func (instance *CIM_PCIController) GetPropertyCapabilities() (value []PCIController_Capabilities, err error) {
	retValue, err := instance.GetProperty("Capabilities")
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
		value = append(value, PCIController_Capabilities(valuetmp))
	}

	return
}

// SetCapabilityDescriptions sets the value of CapabilityDescriptions for the instance
func (instance *CIM_PCIController) SetPropertyCapabilityDescriptions(value []string) (err error) {
	return instance.SetProperty("CapabilityDescriptions", (value))
}

// GetCapabilityDescriptions gets the value of CapabilityDescriptions for the instance
func (instance *CIM_PCIController) GetPropertyCapabilityDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("CapabilityDescriptions")
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

// SetClassCode sets the value of ClassCode for the instance
func (instance *CIM_PCIController) SetPropertyClassCode(value PCIController_ClassCode) (err error) {
	return instance.SetProperty("ClassCode", (value))
}

// GetClassCode gets the value of ClassCode for the instance
func (instance *CIM_PCIController) GetPropertyClassCode() (value PCIController_ClassCode, err error) {
	retValue, err := instance.GetProperty("ClassCode")
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

	value = PCIController_ClassCode(valuetmp)

	return
}

// SetCommandRegister sets the value of CommandRegister for the instance
func (instance *CIM_PCIController) SetPropertyCommandRegister(value uint16) (err error) {
	return instance.SetProperty("CommandRegister", (value))
}

// GetCommandRegister gets the value of CommandRegister for the instance
func (instance *CIM_PCIController) GetPropertyCommandRegister() (value uint16, err error) {
	retValue, err := instance.GetProperty("CommandRegister")
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

// SetDeviceSelectTiming sets the value of DeviceSelectTiming for the instance
func (instance *CIM_PCIController) SetPropertyDeviceSelectTiming(value PCIController_DeviceSelectTiming) (err error) {
	return instance.SetProperty("DeviceSelectTiming", (value))
}

// GetDeviceSelectTiming gets the value of DeviceSelectTiming for the instance
func (instance *CIM_PCIController) GetPropertyDeviceSelectTiming() (value PCIController_DeviceSelectTiming, err error) {
	retValue, err := instance.GetProperty("DeviceSelectTiming")
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

	value = PCIController_DeviceSelectTiming(valuetmp)

	return
}

// SetExpansionROMBaseAddress sets the value of ExpansionROMBaseAddress for the instance
func (instance *CIM_PCIController) SetPropertyExpansionROMBaseAddress(value uint32) (err error) {
	return instance.SetProperty("ExpansionROMBaseAddress", (value))
}

// GetExpansionROMBaseAddress gets the value of ExpansionROMBaseAddress for the instance
func (instance *CIM_PCIController) GetPropertyExpansionROMBaseAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExpansionROMBaseAddress")
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

// SetInterruptPin sets the value of InterruptPin for the instance
func (instance *CIM_PCIController) SetPropertyInterruptPin(value PCIController_InterruptPin) (err error) {
	return instance.SetProperty("InterruptPin", (value))
}

// GetInterruptPin gets the value of InterruptPin for the instance
func (instance *CIM_PCIController) GetPropertyInterruptPin() (value PCIController_InterruptPin, err error) {
	retValue, err := instance.GetProperty("InterruptPin")
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

	value = PCIController_InterruptPin(valuetmp)

	return
}

// SetLatencyTimer sets the value of LatencyTimer for the instance
func (instance *CIM_PCIController) SetPropertyLatencyTimer(value uint8) (err error) {
	return instance.SetProperty("LatencyTimer", (value))
}

// GetLatencyTimer gets the value of LatencyTimer for the instance
func (instance *CIM_PCIController) GetPropertyLatencyTimer() (value uint8, err error) {
	retValue, err := instance.GetProperty("LatencyTimer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSelfTestEnabled sets the value of SelfTestEnabled for the instance
func (instance *CIM_PCIController) SetPropertySelfTestEnabled(value bool) (err error) {
	return instance.SetProperty("SelfTestEnabled", (value))
}

// GetSelfTestEnabled gets the value of SelfTestEnabled for the instance
func (instance *CIM_PCIController) GetPropertySelfTestEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("SelfTestEnabled")
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

//

// <param name="ReturnValue" type="uint8 "></param>
func (instance *CIM_PCIController) BISTExecution() (result uint8, err error) {
	retVal, err := instance.InvokeMethodWithReturn("BISTExecution")
	if err != nil {
		return
	}
	result = uint8(retVal)
	return

}
