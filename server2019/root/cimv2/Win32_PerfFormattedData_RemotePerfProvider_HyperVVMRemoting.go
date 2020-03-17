// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_RemotePerfProvider_HyperVVMRemoting struct
type Win32_PerfFormattedData_RemotePerfProvider_HyperVVMRemoting struct {
	Win32_PerfFormattedData

	//
	ConnectedClients uint32

	//
	UpdatedPixelsPersec uint32
}

// SetConnectedClients sets the value of ConnectedClients for the instance
func (instance *Win32_PerfFormattedData_RemotePerfProvider_HyperVVMRemoting) SetPropertyConnectedClients(value uint32) (err error) {
	return instance.SetProperty("ConnectedClients", value)
}

// GetConnectedClients gets the value of ConnectedClients for the instance
func (instance *Win32_PerfFormattedData_RemotePerfProvider_HyperVVMRemoting) GetPropertyConnectedClients() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectedClients")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpdatedPixelsPersec sets the value of UpdatedPixelsPersec for the instance
func (instance *Win32_PerfFormattedData_RemotePerfProvider_HyperVVMRemoting) SetPropertyUpdatedPixelsPersec(value uint32) (err error) {
	return instance.SetProperty("UpdatedPixelsPersec", value)
}

// GetUpdatedPixelsPersec gets the value of UpdatedPixelsPersec for the instance
func (instance *Win32_PerfFormattedData_RemotePerfProvider_HyperVVMRemoting) GetPropertyUpdatedPixelsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("UpdatedPixelsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
