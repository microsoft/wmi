// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Thread_V1_TypeGroup2 struct
type Thread_V1_TypeGroup2 struct {
	*Thread_V1

	//
	ProcessId uint32

	//
	TThreadId uint32
}

func NewThread_V1_TypeGroup2Ex1(instance *cim.WmiInstance) (newInstance *Thread_V1_TypeGroup2, err error) {
	tmp, err := NewThread_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &Thread_V1_TypeGroup2{
		Thread_V1: tmp,
	}
	return
}

func NewThread_V1_TypeGroup2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Thread_V1_TypeGroup2, err error) {
	tmp, err := NewThread_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Thread_V1_TypeGroup2{
		Thread_V1: tmp,
	}
	return
}

// SetProcessId sets the value of ProcessId for the instance
func (instance *Thread_V1_TypeGroup2) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Thread_V1_TypeGroup2) GetPropertyProcessId() (value uint32, err error) {
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

// SetTThreadId sets the value of TThreadId for the instance
func (instance *Thread_V1_TypeGroup2) SetPropertyTThreadId(value uint32) (err error) {
	return instance.SetProperty("TThreadId", (value))
}

// GetTThreadId gets the value of TThreadId for the instance
func (instance *Thread_V1_TypeGroup2) GetPropertyTThreadId() (value uint32, err error) {
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
