// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_WmiConsumerProviderSinkLoaded struct
type MSFT_WmiConsumerProviderSinkLoaded struct {
	MSFT_WmiConsumerProviderEvent

	//
	Consumer __EventConsumer
}

// SetConsumer sets the value of Consumer for the instance
func (instance *MSFT_WmiConsumerProviderSinkLoaded) SetPropertyConsumer(value __EventConsumer) (err error) {
	return instance.SetProperty("Consumer", value)
}

// GetConsumer gets the value of Consumer for the instance
func (instance *MSFT_WmiConsumerProviderSinkLoaded) GetPropertyConsumer() (value __EventConsumer, err error) {
	retValue, err := instance.GetProperty("Consumer")
	if err != nil {
		return
	}
	value, ok := retValue.(__EventConsumer)
	if !ok {
		// TODO: Set an error
	}
	return
}
