// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_DSCConfigurationOutputWhatIf struct
type MSFT_DSCConfigurationOutputWhatIf struct {
	MSFT_DSCConfigurationOutput

	//
	Message string
}

// SetMessage sets the value of Message for the instance
func (instance *MSFT_DSCConfigurationOutputWhatIf) SetPropertyMessage(value string) (err error) {
	return instance.SetProperty("Message", value)
}

// GetMessage gets the value of Message for the instance
func (instance *MSFT_DSCConfigurationOutputWhatIf) GetPropertyMessage() (value string, err error) {
	retValue, err := instance.GetProperty("Message")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
