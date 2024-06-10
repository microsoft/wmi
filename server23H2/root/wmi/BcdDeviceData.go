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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BcdDeviceData struct
type BcdDeviceData struct {
	*cim.WmiInstance

	// This represents the additional options for the element.
	AdditionalOptions string

	// This identifies the type of device element. This value dictates whether this is a file device element or a partition device element.
	DeviceType BcdDeviceData_DeviceType
}

func NewBcdDeviceDataEx1(instance *cim.WmiInstance) (newInstance *BcdDeviceData, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BcdDeviceData{
		WmiInstance: tmp,
	}
	return
}

func NewBcdDeviceDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BcdDeviceData, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BcdDeviceData{
		WmiInstance: tmp,
	}
	return
}

// SetAdditionalOptions sets the value of AdditionalOptions for the instance
func (instance *BcdDeviceData) SetPropertyAdditionalOptions(value string) (err error) {
	return instance.SetProperty("AdditionalOptions", (value))
}

// GetAdditionalOptions gets the value of AdditionalOptions for the instance
func (instance *BcdDeviceData) GetPropertyAdditionalOptions() (value string, err error) {
	retValue, err := instance.GetProperty("AdditionalOptions")
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

// SetDeviceType sets the value of DeviceType for the instance
func (instance *BcdDeviceData) SetPropertyDeviceType(value BcdDeviceData_DeviceType) (err error) {
	return instance.SetProperty("DeviceType", (value))
}

// GetDeviceType gets the value of DeviceType for the instance
func (instance *BcdDeviceData) GetPropertyDeviceType() (value BcdDeviceData_DeviceType, err error) {
	retValue, err := instance.GetProperty("DeviceType")
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

	value = BcdDeviceData_DeviceType(valuetmp)

	return
}
