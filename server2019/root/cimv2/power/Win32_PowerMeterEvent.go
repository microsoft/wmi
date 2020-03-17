// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

// Win32_PowerMeterEvent struct
type Win32_PowerMeterEvent struct {
	__ExtrinsicEvent

	//
	EventSource CIM_PowerMeter

	//
	EventType uint32
}

// SetEventSource sets the value of EventSource for the instance
func (instance *Win32_PowerMeterEvent) SetPropertyEventSource(value CIM_PowerMeter) (err error) {
	return instance.SetProperty("EventSource", value)
}

// GetEventSource gets the value of EventSource for the instance
func (instance *Win32_PowerMeterEvent) GetPropertyEventSource() (value CIM_PowerMeter, err error) {
	retValue, err := instance.GetProperty("EventSource")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_PowerMeter)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventType sets the value of EventType for the instance
func (instance *Win32_PowerMeterEvent) SetPropertyEventType(value uint32) (err error) {
	return instance.SetProperty("EventType", value)
}

// GetEventType gets the value of EventType for the instance
func (instance *Win32_PowerMeterEvent) GetPropertyEventType() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
