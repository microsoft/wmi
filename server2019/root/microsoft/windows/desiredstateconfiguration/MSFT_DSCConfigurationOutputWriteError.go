// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_DSCConfigurationOutputWriteError struct
type MSFT_DSCConfigurationOutputWriteError struct {
	MSFT_DSCConfigurationOutput

	//
	Error interface{}
}

// SetError sets the value of Error for the instance
func (instance *MSFT_DSCConfigurationOutputWriteError) SetPropertyError(value interface{}) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *MSFT_DSCConfigurationOutputWriteError) GetPropertyError() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
