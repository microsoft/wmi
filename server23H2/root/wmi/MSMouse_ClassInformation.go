// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSMouse_ClassInformation struct
type MSMouse_ClassInformation struct {
	*MSMouse

	//
	Active bool

	//
	DeviceId uint64

	//
	InstanceName string
}

func NewMSMouse_ClassInformationEx1(instance *cim.WmiInstance) (newInstance *MSMouse_ClassInformation, err error) {
	tmp, err := NewMSMouseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMouse_ClassInformation{
		MSMouse: tmp,
	}
	return
}

func NewMSMouse_ClassInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMouse_ClassInformation, err error) {
	tmp, err := NewMSMouseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMouse_ClassInformation{
		MSMouse: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSMouse_ClassInformation) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMouse_ClassInformation) GetPropertyActive() (value bool, err error) {
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

// SetDeviceId sets the value of DeviceId for the instance
func (instance *MSMouse_ClassInformation) SetPropertyDeviceId(value uint64) (err error) {
	return instance.SetProperty("DeviceId", (value))
}

// GetDeviceId gets the value of DeviceId for the instance
func (instance *MSMouse_ClassInformation) GetPropertyDeviceId() (value uint64, err error) {
	retValue, err := instance.GetProperty("DeviceId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSMouse_ClassInformation) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMouse_ClassInformation) GetPropertyInstanceName() (value string, err error) {
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
