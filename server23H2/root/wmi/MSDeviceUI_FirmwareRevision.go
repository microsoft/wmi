// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSDeviceUI_FirmwareRevision struct
type MSDeviceUI_FirmwareRevision struct {
	*MSDeviceUI

	//
	Active bool

	//
	FirmwareRevision string

	//
	InstanceName string
}

func NewMSDeviceUI_FirmwareRevisionEx1(instance *cim.WmiInstance) (newInstance *MSDeviceUI_FirmwareRevision, err error) {
	tmp, err := NewMSDeviceUIEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSDeviceUI_FirmwareRevision{
		MSDeviceUI: tmp,
	}
	return
}

func NewMSDeviceUI_FirmwareRevisionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSDeviceUI_FirmwareRevision, err error) {
	tmp, err := NewMSDeviceUIEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSDeviceUI_FirmwareRevision{
		MSDeviceUI: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSDeviceUI_FirmwareRevision) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSDeviceUI_FirmwareRevision) GetPropertyActive() (value bool, err error) {
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

// SetFirmwareRevision sets the value of FirmwareRevision for the instance
func (instance *MSDeviceUI_FirmwareRevision) SetPropertyFirmwareRevision(value string) (err error) {
	return instance.SetProperty("FirmwareRevision", (value))
}

// GetFirmwareRevision gets the value of FirmwareRevision for the instance
func (instance *MSDeviceUI_FirmwareRevision) GetPropertyFirmwareRevision() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareRevision")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSDeviceUI_FirmwareRevision) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSDeviceUI_FirmwareRevision) GetPropertyInstanceName() (value string, err error) {
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
