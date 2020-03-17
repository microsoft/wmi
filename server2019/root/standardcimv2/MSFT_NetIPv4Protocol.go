// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetIPv4Protocol struct
type MSFT_NetIPv4Protocol struct {
	MSFT_NetBaseIPProtocol

	//
	MinimumMtu uint32
}

// SetMinimumMtu sets the value of MinimumMtu for the instance
func (instance *MSFT_NetIPv4Protocol) SetPropertyMinimumMtu(value uint32) (err error) {
	return instance.SetProperty("MinimumMtu", value)
}

// GetMinimumMtu gets the value of MinimumMtu for the instance
func (instance *MSFT_NetIPv4Protocol) GetPropertyMinimumMtu() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumMtu")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
