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

// ThreadPriority struct
type ThreadPriority struct {
	*Thread_V3

	//
	NewPriority uint8

	//
	OldPriority uint8

	//
	Reserved uint16

	//
	ThreadId uint32
}

func NewThreadPriorityEx1(instance *cim.WmiInstance) (newInstance *ThreadPriority, err error) {
	tmp, err := NewThread_V3Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &ThreadPriority{
		Thread_V3: tmp,
	}
	return
}

func NewThreadPriorityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ThreadPriority, err error) {
	tmp, err := NewThread_V3Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ThreadPriority{
		Thread_V3: tmp,
	}
	return
}

// SetNewPriority sets the value of NewPriority for the instance
func (instance *ThreadPriority) SetPropertyNewPriority(value uint8) (err error) {
	return instance.SetProperty("NewPriority", (value))
}

// GetNewPriority gets the value of NewPriority for the instance
func (instance *ThreadPriority) GetPropertyNewPriority() (value uint8, err error) {
	retValue, err := instance.GetProperty("NewPriority")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetOldPriority sets the value of OldPriority for the instance
func (instance *ThreadPriority) SetPropertyOldPriority(value uint8) (err error) {
	return instance.SetProperty("OldPriority", (value))
}

// GetOldPriority gets the value of OldPriority for the instance
func (instance *ThreadPriority) GetPropertyOldPriority() (value uint8, err error) {
	retValue, err := instance.GetProperty("OldPriority")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetReserved sets the value of Reserved for the instance
func (instance *ThreadPriority) SetPropertyReserved(value uint16) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *ThreadPriority) GetPropertyReserved() (value uint16, err error) {
	retValue, err := instance.GetProperty("Reserved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetThreadId sets the value of ThreadId for the instance
func (instance *ThreadPriority) SetPropertyThreadId(value uint32) (err error) {
	return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *ThreadPriority) GetPropertyThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadId")
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
