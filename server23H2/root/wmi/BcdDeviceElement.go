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

// BcdDeviceElement struct
type BcdDeviceElement struct {
	*BcdElement

	// This field contains information about the device.
	Device BcdDeviceData
}

func NewBcdDeviceElementEx1(instance *cim.WmiInstance) (newInstance *BcdDeviceElement, err error) {
	tmp, err := NewBcdElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceElement{
		BcdElement: tmp,
	}
	return
}

func NewBcdDeviceElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdDeviceElement, err error) {
	tmp, err := NewBcdElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceElement{
		BcdElement: tmp,
	}
	return
}

// SetDevice sets the value of Device for the instance
func (instance *BcdDeviceElement) SetPropertyDevice(value BcdDeviceData) (err error) {
	return instance.SetProperty("Device", (value))
}

// GetDevice gets the value of Device for the instance
func (instance *BcdDeviceElement) GetPropertyDevice() (value BcdDeviceData, err error) {
	retValue, err := instance.GetProperty("Device")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(BcdDeviceData)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " BcdDeviceData is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BcdDeviceData(valuetmp)

	return
}
