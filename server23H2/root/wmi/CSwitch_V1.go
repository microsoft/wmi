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

// CSwitch_V1 struct
type CSwitch_V1 struct {
	*Thread_V1

	//
	NewThreadId uint32

	//
	NewThreadPriority int8

	//
	NewThreadQuantum int8

	//
	OldThreadId uint32

	//
	OldThreadPriority int8

	//
	OldThreadQuantum int8

	//
	OldThreadState int8

	//
	OldThreadWaitIdealProcessor int8

	//
	OldThreadWaitMode int8

	//
	OldThreadWaitReason int8
}

func NewCSwitch_V1Ex1(instance *cim.WmiInstance) (newInstance *CSwitch_V1, err error) {
	tmp, err := NewThread_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &CSwitch_V1{
		Thread_V1: tmp,
	}
	return
}

func NewCSwitch_V1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CSwitch_V1, err error) {
	tmp, err := NewThread_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CSwitch_V1{
		Thread_V1: tmp,
	}
	return
}

// SetNewThreadId sets the value of NewThreadId for the instance
func (instance *CSwitch_V1) SetPropertyNewThreadId(value uint32) (err error) {
	return instance.SetProperty("NewThreadId", (value))
}

// GetNewThreadId gets the value of NewThreadId for the instance
func (instance *CSwitch_V1) GetPropertyNewThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("NewThreadId")
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

// SetNewThreadPriority sets the value of NewThreadPriority for the instance
func (instance *CSwitch_V1) SetPropertyNewThreadPriority(value int8) (err error) {
	return instance.SetProperty("NewThreadPriority", (value))
}

// GetNewThreadPriority gets the value of NewThreadPriority for the instance
func (instance *CSwitch_V1) GetPropertyNewThreadPriority() (value int8, err error) {
	retValue, err := instance.GetProperty("NewThreadPriority")
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

// SetNewThreadQuantum sets the value of NewThreadQuantum for the instance
func (instance *CSwitch_V1) SetPropertyNewThreadQuantum(value int8) (err error) {
	return instance.SetProperty("NewThreadQuantum", (value))
}

// GetNewThreadQuantum gets the value of NewThreadQuantum for the instance
func (instance *CSwitch_V1) GetPropertyNewThreadQuantum() (value int8, err error) {
	retValue, err := instance.GetProperty("NewThreadQuantum")
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

// SetOldThreadId sets the value of OldThreadId for the instance
func (instance *CSwitch_V1) SetPropertyOldThreadId(value uint32) (err error) {
	return instance.SetProperty("OldThreadId", (value))
}

// GetOldThreadId gets the value of OldThreadId for the instance
func (instance *CSwitch_V1) GetPropertyOldThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("OldThreadId")
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

// SetOldThreadPriority sets the value of OldThreadPriority for the instance
func (instance *CSwitch_V1) SetPropertyOldThreadPriority(value int8) (err error) {
	return instance.SetProperty("OldThreadPriority", (value))
}

// GetOldThreadPriority gets the value of OldThreadPriority for the instance
func (instance *CSwitch_V1) GetPropertyOldThreadPriority() (value int8, err error) {
	retValue, err := instance.GetProperty("OldThreadPriority")
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

// SetOldThreadQuantum sets the value of OldThreadQuantum for the instance
func (instance *CSwitch_V1) SetPropertyOldThreadQuantum(value int8) (err error) {
	return instance.SetProperty("OldThreadQuantum", (value))
}

// GetOldThreadQuantum gets the value of OldThreadQuantum for the instance
func (instance *CSwitch_V1) GetPropertyOldThreadQuantum() (value int8, err error) {
	retValue, err := instance.GetProperty("OldThreadQuantum")
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

// SetOldThreadState sets the value of OldThreadState for the instance
func (instance *CSwitch_V1) SetPropertyOldThreadState(value int8) (err error) {
	return instance.SetProperty("OldThreadState", (value))
}

// GetOldThreadState gets the value of OldThreadState for the instance
func (instance *CSwitch_V1) GetPropertyOldThreadState() (value int8, err error) {
	retValue, err := instance.GetProperty("OldThreadState")
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

// SetOldThreadWaitIdealProcessor sets the value of OldThreadWaitIdealProcessor for the instance
func (instance *CSwitch_V1) SetPropertyOldThreadWaitIdealProcessor(value int8) (err error) {
	return instance.SetProperty("OldThreadWaitIdealProcessor", (value))
}

// GetOldThreadWaitIdealProcessor gets the value of OldThreadWaitIdealProcessor for the instance
func (instance *CSwitch_V1) GetPropertyOldThreadWaitIdealProcessor() (value int8, err error) {
	retValue, err := instance.GetProperty("OldThreadWaitIdealProcessor")
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

// SetOldThreadWaitMode sets the value of OldThreadWaitMode for the instance
func (instance *CSwitch_V1) SetPropertyOldThreadWaitMode(value int8) (err error) {
	return instance.SetProperty("OldThreadWaitMode", (value))
}

// GetOldThreadWaitMode gets the value of OldThreadWaitMode for the instance
func (instance *CSwitch_V1) GetPropertyOldThreadWaitMode() (value int8, err error) {
	retValue, err := instance.GetProperty("OldThreadWaitMode")
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

// SetOldThreadWaitReason sets the value of OldThreadWaitReason for the instance
func (instance *CSwitch_V1) SetPropertyOldThreadWaitReason(value int8) (err error) {
	return instance.SetProperty("OldThreadWaitReason", (value))
}

// GetOldThreadWaitReason gets the value of OldThreadWaitReason for the instance
func (instance *CSwitch_V1) GetPropertyOldThreadWaitReason() (value int8, err error) {
	retValue, err := instance.GetProperty("OldThreadWaitReason")
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
