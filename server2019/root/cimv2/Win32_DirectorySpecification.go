// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DirectorySpecification struct
type Win32_DirectorySpecification struct {
	CIM_DirectorySpecification

	//
	DefaultDir string

	//
	Directory string
}

// SetDefaultDir sets the value of DefaultDir for the instance
func (instance *Win32_DirectorySpecification) SetPropertyDefaultDir(value string) (err error) {
	return instance.SetProperty("DefaultDir", value)
}

// GetDefaultDir gets the value of DefaultDir for the instance
func (instance *Win32_DirectorySpecification) GetPropertyDefaultDir() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultDir")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectory sets the value of Directory for the instance
func (instance *Win32_DirectorySpecification) SetPropertyDirectory(value string) (err error) {
	return instance.SetProperty("Directory", value)
}

// GetDirectory gets the value of Directory for the instance
func (instance *Win32_DirectorySpecification) GetPropertyDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("Directory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
