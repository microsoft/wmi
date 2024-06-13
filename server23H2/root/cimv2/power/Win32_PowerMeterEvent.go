// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
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
	return instance.SetProperty("EventSource", (value))
}

// GetEventSource gets the value of EventSource for the instance
func (instance *Win32_PowerMeterEvent) GetPropertyEventSource() (value CIM_PowerMeter, err error) {
	retValue, err := instance.GetProperty("EventSource")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_PowerMeter)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_PowerMeter is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_PowerMeter(valuetmp)

	return
}

// SetEventType sets the value of EventType for the instance
func (instance *Win32_PowerMeterEvent) SetPropertyEventType(value uint32) (err error) {
	return instance.SetProperty("EventType", (value))
}

// GetEventType gets the value of EventType for the instance
func (instance *Win32_PowerMeterEvent) GetPropertyEventType() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
