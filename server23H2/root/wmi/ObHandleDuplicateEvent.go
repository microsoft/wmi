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

// ObHandleDuplicateEvent struct
type ObHandleDuplicateEvent struct {
	*ObTrace

	//
	Object uint32

	//
	ObjectType uint16

	//
	SourceHandle uint32

	//
	TargetHandle uint32

	//
	TargetProcessId uint32
}

func NewObHandleDuplicateEventEx1(instance *cim.WmiInstance) (newInstance *ObHandleDuplicateEvent, err error) {
	tmp, err := NewObTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ObHandleDuplicateEvent{
		ObTrace: tmp,
	}
	return
}

func NewObHandleDuplicateEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ObHandleDuplicateEvent, err error) {
	tmp, err := NewObTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ObHandleDuplicateEvent{
		ObTrace: tmp,
	}
	return
}

// SetObject sets the value of Object for the instance
func (instance *ObHandleDuplicateEvent) SetPropertyObject(value uint32) (err error) {
	return instance.SetProperty("Object", (value))
}

// GetObject gets the value of Object for the instance
func (instance *ObHandleDuplicateEvent) GetPropertyObject() (value uint32, err error) {
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

// SetObjectType sets the value of ObjectType for the instance
func (instance *ObHandleDuplicateEvent) SetPropertyObjectType(value uint16) (err error) {
	return instance.SetProperty("ObjectType", (value))
}

// GetObjectType gets the value of ObjectType for the instance
func (instance *ObHandleDuplicateEvent) GetPropertyObjectType() (value uint16, err error) {
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

// SetSourceHandle sets the value of SourceHandle for the instance
func (instance *ObHandleDuplicateEvent) SetPropertySourceHandle(value uint32) (err error) {
	return instance.SetProperty("SourceHandle", (value))
}

// GetSourceHandle gets the value of SourceHandle for the instance
func (instance *ObHandleDuplicateEvent) GetPropertySourceHandle() (value uint32, err error) {
	retValue, err := instance.GetProperty("SourceHandle")
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

// SetTargetHandle sets the value of TargetHandle for the instance
func (instance *ObHandleDuplicateEvent) SetPropertyTargetHandle(value uint32) (err error) {
	return instance.SetProperty("TargetHandle", (value))
}

// GetTargetHandle gets the value of TargetHandle for the instance
func (instance *ObHandleDuplicateEvent) GetPropertyTargetHandle() (value uint32, err error) {
	retValue, err := instance.GetProperty("TargetHandle")
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

// SetTargetProcessId sets the value of TargetProcessId for the instance
func (instance *ObHandleDuplicateEvent) SetPropertyTargetProcessId(value uint32) (err error) {
	return instance.SetProperty("TargetProcessId", (value))
}

// GetTargetProcessId gets the value of TargetProcessId for the instance
func (instance *ObHandleDuplicateEvent) GetPropertyTargetProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("TargetProcessId")
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
