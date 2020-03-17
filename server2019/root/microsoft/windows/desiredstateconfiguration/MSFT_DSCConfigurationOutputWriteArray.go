// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_DSCConfigurationOutputWriteArray struct
type MSFT_DSCConfigurationOutputWriteArray struct {
	MSFT_DSCConfigurationOutput

	//
	Array []interface{}

	//
	ParameterName string
}

// SetArray sets the value of Array for the instance
func (instance *MSFT_DSCConfigurationOutputWriteArray) SetPropertyArray(value []interface{}) (err error) {
	return instance.SetProperty("Array", value)
}

// GetArray gets the value of Array for the instance
func (instance *MSFT_DSCConfigurationOutputWriteArray) GetPropertyArray() (value []interface{}, err error) {
	retValue, err := instance.GetProperty("Array")
	if err != nil {
		return
	}
	value, ok := retValue.([]interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParameterName sets the value of ParameterName for the instance
func (instance *MSFT_DSCConfigurationOutputWriteArray) SetPropertyParameterName(value string) (err error) {
	return instance.SetProperty("ParameterName", value)
}

// GetParameterName gets the value of ParameterName for the instance
func (instance *MSFT_DSCConfigurationOutputWriteArray) GetPropertyParameterName() (value string, err error) {
	retValue, err := instance.GetProperty("ParameterName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
