// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DeviceMemoryAddress struct
type Win32_DeviceMemoryAddress struct {
	Win32_SystemMemoryResource

	//
	MemoryType string
}

// SetMemoryType sets the value of MemoryType for the instance
func (instance *Win32_DeviceMemoryAddress) SetPropertyMemoryType(value string) (err error) {
	return instance.SetProperty("MemoryType", value)
}

// GetMemoryType gets the value of MemoryType for the instance
func (instance *Win32_DeviceMemoryAddress) GetPropertyMemoryType() (value string, err error) {
	retValue, err := instance.GetProperty("MemoryType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
