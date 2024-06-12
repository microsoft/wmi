// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PowerSettingDefinitionRangeData struct
type Win32_PowerSettingDefinitionRangeData struct {
	*CIM_SettingData

	//
	SettingValue uint32
}

func NewWin32_PowerSettingDefinitionRangeDataEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSettingDefinitionRangeData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingDefinitionRangeData{
		CIM_SettingData: tmp,
	}
	return
}

func NewWin32_PowerSettingDefinitionRangeDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PowerSettingDefinitionRangeData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingDefinitionRangeData{
		CIM_SettingData: tmp,
	}
	return
}

// SetSettingValue sets the value of SettingValue for the instance
func (instance *Win32_PowerSettingDefinitionRangeData) SetPropertySettingValue(value uint32) (err error) {
	return instance.SetProperty("SettingValue", (value))
}

// GetSettingValue gets the value of SettingValue for the instance
func (instance *Win32_PowerSettingDefinitionRangeData) GetPropertySettingValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("SettingValue")
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
