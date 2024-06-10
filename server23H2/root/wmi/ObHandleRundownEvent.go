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

// ObHandleRundownEvent struct
type ObHandleRundownEvent struct {
	*ObTrace

	//
	Handle uint32

	//
	Object uint32

	//
	ObjectName string

	//
	ObjectType uint16

	//
	ProcessId uint32
}

func NewObHandleRundownEventEx1(instance *cim.WmiInstance) (newInstance *ObHandleRundownEvent, err error) {
	tmp, err := NewObTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ObHandleRundownEvent{
		ObTrace: tmp,
	}
	return
}

func NewObHandleRundownEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ObHandleRundownEvent, err error) {
	tmp, err := NewObTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ObHandleRundownEvent{
		ObTrace: tmp,
	}
	return
}

// SetHandle sets the value of Handle for the instance
func (instance *ObHandleRundownEvent) SetPropertyHandle(value uint32) (err error) {
	return instance.SetProperty("Handle", (value))
}

// GetHandle gets the value of Handle for the instance
func (instance *ObHandleRundownEvent) GetPropertyHandle() (value uint32, err error) {
	retValue, err := instance.GetProperty("Handle")
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

// SetObject sets the value of Object for the instance
func (instance *ObHandleRundownEvent) SetPropertyObject(value uint32) (err error) {
	return instance.SetProperty("Object", (value))
}

// GetObject gets the value of Object for the instance
func (instance *ObHandleRundownEvent) GetPropertyObject() (value uint32, err error) {
	retValue, err := instance.GetProperty("Object")
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

// SetObjectName sets the value of ObjectName for the instance
func (instance *ObHandleRundownEvent) SetPropertyObjectName(value string) (err error) {
	return instance.SetProperty("ObjectName", (value))
}

// GetObjectName gets the value of ObjectName for the instance
func (instance *ObHandleRundownEvent) GetPropertyObjectName() (value string, err error) {
	retValue, err := instance.GetProperty("ObjectName")
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

// SetObjectType sets the value of ObjectType for the instance
func (instance *ObHandleRundownEvent) SetPropertyObjectType(value uint16) (err error) {
	return instance.SetProperty("ObjectType", (value))
}

// GetObjectType gets the value of ObjectType for the instance
func (instance *ObHandleRundownEvent) GetPropertyObjectType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ObjectType")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *ObHandleRundownEvent) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *ObHandleRundownEvent) GetPropertyProcessId() (value uint32, err error) {
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
