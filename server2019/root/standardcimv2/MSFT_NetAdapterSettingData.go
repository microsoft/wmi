// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterSettingData struct
type MSFT_NetAdapterSettingData struct {
	*MSFT_NetSettingData

	//
	InterfaceDescription string

	//
	Name string

	//
	Source uint32

	//
	SystemName string
}

func NewMSFT_NetAdapterSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetInterfaceDescription sets the value of InterfaceDescription for the instance
func (instance *MSFT_NetAdapterSettingData) SetPropertyInterfaceDescription(value string) (err error) {
	return instance.SetProperty("InterfaceDescription", (value))
}

// GetInterfaceDescription gets the value of InterfaceDescription for the instance
func (instance *MSFT_NetAdapterSettingData) GetPropertyInterfaceDescription() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceDescription")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_NetAdapterSettingData) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_NetAdapterSettingData) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetSource sets the value of Source for the instance
func (instance *MSFT_NetAdapterSettingData) SetPropertySource(value uint32) (err error) {
	return instance.SetProperty("Source", (value))
}

// GetSource gets the value of Source for the instance
func (instance *MSFT_NetAdapterSettingData) GetPropertySource() (value uint32, err error) {
	retValue, err := instance.GetProperty("Source")
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

// SetSystemName sets the value of SystemName for the instance
func (instance *MSFT_NetAdapterSettingData) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", (value))
}

// GetSystemName gets the value of SystemName for the instance
func (instance *MSFT_NetAdapterSettingData) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
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
