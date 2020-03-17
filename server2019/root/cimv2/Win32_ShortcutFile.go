// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ShortcutFile struct
type Win32_ShortcutFile struct {
	CIM_DataFile

	//
	Target string
}

// SetTarget sets the value of Target for the instance
func (instance *Win32_ShortcutFile) SetPropertyTarget(value string) (err error) {
	return instance.SetProperty("Target", value)
}

// GetTarget gets the value of Target for the instance
func (instance *Win32_ShortcutFile) GetPropertyTarget() (value string, err error) {
	retValue, err := instance.GetProperty("Target")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
