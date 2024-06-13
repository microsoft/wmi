// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime struct
type Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime struct {
	*Win32_PerfRawData

	//
	RuntimeCount16ms uint32

	//
	RuntimeCount16s uint32

	//
	RuntimeCount1min uint32

	//
	RuntimeCount1s uint32

	//
	RuntimeCount256ms uint32

	//
	RuntimeCount4ms uint32

	//
	RuntimeCount4s uint32

	//
	RuntimeCount64ms uint32

	//
	RuntimeCountInfinite uint32
}

func NewWin32_PerfRawData_Counters_StorageManagementWSPSpacesRuntimeEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_StorageManagementWSPSpacesRuntimeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetRuntimeCount16ms sets the value of RuntimeCount16ms for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) SetPropertyRuntimeCount16ms(value uint32) (err error) {
	return instance.SetProperty("RuntimeCount16ms", (value))
}

// GetRuntimeCount16ms gets the value of RuntimeCount16ms for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) GetPropertyRuntimeCount16ms() (value uint32, err error) {
	retValue, err := instance.GetProperty("RuntimeCount16ms")
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

// SetRuntimeCount16s sets the value of RuntimeCount16s for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) SetPropertyRuntimeCount16s(value uint32) (err error) {
	return instance.SetProperty("RuntimeCount16s", (value))
}

// GetRuntimeCount16s gets the value of RuntimeCount16s for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) GetPropertyRuntimeCount16s() (value uint32, err error) {
	retValue, err := instance.GetProperty("RuntimeCount16s")
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

// SetRuntimeCount1min sets the value of RuntimeCount1min for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) SetPropertyRuntimeCount1min(value uint32) (err error) {
	return instance.SetProperty("RuntimeCount1min", (value))
}

// GetRuntimeCount1min gets the value of RuntimeCount1min for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) GetPropertyRuntimeCount1min() (value uint32, err error) {
	retValue, err := instance.GetProperty("RuntimeCount1min")
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

// SetRuntimeCount1s sets the value of RuntimeCount1s for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) SetPropertyRuntimeCount1s(value uint32) (err error) {
	return instance.SetProperty("RuntimeCount1s", (value))
}

// GetRuntimeCount1s gets the value of RuntimeCount1s for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) GetPropertyRuntimeCount1s() (value uint32, err error) {
	retValue, err := instance.GetProperty("RuntimeCount1s")
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

// SetRuntimeCount256ms sets the value of RuntimeCount256ms for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) SetPropertyRuntimeCount256ms(value uint32) (err error) {
	return instance.SetProperty("RuntimeCount256ms", (value))
}

// GetRuntimeCount256ms gets the value of RuntimeCount256ms for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) GetPropertyRuntimeCount256ms() (value uint32, err error) {
	retValue, err := instance.GetProperty("RuntimeCount256ms")
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

// SetRuntimeCount4ms sets the value of RuntimeCount4ms for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) SetPropertyRuntimeCount4ms(value uint32) (err error) {
	return instance.SetProperty("RuntimeCount4ms", (value))
}

// GetRuntimeCount4ms gets the value of RuntimeCount4ms for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) GetPropertyRuntimeCount4ms() (value uint32, err error) {
	retValue, err := instance.GetProperty("RuntimeCount4ms")
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

// SetRuntimeCount4s sets the value of RuntimeCount4s for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) SetPropertyRuntimeCount4s(value uint32) (err error) {
	return instance.SetProperty("RuntimeCount4s", (value))
}

// GetRuntimeCount4s gets the value of RuntimeCount4s for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) GetPropertyRuntimeCount4s() (value uint32, err error) {
	retValue, err := instance.GetProperty("RuntimeCount4s")
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

// SetRuntimeCount64ms sets the value of RuntimeCount64ms for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) SetPropertyRuntimeCount64ms(value uint32) (err error) {
	return instance.SetProperty("RuntimeCount64ms", (value))
}

// GetRuntimeCount64ms gets the value of RuntimeCount64ms for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) GetPropertyRuntimeCount64ms() (value uint32, err error) {
	retValue, err := instance.GetProperty("RuntimeCount64ms")
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

// SetRuntimeCountInfinite sets the value of RuntimeCountInfinite for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) SetPropertyRuntimeCountInfinite(value uint32) (err error) {
	return instance.SetProperty("RuntimeCountInfinite", (value))
}

// GetRuntimeCountInfinite gets the value of RuntimeCountInfinite for the instance
func (instance *Win32_PerfRawData_Counters_StorageManagementWSPSpacesRuntime) GetPropertyRuntimeCountInfinite() (value uint32, err error) {
	retValue, err := instance.GetProperty("RuntimeCountInfinite")
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
