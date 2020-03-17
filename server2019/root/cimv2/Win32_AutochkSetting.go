// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_AutochkSetting struct
type Win32_AutochkSetting struct {
	CIM_Setting

	//
	UserInputDelay uint32
}

// SetUserInputDelay sets the value of UserInputDelay for the instance
func (instance *Win32_AutochkSetting) SetPropertyUserInputDelay(value uint32) (err error) {
	return instance.SetProperty("UserInputDelay", value)
}

// GetUserInputDelay gets the value of UserInputDelay for the instance
func (instance *Win32_AutochkSetting) GetPropertyUserInputDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("UserInputDelay")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
