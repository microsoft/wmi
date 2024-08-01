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

// TP_V2_CBEnqueue struct
type TP_V2_CBEnqueue struct {
	*ThreadPoolTrace_V2

	//
	CallbackContext uint32

	//
	CallbackFunction uint32

	//
	PoolId uint32

	//
	SubProcessTag uint32

	//
	TaskId uint32
}

func NewTP_V2_CBEnqueueEx1(instance *cim.WmiInstance) (newInstance *TP_V2_CBEnqueue, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &TP_V2_CBEnqueue{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

func NewTP_V2_CBEnqueueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TP_V2_CBEnqueue, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TP_V2_CBEnqueue{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

// SetCallbackContext sets the value of CallbackContext for the instance
func (instance *TP_V2_CBEnqueue) SetPropertyCallbackContext(value uint32) (err error) {
	return instance.SetProperty("CallbackContext", (value))
}

// GetCallbackContext gets the value of CallbackContext for the instance
func (instance *TP_V2_CBEnqueue) GetPropertyCallbackContext() (value uint32, err error) {
	retValue, err := instance.GetProperty("CallbackContext")
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

// SetCallbackFunction sets the value of CallbackFunction for the instance
func (instance *TP_V2_CBEnqueue) SetPropertyCallbackFunction(value uint32) (err error) {
	return instance.SetProperty("CallbackFunction", (value))
}

// GetCallbackFunction gets the value of CallbackFunction for the instance
func (instance *TP_V2_CBEnqueue) GetPropertyCallbackFunction() (value uint32, err error) {
	retValue, err := instance.GetProperty("CallbackFunction")
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

// SetPoolId sets the value of PoolId for the instance
func (instance *TP_V2_CBEnqueue) SetPropertyPoolId(value uint32) (err error) {
	return instance.SetProperty("PoolId", (value))
}

// GetPoolId gets the value of PoolId for the instance
func (instance *TP_V2_CBEnqueue) GetPropertyPoolId() (value uint32, err error) {
	retValue, err := instance.GetProperty("PoolId")
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
func (instance *TP_V2_CBEnqueue) SetPropertySubProcessTag(value uint32) (err error) {
	return instance.SetProperty("SubProcessTag", (value))
}

// GetSubProcessTag gets the value of SubProcessTag for the instance
func (instance *TP_V2_CBEnqueue) GetPropertySubProcessTag() (value uint32, err error) {
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

// SetTaskId sets the value of TaskId for the instance
func (instance *TP_V2_CBEnqueue) SetPropertyTaskId(value uint32) (err error) {
	return instance.SetProperty("TaskId", (value))
}

// GetTaskId gets the value of TaskId for the instance
func (instance *TP_V2_CBEnqueue) GetPropertyTaskId() (value uint32, err error) {
	retValue, err := instance.GetProperty("TaskId")
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
