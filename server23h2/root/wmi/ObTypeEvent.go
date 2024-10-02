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

// ObTypeEvent struct
type ObTypeEvent struct {
	*ObTrace

	//
	ObjectType uint16

	//
	Reserved uint16

	//
	TypeName string
}

func NewObTypeEventEx1(instance *cim.WmiInstance) (newInstance *ObTypeEvent, err error) {
	tmp, err := NewObTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ObTypeEvent{
		ObTrace: tmp,
	}
	return
}

func NewObTypeEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ObTypeEvent, err error) {
	tmp, err := NewObTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ObTypeEvent{
		ObTrace: tmp,
	}
	return
}

// SetObjectType sets the value of ObjectType for the instance
func (instance *ObTypeEvent) SetPropertyObjectType(value uint16) (err error) {
	return instance.SetProperty("ObjectType", (value))
}

// GetObjectType gets the value of ObjectType for the instance
func (instance *ObTypeEvent) GetPropertyObjectType() (value uint16, err error) {
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

// SetReserved sets the value of Reserved for the instance
func (instance *ObTypeEvent) SetPropertyReserved(value uint16) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *ObTypeEvent) GetPropertyReserved() (value uint16, err error) {
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

// SetTypeName sets the value of TypeName for the instance
func (instance *ObTypeEvent) SetPropertyTypeName(value string) (err error) {
	return instance.SetProperty("TypeName", (value))
}

// GetTypeName gets the value of TypeName for the instance
func (instance *ObTypeEvent) GetPropertyTypeName() (value string, err error) {
	retValue, err := instance.GetProperty("TypeName")
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
