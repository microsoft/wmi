// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_EnvironmentSpecification struct
type Win32_EnvironmentSpecification struct {
	CIM_Check

	//
	Environment string

	//
	Value string
}

// SetEnvironment sets the value of Environment for the instance
func (instance *Win32_EnvironmentSpecification) SetPropertyEnvironment(value string) (err error) {
	return instance.SetProperty("Environment", value)
}

// GetEnvironment gets the value of Environment for the instance
func (instance *Win32_EnvironmentSpecification) GetPropertyEnvironment() (value string, err error) {
	retValue, err := instance.GetProperty("Environment")
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
func (instance *Win32_EnvironmentSpecification) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *Win32_EnvironmentSpecification) GetPropertyValue() (value string, err error) {
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
