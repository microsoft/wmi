// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Msft_WmiProvider_PutClassAsyncEvent_Pre struct
type Msft_WmiProvider_PutClassAsyncEvent_Pre struct {
	Msft_WmiProvider_OperationEvent_Pre

	//
	ClassObject interface{}

	//
	Flags uint32
}

// SetClassObject sets the value of ClassObject for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Pre) SetPropertyClassObject(value interface{}) (err error) {
	return instance.SetProperty("ClassObject", value)
}

// GetClassObject gets the value of ClassObject for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Pre) GetPropertyClassObject() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ClassObject")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Pre) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *Msft_WmiProvider_PutClassAsyncEvent_Pre) GetPropertyFlags() (value uint32, err error) {
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
