// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters struct
type Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters struct {
	Win32_PerfFormattedData

	//
	BytesReceived uint64

	//
	BytesReceivedPersec uint64

	//
	PacketsDropped uint64

	//
	PacketsDroppedPersec uint64

	//
	PacketsReceived uint64

	//
	PacketsReceivedPersec uint64
}

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) SetPropertyBytesReceived(value uint64) (err error) {
	return instance.SetProperty("BytesReceived", value)
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) GetPropertyBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesReceivedPersec sets the value of BytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) SetPropertyBytesReceivedPersec(value uint64) (err error) {
	return instance.SetProperty("BytesReceivedPersec", value)
}

// GetBytesReceivedPersec gets the value of BytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) GetPropertyBytesReceivedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceivedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsDropped sets the value of PacketsDropped for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) SetPropertyPacketsDropped(value uint64) (err error) {
	return instance.SetProperty("PacketsDropped", value)
}

// GetPacketsDropped gets the value of PacketsDropped for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) GetPropertyPacketsDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsDroppedPersec sets the value of PacketsDroppedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) SetPropertyPacketsDroppedPersec(value uint64) (err error) {
	return instance.SetProperty("PacketsDroppedPersec", value)
}

// GetPacketsDroppedPersec gets the value of PacketsDroppedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) GetPropertyPacketsDroppedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsDroppedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsReceived sets the value of PacketsReceived for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) SetPropertyPacketsReceived(value uint64) (err error) {
	return instance.SetProperty("PacketsReceived", value)
}

// GetPacketsReceived gets the value of PacketsReceived for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) GetPropertyPacketsReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsReceivedPersec sets the value of PacketsReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) SetPropertyPacketsReceivedPersec(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedPersec", value)
}

// GetPacketsReceivedPersec gets the value of PacketsReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_PacketDirectReceiveCounters) GetPropertyPacketsReceivedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
