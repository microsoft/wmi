// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_VFPQoSQueueTotalOutboundNetworkTraffic struct
type Win32_PerfRawData_Counters_VFPQoSQueueTotalOutboundNetworkTraffic struct {
	Win32_PerfRawData

	//
	TotalOutboundBytesDropped uint64

	//
	TotalOutboundPacketsDropped uint64
}

// SetTotalOutboundBytesDropped sets the value of TotalOutboundBytesDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueTotalOutboundNetworkTraffic) SetPropertyTotalOutboundBytesDropped(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundBytesDropped", value)
}

// GetTotalOutboundBytesDropped gets the value of TotalOutboundBytesDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueTotalOutboundNetworkTraffic) GetPropertyTotalOutboundBytesDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundBytesDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalOutboundPacketsDropped sets the value of TotalOutboundPacketsDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueTotalOutboundNetworkTraffic) SetPropertyTotalOutboundPacketsDropped(value uint64) (err error) {
	return instance.SetProperty("TotalOutboundPacketsDropped", value)
}

// GetTotalOutboundPacketsDropped gets the value of TotalOutboundPacketsDropped for the instance
func (instance *Win32_PerfRawData_Counters_VFPQoSQueueTotalOutboundNetworkTraffic) GetPropertyTotalOutboundPacketsDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalOutboundPacketsDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
