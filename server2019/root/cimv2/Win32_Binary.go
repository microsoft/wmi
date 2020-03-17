// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Binary struct
type Win32_Binary struct {
	Win32_MSIResource

	//
	Data string

	//
	Name string

	//
	ProductCode string
}

// SetData sets the value of Data for the instance
func (instance *Win32_Binary) SetPropertyData(value string) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *Win32_Binary) GetPropertyData() (value string, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *Win32_Binary) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Win32_Binary) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProductCode sets the value of ProductCode for the instance
func (instance *Win32_Binary) SetPropertyProductCode(value string) (err error) {
	return instance.SetProperty("ProductCode", value)
}

// GetProductCode gets the value of ProductCode for the instance
func (instance *Win32_Binary) GetPropertyProductCode() (value string, err error) {
	retValue, err := instance.GetProperty("ProductCode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
