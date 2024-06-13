// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_BootSourceSettingData struct
type Msvm_BootSourceSettingData struct {
	*CIM_SettingData

	//
	BootSourceDescription string

	//
	BootSourceType uint32

	//
	FirmwareDevicePath string

	//
	OptionalData []uint8

	//
	OtherLocation string
}

func NewMsvm_BootSourceSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_BootSourceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_BootSourceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_BootSourceSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_BootSourceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_BootSourceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetBootSourceDescription sets the value of BootSourceDescription for the instance
func (instance *Msvm_BootSourceSettingData) SetPropertyBootSourceDescription(value string) (err error) {
	return instance.SetProperty("BootSourceDescription", (value))
}

// GetBootSourceDescription gets the value of BootSourceDescription for the instance
func (instance *Msvm_BootSourceSettingData) GetPropertyBootSourceDescription() (value string, err error) {
	retValue, err := instance.GetProperty("BootSourceDescription")
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

// SetBootSourceType sets the value of BootSourceType for the instance
func (instance *Msvm_BootSourceSettingData) SetPropertyBootSourceType(value uint32) (err error) {
	return instance.SetProperty("BootSourceType", (value))
}

// GetBootSourceType gets the value of BootSourceType for the instance
func (instance *Msvm_BootSourceSettingData) GetPropertyBootSourceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("BootSourceType")
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

// SetFirmwareDevicePath sets the value of FirmwareDevicePath for the instance
func (instance *Msvm_BootSourceSettingData) SetPropertyFirmwareDevicePath(value string) (err error) {
	return instance.SetProperty("FirmwareDevicePath", (value))
}

// GetFirmwareDevicePath gets the value of FirmwareDevicePath for the instance
func (instance *Msvm_BootSourceSettingData) GetPropertyFirmwareDevicePath() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareDevicePath")
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

// SetOptionalData sets the value of OptionalData for the instance
func (instance *Msvm_BootSourceSettingData) SetPropertyOptionalData(value []uint8) (err error) {
	return instance.SetProperty("OptionalData", (value))
}

// GetOptionalData gets the value of OptionalData for the instance
func (instance *Msvm_BootSourceSettingData) GetPropertyOptionalData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("OptionalData")
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

// SetOtherLocation sets the value of OtherLocation for the instance
func (instance *Msvm_BootSourceSettingData) SetPropertyOtherLocation(value string) (err error) {
	return instance.SetProperty("OtherLocation", (value))
}

// GetOtherLocation gets the value of OtherLocation for the instance
func (instance *Msvm_BootSourceSettingData) GetPropertyOtherLocation() (value string, err error) {
	retValue, err := instance.GetProperty("OtherLocation")
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
