// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_MicrosoftWindowsMPTFCounters_MPTFInformation struct
type Win32_PerfFormattedData_MicrosoftWindowsMPTFCounters_MPTFInformation struct {
	*Win32_PerfFormattedData

	//
	Temperature uint32
}

func NewWin32_PerfFormattedData_MicrosoftWindowsMPTFCounters_MPTFInformationEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MicrosoftWindowsMPTFCounters_MPTFInformation, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MicrosoftWindowsMPTFCounters_MPTFInformation{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MicrosoftWindowsMPTFCounters_MPTFInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MicrosoftWindowsMPTFCounters_MPTFInformation, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MicrosoftWindowsMPTFCounters_MPTFInformation{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetTemperature sets the value of Temperature for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsMPTFCounters_MPTFInformation) SetPropertyTemperature(value uint32) (err error) {
	return instance.SetProperty("Temperature", (value))
}

// GetTemperature gets the value of Temperature for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsMPTFCounters_MPTFInformation) GetPropertyTemperature() (value uint32, err error) {
	retValue, err := instance.GetProperty("Temperature")
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
