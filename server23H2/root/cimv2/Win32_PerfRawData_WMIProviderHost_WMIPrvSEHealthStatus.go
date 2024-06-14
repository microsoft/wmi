// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus struct
type Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus struct {
	*Win32_PerfRawData

	//
	CPUPercentConsumption uint64

	//
	HandleCount uint32

	//
	PrivatePageMemory uint64

	//
	ProcessID uint32

	//
	ThreadCount uint32

	//
	TimestampOfHealthStatus uint64
}

func NewWin32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatusEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetCPUPercentConsumption sets the value of CPUPercentConsumption for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) SetPropertyCPUPercentConsumption(value uint64) (err error) {
	return instance.SetProperty("CPUPercentConsumption", (value))
}

// GetCPUPercentConsumption gets the value of CPUPercentConsumption for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) GetPropertyCPUPercentConsumption() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUPercentConsumption")
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

// SetHandleCount sets the value of HandleCount for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) SetPropertyHandleCount(value uint32) (err error) {
	return instance.SetProperty("HandleCount", (value))
}

// GetHandleCount gets the value of HandleCount for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) GetPropertyHandleCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("HandleCount")
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

// SetPrivatePageMemory sets the value of PrivatePageMemory for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) SetPropertyPrivatePageMemory(value uint64) (err error) {
	return instance.SetProperty("PrivatePageMemory", (value))
}

// GetPrivatePageMemory gets the value of PrivatePageMemory for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) GetPropertyPrivatePageMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("PrivatePageMemory")
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

// SetProcessID sets the value of ProcessID for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) SetPropertyProcessID(value uint32) (err error) {
	return instance.SetProperty("ProcessID", (value))
}

// GetProcessID gets the value of ProcessID for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) GetPropertyProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessID")
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

// SetThreadCount sets the value of ThreadCount for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) SetPropertyThreadCount(value uint32) (err error) {
	return instance.SetProperty("ThreadCount", (value))
}

// GetThreadCount gets the value of ThreadCount for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) GetPropertyThreadCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadCount")
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

// SetTimestampOfHealthStatus sets the value of TimestampOfHealthStatus for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) SetPropertyTimestampOfHealthStatus(value uint64) (err error) {
	return instance.SetProperty("TimestampOfHealthStatus", (value))
}

// GetTimestampOfHealthStatus gets the value of TimestampOfHealthStatus for the instance
func (instance *Win32_PerfRawData_WMIProviderHost_WMIPrvSEHealthStatus) GetPropertyTimestampOfHealthStatus() (value uint64, err error) {
	retValue, err := instance.GetProperty("TimestampOfHealthStatus")
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
