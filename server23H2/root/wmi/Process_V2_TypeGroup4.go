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

// Process_V2_TypeGroup4 struct
type Process_V2_TypeGroup4 struct {
	*Process_V2

	//
	Count uint32

	//
	Object uint32

	//
	ProcessId uint32

	//
	Tag uint32
}

func NewProcess_V2_TypeGroup4Ex1(instance *cim.WmiInstance) (newInstance *Process_V2_TypeGroup4, err error) {
	tmp, err := NewProcess_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &Process_V2_TypeGroup4{
		Process_V2: tmp,
	}
	return
}

func NewProcess_V2_TypeGroup4Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Process_V2_TypeGroup4, err error) {
	tmp, err := NewProcess_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Process_V2_TypeGroup4{
		Process_V2: tmp,
	}
	return
}

// SetCount sets the value of Count for the instance
func (instance *Process_V2_TypeGroup4) SetPropertyCount(value uint32) (err error) {
	return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *Process_V2_TypeGroup4) GetPropertyCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("Count")
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
func (instance *Process_V2_TypeGroup4) SetPropertyObject(value uint32) (err error) {
	return instance.SetProperty("Object", (value))
}

// GetObject gets the value of Object for the instance
func (instance *Process_V2_TypeGroup4) GetPropertyObject() (value uint32, err error) {
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *Process_V2_TypeGroup4) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Process_V2_TypeGroup4) GetPropertyProcessId() (value uint32, err error) {
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

// SetTag sets the value of Tag for the instance
func (instance *Process_V2_TypeGroup4) SetPropertyTag(value uint32) (err error) {
	return instance.SetProperty("Tag", (value))
}

// GetTag gets the value of Tag for the instance
func (instance *Process_V2_TypeGroup4) GetPropertyTag() (value uint32, err error) {
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
