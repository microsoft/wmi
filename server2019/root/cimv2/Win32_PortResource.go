// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PortResource struct
type Win32_PortResource struct {
	Win32_SystemMemoryResource

	//
	Alias bool
}

// SetAlias sets the value of Alias for the instance
func (instance *Win32_PortResource) SetPropertyAlias(value bool) (err error) {
	return instance.SetProperty("Alias", value)
}

// GetAlias gets the value of Alias for the instance
func (instance *Win32_PortResource) GetPropertyAlias() (value bool, err error) {
	retValue, err := instance.GetProperty("Alias")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
