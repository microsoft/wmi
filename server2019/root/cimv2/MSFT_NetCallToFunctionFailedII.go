// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetCallToFunctionFailedII struct
type MSFT_NetCallToFunctionFailedII struct {
	MSFT_SCMEventLogEvent

	//
	Argument string

	//
	Error uint32

	//
	FunctionName string
}

// SetArgument sets the value of Argument for the instance
func (instance *MSFT_NetCallToFunctionFailedII) SetPropertyArgument(value string) (err error) {
	return instance.SetProperty("Argument", value)
}

// GetArgument gets the value of Argument for the instance
func (instance *MSFT_NetCallToFunctionFailedII) GetPropertyArgument() (value string, err error) {
	retValue, err := instance.GetProperty("Argument")
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
func (instance *MSFT_NetCallToFunctionFailedII) SetPropertyError(value uint32) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *MSFT_NetCallToFunctionFailedII) GetPropertyError() (value uint32, err error) {
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

// SetFunctionName sets the value of FunctionName for the instance
func (instance *MSFT_NetCallToFunctionFailedII) SetPropertyFunctionName(value string) (err error) {
	return instance.SetProperty("FunctionName", value)
}

// GetFunctionName gets the value of FunctionName for the instance
func (instance *MSFT_NetCallToFunctionFailedII) GetPropertyFunctionName() (value string, err error) {
	retValue, err := instance.GetProperty("FunctionName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
