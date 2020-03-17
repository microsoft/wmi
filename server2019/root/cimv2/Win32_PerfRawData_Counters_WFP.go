// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_WFP struct
type Win32_PerfRawData_Counters_WFP struct {
	Win32_PerfRawData

	//
	ProviderCount uint32
}

// SetProviderCount sets the value of ProviderCount for the instance
func (instance *Win32_PerfRawData_Counters_WFP) SetPropertyProviderCount(value uint32) (err error) {
	return instance.SetProperty("ProviderCount", value)
}

// GetProviderCount gets the value of ProviderCount for the instance
func (instance *Win32_PerfRawData_Counters_WFP) GetPropertyProviderCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProviderCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
