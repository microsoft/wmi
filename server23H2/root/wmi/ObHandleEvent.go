// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ObHandleEvent struct
type ObHandleEvent struct {
	*ObTrace

	//
	Handle uint32

	//
	Object uint32

	//
	ObjectName string

	//
	ObjectType uint16
}

func NewObHandleEventEx1(instance *cim.WmiInstance) (newInstance *ObHandleEvent, err error) {
	tmp, err := NewObTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ObHandleEvent{
		ObTrace: tmp,
	}
	return
}

func NewObHandleEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ObHandleEvent, err error) {
	tmp, err := NewObTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ObHandleEvent{
		ObTrace: tmp,
	}
	return
}

// SetHandle sets the value of Handle for the instance
func (instance *ObHandleEvent) SetPropertyHandle(value uint32) (err error) {
	return instance.SetProperty("Handle", (value))
}

// GetHandle gets the value of Handle for the instance
func (instance *ObHandleEvent) GetPropertyHandle() (value uint32, err error) {
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
func (instance *ObHandleEvent) SetPropertyObject(value uint32) (err error) {
	return instance.SetProperty("Object", (value))
}

// GetObject gets the value of Object for the instance
func (instance *ObHandleEvent) GetPropertyObject() (value uint32, err error) {
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
func (instance *ObHandleEvent) SetPropertyObjectName(value string) (err error) {
	return instance.SetProperty("ObjectName", (value))
}

// GetObjectName gets the value of ObjectName for the instance
func (instance *ObHandleEvent) GetPropertyObjectName() (value string, err error) {
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
func (instance *ObHandleEvent) SetPropertyObjectType(value uint16) (err error) {
	return instance.SetProperty("ObjectType", (value))
}

// GetObjectType gets the value of ObjectType for the instance
func (instance *ObHandleEvent) GetPropertyObjectType() (value uint16, err error) {
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
