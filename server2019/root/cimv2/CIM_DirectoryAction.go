// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_DirectoryAction struct
type CIM_DirectoryAction struct {
	CIM_Action

	//
	DirectoryName string
}

// SetDirectoryName sets the value of DirectoryName for the instance
func (instance *CIM_DirectoryAction) SetPropertyDirectoryName(value string) (err error) {
	return instance.SetProperty("DirectoryName", value)
}

// GetDirectoryName gets the value of DirectoryName for the instance
func (instance *CIM_DirectoryAction) GetPropertyDirectoryName() (value string, err error) {
	retValue, err := instance.GetProperty("DirectoryName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
