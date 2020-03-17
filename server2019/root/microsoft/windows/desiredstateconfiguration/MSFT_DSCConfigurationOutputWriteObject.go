// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_DSCConfigurationOutputWriteObject struct
type MSFT_DSCConfigurationOutputWriteObject struct {
	MSFT_DSCConfigurationOutput

	//
	Object interface{}
}

// SetObject sets the value of Object for the instance
func (instance *MSFT_DSCConfigurationOutputWriteObject) SetPropertyObject(value interface{}) (err error) {
	return instance.SetProperty("Object", value)
}

// GetObject gets the value of Object for the instance
func (instance *MSFT_DSCConfigurationOutputWriteObject) GetPropertyObject() (value interface{}, err error) {
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
