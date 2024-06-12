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

// StackWalk_TypeGroup1 struct
type StackWalk_TypeGroup1 struct {
	*StackWalk

	//
	key uint32

	//
	StackFrame []uint32
}

func NewStackWalk_TypeGroup1Ex1(instance *cim.WmiInstance) (newInstance *StackWalk_TypeGroup1, err error) {
	tmp, err := NewStackWalkEx1(instance)

	if err != nil {
		return
	}
	newInstance = &StackWalk_TypeGroup1{
		StackWalk: tmp,
	}
	return
}

func NewStackWalk_TypeGroup1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *StackWalk_TypeGroup1, err error) {
	tmp, err := NewStackWalkEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &StackWalk_TypeGroup1{
		StackWalk: tmp,
	}
	return
}

// Setkey sets the value of key for the instance
func (instance *StackWalk_TypeGroup1) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *StackWalk_TypeGroup1) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
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

// SetStackFrame sets the value of StackFrame for the instance
func (instance *StackWalk_TypeGroup1) SetPropertyStackFrame(value []uint32) (err error) {
	return instance.SetProperty("StackFrame", (value))
}

// GetStackFrame gets the value of StackFrame for the instance
func (instance *StackWalk_TypeGroup1) GetPropertyStackFrame() (value []uint32, err error) {
	retValue, err := instance.GetProperty("StackFrame")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}
