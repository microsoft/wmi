// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_PerfOS_PagingFile struct
type Win32_PerfFormattedData_PerfOS_PagingFile struct {
	Win32_PerfFormattedData

	//
	PercentUsage uint32

	//
	PercentUsagePeak uint32
}

// SetPercentUsage sets the value of PercentUsage for the instance
func (instance *Win32_PerfFormattedData_PerfOS_PagingFile) SetPropertyPercentUsage(value uint32) (err error) {
	return instance.SetProperty("PercentUsage", value)
}

// GetPercentUsage gets the value of PercentUsage for the instance
func (instance *Win32_PerfFormattedData_PerfOS_PagingFile) GetPropertyPercentUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentUsagePeak sets the value of PercentUsagePeak for the instance
func (instance *Win32_PerfFormattedData_PerfOS_PagingFile) SetPropertyPercentUsagePeak(value uint32) (err error) {
	return instance.SetProperty("PercentUsagePeak", value)
}

// GetPercentUsagePeak gets the value of PercentUsagePeak for the instance
func (instance *Win32_PerfFormattedData_PerfOS_PagingFile) GetPropertyPercentUsagePeak() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentUsagePeak")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
