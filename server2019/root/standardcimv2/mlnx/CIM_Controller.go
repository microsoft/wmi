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

// CIM_Controller struct
type CIM_Controller struct {
	*CIM_LogicalDevice

	//
	MaxNumberControlled uint32

	//
	ProtocolDescription string

	//
	ProtocolSupported Controller_ProtocolSupported

	//
	TimeOfLastReset string
}

func NewCIM_ControllerEx1(instance *cim.WmiInstance) (newInstance *CIM_Controller, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Controller{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewCIM_ControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Controller, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Controller{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetMaxNumberControlled sets the value of MaxNumberControlled for the instance
func (instance *CIM_Controller) SetPropertyMaxNumberControlled(value uint32) (err error) {
	return instance.SetProperty("MaxNumberControlled", (value))
}

// GetMaxNumberControlled gets the value of MaxNumberControlled for the instance
func (instance *CIM_Controller) GetPropertyMaxNumberControlled() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumberControlled")
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

// SetProtocolDescription sets the value of ProtocolDescription for the instance
func (instance *CIM_Controller) SetPropertyProtocolDescription(value string) (err error) {
	return instance.SetProperty("ProtocolDescription", (value))
}

// GetProtocolDescription gets the value of ProtocolDescription for the instance
func (instance *CIM_Controller) GetPropertyProtocolDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ProtocolDescription")
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

// SetProtocolSupported sets the value of ProtocolSupported for the instance
func (instance *CIM_Controller) SetPropertyProtocolSupported(value Controller_ProtocolSupported) (err error) {
	return instance.SetProperty("ProtocolSupported", (value))
}

// GetProtocolSupported gets the value of ProtocolSupported for the instance
func (instance *CIM_Controller) GetPropertyProtocolSupported() (value Controller_ProtocolSupported, err error) {
	retValue, err := instance.GetProperty("ProtocolSupported")
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

	value = Controller_ProtocolSupported(valuetmp)

	return
}

// SetTimeOfLastReset sets the value of TimeOfLastReset for the instance
func (instance *CIM_Controller) SetPropertyTimeOfLastReset(value string) (err error) {
	return instance.SetProperty("TimeOfLastReset", (value))
}

// GetTimeOfLastReset gets the value of TimeOfLastReset for the instance
func (instance *CIM_Controller) GetPropertyTimeOfLastReset() (value string, err error) {
	retValue, err := instance.GetProperty("TimeOfLastReset")
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
