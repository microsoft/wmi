// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// UmsScheduledPark struct
type UmsScheduledPark struct {
	*UmsEvent

	//
	ParkFlags uint32

	//
	ProcessId uint32

	//
	ScheduledThreadId uint32
}

func NewUmsScheduledParkEx1(instance *cim.WmiInstance) (newInstance *UmsScheduledPark, err error) {
	tmp, err := NewUmsEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &UmsScheduledPark{
		UmsEvent: tmp,
	}
	return
}

func NewUmsScheduledParkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *UmsScheduledPark, err error) {
	tmp, err := NewUmsEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &UmsScheduledPark{
		UmsEvent: tmp,
	}
	return
}

// SetParkFlags sets the value of ParkFlags for the instance
func (instance *UmsScheduledPark) SetPropertyParkFlags(value uint32) (err error) {
	return instance.SetProperty("ParkFlags", (value))
}

// GetParkFlags gets the value of ParkFlags for the instance
func (instance *UmsScheduledPark) GetPropertyParkFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("ParkFlags")
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
func (instance *UmsScheduledPark) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *UmsScheduledPark) GetPropertyProcessId() (value uint32, err error) {
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
func (instance *UmsScheduledPark) SetPropertyScheduledThreadId(value uint32) (err error) {
	return instance.SetProperty("ScheduledThreadId", (value))
}

// GetScheduledThreadId gets the value of ScheduledThreadId for the instance
func (instance *UmsScheduledPark) GetPropertyScheduledThreadId() (value uint32, err error) {
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
