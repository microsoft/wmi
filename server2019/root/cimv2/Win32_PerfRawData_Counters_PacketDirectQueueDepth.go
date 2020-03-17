// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_PacketDirectQueueDepth struct
type Win32_PerfRawData_Counters_PacketDirectQueueDepth struct {
	Win32_PerfRawData

	//
	AverageQueueDepth uint32

	//
	PercentAverageQueueUtilization uint32
}

// SetAverageQueueDepth sets the value of AverageQueueDepth for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectQueueDepth) SetPropertyAverageQueueDepth(value uint32) (err error) {
	return instance.SetProperty("AverageQueueDepth", value)
}

// GetAverageQueueDepth gets the value of AverageQueueDepth for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectQueueDepth) GetPropertyAverageQueueDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("AverageQueueDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentAverageQueueUtilization sets the value of PercentAverageQueueUtilization for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectQueueDepth) SetPropertyPercentAverageQueueUtilization(value uint32) (err error) {
	return instance.SetProperty("PercentAverageQueueUtilization", value)
}

// GetPercentAverageQueueUtilization gets the value of PercentAverageQueueUtilization for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectQueueDepth) GetPropertyPercentAverageQueueUtilization() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentAverageQueueUtilization")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
