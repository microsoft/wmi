// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ODBCSourceAttribute struct
type Win32_ODBCSourceAttribute struct {
	CIM_Setting

	//
	Attribute string

	//
	DataSource string

	//
	Value string
}

// SetAttribute sets the value of Attribute for the instance
func (instance *Win32_ODBCSourceAttribute) SetPropertyAttribute(value string) (err error) {
	return instance.SetProperty("Attribute", value)
}

// GetAttribute gets the value of Attribute for the instance
func (instance *Win32_ODBCSourceAttribute) GetPropertyAttribute() (value string, err error) {
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

// SetDataSource sets the value of DataSource for the instance
func (instance *Win32_ODBCSourceAttribute) SetPropertyDataSource(value string) (err error) {
	return instance.SetProperty("DataSource", value)
}

// GetDataSource gets the value of DataSource for the instance
func (instance *Win32_ODBCSourceAttribute) GetPropertyDataSource() (value string, err error) {
	retValue, err := instance.GetProperty("DataSource")
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
func (instance *Win32_ODBCSourceAttribute) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *Win32_ODBCSourceAttribute) GetPropertyValue() (value string, err error) {
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
