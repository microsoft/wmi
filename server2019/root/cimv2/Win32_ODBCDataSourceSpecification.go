// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ODBCDataSourceSpecification struct
type Win32_ODBCDataSourceSpecification struct {
	CIM_Check

	//
	DataSource string

	//
	DriverDescription string

	//
	Registration string
}

// SetDataSource sets the value of DataSource for the instance
func (instance *Win32_ODBCDataSourceSpecification) SetPropertyDataSource(value string) (err error) {
	return instance.SetProperty("DataSource", value)
}

// GetDataSource gets the value of DataSource for the instance
func (instance *Win32_ODBCDataSourceSpecification) GetPropertyDataSource() (value string, err error) {
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

// SetDriverDescription sets the value of DriverDescription for the instance
func (instance *Win32_ODBCDataSourceSpecification) SetPropertyDriverDescription(value string) (err error) {
	return instance.SetProperty("DriverDescription", value)
}

// GetDriverDescription gets the value of DriverDescription for the instance
func (instance *Win32_ODBCDataSourceSpecification) GetPropertyDriverDescription() (value string, err error) {
	retValue, err := instance.GetProperty("DriverDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistration sets the value of Registration for the instance
func (instance *Win32_ODBCDataSourceSpecification) SetPropertyRegistration(value string) (err error) {
	return instance.SetProperty("Registration", value)
}

// GetRegistration gets the value of Registration for the instance
func (instance *Win32_ODBCDataSourceSpecification) GetPropertyRegistration() (value string, err error) {
	retValue, err := instance.GetProperty("Registration")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
