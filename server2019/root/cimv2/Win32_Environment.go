// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Environment struct
type Win32_Environment struct {
	CIM_SystemResource

	//
	SystemVariable bool

	//
	UserName string

	//
	VariableValue string
}

// SetSystemVariable sets the value of SystemVariable for the instance
func (instance *Win32_Environment) SetPropertySystemVariable(value bool) (err error) {
	return instance.SetProperty("SystemVariable", value)
}

// GetSystemVariable gets the value of SystemVariable for the instance
func (instance *Win32_Environment) GetPropertySystemVariable() (value bool, err error) {
	retValue, err := instance.GetProperty("SystemVariable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserName sets the value of UserName for the instance
func (instance *Win32_Environment) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", value)
}

// GetUserName gets the value of UserName for the instance
func (instance *Win32_Environment) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVariableValue sets the value of VariableValue for the instance
func (instance *Win32_Environment) SetPropertyVariableValue(value string) (err error) {
	return instance.SetProperty("VariableValue", value)
}

// GetVariableValue gets the value of VariableValue for the instance
func (instance *Win32_Environment) GetPropertyVariableValue() (value string, err error) {
	retValue, err := instance.GetProperty("VariableValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
