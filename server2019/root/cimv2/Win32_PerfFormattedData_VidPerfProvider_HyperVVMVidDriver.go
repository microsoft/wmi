// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidDriver struct
type Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidDriver struct {
	Win32_PerfFormattedData

	//
	VidPartitions uint64
}

// SetVidPartitions sets the value of VidPartitions for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidDriver) SetPropertyVidPartitions(value uint64) (err error) {
	return instance.SetProperty("VidPartitions", value)
}

// GetVidPartitions gets the value of VidPartitions for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidDriver) GetPropertyVidPartitions() (value uint64, err error) {
	retValue, err := instance.GetProperty("VidPartitions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
