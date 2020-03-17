// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_WmiThreadPoolEvent struct
type MSFT_WmiThreadPoolEvent struct {
	MSFT_WmiEssEvent

	//
	ThreadId uint32
}

// SetThreadId sets the value of ThreadId for the instance
func (instance *MSFT_WmiThreadPoolEvent) SetPropertyThreadId(value uint32) (err error) {
	return instance.SetProperty("ThreadId", value)
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *MSFT_WmiThreadPoolEvent) GetPropertyThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
