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

// SessionPoolAllocFree struct
type SessionPoolAllocFree struct {
	*PoolTrace

	//
	Entry uint32

	//
	NumberOfBytes interface{}

	//
	SessionId uint32

	//
	Tag uint32

	//
	Type uint32
}

func NewSessionPoolAllocFreeEx1(instance *cim.WmiInstance) (newInstance *SessionPoolAllocFree, err error) {
	tmp, err := NewPoolTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SessionPoolAllocFree{
		PoolTrace: tmp,
	}
	return
}

func NewSessionPoolAllocFreeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SessionPoolAllocFree, err error) {
	tmp, err := NewPoolTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SessionPoolAllocFree{
		PoolTrace: tmp,
	}
	return
}

// SetEntry sets the value of Entry for the instance
func (instance *SessionPoolAllocFree) SetPropertyEntry(value uint32) (err error) {
	return instance.SetProperty("Entry", (value))
}

// GetEntry gets the value of Entry for the instance
func (instance *SessionPoolAllocFree) GetPropertyEntry() (value uint32, err error) {
	retValue, err := instance.GetProperty("Entry")
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

// SetNumberOfBytes sets the value of NumberOfBytes for the instance
func (instance *SessionPoolAllocFree) SetPropertyNumberOfBytes(value interface{}) (err error) {
	return instance.SetProperty("NumberOfBytes", (value))
}

// GetNumberOfBytes gets the value of NumberOfBytes for the instance
func (instance *SessionPoolAllocFree) GetPropertyNumberOfBytes() (value interface{}, err error) {
	retValue, err := instance.GetProperty("NumberOfBytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetSessionId sets the value of SessionId for the instance
func (instance *SessionPoolAllocFree) SetPropertySessionId(value uint32) (err error) {
	return instance.SetProperty("SessionId", (value))
}

// GetSessionId gets the value of SessionId for the instance
func (instance *SessionPoolAllocFree) GetPropertySessionId() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionId")
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

// SetTag sets the value of Tag for the instance
func (instance *SessionPoolAllocFree) SetPropertyTag(value uint32) (err error) {
	return instance.SetProperty("Tag", (value))
}

// GetTag gets the value of Tag for the instance
func (instance *SessionPoolAllocFree) GetPropertyTag() (value uint32, err error) {
	retValue, err := instance.GetProperty("Tag")
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

// SetType sets the value of Type for the instance
func (instance *SessionPoolAllocFree) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *SessionPoolAllocFree) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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
