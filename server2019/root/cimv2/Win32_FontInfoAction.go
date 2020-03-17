// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_FontInfoAction struct
type Win32_FontInfoAction struct {
	CIM_Action

	//
	File string

	//
	FontTitle string
}

// SetFile sets the value of File for the instance
func (instance *Win32_FontInfoAction) SetPropertyFile(value string) (err error) {
	return instance.SetProperty("File", value)
}

// GetFile gets the value of File for the instance
func (instance *Win32_FontInfoAction) GetPropertyFile() (value string, err error) {
	retValue, err := instance.GetProperty("File")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFontTitle sets the value of FontTitle for the instance
func (instance *Win32_FontInfoAction) SetPropertyFontTitle(value string) (err error) {
	return instance.SetProperty("FontTitle", value)
}

// GetFontTitle gets the value of FontTitle for the instance
func (instance *Win32_FontInfoAction) GetPropertyFontTitle() (value string, err error) {
	retValue, err := instance.GetProperty("FontTitle")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
