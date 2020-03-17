// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetReadfileTimeout struct
type MSFT_NetReadfileTimeout struct {
	MSFT_SCMEventLogEvent

	//
	Milliseconds uint32
}

// SetMilliseconds sets the value of Milliseconds for the instance
func (instance *MSFT_NetReadfileTimeout) SetPropertyMilliseconds(value uint32) (err error) {
	return instance.SetProperty("Milliseconds", value)
}

// GetMilliseconds gets the value of Milliseconds for the instance
func (instance *MSFT_NetReadfileTimeout) GetPropertyMilliseconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("Milliseconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
