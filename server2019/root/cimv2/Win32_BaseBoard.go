// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_BaseBoard struct
type Win32_BaseBoard struct {
	CIM_Card

	//
	ConfigOptions []string

	//
	Product string
}

// SetConfigOptions sets the value of ConfigOptions for the instance
func (instance *Win32_BaseBoard) SetPropertyConfigOptions(value []string) (err error) {
	return instance.SetProperty("ConfigOptions", value)
}

// GetConfigOptions gets the value of ConfigOptions for the instance
func (instance *Win32_BaseBoard) GetPropertyConfigOptions() (value []string, err error) {
	retValue, err := instance.GetProperty("ConfigOptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProduct sets the value of Product for the instance
func (instance *Win32_BaseBoard) SetPropertyProduct(value string) (err error) {
	return instance.SetProperty("Product", value)
}

// GetProduct gets the value of Product for the instance
func (instance *Win32_BaseBoard) GetPropertyProduct() (value string, err error) {
	retValue, err := instance.GetProperty("Product")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
