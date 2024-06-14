// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFC_TargetEvent struct
type MSFC_TargetEvent struct {
	*WMIEvent

	//
	Active bool

	//
	DiscoveredPortWWN []uint8

	//
	EventType uint32

	//
	InstanceName string

	//
	PortWWN []uint8
}

func NewMSFC_TargetEventEx1(instance *cim.WmiInstance) (newInstance *MSFC_TargetEvent, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFC_TargetEvent{
		WMIEvent: tmp,
	}
	return
}

func NewMSFC_TargetEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_TargetEvent, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_TargetEvent{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_TargetEvent) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_TargetEvent) GetPropertyActive() (value bool, err error) {
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

// SetDiscoveredPortWWN sets the value of DiscoveredPortWWN for the instance
func (instance *MSFC_TargetEvent) SetPropertyDiscoveredPortWWN(value []uint8) (err error) {
	return instance.SetProperty("DiscoveredPortWWN", (value))
}

// GetDiscoveredPortWWN gets the value of DiscoveredPortWWN for the instance
func (instance *MSFC_TargetEvent) GetPropertyDiscoveredPortWWN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("DiscoveredPortWWN")
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
func (instance *MSFC_TargetEvent) SetPropertyEventType(value uint32) (err error) {
	return instance.SetProperty("EventType", (value))
}

// GetEventType gets the value of EventType for the instance
func (instance *MSFC_TargetEvent) GetPropertyEventType() (value uint32, err error) {
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
func (instance *MSFC_TargetEvent) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_TargetEvent) GetPropertyInstanceName() (value string, err error) {
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

// SetPortWWN sets the value of PortWWN for the instance
func (instance *MSFC_TargetEvent) SetPropertyPortWWN(value []uint8) (err error) {
	return instance.SetProperty("PortWWN", (value))
}

// GetPortWWN gets the value of PortWWN for the instance
func (instance *MSFC_TargetEvent) GetPropertyPortWWN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("PortWWN")
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
