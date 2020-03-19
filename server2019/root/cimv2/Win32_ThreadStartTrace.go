// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ThreadStartTrace struct
type Win32_ThreadStartTrace struct {
	*Win32_ThreadTrace

	//
	StackBase uint64

	//
	StackLimit uint64

	//
	StartAddr uint64

	//
	UserStackBase uint64

	//
	UserStackLimit uint64

	//
	WaitMode uint32

	//
	Win32StartAddr uint64
}

func NewWin32_ThreadStartTraceEx1(instance *cim.WmiInstance) (newInstance *Win32_ThreadStartTrace, err error) {
	tmp, err := NewWin32_ThreadTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_ThreadStartTrace{
		Win32_ThreadTrace: tmp,
	}
	return
}

func NewWin32_ThreadStartTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ThreadStartTrace, err error) {
	tmp, err := NewWin32_ThreadTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ThreadStartTrace{
		Win32_ThreadTrace: tmp,
	}
	return
}

// SetStackBase sets the value of StackBase for the instance
func (instance *Win32_ThreadStartTrace) SetPropertyStackBase(value uint64) (err error) {
	return instance.SetProperty("StackBase", value)
}

// GetStackBase gets the value of StackBase for the instance
func (instance *Win32_ThreadStartTrace) GetPropertyStackBase() (value uint64, err error) {
	retValue, err := instance.GetProperty("StackBase")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStackLimit sets the value of StackLimit for the instance
func (instance *Win32_ThreadStartTrace) SetPropertyStackLimit(value uint64) (err error) {
	return instance.SetProperty("StackLimit", value)
}

// GetStackLimit gets the value of StackLimit for the instance
func (instance *Win32_ThreadStartTrace) GetPropertyStackLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("StackLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartAddr sets the value of StartAddr for the instance
func (instance *Win32_ThreadStartTrace) SetPropertyStartAddr(value uint64) (err error) {
	return instance.SetProperty("StartAddr", value)
}

// GetStartAddr gets the value of StartAddr for the instance
func (instance *Win32_ThreadStartTrace) GetPropertyStartAddr() (value uint64, err error) {
	retValue, err := instance.GetProperty("StartAddr")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserStackBase sets the value of UserStackBase for the instance
func (instance *Win32_ThreadStartTrace) SetPropertyUserStackBase(value uint64) (err error) {
	return instance.SetProperty("UserStackBase", value)
}

// GetUserStackBase gets the value of UserStackBase for the instance
func (instance *Win32_ThreadStartTrace) GetPropertyUserStackBase() (value uint64, err error) {
	retValue, err := instance.GetProperty("UserStackBase")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserStackLimit sets the value of UserStackLimit for the instance
func (instance *Win32_ThreadStartTrace) SetPropertyUserStackLimit(value uint64) (err error) {
	return instance.SetProperty("UserStackLimit", value)
}

// GetUserStackLimit gets the value of UserStackLimit for the instance
func (instance *Win32_ThreadStartTrace) GetPropertyUserStackLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("UserStackLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWaitMode sets the value of WaitMode for the instance
func (instance *Win32_ThreadStartTrace) SetPropertyWaitMode(value uint32) (err error) {
	return instance.SetProperty("WaitMode", value)
}

// GetWaitMode gets the value of WaitMode for the instance
func (instance *Win32_ThreadStartTrace) GetPropertyWaitMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("WaitMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWin32StartAddr sets the value of Win32StartAddr for the instance
func (instance *Win32_ThreadStartTrace) SetPropertyWin32StartAddr(value uint64) (err error) {
	return instance.SetProperty("Win32StartAddr", value)
}

// GetWin32StartAddr gets the value of Win32StartAddr for the instance
func (instance *Win32_ThreadStartTrace) GetPropertyWin32StartAddr() (value uint64, err error) {
	retValue, err := instance.GetProperty("Win32StartAddr")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
