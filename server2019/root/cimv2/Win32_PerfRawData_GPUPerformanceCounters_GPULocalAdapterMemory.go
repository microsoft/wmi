// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory struct
type Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory struct {
	Win32_PerfRawData

	//
	LocalUsage uint64
}

// SetLocalUsage sets the value of LocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory) SetPropertyLocalUsage(value uint64) (err error) {
	return instance.SetProperty("LocalUsage", value)
}

// GetLocalUsage gets the value of LocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPULocalAdapterMemory) GetPropertyLocalUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("LocalUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
