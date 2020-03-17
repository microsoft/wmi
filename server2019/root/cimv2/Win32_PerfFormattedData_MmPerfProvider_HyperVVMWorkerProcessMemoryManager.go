// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_MmPerfProvider_HyperVVMWorkerProcessMemoryManager struct
type Win32_PerfFormattedData_MmPerfProvider_HyperVVMWorkerProcessMemoryManager struct {
	Win32_PerfFormattedData

	//
	MemoryBlockCount uint64
}

// SetMemoryBlockCount sets the value of MemoryBlockCount for the instance
func (instance *Win32_PerfFormattedData_MmPerfProvider_HyperVVMWorkerProcessMemoryManager) SetPropertyMemoryBlockCount(value uint64) (err error) {
	return instance.SetProperty("MemoryBlockCount", value)
}

// GetMemoryBlockCount gets the value of MemoryBlockCount for the instance
func (instance *Win32_PerfFormattedData_MmPerfProvider_HyperVVMWorkerProcessMemoryManager) GetPropertyMemoryBlockCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryBlockCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
