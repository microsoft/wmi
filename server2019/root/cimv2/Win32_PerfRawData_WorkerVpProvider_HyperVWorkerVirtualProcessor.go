// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor struct
type Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor struct {
	Win32_PerfRawData

	//
	InterceptDelayTimems uint64

	//
	InterceptsDelayed uint64
}

// SetInterceptDelayTimems sets the value of InterceptDelayTimems for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) SetPropertyInterceptDelayTimems(value uint64) (err error) {
	return instance.SetProperty("InterceptDelayTimems", value)
}

// GetInterceptDelayTimems gets the value of InterceptDelayTimems for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) GetPropertyInterceptDelayTimems() (value uint64, err error) {
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
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) SetPropertyInterceptsDelayed(value uint64) (err error) {
	return instance.SetProperty("InterceptsDelayed", value)
}

// GetInterceptsDelayed gets the value of InterceptsDelayed for the instance
func (instance *Win32_PerfRawData_WorkerVpProvider_HyperVWorkerVirtualProcessor) GetPropertyInterceptsDelayed() (value uint64, err error) {
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
