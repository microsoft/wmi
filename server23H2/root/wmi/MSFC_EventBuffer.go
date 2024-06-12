// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFC_EventBuffer struct
type MSFC_EventBuffer struct {
	*cim.WmiInstance

	//
	EventInfo []uint32

	//
	EventType uint32
}

func NewMSFC_EventBufferEx1(instance *cim.WmiInstance) (newInstance *MSFC_EventBuffer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_EventBuffer{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_EventBufferEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_EventBuffer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_EventBuffer{
		WmiInstance: tmp,
	}
	return
}

// SetEventInfo sets the value of EventInfo for the instance
func (instance *MSFC_EventBuffer) SetPropertyEventInfo(value []uint32) (err error) {
	return instance.SetProperty("EventInfo", (value))
}

// GetEventInfo gets the value of EventInfo for the instance
func (instance *MSFC_EventBuffer) GetPropertyEventInfo() (value []uint32, err error) {
	retValue, err := instance.GetProperty("EventInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetEventType sets the value of EventType for the instance
func (instance *MSFC_EventBuffer) SetPropertyEventType(value uint32) (err error) {
	return instance.SetProperty("EventType", (value))
}

// GetEventType gets the value of EventType for the instance
func (instance *MSFC_EventBuffer) GetPropertyEventType() (value uint32, err error) {
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
