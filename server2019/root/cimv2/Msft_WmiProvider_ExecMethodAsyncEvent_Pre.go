// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Msft_WmiProvider_ExecMethodAsyncEvent_Pre struct
type Msft_WmiProvider_ExecMethodAsyncEvent_Pre struct {
	Msft_WmiProvider_OperationEvent_Pre

	//
	Flags uint32

	//
	InputParameters interface{}

	//
	MethodName string

	//
	ObjectPath string
}

// SetFlags sets the value of Flags for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Pre) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Pre) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInputParameters sets the value of InputParameters for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Pre) SetPropertyInputParameters(value interface{}) (err error) {
	return instance.SetProperty("InputParameters", value)
}

// GetInputParameters gets the value of InputParameters for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Pre) GetPropertyInputParameters() (value interface{}, err error) {
	retValue, err := instance.GetProperty("InputParameters")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMethodName sets the value of MethodName for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Pre) SetPropertyMethodName(value string) (err error) {
	return instance.SetProperty("MethodName", value)
}

// GetMethodName gets the value of MethodName for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Pre) GetPropertyMethodName() (value string, err error) {
	retValue, err := instance.GetProperty("MethodName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetObjectPath sets the value of ObjectPath for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Pre) SetPropertyObjectPath(value string) (err error) {
	return instance.SetProperty("ObjectPath", value)
}

// GetObjectPath gets the value of ObjectPath for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Pre) GetPropertyObjectPath() (value string, err error) {
	retValue, err := instance.GetProperty("ObjectPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
