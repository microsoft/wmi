// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_RemoveDirectoryAction struct
type CIM_RemoveDirectoryAction struct {
	CIM_DirectoryAction

	//
	MustBeEmpty bool
}

// SetMustBeEmpty sets the value of MustBeEmpty for the instance
func (instance *CIM_RemoveDirectoryAction) SetPropertyMustBeEmpty(value bool) (err error) {
	return instance.SetProperty("MustBeEmpty", value)
}

// GetMustBeEmpty gets the value of MustBeEmpty for the instance
func (instance *CIM_RemoveDirectoryAction) GetPropertyMustBeEmpty() (value bool, err error) {
	retValue, err := instance.GetProperty("MustBeEmpty")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
