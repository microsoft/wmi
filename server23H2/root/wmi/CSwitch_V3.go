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

// CSwitch_V3 struct
type CSwitch_V3 struct {
	*Thread_V3

	//
	NewThreadId uint32

	//
	NewThreadPriority int8

	//
	NewThreadWaitTime uint32

	//
	OldThreadId uint32

	//
	OldThreadPriority int8

	//
	OldThreadState int8

	//
	OldThreadWaitIdealProcessor int8

	//
	OldThreadWaitReason int8

	//
	PreviousCState uint8

	//
	Reserved uint32

	//
	SpareByte int8

	//
	ThreadFlags int8
}

func NewCSwitch_V3Ex1(instance *cim.WmiInstance) (newInstance *CSwitch_V3, err error) {
	tmp, err := NewThread_V3Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &CSwitch_V3{
		Thread_V3: tmp,
	}
	return
}

func NewCSwitch_V3Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CSwitch_V3, err error) {
	tmp, err := NewThread_V3Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CSwitch_V3{
		Thread_V3: tmp,
	}
	return
}

// SetNewThreadId sets the value of NewThreadId for the instance
func (instance *CSwitch_V3) SetPropertyNewThreadId(value uint32) (err error) {
	return instance.SetProperty("NewThreadId", (value))
}

// GetNewThreadId gets the value of NewThreadId for the instance
func (instance *CSwitch_V3) GetPropertyNewThreadId() (value uint32, err error) {
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
func (instance *CSwitch_V3) SetPropertyNewThreadPriority(value int8) (err error) {
	return instance.SetProperty("NewThreadPriority", (value))
}

// GetNewThreadPriority gets the value of NewThreadPriority for the instance
func (instance *CSwitch_V3) GetPropertyNewThreadPriority() (value int8, err error) {
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

// SetNewThreadWaitTime sets the value of NewThreadWaitTime for the instance
func (instance *CSwitch_V3) SetPropertyNewThreadWaitTime(value uint32) (err error) {
	return instance.SetProperty("NewThreadWaitTime", (value))
}

// GetNewThreadWaitTime gets the value of NewThreadWaitTime for the instance
func (instance *CSwitch_V3) GetPropertyNewThreadWaitTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("NewThreadWaitTime")
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

// SetOldThreadId sets the value of OldThreadId for the instance
func (instance *CSwitch_V3) SetPropertyOldThreadId(value uint32) (err error) {
	return instance.SetProperty("OldThreadId", (value))
}

// GetOldThreadId gets the value of OldThreadId for the instance
func (instance *CSwitch_V3) GetPropertyOldThreadId() (value uint32, err error) {
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
func (instance *CSwitch_V3) SetPropertyOldThreadPriority(value int8) (err error) {
	return instance.SetProperty("OldThreadPriority", (value))
}

// GetOldThreadPriority gets the value of OldThreadPriority for the instance
func (instance *CSwitch_V3) GetPropertyOldThreadPriority() (value int8, err error) {
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

// SetOldThreadState sets the value of OldThreadState for the instance
func (instance *CSwitch_V3) SetPropertyOldThreadState(value int8) (err error) {
	return instance.SetProperty("OldThreadState", (value))
}

// GetOldThreadState gets the value of OldThreadState for the instance
func (instance *CSwitch_V3) GetPropertyOldThreadState() (value int8, err error) {
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
func (instance *CSwitch_V3) SetPropertyOldThreadWaitIdealProcessor(value int8) (err error) {
	return instance.SetProperty("OldThreadWaitIdealProcessor", (value))
}

// GetOldThreadWaitIdealProcessor gets the value of OldThreadWaitIdealProcessor for the instance
func (instance *CSwitch_V3) GetPropertyOldThreadWaitIdealProcessor() (value int8, err error) {
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

// SetOldThreadWaitReason sets the value of OldThreadWaitReason for the instance
func (instance *CSwitch_V3) SetPropertyOldThreadWaitReason(value int8) (err error) {
	return instance.SetProperty("OldThreadWaitReason", (value))
}

// GetOldThreadWaitReason gets the value of OldThreadWaitReason for the instance
func (instance *CSwitch_V3) GetPropertyOldThreadWaitReason() (value int8, err error) {
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

// SetPreviousCState sets the value of PreviousCState for the instance
func (instance *CSwitch_V3) SetPropertyPreviousCState(value uint8) (err error) {
	return instance.SetProperty("PreviousCState", (value))
}

// GetPreviousCState gets the value of PreviousCState for the instance
func (instance *CSwitch_V3) GetPropertyPreviousCState() (value uint8, err error) {
	retValue, err := instance.GetProperty("PreviousCState")
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

// SetReserved sets the value of Reserved for the instance
func (instance *CSwitch_V3) SetPropertyReserved(value uint32) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *CSwitch_V3) GetPropertyReserved() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetSpareByte sets the value of SpareByte for the instance
func (instance *CSwitch_V3) SetPropertySpareByte(value int8) (err error) {
	return instance.SetProperty("SpareByte", (value))
}

// GetSpareByte gets the value of SpareByte for the instance
func (instance *CSwitch_V3) GetPropertySpareByte() (value int8, err error) {
	retValue, err := instance.GetProperty("SpareByte")
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

// SetThreadFlags sets the value of ThreadFlags for the instance
func (instance *CSwitch_V3) SetPropertyThreadFlags(value int8) (err error) {
	return instance.SetProperty("ThreadFlags", (value))
}

// GetThreadFlags gets the value of ThreadFlags for the instance
func (instance *CSwitch_V3) GetPropertyThreadFlags() (value int8, err error) {
	retValue, err := instance.GetProperty("ThreadFlags")
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
