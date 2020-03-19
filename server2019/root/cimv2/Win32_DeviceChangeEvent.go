// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_DeviceChangeEvent struct
type Win32_DeviceChangeEvent struct {
	*__ExtrinsicEvent

	//
	EventType uint16
}

func NewWin32_DeviceChangeEventEx1(instance *cim.WmiInstance) (newInstance *Win32_DeviceChangeEvent, err error) {
	tmp, err := New__ExtrinsicEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_DeviceChangeEvent{
		__ExtrinsicEvent: tmp,
	}
	return
}

func NewWin32_DeviceChangeEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_DeviceChangeEvent, err error) {
	tmp, err := New__ExtrinsicEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_DeviceChangeEvent{
		__ExtrinsicEvent: tmp,
	}
	return
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
