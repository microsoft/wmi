// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MLNX_PCIDeviceSettingData struct
type MLNX_PCIDeviceSettingData struct {
	*CIM_SettingData

	//
	Name string

	//
	Source PCIDeviceSettingData_Source

	//
	SystemName string
}

func NewMLNX_PCIDeviceSettingDataEx1(instance *cim.WmiInstance) (newInstance *MLNX_PCIDeviceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_PCIDeviceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMLNX_PCIDeviceSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_PCIDeviceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_PCIDeviceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MLNX_PCIDeviceSettingData) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MLNX_PCIDeviceSettingData) GetPropertyName() (value string, err error) {
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
func (instance *MLNX_PCIDeviceSettingData) SetPropertySource(value PCIDeviceSettingData_Source) (err error) {
	return instance.SetProperty("Source", (value))
}

// GetSource gets the value of Source for the instance
func (instance *MLNX_PCIDeviceSettingData) GetPropertySource() (value PCIDeviceSettingData_Source, err error) {
	retValue, err := instance.GetProperty("Source")
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

	value = PCIDeviceSettingData_Source(valuetmp)

	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *MLNX_PCIDeviceSettingData) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", (value))
}

// GetSystemName gets the value of SystemName for the instance
func (instance *MLNX_PCIDeviceSettingData) GetPropertySystemName() (value string, err error) {
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
