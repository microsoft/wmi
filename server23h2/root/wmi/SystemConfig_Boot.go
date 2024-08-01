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

// SystemConfig_Boot struct
type SystemConfig_Boot struct {
	*SystemConfig_V2

	//
	BootFlags uint64

	//
	FirmwareType uint32

	//
	Reserved1 uint8

	//
	Reserved2 uint8

	//
	SecureBootCapable uint8

	//
	SecureBootEnabled uint8
}

func NewSystemConfig_BootEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_Boot, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_Boot{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_BootEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_Boot, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_Boot{
		SystemConfig_V2: tmp,
	}
	return
}

// SetBootFlags sets the value of BootFlags for the instance
func (instance *SystemConfig_Boot) SetPropertyBootFlags(value uint64) (err error) {
	return instance.SetProperty("BootFlags", (value))
}

// GetBootFlags gets the value of BootFlags for the instance
func (instance *SystemConfig_Boot) GetPropertyBootFlags() (value uint64, err error) {
	retValue, err := instance.GetProperty("BootFlags")
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

// SetFirmwareType sets the value of FirmwareType for the instance
func (instance *SystemConfig_Boot) SetPropertyFirmwareType(value uint32) (err error) {
	return instance.SetProperty("FirmwareType", (value))
}

// GetFirmwareType gets the value of FirmwareType for the instance
func (instance *SystemConfig_Boot) GetPropertyFirmwareType() (value uint32, err error) {
	retValue, err := instance.GetProperty("FirmwareType")
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

// SetReserved1 sets the value of Reserved1 for the instance
func (instance *SystemConfig_Boot) SetPropertyReserved1(value uint8) (err error) {
	return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *SystemConfig_Boot) GetPropertyReserved1() (value uint8, err error) {
	retValue, err := instance.GetProperty("Reserved1")
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

// SetReserved2 sets the value of Reserved2 for the instance
func (instance *SystemConfig_Boot) SetPropertyReserved2(value uint8) (err error) {
	return instance.SetProperty("Reserved2", (value))
}

// GetReserved2 gets the value of Reserved2 for the instance
func (instance *SystemConfig_Boot) GetPropertyReserved2() (value uint8, err error) {
	retValue, err := instance.GetProperty("Reserved2")
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

// SetSecureBootCapable sets the value of SecureBootCapable for the instance
func (instance *SystemConfig_Boot) SetPropertySecureBootCapable(value uint8) (err error) {
	return instance.SetProperty("SecureBootCapable", (value))
}

// GetSecureBootCapable gets the value of SecureBootCapable for the instance
func (instance *SystemConfig_Boot) GetPropertySecureBootCapable() (value uint8, err error) {
	retValue, err := instance.GetProperty("SecureBootCapable")
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

// SetSecureBootEnabled sets the value of SecureBootEnabled for the instance
func (instance *SystemConfig_Boot) SetPropertySecureBootEnabled(value uint8) (err error) {
	return instance.SetProperty("SecureBootEnabled", (value))
}

// GetSecureBootEnabled gets the value of SecureBootEnabled for the instance
func (instance *SystemConfig_Boot) GetPropertySecureBootEnabled() (value uint8, err error) {
	retValue, err := instance.GetProperty("SecureBootEnabled")
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
