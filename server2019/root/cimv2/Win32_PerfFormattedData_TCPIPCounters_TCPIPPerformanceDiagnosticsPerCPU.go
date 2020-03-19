// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU struct
type Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU struct {
	*Win32_PerfFormattedData

	//
	TCPcurrentconnections uint32
}

func NewWin32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPUEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPUEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetTCPcurrentconnections sets the value of TCPcurrentconnections for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU) SetPropertyTCPcurrentconnections(value uint32) (err error) {
	return instance.SetProperty("TCPcurrentconnections", value)
}

// GetTCPcurrentconnections gets the value of TCPcurrentconnections for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsPerCPU) GetPropertyTCPcurrentconnections() (value uint32, err error) {
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
