// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_MemoryArray struct
type Win32_MemoryArray struct {
	Win32_SMBIOSMemory

	//
	ErrorGranularity uint16
}

// SetErrorGranularity sets the value of ErrorGranularity for the instance
func (instance *Win32_MemoryArray) SetPropertyErrorGranularity(value uint16) (err error) {
	return instance.SetProperty("ErrorGranularity", value)
}

// GetErrorGranularity gets the value of ErrorGranularity for the instance
func (instance *Win32_MemoryArray) GetPropertyErrorGranularity() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorGranularity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
