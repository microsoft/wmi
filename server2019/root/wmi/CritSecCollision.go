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

// CritSecCollision struct
type CritSecCollision struct {
	*CritSecTrace

	//
	CritSecAddr uint32

	//
	LockCount uint32

	//
	OwningThread uint32

	//
	SpinCount uint32
}

func NewCritSecCollisionEx1(instance *cim.WmiInstance) (newInstance *CritSecCollision, err error) {
	tmp, err := NewCritSecTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CritSecCollision{
		CritSecTrace: tmp,
	}
	return
}

func NewCritSecCollisionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CritSecCollision, err error) {
	tmp, err := NewCritSecTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CritSecCollision{
		CritSecTrace: tmp,
	}
	return
}

// SetCritSecAddr sets the value of CritSecAddr for the instance
func (instance *CritSecCollision) SetPropertyCritSecAddr(value uint32) (err error) {
	return instance.SetProperty("CritSecAddr", (value))
}

// GetCritSecAddr gets the value of CritSecAddr for the instance
func (instance *CritSecCollision) GetPropertyCritSecAddr() (value uint32, err error) {
	retValue, err := instance.GetProperty("CritSecAddr")
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

// SetLockCount sets the value of LockCount for the instance
func (instance *CritSecCollision) SetPropertyLockCount(value uint32) (err error) {
	return instance.SetProperty("LockCount", (value))
}

// GetLockCount gets the value of LockCount for the instance
func (instance *CritSecCollision) GetPropertyLockCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("LockCount")
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

// SetOwningThread sets the value of OwningThread for the instance
func (instance *CritSecCollision) SetPropertyOwningThread(value uint32) (err error) {
	return instance.SetProperty("OwningThread", (value))
}

// GetOwningThread gets the value of OwningThread for the instance
func (instance *CritSecCollision) GetPropertyOwningThread() (value uint32, err error) {
	retValue, err := instance.GetProperty("OwningThread")
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

// SetSpinCount sets the value of SpinCount for the instance
func (instance *CritSecCollision) SetPropertySpinCount(value uint32) (err error) {
	return instance.SetProperty("SpinCount", (value))
}

// GetSpinCount gets the value of SpinCount for the instance
func (instance *CritSecCollision) GetPropertySpinCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SpinCount")
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
