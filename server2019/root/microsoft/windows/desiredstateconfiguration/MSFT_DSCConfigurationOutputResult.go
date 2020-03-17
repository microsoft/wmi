// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_DSCConfigurationOutputResult struct
type MSFT_DSCConfigurationOutputResult struct {
	MSFT_DSCConfigurationOutput

	//
	Object interface{}

	//
	Result uint32
}

// SetObject sets the value of Object for the instance
func (instance *MSFT_DSCConfigurationOutputResult) SetPropertyObject(value interface{}) (err error) {
	return instance.SetProperty("Object", value)
}

// GetObject gets the value of Object for the instance
func (instance *MSFT_DSCConfigurationOutputResult) GetPropertyObject() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Object")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResult sets the value of Result for the instance
func (instance *MSFT_DSCConfigurationOutputResult) SetPropertyResult(value uint32) (err error) {
	return instance.SetProperty("Result", value)
}

// GetResult gets the value of Result for the instance
func (instance *MSFT_DSCConfigurationOutputResult) GetPropertyResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("Result")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
