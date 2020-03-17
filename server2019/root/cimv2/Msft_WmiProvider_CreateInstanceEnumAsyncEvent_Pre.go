// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Msft_WmiProvider_CreateInstanceEnumAsyncEvent_Pre struct
type Msft_WmiProvider_CreateInstanceEnumAsyncEvent_Pre struct {
	Msft_WmiProvider_OperationEvent_Pre

	//
	ClassName string

	//
	Flags uint32
}

// SetClassName sets the value of ClassName for the instance
func (instance *Msft_WmiProvider_CreateInstanceEnumAsyncEvent_Pre) SetPropertyClassName(value string) (err error) {
	return instance.SetProperty("ClassName", value)
}

// GetClassName gets the value of ClassName for the instance
func (instance *Msft_WmiProvider_CreateInstanceEnumAsyncEvent_Pre) GetPropertyClassName() (value string, err error) {
	retValue, err := instance.GetProperty("ClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *Msft_WmiProvider_CreateInstanceEnumAsyncEvent_Pre) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *Msft_WmiProvider_CreateInstanceEnumAsyncEvent_Pre) GetPropertyFlags() (value uint32, err error) {
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
