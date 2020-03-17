// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetServiceControlSuccess struct
type MSFT_NetServiceControlSuccess struct {
	MSFT_SCMEventLogEvent

	//
	Control string

	//
	Service string

	//
	sid string
}

// SetControl sets the value of Control for the instance
func (instance *MSFT_NetServiceControlSuccess) SetPropertyControl(value string) (err error) {
	return instance.SetProperty("Control", value)
}

// GetControl gets the value of Control for the instance
func (instance *MSFT_NetServiceControlSuccess) GetPropertyControl() (value string, err error) {
	retValue, err := instance.GetProperty("Control")
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
func (instance *MSFT_NetServiceControlSuccess) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceControlSuccess) GetPropertyService() (value string, err error) {
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

// Setsid sets the value of sid for the instance
func (instance *MSFT_NetServiceControlSuccess) SetPropertysid(value string) (err error) {
	return instance.SetProperty("sid", value)
}

// Getsid gets the value of sid for the instance
func (instance *MSFT_NetServiceControlSuccess) GetPropertysid() (value string, err error) {
	retValue, err := instance.GetProperty("sid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
