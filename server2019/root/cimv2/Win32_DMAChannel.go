// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DMAChannel struct
type Win32_DMAChannel struct {
	CIM_DMA

	//
	Port uint32
}

// SetPort sets the value of Port for the instance
func (instance *Win32_DMAChannel) SetPropertyPort(value uint32) (err error) {
	return instance.SetProperty("Port", value)
}

// GetPort gets the value of Port for the instance
func (instance *Win32_DMAChannel) GetPropertyPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("Port")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
