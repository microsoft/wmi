// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_DdmCounterProvider_RAS struct
type Win32_PerfFormattedData_DdmCounterProvider_RAS struct {
	Win32_PerfFormattedData

	//
	BytesReceivedByDisconnectedClients uint64

	//
	BytesTransmittedByDisconnectedClients uint64

	//
	FailedAuthentications uint32

	//
	MaxClients uint32

	//
	TotalClients uint32
}

// SetBytesReceivedByDisconnectedClients sets the value of BytesReceivedByDisconnectedClients for the instance
func (instance *Win32_PerfFormattedData_DdmCounterProvider_RAS) SetPropertyBytesReceivedByDisconnectedClients(value uint64) (err error) {
	return instance.SetProperty("BytesReceivedByDisconnectedClients", value)
}

// GetBytesReceivedByDisconnectedClients gets the value of BytesReceivedByDisconnectedClients for the instance
func (instance *Win32_PerfFormattedData_DdmCounterProvider_RAS) GetPropertyBytesReceivedByDisconnectedClients() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceivedByDisconnectedClients")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesTransmittedByDisconnectedClients sets the value of BytesTransmittedByDisconnectedClients for the instance
func (instance *Win32_PerfFormattedData_DdmCounterProvider_RAS) SetPropertyBytesTransmittedByDisconnectedClients(value uint64) (err error) {
	return instance.SetProperty("BytesTransmittedByDisconnectedClients", value)
}

// GetBytesTransmittedByDisconnectedClients gets the value of BytesTransmittedByDisconnectedClients for the instance
func (instance *Win32_PerfFormattedData_DdmCounterProvider_RAS) GetPropertyBytesTransmittedByDisconnectedClients() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesTransmittedByDisconnectedClients")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailedAuthentications sets the value of FailedAuthentications for the instance
func (instance *Win32_PerfFormattedData_DdmCounterProvider_RAS) SetPropertyFailedAuthentications(value uint32) (err error) {
	return instance.SetProperty("FailedAuthentications", value)
}

// GetFailedAuthentications gets the value of FailedAuthentications for the instance
func (instance *Win32_PerfFormattedData_DdmCounterProvider_RAS) GetPropertyFailedAuthentications() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailedAuthentications")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxClients sets the value of MaxClients for the instance
func (instance *Win32_PerfFormattedData_DdmCounterProvider_RAS) SetPropertyMaxClients(value uint32) (err error) {
	return instance.SetProperty("MaxClients", value)
}

// GetMaxClients gets the value of MaxClients for the instance
func (instance *Win32_PerfFormattedData_DdmCounterProvider_RAS) GetPropertyMaxClients() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxClients")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalClients sets the value of TotalClients for the instance
func (instance *Win32_PerfFormattedData_DdmCounterProvider_RAS) SetPropertyTotalClients(value uint32) (err error) {
	return instance.SetProperty("TotalClients", value)
}

// GetTotalClients gets the value of TotalClients for the instance
func (instance *Win32_PerfFormattedData_DdmCounterProvider_RAS) GetPropertyTotalClients() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalClients")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
