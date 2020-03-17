// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_CIMLogicalDeviceCIMDataFile struct
type Win32_CIMLogicalDeviceCIMDataFile struct {
	CIM_Dependency

	//
	Purpose uint16

	//
	PurposeDescription string
}

// SetPurpose sets the value of Purpose for the instance
func (instance *Win32_CIMLogicalDeviceCIMDataFile) SetPropertyPurpose(value uint16) (err error) {
	return instance.SetProperty("Purpose", value)
}

// GetPurpose gets the value of Purpose for the instance
func (instance *Win32_CIMLogicalDeviceCIMDataFile) GetPropertyPurpose() (value uint16, err error) {
	retValue, err := instance.GetProperty("Purpose")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPurposeDescription sets the value of PurposeDescription for the instance
func (instance *Win32_CIMLogicalDeviceCIMDataFile) SetPropertyPurposeDescription(value string) (err error) {
	return instance.SetProperty("PurposeDescription", value)
}

// GetPurposeDescription gets the value of PurposeDescription for the instance
func (instance *Win32_CIMLogicalDeviceCIMDataFile) GetPropertyPurposeDescription() (value string, err error) {
	retValue, err := instance.GetProperty("PurposeDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
