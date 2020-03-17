// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Msft_WmiProvider_ExecMethodAsyncEvent_Post struct
type Msft_WmiProvider_ExecMethodAsyncEvent_Post struct {
	Msft_WmiProvider_OperationEvent_Post

	//
	Flags uint32

	//
	InputParameters interface{}

	//
	MethodName string

	//
	ObjectParameter interface{}

	//
	ObjectPath string

	//
	ResultCode uint32

	//
	StringParameter string
}

// SetFlags sets the value of Flags for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) GetPropertyFlags() (value uint32, err error) {
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
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) SetPropertyInputParameters(value interface{}) (err error) {
	return instance.SetProperty("InputParameters", value)
}

// GetInputParameters gets the value of InputParameters for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) GetPropertyInputParameters() (value interface{}, err error) {
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
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) SetPropertyMethodName(value string) (err error) {
	return instance.SetProperty("MethodName", value)
}

// GetMethodName gets the value of MethodName for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) GetPropertyMethodName() (value string, err error) {
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

// SetObjectParameter sets the value of ObjectParameter for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) SetPropertyObjectParameter(value interface{}) (err error) {
	return instance.SetProperty("ObjectParameter", value)
}

// GetObjectParameter gets the value of ObjectParameter for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) GetPropertyObjectParameter() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ObjectParameter")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetObjectPath sets the value of ObjectPath for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) SetPropertyObjectPath(value string) (err error) {
	return instance.SetProperty("ObjectPath", value)
}

// GetObjectPath gets the value of ObjectPath for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) GetPropertyObjectPath() (value string, err error) {
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

// SetResultCode sets the value of ResultCode for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) SetPropertyResultCode(value uint32) (err error) {
	return instance.SetProperty("ResultCode", value)
}

// GetResultCode gets the value of ResultCode for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) GetPropertyResultCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResultCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStringParameter sets the value of StringParameter for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) SetPropertyStringParameter(value string) (err error) {
	return instance.SetProperty("StringParameter", value)
}

// GetStringParameter gets the value of StringParameter for the instance
func (instance *Msft_WmiProvider_ExecMethodAsyncEvent_Post) GetPropertyStringParameter() (value string, err error) {
	retValue, err := instance.GetProperty("StringParameter")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
