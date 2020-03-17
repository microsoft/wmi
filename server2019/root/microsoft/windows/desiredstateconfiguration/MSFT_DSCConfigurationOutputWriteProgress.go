// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_DSCConfigurationOutputWriteProgress struct
type MSFT_DSCConfigurationOutputWriteProgress struct {
	MSFT_DSCConfigurationOutput

	//
	Activity string

	//
	CurrentOperation string

	//
	PercentComplete uint32

	//
	SecondsRemaining uint32

	//
	StatusDescription string
}

// SetActivity sets the value of Activity for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) SetPropertyActivity(value string) (err error) {
	return instance.SetProperty("Activity", value)
}

// GetActivity gets the value of Activity for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) GetPropertyActivity() (value string, err error) {
	retValue, err := instance.GetProperty("Activity")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentOperation sets the value of CurrentOperation for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) SetPropertyCurrentOperation(value string) (err error) {
	return instance.SetProperty("CurrentOperation", value)
}

// GetCurrentOperation gets the value of CurrentOperation for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) GetPropertyCurrentOperation() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentOperation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentComplete sets the value of PercentComplete for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) SetPropertyPercentComplete(value uint32) (err error) {
	return instance.SetProperty("PercentComplete", value)
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) GetPropertyPercentComplete() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentComplete")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecondsRemaining sets the value of SecondsRemaining for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) SetPropertySecondsRemaining(value uint32) (err error) {
	return instance.SetProperty("SecondsRemaining", value)
}

// GetSecondsRemaining gets the value of SecondsRemaining for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) GetPropertySecondsRemaining() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecondsRemaining")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatusDescription sets the value of StatusDescription for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) SetPropertyStatusDescription(value string) (err error) {
	return instance.SetProperty("StatusDescription", value)
}

// GetStatusDescription gets the value of StatusDescription for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) GetPropertyStatusDescription() (value string, err error) {
	retValue, err := instance.GetProperty("StatusDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
