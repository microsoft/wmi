// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Msft_WmiProvider_PutInstanceAsyncEvent_Post struct
type Msft_WmiProvider_PutInstanceAsyncEvent_Post struct {
	Msft_WmiProvider_OperationEvent_Post

	//
	Flags uint32

	//
	InstanceObject interface{}

	//
	ObjectParameter interface{}

	//
	ResultCode uint32

	//
	StringParameter string
}

// SetFlags sets the value of Flags for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Post) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Post) GetPropertyFlags() (value uint32, err error) {
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

// SetInstanceObject sets the value of InstanceObject for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Post) SetPropertyInstanceObject(value interface{}) (err error) {
	return instance.SetProperty("InstanceObject", value)
}

// GetInstanceObject gets the value of InstanceObject for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Post) GetPropertyInstanceObject() (value interface{}, err error) {
	retValue, err := instance.GetProperty("InstanceObject")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetObjectParameter sets the value of ObjectParameter for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Post) SetPropertyObjectParameter(value interface{}) (err error) {
	return instance.SetProperty("ObjectParameter", value)
}

// GetObjectParameter gets the value of ObjectParameter for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Post) GetPropertyObjectParameter() (value interface{}, err error) {
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

// SetResultCode sets the value of ResultCode for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Post) SetPropertyResultCode(value uint32) (err error) {
	return instance.SetProperty("ResultCode", value)
}

// GetResultCode gets the value of ResultCode for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Post) GetPropertyResultCode() (value uint32, err error) {
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
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Post) SetPropertyStringParameter(value string) (err error) {
	return instance.SetProperty("StringParameter", value)
}

// GetStringParameter gets the value of StringParameter for the instance
func (instance *Msft_WmiProvider_PutInstanceAsyncEvent_Post) GetPropertyStringParameter() (value string, err error) {
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
