// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ComputerSystemEvent struct
type Win32_ComputerSystemEvent struct {
	__ExtrinsicEvent

	//
	MachineName string
}

// SetMachineName sets the value of MachineName for the instance
func (instance *Win32_ComputerSystemEvent) SetPropertyMachineName(value string) (err error) {
	return instance.SetProperty("MachineName", value)
}

// GetMachineName gets the value of MachineName for the instance
func (instance *Win32_ComputerSystemEvent) GetPropertyMachineName() (value string, err error) {
	retValue, err := instance.GetProperty("MachineName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
