// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_DSCConfigurationOutputWriteMessage struct
type MSFT_DSCConfigurationOutputWriteMessage struct {
	MSFT_DSCConfigurationOutput

	//
	Channel uint32

	//
	Message string
}

// SetChannel sets the value of Channel for the instance
func (instance *MSFT_DSCConfigurationOutputWriteMessage) SetPropertyChannel(value uint32) (err error) {
	return instance.SetProperty("Channel", value)
}

// GetChannel gets the value of Channel for the instance
func (instance *MSFT_DSCConfigurationOutputWriteMessage) GetPropertyChannel() (value uint32, err error) {
	retValue, err := instance.GetProperty("Channel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMessage sets the value of Message for the instance
func (instance *MSFT_DSCConfigurationOutputWriteMessage) SetPropertyMessage(value string) (err error) {
	return instance.SetProperty("Message", value)
}

// GetMessage gets the value of Message for the instance
func (instance *MSFT_DSCConfigurationOutputWriteMessage) GetPropertyMessage() (value string, err error) {
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
