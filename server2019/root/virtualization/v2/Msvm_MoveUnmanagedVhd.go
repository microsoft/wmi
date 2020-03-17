// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_MoveUnmanagedVhd struct
type Msvm_MoveUnmanagedVhd struct {
	CIM_ManagedElement

	//
	VhdDestinationPath string

	//
	VhdSourcePath string
}

// SetVhdDestinationPath sets the value of VhdDestinationPath for the instance
func (instance *Msvm_MoveUnmanagedVhd) SetPropertyVhdDestinationPath(value string) (err error) {
	return instance.SetProperty("VhdDestinationPath", value)
}

// GetVhdDestinationPath gets the value of VhdDestinationPath for the instance
func (instance *Msvm_MoveUnmanagedVhd) GetPropertyVhdDestinationPath() (value string, err error) {
	retValue, err := instance.GetProperty("VhdDestinationPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVhdSourcePath sets the value of VhdSourcePath for the instance
func (instance *Msvm_MoveUnmanagedVhd) SetPropertyVhdSourcePath(value string) (err error) {
	return instance.SetProperty("VhdSourcePath", value)
}

// GetVhdSourcePath gets the value of VhdSourcePath for the instance
func (instance *Msvm_MoveUnmanagedVhd) GetPropertyVhdSourcePath() (value string, err error) {
	retValue, err := instance.GetProperty("VhdSourcePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
