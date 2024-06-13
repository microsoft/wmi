// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFC_LinkEvent struct
type MSFC_LinkEvent struct {
	*WMIEvent

	//
	Active bool

	//
	AdapterWWN []uint8

	//
	EventType uint32

	//
	InstanceName string

	//
	RLIRBuffer []uint8

	//
	RLIRBufferSize uint32
}

func NewMSFC_LinkEventEx1(instance *cim.WmiInstance) (newInstance *MSFC_LinkEvent, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFC_LinkEvent{
		WMIEvent: tmp,
	}
	return
}

func NewMSFC_LinkEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_LinkEvent, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_LinkEvent{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_LinkEvent) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_LinkEvent) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetAdapterWWN sets the value of AdapterWWN for the instance
func (instance *MSFC_LinkEvent) SetPropertyAdapterWWN(value []uint8) (err error) {
	return instance.SetProperty("AdapterWWN", (value))
}

// GetAdapterWWN gets the value of AdapterWWN for the instance
func (instance *MSFC_LinkEvent) GetPropertyAdapterWWN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("AdapterWWN")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetEventType sets the value of EventType for the instance
func (instance *MSFC_LinkEvent) SetPropertyEventType(value uint32) (err error) {
	return instance.SetProperty("EventType", (value))
}

// GetEventType gets the value of EventType for the instance
func (instance *MSFC_LinkEvent) GetPropertyEventType() (value uint32, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSFC_LinkEvent) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_LinkEvent) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetRLIRBuffer sets the value of RLIRBuffer for the instance
func (instance *MSFC_LinkEvent) SetPropertyRLIRBuffer(value []uint8) (err error) {
	return instance.SetProperty("RLIRBuffer", (value))
}

// GetRLIRBuffer gets the value of RLIRBuffer for the instance
func (instance *MSFC_LinkEvent) GetPropertyRLIRBuffer() (value []uint8, err error) {
	retValue, err := instance.GetProperty("RLIRBuffer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetRLIRBufferSize sets the value of RLIRBufferSize for the instance
func (instance *MSFC_LinkEvent) SetPropertyRLIRBufferSize(value uint32) (err error) {
	return instance.SetProperty("RLIRBufferSize", (value))
}

// GetRLIRBufferSize gets the value of RLIRBufferSize for the instance
func (instance *MSFC_LinkEvent) GetPropertyRLIRBufferSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("RLIRBufferSize")
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
