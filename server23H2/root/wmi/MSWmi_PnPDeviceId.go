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

// MSWmi_PnPDeviceId struct
type MSWmi_PnPDeviceId struct {
	*MS_WmiInternal

	//
	Active bool

	//
	InstanceName string

	//
	PnPDeviceId string
}

func NewMSWmi_PnPDeviceIdEx1(instance *cim.WmiInstance) (newInstance *MSWmi_PnPDeviceId, err error) {
	tmp, err := NewMS_WmiInternalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSWmi_PnPDeviceId{
		MS_WmiInternal: tmp,
	}
	return
}

func NewMSWmi_PnPDeviceIdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSWmi_PnPDeviceId, err error) {
	tmp, err := NewMS_WmiInternalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSWmi_PnPDeviceId{
		MS_WmiInternal: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSWmi_PnPDeviceId) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSWmi_PnPDeviceId) GetPropertyActive() (value bool, err error) {
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
func (instance *MSWmi_PnPDeviceId) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSWmi_PnPDeviceId) GetPropertyInstanceName() (value string, err error) {
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

// SetPnPDeviceId sets the value of PnPDeviceId for the instance
func (instance *MSWmi_PnPDeviceId) SetPropertyPnPDeviceId(value string) (err error) {
	return instance.SetProperty("PnPDeviceId", (value))
}

// GetPnPDeviceId gets the value of PnPDeviceId for the instance
func (instance *MSWmi_PnPDeviceId) GetPropertyPnPDeviceId() (value string, err error) {
	retValue, err := instance.GetProperty("PnPDeviceId")
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
