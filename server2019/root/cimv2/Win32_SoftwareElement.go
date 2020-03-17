// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_SoftwareElement struct
type Win32_SoftwareElement struct {
	CIM_SoftwareElement

	//
	Attributes uint16

	//
	InstallState int16

	//
	Path string
}

// SetAttributes sets the value of Attributes for the instance
func (instance *Win32_SoftwareElement) SetPropertyAttributes(value uint16) (err error) {
	return instance.SetProperty("Attributes", value)
}

// GetAttributes gets the value of Attributes for the instance
func (instance *Win32_SoftwareElement) GetPropertyAttributes() (value uint16, err error) {
	retValue, err := instance.GetProperty("Attributes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstallState sets the value of InstallState for the instance
func (instance *Win32_SoftwareElement) SetPropertyInstallState(value int16) (err error) {
	return instance.SetProperty("InstallState", value)
}

// GetInstallState gets the value of InstallState for the instance
func (instance *Win32_SoftwareElement) GetPropertyInstallState() (value int16, err error) {
	retValue, err := instance.GetProperty("InstallState")
	if err != nil {
		return
	}
	value, ok := retValue.(int16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *Win32_SoftwareElement) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *Win32_SoftwareElement) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
