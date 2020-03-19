// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PerfFormattedData_PerfOS_NUMANodeMemory struct
type Win32_PerfFormattedData_PerfOS_NUMANodeMemory struct {
	*Win32_PerfFormattedData

	//
	FreeAndZeroPageListMBytes uint32

	//
	TotalMBytes uint32
}

func NewWin32_PerfFormattedData_PerfOS_NUMANodeMemoryEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_PerfOS_NUMANodeMemory, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_PerfOS_NUMANodeMemory{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_PerfOS_NUMANodeMemoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_PerfOS_NUMANodeMemory, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_PerfOS_NUMANodeMemory{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetFreeAndZeroPageListMBytes sets the value of FreeAndZeroPageListMBytes for the instance
func (instance *Win32_PerfFormattedData_PerfOS_NUMANodeMemory) SetPropertyFreeAndZeroPageListMBytes(value uint32) (err error) {
	return instance.SetProperty("FreeAndZeroPageListMBytes", value)
}

// GetFreeAndZeroPageListMBytes gets the value of FreeAndZeroPageListMBytes for the instance
func (instance *Win32_PerfFormattedData_PerfOS_NUMANodeMemory) GetPropertyFreeAndZeroPageListMBytes() (value uint32, err error) {
	retValue, err := instance.GetProperty("FreeAndZeroPageListMBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalMBytes sets the value of TotalMBytes for the instance
func (instance *Win32_PerfFormattedData_PerfOS_NUMANodeMemory) SetPropertyTotalMBytes(value uint32) (err error) {
	return instance.SetProperty("TotalMBytes", value)
}

// GetTotalMBytes gets the value of TotalMBytes for the instance
func (instance *Win32_PerfFormattedData_PerfOS_NUMANodeMemory) GetPropertyTotalMBytes() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalMBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
