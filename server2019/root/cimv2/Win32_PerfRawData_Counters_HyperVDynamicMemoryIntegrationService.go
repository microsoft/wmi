// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_HyperVDynamicMemoryIntegrationService struct
type Win32_PerfRawData_Counters_HyperVDynamicMemoryIntegrationService struct {
	Win32_PerfRawData

	//
	MaximumMemoryMbytes uint64
}

// SetMaximumMemoryMbytes sets the value of MaximumMemoryMbytes for the instance
func (instance *Win32_PerfRawData_Counters_HyperVDynamicMemoryIntegrationService) SetPropertyMaximumMemoryMbytes(value uint64) (err error) {
	return instance.SetProperty("MaximumMemoryMbytes", value)
}

// GetMaximumMemoryMbytes gets the value of MaximumMemoryMbytes for the instance
func (instance *Win32_PerfRawData_Counters_HyperVDynamicMemoryIntegrationService) GetPropertyMaximumMemoryMbytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaximumMemoryMbytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
