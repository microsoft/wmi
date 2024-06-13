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

// SystemConfig_V0_PnP struct
type SystemConfig_V0_PnP struct {
	*SystemConfig_V0

	//
	DescriptionLength uint32

	//
	DeviceDescription string

	//
	DeviceID string

	//
	FriendlyName string

	//
	FriendlyNameLength uint32

	//
	IDLength uint32
}

func NewSystemConfig_V0_PnPEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V0_PnP, err error) {
	tmp, err := NewSystemConfig_V0Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V0_PnP{
		SystemConfig_V0: tmp,
	}
	return
}

func NewSystemConfig_V0_PnPEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V0_PnP, err error) {
	tmp, err := NewSystemConfig_V0Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V0_PnP{
		SystemConfig_V0: tmp,
	}
	return
}

// SetDescriptionLength sets the value of DescriptionLength for the instance
func (instance *SystemConfig_V0_PnP) SetPropertyDescriptionLength(value uint32) (err error) {
	return instance.SetProperty("DescriptionLength", (value))
}

// GetDescriptionLength gets the value of DescriptionLength for the instance
func (instance *SystemConfig_V0_PnP) GetPropertyDescriptionLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("DescriptionLength")
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

// SetDeviceDescription sets the value of DeviceDescription for the instance
func (instance *SystemConfig_V0_PnP) SetPropertyDeviceDescription(value string) (err error) {
	return instance.SetProperty("DeviceDescription", (value))
}

// GetDeviceDescription gets the value of DeviceDescription for the instance
func (instance *SystemConfig_V0_PnP) GetPropertyDeviceDescription() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceDescription")
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

// SetDeviceID sets the value of DeviceID for the instance
func (instance *SystemConfig_V0_PnP) SetPropertyDeviceID(value string) (err error) {
	return instance.SetProperty("DeviceID", (value))
}

// GetDeviceID gets the value of DeviceID for the instance
func (instance *SystemConfig_V0_PnP) GetPropertyDeviceID() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceID")
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

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *SystemConfig_V0_PnP) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *SystemConfig_V0_PnP) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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

// SetFriendlyNameLength sets the value of FriendlyNameLength for the instance
func (instance *SystemConfig_V0_PnP) SetPropertyFriendlyNameLength(value uint32) (err error) {
	return instance.SetProperty("FriendlyNameLength", (value))
}

// GetFriendlyNameLength gets the value of FriendlyNameLength for the instance
func (instance *SystemConfig_V0_PnP) GetPropertyFriendlyNameLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("FriendlyNameLength")
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

// SetIDLength sets the value of IDLength for the instance
func (instance *SystemConfig_V0_PnP) SetPropertyIDLength(value uint32) (err error) {
	return instance.SetProperty("IDLength", (value))
}

// GetIDLength gets the value of IDLength for the instance
func (instance *SystemConfig_V0_PnP) GetPropertyIDLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("IDLength")
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
