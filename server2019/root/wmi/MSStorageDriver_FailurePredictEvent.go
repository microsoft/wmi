// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSStorageDriver_FailurePredictEvent struct
type MSStorageDriver_FailurePredictEvent struct {
	*WMIEvent

	//
	Active bool

	//
	InstanceName string

	//
	Length uint32

	//
	VendorSpecific []uint8
}

func NewMSStorageDriver_FailurePredictEventEx1(instance *cim.WmiInstance) (newInstance *MSStorageDriver_FailurePredictEvent, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver_FailurePredictEvent{
		WMIEvent: tmp,
	}
	return
}

func NewMSStorageDriver_FailurePredictEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSStorageDriver_FailurePredictEvent, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver_FailurePredictEvent{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSStorageDriver_FailurePredictEvent) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSStorageDriver_FailurePredictEvent) GetPropertyActive() (value bool, err error) {
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
func (instance *MSStorageDriver_FailurePredictEvent) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSStorageDriver_FailurePredictEvent) GetPropertyInstanceName() (value string, err error) {
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

// SetLength sets the value of Length for the instance
func (instance *MSStorageDriver_FailurePredictEvent) SetPropertyLength(value uint32) (err error) {
	return instance.SetProperty("Length", (value))
}

// GetLength gets the value of Length for the instance
func (instance *MSStorageDriver_FailurePredictEvent) GetPropertyLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("Length")
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

// SetVendorSpecific sets the value of VendorSpecific for the instance
func (instance *MSStorageDriver_FailurePredictEvent) SetPropertyVendorSpecific(value []uint8) (err error) {
	return instance.SetProperty("VendorSpecific", (value))
}

// GetVendorSpecific gets the value of VendorSpecific for the instance
func (instance *MSStorageDriver_FailurePredictEvent) GetPropertyVendorSpecific() (value []uint8, err error) {
	retValue, err := instance.GetProperty("VendorSpecific")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}
