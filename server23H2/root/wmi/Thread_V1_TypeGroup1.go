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

// Thread_V1_TypeGroup1 struct
type Thread_V1_TypeGroup1 struct {
	*Thread_V1

	//
	ProcessId uint32

	//
	StackBase uint32

	//
	StackLimit uint32

	//
	StartAddr uint32

	//
	TThreadId uint32

	//
	UserStackBase uint32

	//
	UserStackLimit uint32

	//
	WaitMode int8

	//
	Win32StartAddr uint32
}

func NewThread_V1_TypeGroup1Ex1(instance *cim.WmiInstance) (newInstance *Thread_V1_TypeGroup1, err error) {
	tmp, err := NewThread_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &Thread_V1_TypeGroup1{
		Thread_V1: tmp,
	}
	return
}

func NewThread_V1_TypeGroup1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Thread_V1_TypeGroup1, err error) {
	tmp, err := NewThread_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Thread_V1_TypeGroup1{
		Thread_V1: tmp,
	}
	return
}

// SetProcessId sets the value of ProcessId for the instance
func (instance *Thread_V1_TypeGroup1) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Thread_V1_TypeGroup1) GetPropertyProcessId() (value uint32, err error) {
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

// SetStackBase sets the value of StackBase for the instance
func (instance *Thread_V1_TypeGroup1) SetPropertyStackBase(value uint32) (err error) {
	return instance.SetProperty("StackBase", (value))
}

// GetStackBase gets the value of StackBase for the instance
func (instance *Thread_V1_TypeGroup1) GetPropertyStackBase() (value uint32, err error) {
	retValue, err := instance.GetProperty("StackBase")
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

// SetStackLimit sets the value of StackLimit for the instance
func (instance *Thread_V1_TypeGroup1) SetPropertyStackLimit(value uint32) (err error) {
	return instance.SetProperty("StackLimit", (value))
}

// GetStackLimit gets the value of StackLimit for the instance
func (instance *Thread_V1_TypeGroup1) GetPropertyStackLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("StackLimit")
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

// SetStartAddr sets the value of StartAddr for the instance
func (instance *Thread_V1_TypeGroup1) SetPropertyStartAddr(value uint32) (err error) {
	return instance.SetProperty("StartAddr", (value))
}

// GetStartAddr gets the value of StartAddr for the instance
func (instance *Thread_V1_TypeGroup1) GetPropertyStartAddr() (value uint32, err error) {
	retValue, err := instance.GetProperty("StartAddr")
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

// SetTThreadId sets the value of TThreadId for the instance
func (instance *Thread_V1_TypeGroup1) SetPropertyTThreadId(value uint32) (err error) {
	return instance.SetProperty("TThreadId", (value))
}

// GetTThreadId gets the value of TThreadId for the instance
func (instance *Thread_V1_TypeGroup1) GetPropertyTThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("TThreadId")
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

// SetUserStackBase sets the value of UserStackBase for the instance
func (instance *Thread_V1_TypeGroup1) SetPropertyUserStackBase(value uint32) (err error) {
	return instance.SetProperty("UserStackBase", (value))
}

// GetUserStackBase gets the value of UserStackBase for the instance
func (instance *Thread_V1_TypeGroup1) GetPropertyUserStackBase() (value uint32, err error) {
	retValue, err := instance.GetProperty("UserStackBase")
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

// SetUserStackLimit sets the value of UserStackLimit for the instance
func (instance *Thread_V1_TypeGroup1) SetPropertyUserStackLimit(value uint32) (err error) {
	return instance.SetProperty("UserStackLimit", (value))
}

// GetUserStackLimit gets the value of UserStackLimit for the instance
func (instance *Thread_V1_TypeGroup1) GetPropertyUserStackLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("UserStackLimit")
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

// SetWaitMode sets the value of WaitMode for the instance
func (instance *Thread_V1_TypeGroup1) SetPropertyWaitMode(value int8) (err error) {
	return instance.SetProperty("WaitMode", (value))
}

// GetWaitMode gets the value of WaitMode for the instance
func (instance *Thread_V1_TypeGroup1) GetPropertyWaitMode() (value int8, err error) {
	retValue, err := instance.GetProperty("WaitMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int8(valuetmp)

	return
}

// SetWin32StartAddr sets the value of Win32StartAddr for the instance
func (instance *Thread_V1_TypeGroup1) SetPropertyWin32StartAddr(value uint32) (err error) {
	return instance.SetProperty("Win32StartAddr", (value))
}

// GetWin32StartAddr gets the value of Win32StartAddr for the instance
func (instance *Thread_V1_TypeGroup1) GetPropertyWin32StartAddr() (value uint32, err error) {
	retValue, err := instance.GetProperty("Win32StartAddr")
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
