// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_DirectorySpecification struct
type CIM_DirectorySpecification struct {
	CIM_Check

	//
	DirectoryPath string

	//
	DirectoryType uint16
}

// SetDirectoryPath sets the value of DirectoryPath for the instance
func (instance *CIM_DirectorySpecification) SetPropertyDirectoryPath(value string) (err error) {
	return instance.SetProperty("DirectoryPath", value)
}

// GetDirectoryPath gets the value of DirectoryPath for the instance
func (instance *CIM_DirectorySpecification) GetPropertyDirectoryPath() (value string, err error) {
	retValue, err := instance.GetProperty("DirectoryPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectoryType sets the value of DirectoryType for the instance
func (instance *CIM_DirectorySpecification) SetPropertyDirectoryType(value uint16) (err error) {
	return instance.SetProperty("DirectoryType", value)
}

// GetDirectoryType gets the value of DirectoryType for the instance
func (instance *CIM_DirectorySpecification) GetPropertyDirectoryType() (value uint16, err error) {
	retValue, err := instance.GetProperty("DirectoryType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
