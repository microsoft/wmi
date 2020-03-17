// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Property struct
type Win32_Property struct {
	Win32_MSIResource

	//
	ProductCode string

	//
	Property string

	//
	Value string
}

// SetProductCode sets the value of ProductCode for the instance
func (instance *Win32_Property) SetPropertyProductCode(value string) (err error) {
	return instance.SetProperty("ProductCode", value)
}

// GetProductCode gets the value of ProductCode for the instance
func (instance *Win32_Property) GetPropertyProductCode() (value string, err error) {
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

// SetProperty sets the value of Property for the instance
func (instance *Win32_Property) SetPropertyProperty(value string) (err error) {
	return instance.SetProperty("Property", value)
}

// GetProperty gets the value of Property for the instance
func (instance *Win32_Property) GetPropertyProperty() (value string, err error) {
	retValue, err := instance.GetProperty("Property")
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
func (instance *Win32_Property) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *Win32_Property) GetPropertyValue() (value string, err error) {
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
