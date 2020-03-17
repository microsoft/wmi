// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DeviceChangeEvent struct
type Win32_DeviceChangeEvent struct {
	__ExtrinsicEvent

	//
	EventType uint16
}

// SetEventType sets the value of EventType for the instance
func (instance *Win32_DeviceChangeEvent) SetPropertyEventType(value uint16) (err error) {
	return instance.SetProperty("EventType", value)
}

// GetEventType gets the value of EventType for the instance
func (instance *Win32_DeviceChangeEvent) GetPropertyEventType() (value uint16, err error) {
	retValue, err := instance.GetProperty("EventType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
