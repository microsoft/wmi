// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_RemoteDesktopConnectionBrokerRedirectorPerformanceCounterProvider_RemoteDesktopConnectionBrokerRedirectorCounterset struct
type Win32_PerfRawData_RemoteDesktopConnectionBrokerRedirectorPerformanceCounterProvider_RemoteDesktopConnectionBrokerRedirectorCounterset struct {
	Win32_PerfRawData

	//
	Connectiontime uint64

	//
	Contextacquisitionwaittime uint64

	//
	RPCContext uint64

	//
	ThreadswaitingforRPCContext uint64
}

// SetConnectiontime sets the value of Connectiontime for the instance
func (instance *Win32_PerfRawData_RemoteDesktopConnectionBrokerRedirectorPerformanceCounterProvider_RemoteDesktopConnectionBrokerRedirectorCounterset) SetPropertyConnectiontime(value uint64) (err error) {
	return instance.SetProperty("Connectiontime", value)
}

// GetConnectiontime gets the value of Connectiontime for the instance
func (instance *Win32_PerfRawData_RemoteDesktopConnectionBrokerRedirectorPerformanceCounterProvider_RemoteDesktopConnectionBrokerRedirectorCounterset) GetPropertyConnectiontime() (value uint64, err error) {
	retValue, err := instance.GetProperty("Connectiontime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContextacquisitionwaittime sets the value of Contextacquisitionwaittime for the instance
func (instance *Win32_PerfRawData_RemoteDesktopConnectionBrokerRedirectorPerformanceCounterProvider_RemoteDesktopConnectionBrokerRedirectorCounterset) SetPropertyContextacquisitionwaittime(value uint64) (err error) {
	return instance.SetProperty("Contextacquisitionwaittime", value)
}

// GetContextacquisitionwaittime gets the value of Contextacquisitionwaittime for the instance
func (instance *Win32_PerfRawData_RemoteDesktopConnectionBrokerRedirectorPerformanceCounterProvider_RemoteDesktopConnectionBrokerRedirectorCounterset) GetPropertyContextacquisitionwaittime() (value uint64, err error) {
	retValue, err := instance.GetProperty("Contextacquisitionwaittime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRPCContext sets the value of RPCContext for the instance
func (instance *Win32_PerfRawData_RemoteDesktopConnectionBrokerRedirectorPerformanceCounterProvider_RemoteDesktopConnectionBrokerRedirectorCounterset) SetPropertyRPCContext(value uint64) (err error) {
	return instance.SetProperty("RPCContext", value)
}

// GetRPCContext gets the value of RPCContext for the instance
func (instance *Win32_PerfRawData_RemoteDesktopConnectionBrokerRedirectorPerformanceCounterProvider_RemoteDesktopConnectionBrokerRedirectorCounterset) GetPropertyRPCContext() (value uint64, err error) {
	retValue, err := instance.GetProperty("RPCContext")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreadswaitingforRPCContext sets the value of ThreadswaitingforRPCContext for the instance
func (instance *Win32_PerfRawData_RemoteDesktopConnectionBrokerRedirectorPerformanceCounterProvider_RemoteDesktopConnectionBrokerRedirectorCounterset) SetPropertyThreadswaitingforRPCContext(value uint64) (err error) {
	return instance.SetProperty("ThreadswaitingforRPCContext", value)
}

// GetThreadswaitingforRPCContext gets the value of ThreadswaitingforRPCContext for the instance
func (instance *Win32_PerfRawData_RemoteDesktopConnectionBrokerRedirectorPerformanceCounterProvider_RemoteDesktopConnectionBrokerRedirectorCounterset) GetPropertyThreadswaitingforRPCContext() (value uint64, err error) {
	retValue, err := instance.GetProperty("ThreadswaitingforRPCContext")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
