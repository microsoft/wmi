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

// MLNX_PCIDeviceCapabilities struct
type MLNX_PCIDeviceCapabilities struct {
	*CIM_Capabilities

	//
	DualPort bool

	//
	Name string

	//
	PortOneAutoSense bool

	//
	PortOneDefault bool

	//
	PortOneDoSenseAllowed bool

	//
	PortOneEth bool

	//
	PortOneIb bool

	//
	PortTwoAutoSenseCap bool

	//
	PortTwoDefault bool

	//
	PortTwoDoSenseAllowed bool

	//
	PortTwoEth bool

	//
	PortTwoIb bool

	//
	SystemName string
}

func NewMLNX_PCIDeviceCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MLNX_PCIDeviceCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_PCIDeviceCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

func NewMLNX_PCIDeviceCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_PCIDeviceCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_PCIDeviceCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

// SetDualPort sets the value of DualPort for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyDualPort(value bool) (err error) {
	return instance.SetProperty("DualPort", (value))
}

// GetDualPort gets the value of DualPort for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyDualPort() (value bool, err error) {
	retValue, err := instance.GetProperty("DualPort")
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

// SetName sets the value of Name for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetPortOneAutoSense sets the value of PortOneAutoSense for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyPortOneAutoSense(value bool) (err error) {
	return instance.SetProperty("PortOneAutoSense", (value))
}

// GetPortOneAutoSense gets the value of PortOneAutoSense for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyPortOneAutoSense() (value bool, err error) {
	retValue, err := instance.GetProperty("PortOneAutoSense")
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

// SetPortOneDefault sets the value of PortOneDefault for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyPortOneDefault(value bool) (err error) {
	return instance.SetProperty("PortOneDefault", (value))
}

// GetPortOneDefault gets the value of PortOneDefault for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyPortOneDefault() (value bool, err error) {
	retValue, err := instance.GetProperty("PortOneDefault")
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

// SetPortOneDoSenseAllowed sets the value of PortOneDoSenseAllowed for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyPortOneDoSenseAllowed(value bool) (err error) {
	return instance.SetProperty("PortOneDoSenseAllowed", (value))
}

// GetPortOneDoSenseAllowed gets the value of PortOneDoSenseAllowed for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyPortOneDoSenseAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("PortOneDoSenseAllowed")
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

// SetPortOneEth sets the value of PortOneEth for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyPortOneEth(value bool) (err error) {
	return instance.SetProperty("PortOneEth", (value))
}

// GetPortOneEth gets the value of PortOneEth for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyPortOneEth() (value bool, err error) {
	retValue, err := instance.GetProperty("PortOneEth")
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

// SetPortOneIb sets the value of PortOneIb for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyPortOneIb(value bool) (err error) {
	return instance.SetProperty("PortOneIb", (value))
}

// GetPortOneIb gets the value of PortOneIb for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyPortOneIb() (value bool, err error) {
	retValue, err := instance.GetProperty("PortOneIb")
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

// SetPortTwoAutoSenseCap sets the value of PortTwoAutoSenseCap for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyPortTwoAutoSenseCap(value bool) (err error) {
	return instance.SetProperty("PortTwoAutoSenseCap", (value))
}

// GetPortTwoAutoSenseCap gets the value of PortTwoAutoSenseCap for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyPortTwoAutoSenseCap() (value bool, err error) {
	retValue, err := instance.GetProperty("PortTwoAutoSenseCap")
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

// SetPortTwoDefault sets the value of PortTwoDefault for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyPortTwoDefault(value bool) (err error) {
	return instance.SetProperty("PortTwoDefault", (value))
}

// GetPortTwoDefault gets the value of PortTwoDefault for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyPortTwoDefault() (value bool, err error) {
	retValue, err := instance.GetProperty("PortTwoDefault")
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

// SetPortTwoDoSenseAllowed sets the value of PortTwoDoSenseAllowed for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyPortTwoDoSenseAllowed(value bool) (err error) {
	return instance.SetProperty("PortTwoDoSenseAllowed", (value))
}

// GetPortTwoDoSenseAllowed gets the value of PortTwoDoSenseAllowed for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyPortTwoDoSenseAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("PortTwoDoSenseAllowed")
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

// SetPortTwoEth sets the value of PortTwoEth for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyPortTwoEth(value bool) (err error) {
	return instance.SetProperty("PortTwoEth", (value))
}

// GetPortTwoEth gets the value of PortTwoEth for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyPortTwoEth() (value bool, err error) {
	retValue, err := instance.GetProperty("PortTwoEth")
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

// SetPortTwoIb sets the value of PortTwoIb for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertyPortTwoIb(value bool) (err error) {
	return instance.SetProperty("PortTwoIb", (value))
}

// GetPortTwoIb gets the value of PortTwoIb for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertyPortTwoIb() (value bool, err error) {
	retValue, err := instance.GetProperty("PortTwoIb")
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

// SetSystemName sets the value of SystemName for the instance
func (instance *MLNX_PCIDeviceCapabilities) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", (value))
}

// GetSystemName gets the value of SystemName for the instance
func (instance *MLNX_PCIDeviceCapabilities) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
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
