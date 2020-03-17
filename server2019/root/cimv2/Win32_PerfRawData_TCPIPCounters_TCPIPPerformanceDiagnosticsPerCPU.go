// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU struct
type Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU struct {
	Win32_PerfRawData

	//
	TCPcurrentconnections uint32
}

// SetTCPcurrentconnections sets the value of TCPcurrentconnections for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU) SetPropertyTCPcurrentconnections(value uint32) (err error) {
	return instance.SetProperty("TCPcurrentconnections", value)
}

// GetTCPcurrentconnections gets the value of TCPcurrentconnections for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU) GetPropertyTCPcurrentconnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPcurrentconnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
