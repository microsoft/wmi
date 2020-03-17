// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_PerfOS_NUMANodeMemory struct
type Win32_PerfFormattedData_PerfOS_NUMANodeMemory struct {
	Win32_PerfFormattedData

	//
	FreeAndZeroPageListMBytes uint32

	//
	TotalMBytes uint32
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
