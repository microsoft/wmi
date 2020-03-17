// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_VidPerfProvider_HyperVVMVidNumaNode struct
type Win32_PerfRawData_VidPerfProvider_HyperVVMVidNumaNode struct {
	Win32_PerfRawData

	//
	PageCount uint64

	//
	ProcessorCount uint64
}

// SetPageCount sets the value of PageCount for the instance
func (instance *Win32_PerfRawData_VidPerfProvider_HyperVVMVidNumaNode) SetPropertyPageCount(value uint64) (err error) {
	return instance.SetProperty("PageCount", value)
}

// GetPageCount gets the value of PageCount for the instance
func (instance *Win32_PerfRawData_VidPerfProvider_HyperVVMVidNumaNode) GetPropertyPageCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("PageCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorCount sets the value of ProcessorCount for the instance
func (instance *Win32_PerfRawData_VidPerfProvider_HyperVVMVidNumaNode) SetPropertyProcessorCount(value uint64) (err error) {
	return instance.SetProperty("ProcessorCount", value)
}

// GetProcessorCount gets the value of ProcessorCount for the instance
func (instance *Win32_PerfRawData_VidPerfProvider_HyperVVMVidNumaNode) GetPropertyProcessorCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ProcessorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
