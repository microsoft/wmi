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

// SystemConfig_V2_DeviceFamily struct
type SystemConfig_V2_DeviceFamily struct {
	*SystemConfig_V2

	//
	DeviceFamily uint32

	//
	DeviceForm uint32

	//
	UAPInfo uint64
}

func NewSystemConfig_V2_DeviceFamilyEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_DeviceFamily, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_DeviceFamily{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_DeviceFamilyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_DeviceFamily, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_DeviceFamily{
		SystemConfig_V2: tmp,
	}
	return
}

// SetDeviceFamily sets the value of DeviceFamily for the instance
func (instance *SystemConfig_V2_DeviceFamily) SetPropertyDeviceFamily(value uint32) (err error) {
	return instance.SetProperty("DeviceFamily", (value))
}

// GetDeviceFamily gets the value of DeviceFamily for the instance
func (instance *SystemConfig_V2_DeviceFamily) GetPropertyDeviceFamily() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceFamily")
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

// SetDeviceForm sets the value of DeviceForm for the instance
func (instance *SystemConfig_V2_DeviceFamily) SetPropertyDeviceForm(value uint32) (err error) {
	return instance.SetProperty("DeviceForm", (value))
}

// GetDeviceForm gets the value of DeviceForm for the instance
func (instance *SystemConfig_V2_DeviceFamily) GetPropertyDeviceForm() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceForm")
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

// SetUAPInfo sets the value of UAPInfo for the instance
func (instance *SystemConfig_V2_DeviceFamily) SetPropertyUAPInfo(value uint64) (err error) {
	return instance.SetProperty("UAPInfo", (value))
}

// GetUAPInfo gets the value of UAPInfo for the instance
func (instance *SystemConfig_V2_DeviceFamily) GetPropertyUAPInfo() (value uint64, err error) {
	retValue, err := instance.GetProperty("UAPInfo")
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
