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

// Thread_TypeGroup1 struct
type Thread_TypeGroup1 struct {
	*Thread_V4

	//
	Affinity uint32

	//
	BasePriority uint8

	//
	IoPriority uint8

	//
	PagePriority uint8

	//
	ProcessId uint32

	//
	StackBase uint32

	//
	StackLimit uint32

	//
	SubProcessTag uint32

	//
	TebBase uint32

	//
	ThreadFlags uint8

	//
	ThreadName string

	//
	TThreadId uint32

	//
	UserStackBase uint32

	//
	UserStackLimit uint32

	//
	Win32StartAddr uint32
}

func NewThread_TypeGroup1Ex1(instance *cim.WmiInstance) (newInstance *Thread_TypeGroup1, err error) {
	tmp, err := NewThread_V4Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &Thread_TypeGroup1{
		Thread_V4: tmp,
	}
	return
}

func NewThread_TypeGroup1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Thread_TypeGroup1, err error) {
	tmp, err := NewThread_V4Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Thread_TypeGroup1{
		Thread_V4: tmp,
	}
	return
}

// SetAffinity sets the value of Affinity for the instance
func (instance *Thread_TypeGroup1) SetPropertyAffinity(value uint32) (err error) {
	return instance.SetProperty("Affinity", (value))
}

// GetAffinity gets the value of Affinity for the instance
func (instance *Thread_TypeGroup1) GetPropertyAffinity() (value uint32, err error) {
	retValue, err := instance.GetProperty("Affinity")
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

// SetBasePriority sets the value of BasePriority for the instance
func (instance *Thread_TypeGroup1) SetPropertyBasePriority(value uint8) (err error) {
	return instance.SetProperty("BasePriority", (value))
}

// GetBasePriority gets the value of BasePriority for the instance
func (instance *Thread_TypeGroup1) GetPropertyBasePriority() (value uint8, err error) {
	retValue, err := instance.GetProperty("BasePriority")
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

// SetIoPriority sets the value of IoPriority for the instance
func (instance *Thread_TypeGroup1) SetPropertyIoPriority(value uint8) (err error) {
	return instance.SetProperty("IoPriority", (value))
}

// GetIoPriority gets the value of IoPriority for the instance
func (instance *Thread_TypeGroup1) GetPropertyIoPriority() (value uint8, err error) {
	retValue, err := instance.GetProperty("IoPriority")
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

// SetPagePriority sets the value of PagePriority for the instance
func (instance *Thread_TypeGroup1) SetPropertyPagePriority(value uint8) (err error) {
	return instance.SetProperty("PagePriority", (value))
}

// GetPagePriority gets the value of PagePriority for the instance
func (instance *Thread_TypeGroup1) GetPropertyPagePriority() (value uint8, err error) {
	retValue, err := instance.GetProperty("PagePriority")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *Thread_TypeGroup1) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Thread_TypeGroup1) GetPropertyProcessId() (value uint32, err error) {
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
func (instance *Thread_TypeGroup1) SetPropertyStackBase(value uint32) (err error) {
	return instance.SetProperty("StackBase", (value))
}

// GetStackBase gets the value of StackBase for the instance
func (instance *Thread_TypeGroup1) GetPropertyStackBase() (value uint32, err error) {
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
func (instance *Thread_TypeGroup1) SetPropertyStackLimit(value uint32) (err error) {
	return instance.SetProperty("StackLimit", (value))
}

// GetStackLimit gets the value of StackLimit for the instance
func (instance *Thread_TypeGroup1) GetPropertyStackLimit() (value uint32, err error) {
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

// SetSubProcessTag sets the value of SubProcessTag for the instance
func (instance *Thread_TypeGroup1) SetPropertySubProcessTag(value uint32) (err error) {
	return instance.SetProperty("SubProcessTag", (value))
}

// GetSubProcessTag gets the value of SubProcessTag for the instance
func (instance *Thread_TypeGroup1) GetPropertySubProcessTag() (value uint32, err error) {
	retValue, err := instance.GetProperty("SubProcessTag")
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

// SetTebBase sets the value of TebBase for the instance
func (instance *Thread_TypeGroup1) SetPropertyTebBase(value uint32) (err error) {
	return instance.SetProperty("TebBase", (value))
}

// GetTebBase gets the value of TebBase for the instance
func (instance *Thread_TypeGroup1) GetPropertyTebBase() (value uint32, err error) {
	retValue, err := instance.GetProperty("TebBase")
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

// SetThreadFlags sets the value of ThreadFlags for the instance
func (instance *Thread_TypeGroup1) SetPropertyThreadFlags(value uint8) (err error) {
	return instance.SetProperty("ThreadFlags", (value))
}

// GetThreadFlags gets the value of ThreadFlags for the instance
func (instance *Thread_TypeGroup1) GetPropertyThreadFlags() (value uint8, err error) {
	retValue, err := instance.GetProperty("ThreadFlags")
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

// SetThreadName sets the value of ThreadName for the instance
func (instance *Thread_TypeGroup1) SetPropertyThreadName(value string) (err error) {
	return instance.SetProperty("ThreadName", (value))
}

// GetThreadName gets the value of ThreadName for the instance
func (instance *Thread_TypeGroup1) GetPropertyThreadName() (value string, err error) {
	retValue, err := instance.GetProperty("ThreadName")
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

// SetTThreadId sets the value of TThreadId for the instance
func (instance *Thread_TypeGroup1) SetPropertyTThreadId(value uint32) (err error) {
	return instance.SetProperty("TThreadId", (value))
}

// GetTThreadId gets the value of TThreadId for the instance
func (instance *Thread_TypeGroup1) GetPropertyTThreadId() (value uint32, err error) {
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
func (instance *Thread_TypeGroup1) SetPropertyUserStackBase(value uint32) (err error) {
	return instance.SetProperty("UserStackBase", (value))
}

// GetUserStackBase gets the value of UserStackBase for the instance
func (instance *Thread_TypeGroup1) GetPropertyUserStackBase() (value uint32, err error) {
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
func (instance *Thread_TypeGroup1) SetPropertyUserStackLimit(value uint32) (err error) {
	return instance.SetProperty("UserStackLimit", (value))
}

// GetUserStackLimit gets the value of UserStackLimit for the instance
func (instance *Thread_TypeGroup1) GetPropertyUserStackLimit() (value uint32, err error) {
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

// SetWin32StartAddr sets the value of Win32StartAddr for the instance
func (instance *Thread_TypeGroup1) SetPropertyWin32StartAddr(value uint32) (err error) {
	return instance.SetProperty("Win32StartAddr", (value))
}

// GetWin32StartAddr gets the value of Win32StartAddr for the instance
func (instance *Thread_TypeGroup1) GetPropertyWin32StartAddr() (value uint32, err error) {
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
