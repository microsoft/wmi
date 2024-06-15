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

// WorkerThread struct
type WorkerThread struct {
	*Thread_V2

	//
	StartTime uint64

	//
	ThreadRoutine uint32

	//
	TThreadId uint32
}

func NewWorkerThreadEx1(instance *cim.WmiInstance) (newInstance *WorkerThread, err error) {
	tmp, err := NewThread_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &WorkerThread{
		Thread_V2: tmp,
	}
	return
}

func NewWorkerThreadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WorkerThread, err error) {
	tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WorkerThread{
		Thread_V2: tmp,
	}
	return
}

// SetStartTime sets the value of StartTime for the instance
func (instance *WorkerThread) SetPropertyStartTime(value uint64) (err error) {
	return instance.SetProperty("StartTime", (value))
}

// GetStartTime gets the value of StartTime for the instance
func (instance *WorkerThread) GetPropertyStartTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("StartTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetThreadRoutine sets the value of ThreadRoutine for the instance
func (instance *WorkerThread) SetPropertyThreadRoutine(value uint32) (err error) {
	return instance.SetProperty("ThreadRoutine", (value))
}

// GetThreadRoutine gets the value of ThreadRoutine for the instance
func (instance *WorkerThread) GetPropertyThreadRoutine() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadRoutine")
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
func (instance *WorkerThread) SetPropertyTThreadId(value uint32) (err error) {
	return instance.SetProperty("TThreadId", (value))
}

// GetTThreadId gets the value of TThreadId for the instance
func (instance *WorkerThread) GetPropertyTThreadId() (value uint32, err error) {
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
