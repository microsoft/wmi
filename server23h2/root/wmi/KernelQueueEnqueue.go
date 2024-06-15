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

// KernelQueueEnqueue struct
type KernelQueueEnqueue struct {
	*Thread_V2

	//
	Entry uint32

	//
	ThreadId uint32
}

func NewKernelQueueEnqueueEx1(instance *cim.WmiInstance) (newInstance *KernelQueueEnqueue, err error) {
	tmp, err := NewThread_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &KernelQueueEnqueue{
		Thread_V2: tmp,
	}
	return
}

func NewKernelQueueEnqueueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KernelQueueEnqueue, err error) {
	tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KernelQueueEnqueue{
		Thread_V2: tmp,
	}
	return
}

// SetEntry sets the value of Entry for the instance
func (instance *KernelQueueEnqueue) SetPropertyEntry(value uint32) (err error) {
	return instance.SetProperty("Entry", (value))
}

// GetEntry gets the value of Entry for the instance
func (instance *KernelQueueEnqueue) GetPropertyEntry() (value uint32, err error) {
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

// SetThreadId sets the value of ThreadId for the instance
func (instance *KernelQueueEnqueue) SetPropertyThreadId(value uint32) (err error) {
	return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *KernelQueueEnqueue) GetPropertyThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadId")
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
