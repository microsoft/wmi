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

// UmsDirectedSwitchEnd struct
type UmsDirectedSwitchEnd struct {
	*UmsEvent

	//
	PrimaryThreadId uint32

	//
	ProcessId uint32

	//
	ScheduledThreadId uint32

	//
	SwitchFlags uint32
}

func NewUmsDirectedSwitchEndEx1(instance *cim.WmiInstance) (newInstance *UmsDirectedSwitchEnd, err error) {
	tmp, err := NewUmsEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &UmsDirectedSwitchEnd{
		UmsEvent: tmp,
	}
	return
}

func NewUmsDirectedSwitchEndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *UmsDirectedSwitchEnd, err error) {
	tmp, err := NewUmsEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &UmsDirectedSwitchEnd{
		UmsEvent: tmp,
	}
	return
}

// SetPrimaryThreadId sets the value of PrimaryThreadId for the instance
func (instance *UmsDirectedSwitchEnd) SetPropertyPrimaryThreadId(value uint32) (err error) {
	return instance.SetProperty("PrimaryThreadId", (value))
}

// GetPrimaryThreadId gets the value of PrimaryThreadId for the instance
func (instance *UmsDirectedSwitchEnd) GetPropertyPrimaryThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("PrimaryThreadId")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *UmsDirectedSwitchEnd) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *UmsDirectedSwitchEnd) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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

// SetScheduledThreadId sets the value of ScheduledThreadId for the instance
func (instance *UmsDirectedSwitchEnd) SetPropertyScheduledThreadId(value uint32) (err error) {
	return instance.SetProperty("ScheduledThreadId", (value))
}

// GetScheduledThreadId gets the value of ScheduledThreadId for the instance
func (instance *UmsDirectedSwitchEnd) GetPropertyScheduledThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScheduledThreadId")
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

// SetSwitchFlags sets the value of SwitchFlags for the instance
func (instance *UmsDirectedSwitchEnd) SetPropertySwitchFlags(value uint32) (err error) {
	return instance.SetProperty("SwitchFlags", (value))
}

// GetSwitchFlags gets the value of SwitchFlags for the instance
func (instance *UmsDirectedSwitchEnd) GetPropertySwitchFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("SwitchFlags")
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
