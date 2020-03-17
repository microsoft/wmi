// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2 struct
type Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2 struct {
	Win32_PerfFormattedData

	//
	UnacknowledgedMessageCount uint64
}

// SetUnacknowledgedMessageCount sets the value of UnacknowledgedMessageCount for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2) SetPropertyUnacknowledgedMessageCount(value uint64) (err error) {
	return instance.SetProperty("UnacknowledgedMessageCount", value)
}

// GetUnacknowledgedMessageCount gets the value of UnacknowledgedMessageCount for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterGoodEnoughMulticastMessages2) GetPropertyUnacknowledgedMessageCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("UnacknowledgedMessageCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
