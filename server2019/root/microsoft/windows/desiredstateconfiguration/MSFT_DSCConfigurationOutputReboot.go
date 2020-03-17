// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_DSCConfigurationOutputReboot struct
type MSFT_DSCConfigurationOutputReboot struct {
	MSFT_DSCConfigurationOutput

	//
	Automatic bool
}

// SetAutomatic sets the value of Automatic for the instance
func (instance *MSFT_DSCConfigurationOutputReboot) SetPropertyAutomatic(value bool) (err error) {
	return instance.SetProperty("Automatic", value)
}

// GetAutomatic gets the value of Automatic for the instance
func (instance *MSFT_DSCConfigurationOutputReboot) GetPropertyAutomatic() (value bool, err error) {
	retValue, err := instance.GetProperty("Automatic")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
