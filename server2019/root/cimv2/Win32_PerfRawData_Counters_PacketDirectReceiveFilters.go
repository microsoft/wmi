// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_PacketDirectReceiveFilters struct
type Win32_PerfRawData_Counters_PacketDirectReceiveFilters struct {
	Win32_PerfRawData

	//
	BytesMatched uint64

	//
	BytesMatchedPersec uint64

	//
	PacketsMatched uint64

	//
	PacketsMatchedPersec uint64
}

// SetBytesMatched sets the value of BytesMatched for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectReceiveFilters) SetPropertyBytesMatched(value uint64) (err error) {
	return instance.SetProperty("BytesMatched", value)
}

// GetBytesMatched gets the value of BytesMatched for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectReceiveFilters) GetPropertyBytesMatched() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesMatched")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesMatchedPersec sets the value of BytesMatchedPersec for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectReceiveFilters) SetPropertyBytesMatchedPersec(value uint64) (err error) {
	return instance.SetProperty("BytesMatchedPersec", value)
}

// GetBytesMatchedPersec gets the value of BytesMatchedPersec for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectReceiveFilters) GetPropertyBytesMatchedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesMatchedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsMatched sets the value of PacketsMatched for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectReceiveFilters) SetPropertyPacketsMatched(value uint64) (err error) {
	return instance.SetProperty("PacketsMatched", value)
}

// GetPacketsMatched gets the value of PacketsMatched for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectReceiveFilters) GetPropertyPacketsMatched() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsMatched")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsMatchedPersec sets the value of PacketsMatchedPersec for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectReceiveFilters) SetPropertyPacketsMatchedPersec(value uint64) (err error) {
	return instance.SetProperty("PacketsMatchedPersec", value)
}

// GetPacketsMatchedPersec gets the value of PacketsMatchedPersec for the instance
func (instance *Win32_PerfRawData_Counters_PacketDirectReceiveFilters) GetPropertyPacketsMatchedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsMatchedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
