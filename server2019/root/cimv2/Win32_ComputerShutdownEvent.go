// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ComputerShutdownEvent struct
type Win32_ComputerShutdownEvent struct {
	Win32_ComputerSystemEvent

	//
	Type uint32
}

// SetType sets the value of Type for the instance
func (instance *Win32_ComputerShutdownEvent) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *Win32_ComputerShutdownEvent) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
