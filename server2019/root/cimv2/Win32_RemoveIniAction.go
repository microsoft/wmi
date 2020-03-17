// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_RemoveIniAction struct
type Win32_RemoveIniAction struct {
	CIM_Action

	//
	Action uint16

	//
	key string

	//
	Section string

	//
	Value string
}

// SetAction sets the value of Action for the instance
func (instance *Win32_RemoveIniAction) SetPropertyAction(value uint16) (err error) {
	return instance.SetProperty("Action", value)
}

// GetAction gets the value of Action for the instance
func (instance *Win32_RemoveIniAction) GetPropertyAction() (value uint16, err error) {
	retValue, err := instance.GetProperty("Action")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setkey sets the value of key for the instance
func (instance *Win32_RemoveIniAction) SetPropertykey(value string) (err error) {
	return instance.SetProperty("key", value)
}

// Getkey gets the value of key for the instance
func (instance *Win32_RemoveIniAction) GetPropertykey() (value string, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSection sets the value of Section for the instance
func (instance *Win32_RemoveIniAction) SetPropertySection(value string) (err error) {
	return instance.SetProperty("Section", value)
}

// GetSection gets the value of Section for the instance
func (instance *Win32_RemoveIniAction) GetPropertySection() (value string, err error) {
	retValue, err := instance.GetProperty("Section")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValue sets the value of Value for the instance
func (instance *Win32_RemoveIniAction) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *Win32_RemoveIniAction) GetPropertyValue() (value string, err error) {
	retValue, err := instance.GetProperty("Value")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
