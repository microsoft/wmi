// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP struct
type Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP struct {
	Win32_PerfFormattedData

	//
	DroppedDatagrams uint32

	//
	DroppedDatagramsPersec uint32

	//
	RejectedConnections uint32

	//
	RejectedConnectionsPersec uint32
}

// SetDroppedDatagrams sets the value of DroppedDatagrams for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) SetPropertyDroppedDatagrams(value uint32) (err error) {
	return instance.SetProperty("DroppedDatagrams", value)
}

// GetDroppedDatagrams gets the value of DroppedDatagrams for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) GetPropertyDroppedDatagrams() (value uint32, err error) {
	retValue, err := instance.GetProperty("DroppedDatagrams")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDroppedDatagramsPersec sets the value of DroppedDatagramsPersec for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) SetPropertyDroppedDatagramsPersec(value uint32) (err error) {
	return instance.SetProperty("DroppedDatagramsPersec", value)
}

// GetDroppedDatagramsPersec gets the value of DroppedDatagramsPersec for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) GetPropertyDroppedDatagramsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DroppedDatagramsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRejectedConnections sets the value of RejectedConnections for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) SetPropertyRejectedConnections(value uint32) (err error) {
	return instance.SetProperty("RejectedConnections", value)
}

// GetRejectedConnections gets the value of RejectedConnections for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) GetPropertyRejectedConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("RejectedConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRejectedConnectionsPersec sets the value of RejectedConnectionsPersec for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) SetPropertyRejectedConnectionsPersec(value uint32) (err error) {
	return instance.SetProperty("RejectedConnectionsPersec", value)
}

// GetRejectedConnectionsPersec gets the value of RejectedConnectionsPersec for the instance
func (instance *Win32_PerfFormattedData_AFDCounters_MicrosoftWinsockBSP) GetPropertyRejectedConnectionsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("RejectedConnectionsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
