// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory struct
type Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory struct {
	Win32_PerfRawData

	//
	DedicatedUsage uint64

	//
	LocalUsage uint64

	//
	NonLocalUsage uint64

	//
	SharedUsage uint64

	//
	TotalCommitted uint64
}

// SetDedicatedUsage sets the value of DedicatedUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory) SetPropertyDedicatedUsage(value uint64) (err error) {
	return instance.SetProperty("DedicatedUsage", value)
}

// GetDedicatedUsage gets the value of DedicatedUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory) GetPropertyDedicatedUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("DedicatedUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalUsage sets the value of LocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory) SetPropertyLocalUsage(value uint64) (err error) {
	return instance.SetProperty("LocalUsage", value)
}

// GetLocalUsage gets the value of LocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory) GetPropertyLocalUsage() (value uint64, err error) {
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

// SetNonLocalUsage sets the value of NonLocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory) SetPropertyNonLocalUsage(value uint64) (err error) {
	return instance.SetProperty("NonLocalUsage", value)
}

// GetNonLocalUsage gets the value of NonLocalUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory) GetPropertyNonLocalUsage() (value uint64, err error) {
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

// SetSharedUsage sets the value of SharedUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory) SetPropertySharedUsage(value uint64) (err error) {
	return instance.SetProperty("SharedUsage", value)
}

// GetSharedUsage gets the value of SharedUsage for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory) GetPropertySharedUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("SharedUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalCommitted sets the value of TotalCommitted for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory) SetPropertyTotalCommitted(value uint64) (err error) {
	return instance.SetProperty("TotalCommitted", value)
}

// GetTotalCommitted gets the value of TotalCommitted for the instance
func (instance *Win32_PerfRawData_GPUPerformanceCounters_GPUProcessMemory) GetPropertyTotalCommitted() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalCommitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
