// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PrivilegesStatus struct
type Win32_PrivilegesStatus struct {
	__ExtendedStatus

	//
	PrivilegesNotHeld []string

	//
	PrivilegesRequired []string
}

// SetPrivilegesNotHeld sets the value of PrivilegesNotHeld for the instance
func (instance *Win32_PrivilegesStatus) SetPropertyPrivilegesNotHeld(value []string) (err error) {
	return instance.SetProperty("PrivilegesNotHeld", value)
}

// GetPrivilegesNotHeld gets the value of PrivilegesNotHeld for the instance
func (instance *Win32_PrivilegesStatus) GetPropertyPrivilegesNotHeld() (value []string, err error) {
	retValue, err := instance.GetProperty("PrivilegesNotHeld")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrivilegesRequired sets the value of PrivilegesRequired for the instance
func (instance *Win32_PrivilegesStatus) SetPropertyPrivilegesRequired(value []string) (err error) {
	return instance.SetProperty("PrivilegesRequired", value)
}

// GetPrivilegesRequired gets the value of PrivilegesRequired for the instance
func (instance *Win32_PrivilegesStatus) GetPropertyPrivilegesRequired() (value []string, err error) {
	retValue, err := instance.GetProperty("PrivilegesRequired")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
