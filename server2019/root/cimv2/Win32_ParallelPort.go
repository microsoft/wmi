// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ParallelPort struct
type Win32_ParallelPort struct {
	CIM_ParallelController

	//
	OSAutoDiscovered bool
}

// SetOSAutoDiscovered sets the value of OSAutoDiscovered for the instance
func (instance *Win32_ParallelPort) SetPropertyOSAutoDiscovered(value bool) (err error) {
	return instance.SetProperty("OSAutoDiscovered", value)
}

// GetOSAutoDiscovered gets the value of OSAutoDiscovered for the instance
func (instance *Win32_ParallelPort) GetPropertyOSAutoDiscovered() (value bool, err error) {
	retValue, err := instance.GetProperty("OSAutoDiscovered")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
