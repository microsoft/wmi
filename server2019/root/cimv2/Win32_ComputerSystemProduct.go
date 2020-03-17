// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ComputerSystemProduct struct
type Win32_ComputerSystemProduct struct {
	CIM_Product

	//
	UUID string
}

// SetUUID sets the value of UUID for the instance
func (instance *Win32_ComputerSystemProduct) SetPropertyUUID(value string) (err error) {
	return instance.SetProperty("UUID", value)
}

// GetUUID gets the value of UUID for the instance
func (instance *Win32_ComputerSystemProduct) GetPropertyUUID() (value string, err error) {
	retValue, err := instance.GetProperty("UUID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
