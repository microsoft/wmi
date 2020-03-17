// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ODBCAttribute struct
type Win32_ODBCAttribute struct {
	CIM_Setting

	//
	Attribute string

	//
	Driver string

	//
	Value string
}

// SetAttribute sets the value of Attribute for the instance
func (instance *Win32_ODBCAttribute) SetPropertyAttribute(value string) (err error) {
	return instance.SetProperty("Attribute", value)
}

// GetAttribute gets the value of Attribute for the instance
func (instance *Win32_ODBCAttribute) GetPropertyAttribute() (value string, err error) {
	retValue, err := instance.GetProperty("Attribute")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriver sets the value of Driver for the instance
func (instance *Win32_ODBCAttribute) SetPropertyDriver(value string) (err error) {
	return instance.SetProperty("Driver", value)
}

// GetDriver gets the value of Driver for the instance
func (instance *Win32_ODBCAttribute) GetPropertyDriver() (value string, err error) {
	retValue, err := instance.GetProperty("Driver")
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
func (instance *Win32_ODBCAttribute) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *Win32_ODBCAttribute) GetPropertyValue() (value string, err error) {
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
