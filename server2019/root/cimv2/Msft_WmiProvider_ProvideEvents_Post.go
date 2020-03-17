// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Msft_WmiProvider_ProvideEvents_Post struct
type Msft_WmiProvider_ProvideEvents_Post struct {
	Msft_WmiProvider_OperationEvent_Post

	//
	Flags uint32

	//
	Result uint32
}

// SetFlags sets the value of Flags for the instance
func (instance *Msft_WmiProvider_ProvideEvents_Post) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *Msft_WmiProvider_ProvideEvents_Post) GetPropertyFlags() (value uint32, err error) {
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

// SetResult sets the value of Result for the instance
func (instance *Msft_WmiProvider_ProvideEvents_Post) SetPropertyResult(value uint32) (err error) {
	return instance.SetProperty("Result", value)
}

// GetResult gets the value of Result for the instance
func (instance *Msft_WmiProvider_ProvideEvents_Post) GetPropertyResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("Result")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
