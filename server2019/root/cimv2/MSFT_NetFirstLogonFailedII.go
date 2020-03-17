// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetFirstLogonFailedII struct
type MSFT_NetFirstLogonFailedII struct {
	MSFT_SCMEventLogEvent

	//
	Account string

	//
	Error uint32

	//
	Service string
}

// SetAccount sets the value of Account for the instance
func (instance *MSFT_NetFirstLogonFailedII) SetPropertyAccount(value string) (err error) {
	return instance.SetProperty("Account", value)
}

// GetAccount gets the value of Account for the instance
func (instance *MSFT_NetFirstLogonFailedII) GetPropertyAccount() (value string, err error) {
	retValue, err := instance.GetProperty("Account")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetError sets the value of Error for the instance
func (instance *MSFT_NetFirstLogonFailedII) SetPropertyError(value uint32) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *MSFT_NetFirstLogonFailedII) GetPropertyError() (value uint32, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetService sets the value of Service for the instance
func (instance *MSFT_NetFirstLogonFailedII) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetFirstLogonFailedII) GetPropertyService() (value string, err error) {
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
