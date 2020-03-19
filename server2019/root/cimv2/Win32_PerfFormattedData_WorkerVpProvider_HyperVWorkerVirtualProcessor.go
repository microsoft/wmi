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

// Win32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessor struct
type Win32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessor struct {
	*Win32_PerfFormattedData

	//
	InterceptDelayTimems uint64

	//
	InterceptsDelayed uint64
}

func NewWin32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessorEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessor, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessor{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessor, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessor{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetInterceptDelayTimems sets the value of InterceptDelayTimems for the instance
func (instance *Win32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessor) SetPropertyInterceptDelayTimems(value uint64) (err error) {
	return instance.SetProperty("InterceptDelayTimems", value)
}

// GetInterceptDelayTimems gets the value of InterceptDelayTimems for the instance
func (instance *Win32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessor) GetPropertyInterceptDelayTimems() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterceptDelayTimems")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterceptsDelayed sets the value of InterceptsDelayed for the instance
func (instance *Win32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessor) SetPropertyInterceptsDelayed(value uint64) (err error) {
	return instance.SetProperty("InterceptsDelayed", value)
}

// GetInterceptsDelayed gets the value of InterceptsDelayed for the instance
func (instance *Win32_PerfFormattedData_WorkerVpProvider_HyperVWorkerVirtualProcessor) GetPropertyInterceptsDelayed() (value uint64, err error) {
	retValue, err := instance.GetProperty("InterceptsDelayed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
