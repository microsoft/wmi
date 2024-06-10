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

// AutoBoostSetFloor struct
type AutoBoostSetFloor struct {
	*Thread_V2

	//
	BoostFlags uint8

	//
	IoPriorities uint8

	//
	Lock uint32

	//
	NewCpuPriorityFloor uint8

	//
	OldCpuPriority uint8

	//
	ThreadId uint32
}

func NewAutoBoostSetFloorEx1(instance *cim.WmiInstance) (newInstance *AutoBoostSetFloor, err error) {
	tmp, err := NewThread_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &AutoBoostSetFloor{
		Thread_V2: tmp,
	}
	return
}

func NewAutoBoostSetFloorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AutoBoostSetFloor, err error) {
	tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AutoBoostSetFloor{
		Thread_V2: tmp,
	}
	return
}

// SetBoostFlags sets the value of BoostFlags for the instance
func (instance *AutoBoostSetFloor) SetPropertyBoostFlags(value uint8) (err error) {
	return instance.SetProperty("BoostFlags", (value))
}

// GetBoostFlags gets the value of BoostFlags for the instance
func (instance *AutoBoostSetFloor) GetPropertyBoostFlags() (value uint8, err error) {
	retValue, err := instance.GetProperty("BoostFlags")
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

// SetIoPriorities sets the value of IoPriorities for the instance
func (instance *AutoBoostSetFloor) SetPropertyIoPriorities(value uint8) (err error) {
	return instance.SetProperty("IoPriorities", (value))
}

// GetIoPriorities gets the value of IoPriorities for the instance
func (instance *AutoBoostSetFloor) GetPropertyIoPriorities() (value uint8, err error) {
	retValue, err := instance.GetProperty("IoPriorities")
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

// SetLock sets the value of Lock for the instance
func (instance *AutoBoostSetFloor) SetPropertyLock(value uint32) (err error) {
	return instance.SetProperty("Lock", (value))
}

// GetLock gets the value of Lock for the instance
func (instance *AutoBoostSetFloor) GetPropertyLock() (value uint32, err error) {
	retValue, err := instance.GetProperty("Lock")
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

// SetNewCpuPriorityFloor sets the value of NewCpuPriorityFloor for the instance
func (instance *AutoBoostSetFloor) SetPropertyNewCpuPriorityFloor(value uint8) (err error) {
	return instance.SetProperty("NewCpuPriorityFloor", (value))
}

// GetNewCpuPriorityFloor gets the value of NewCpuPriorityFloor for the instance
func (instance *AutoBoostSetFloor) GetPropertyNewCpuPriorityFloor() (value uint8, err error) {
	retValue, err := instance.GetProperty("NewCpuPriorityFloor")
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

// SetOldCpuPriority sets the value of OldCpuPriority for the instance
func (instance *AutoBoostSetFloor) SetPropertyOldCpuPriority(value uint8) (err error) {
	return instance.SetProperty("OldCpuPriority", (value))
}

// GetOldCpuPriority gets the value of OldCpuPriority for the instance
func (instance *AutoBoostSetFloor) GetPropertyOldCpuPriority() (value uint8, err error) {
	retValue, err := instance.GetProperty("OldCpuPriority")
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

// SetThreadId sets the value of ThreadId for the instance
func (instance *AutoBoostSetFloor) SetPropertyThreadId(value uint32) (err error) {
	return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *AutoBoostSetFloor) GetPropertyThreadId() (value uint32, err error) {
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
