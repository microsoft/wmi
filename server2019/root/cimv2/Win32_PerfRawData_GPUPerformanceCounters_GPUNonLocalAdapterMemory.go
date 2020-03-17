// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory struct
type Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory struct {
	Win32_PerfRawData

	//
	NonLocalUsage uint64
}

// SetNonLocalUsage sets the value of NonLocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory) SetPropertyNonLocalUsage(value uint64) (err error) {
	return instance.SetProperty("NonLocalUsage", value)
}

// GetNonLocalUsage gets the value of NonLocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUNonLocalAdapterMemory) GetPropertyNonLocalUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonLocalUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
