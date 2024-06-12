// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSPower_DeviceEnable struct
type MSPower_DeviceEnable struct {
	*MSPower

	//
	Active bool

	//
	Enable bool

	//
	InstanceName string
}

func NewMSPower_DeviceEnableEx1(instance *cim.WmiInstance) (newInstance *MSPower_DeviceEnable, err error) {
	tmp, err := NewMSPowerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSPower_DeviceEnable{
		MSPower: tmp,
	}
	return
}

func NewMSPower_DeviceEnableEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSPower_DeviceEnable, err error) {
	tmp, err := NewMSPowerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSPower_DeviceEnable{
		MSPower: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSPower_DeviceEnable) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSPower_DeviceEnable) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetEnable sets the value of Enable for the instance
func (instance *MSPower_DeviceEnable) SetPropertyEnable(value bool) (err error) {
	return instance.SetProperty("Enable", (value))
}

// GetEnable gets the value of Enable for the instance
func (instance *MSPower_DeviceEnable) GetPropertyEnable() (value bool, err error) {
	retValue, err := instance.GetProperty("Enable")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSPower_DeviceEnable) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSPower_DeviceEnable) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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
