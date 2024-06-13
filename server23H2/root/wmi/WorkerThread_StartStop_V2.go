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

// WorkerThread_StartStop_V2 struct
type WorkerThread_StartStop_V2 struct {
	*Thread_V2

	//
	CallbackRoutine uint32
}

func NewWorkerThread_StartStop_V2Ex1(instance *cim.WmiInstance) (newInstance *WorkerThread_StartStop_V2, err error) {
	tmp, err := NewThread_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &WorkerThread_StartStop_V2{
		Thread_V2: tmp,
	}
	return
}

func NewWorkerThread_StartStop_V2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WorkerThread_StartStop_V2, err error) {
	tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WorkerThread_StartStop_V2{
		Thread_V2: tmp,
	}
	return
}

// SetCallbackRoutine sets the value of CallbackRoutine for the instance
func (instance *WorkerThread_StartStop_V2) SetPropertyCallbackRoutine(value uint32) (err error) {
	return instance.SetProperty("CallbackRoutine", (value))
}

// GetCallbackRoutine gets the value of CallbackRoutine for the instance
func (instance *WorkerThread_StartStop_V2) GetPropertyCallbackRoutine() (value uint32, err error) {
	retValue, err := instance.GetProperty("CallbackRoutine")
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
