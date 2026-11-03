// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HDAudioBus_Enum struct
type HDAudioBus_Enum struct {
	*HDAudioBusEvent

	//
	Active bool

	//
	InstanceName string

	//
	NewDeviceState uint8

	//
	OldDeviceState uint8
}

func NewHDAudioBus_EnumEx1(instance *cim.WmiInstance) (newInstance *HDAudioBus_Enum, err error) {
	tmp, err := NewHDAudioBusEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HDAudioBus_Enum{
		HDAudioBusEvent: tmp,
	}
	return
}

func NewHDAudioBus_EnumEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HDAudioBus_Enum, err error) {
	tmp, err := NewHDAudioBusEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HDAudioBus_Enum{
		HDAudioBusEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *HDAudioBus_Enum) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *HDAudioBus_Enum) GetPropertyActive() (value bool, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *HDAudioBus_Enum) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *HDAudioBus_Enum) GetPropertyInstanceName() (value string, err error) {
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

// SetNewDeviceState sets the value of NewDeviceState for the instance
func (instance *HDAudioBus_Enum) SetPropertyNewDeviceState(value uint8) (err error) {
	return instance.SetProperty("NewDeviceState", (value))
}

// GetNewDeviceState gets the value of NewDeviceState for the instance
func (instance *HDAudioBus_Enum) GetPropertyNewDeviceState() (value uint8, err error) {
	retValue, err := instance.GetProperty("NewDeviceState")
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

// SetOldDeviceState sets the value of OldDeviceState for the instance
func (instance *HDAudioBus_Enum) SetPropertyOldDeviceState(value uint8) (err error) {
	return instance.SetProperty("OldDeviceState", (value))
}

// GetOldDeviceState gets the value of OldDeviceState for the instance
func (instance *HDAudioBus_Enum) GetPropertyOldDeviceState() (value uint8, err error) {
	retValue, err := instance.GetProperty("OldDeviceState")
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
