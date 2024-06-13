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

// TPCBCancel struct
type TPCBCancel struct {
	*ThreadPoolTrace

	//
	CallbackContext uint32

	//
	CallbackFunction uint32

	//
	CancelCount uint32

	//
	PoolId uint32

	//
	SubProcessTag uint32

	//
	TaskId uint32
}

func NewTPCBCancelEx1(instance *cim.WmiInstance) (newInstance *TPCBCancel, err error) {
	tmp, err := NewThreadPoolTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TPCBCancel{
		ThreadPoolTrace: tmp,
	}
	return
}

func NewTPCBCancelEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TPCBCancel, err error) {
	tmp, err := NewThreadPoolTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TPCBCancel{
		ThreadPoolTrace: tmp,
	}
	return
}

// SetCallbackContext sets the value of CallbackContext for the instance
func (instance *TPCBCancel) SetPropertyCallbackContext(value uint32) (err error) {
	return instance.SetProperty("CallbackContext", (value))
}

// GetCallbackContext gets the value of CallbackContext for the instance
func (instance *TPCBCancel) GetPropertyCallbackContext() (value uint32, err error) {
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
func (instance *TPCBCancel) SetPropertyCallbackFunction(value uint32) (err error) {
	return instance.SetProperty("CallbackFunction", (value))
}

// GetCallbackFunction gets the value of CallbackFunction for the instance
func (instance *TPCBCancel) GetPropertyCallbackFunction() (value uint32, err error) {
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

// SetCancelCount sets the value of CancelCount for the instance
func (instance *TPCBCancel) SetPropertyCancelCount(value uint32) (err error) {
	return instance.SetProperty("CancelCount", (value))
}

// GetCancelCount gets the value of CancelCount for the instance
func (instance *TPCBCancel) GetPropertyCancelCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("CancelCount")
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
func (instance *TPCBCancel) SetPropertyPoolId(value uint32) (err error) {
	return instance.SetProperty("PoolId", (value))
}

// GetPoolId gets the value of PoolId for the instance
func (instance *TPCBCancel) GetPropertyPoolId() (value uint32, err error) {
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
func (instance *TPCBCancel) SetPropertySubProcessTag(value uint32) (err error) {
	return instance.SetProperty("SubProcessTag", (value))
}

// GetSubProcessTag gets the value of SubProcessTag for the instance
func (instance *TPCBCancel) GetPropertySubProcessTag() (value uint32, err error) {
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
func (instance *TPCBCancel) SetPropertyTaskId(value uint32) (err error) {
	return instance.SetProperty("TaskId", (value))
}

// GetTaskId gets the value of TaskId for the instance
func (instance *TPCBCancel) GetPropertyTaskId() (value uint32, err error) {
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
