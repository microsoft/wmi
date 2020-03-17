// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_OptionalFeature struct
type Win32_OptionalFeature struct {
	CIM_LogicalElement

	//
	InstallState uint32
}

// SetInstallState sets the value of InstallState for the instance
func (instance *Win32_OptionalFeature) SetPropertyInstallState(value uint32) (err error) {
	return instance.SetProperty("InstallState", value)
}

// GetInstallState gets the value of InstallState for the instance
func (instance *Win32_OptionalFeature) GetPropertyInstallState() (value uint32, err error) {
	retValue, err := instance.GetProperty("InstallState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
