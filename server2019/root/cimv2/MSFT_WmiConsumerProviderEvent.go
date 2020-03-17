// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_WmiConsumerProviderEvent struct
type MSFT_WmiConsumerProviderEvent struct {
	MSFT_WmiProviderEvent

	//
	Machine string
}

// SetMachine sets the value of Machine for the instance
func (instance *MSFT_WmiConsumerProviderEvent) SetPropertyMachine(value string) (err error) {
	return instance.SetProperty("Machine", value)
}

// GetMachine gets the value of Machine for the instance
func (instance *MSFT_WmiConsumerProviderEvent) GetPropertyMachine() (value string, err error) {
	retValue, err := instance.GetProperty("Machine")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
