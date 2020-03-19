// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PowerMeterEvent struct
type Win32_PowerMeterEvent struct {
	*__ExtrinsicEvent

	//
	EventSource CIM_PowerMeter

	//
	EventType uint32
}

func NewWin32_PowerMeterEventEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerMeterEvent, err error) {
	tmp, err := New__ExtrinsicEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerMeterEvent{
		__ExtrinsicEvent: tmp,
	}
	return
}

func NewWin32_PowerMeterEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PowerMeterEvent, err error) {
	tmp, err := New__ExtrinsicEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerMeterEvent{
		__ExtrinsicEvent: tmp,
	}
	return
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
