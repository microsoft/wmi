// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_KPSSVC struct
type Win32_PerfRawData_Counters_KPSSVC struct {
	Win32_PerfRawData

	//
	FailedRequests uint32

	//
	IncomingArmoredRequests uint32

	//
	IncomingPasswordChangeRequests uint32

	//
	IncomingRequests uint32
}

// SetFailedRequests sets the value of FailedRequests for the instance
func (instance *Win32_PerfRawData_Counters_KPSSVC) SetPropertyFailedRequests(value uint32) (err error) {
	return instance.SetProperty("FailedRequests", value)
}

// GetFailedRequests gets the value of FailedRequests for the instance
func (instance *Win32_PerfRawData_Counters_KPSSVC) GetPropertyFailedRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailedRequests")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIncomingArmoredRequests sets the value of IncomingArmoredRequests for the instance
func (instance *Win32_PerfRawData_Counters_KPSSVC) SetPropertyIncomingArmoredRequests(value uint32) (err error) {
	return instance.SetProperty("IncomingArmoredRequests", value)
}

// GetIncomingArmoredRequests gets the value of IncomingArmoredRequests for the instance
func (instance *Win32_PerfRawData_Counters_KPSSVC) GetPropertyIncomingArmoredRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("IncomingArmoredRequests")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIncomingPasswordChangeRequests sets the value of IncomingPasswordChangeRequests for the instance
func (instance *Win32_PerfRawData_Counters_KPSSVC) SetPropertyIncomingPasswordChangeRequests(value uint32) (err error) {
	return instance.SetProperty("IncomingPasswordChangeRequests", value)
}

// GetIncomingPasswordChangeRequests gets the value of IncomingPasswordChangeRequests for the instance
func (instance *Win32_PerfRawData_Counters_KPSSVC) GetPropertyIncomingPasswordChangeRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("IncomingPasswordChangeRequests")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIncomingRequests sets the value of IncomingRequests for the instance
func (instance *Win32_PerfRawData_Counters_KPSSVC) SetPropertyIncomingRequests(value uint32) (err error) {
	return instance.SetProperty("IncomingRequests", value)
}

// GetIncomingRequests gets the value of IncomingRequests for the instance
func (instance *Win32_PerfRawData_Counters_KPSSVC) GetPropertyIncomingRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("IncomingRequests")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
