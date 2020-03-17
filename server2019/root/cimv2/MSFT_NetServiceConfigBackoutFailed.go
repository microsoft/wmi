// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetServiceConfigBackoutFailed struct
type MSFT_NetServiceConfigBackoutFailed struct {
	MSFT_SCMEventLogEvent

	//
	ConfigField string

	//
	Service string
}

// SetConfigField sets the value of ConfigField for the instance
func (instance *MSFT_NetServiceConfigBackoutFailed) SetPropertyConfigField(value string) (err error) {
	return instance.SetProperty("ConfigField", value)
}

// GetConfigField gets the value of ConfigField for the instance
func (instance *MSFT_NetServiceConfigBackoutFailed) GetPropertyConfigField() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigField")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetService sets the value of Service for the instance
func (instance *MSFT_NetServiceConfigBackoutFailed) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceConfigBackoutFailed) GetPropertyService() (value string, err error) {
	retValue, err := instance.GetProperty("Service")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
