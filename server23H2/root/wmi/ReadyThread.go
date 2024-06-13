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

// ReadyThread struct
type ReadyThread struct {
	*Thread_V2

	//
	AdjustIncrement int8

	//
	AdjustReason int8

	//
	Flag int8

	//
	Reserved int8

	//
	TThreadId uint32
}

func NewReadyThreadEx1(instance *cim.WmiInstance) (newInstance *ReadyThread, err error) {
	tmp, err := NewThread_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &ReadyThread{
		Thread_V2: tmp,
	}
	return
}

func NewReadyThreadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ReadyThread, err error) {
	tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ReadyThread{
		Thread_V2: tmp,
	}
	return
}

// SetAdjustIncrement sets the value of AdjustIncrement for the instance
func (instance *ReadyThread) SetPropertyAdjustIncrement(value int8) (err error) {
	return instance.SetProperty("AdjustIncrement", (value))
}

// GetAdjustIncrement gets the value of AdjustIncrement for the instance
func (instance *ReadyThread) GetPropertyAdjustIncrement() (value int8, err error) {
	retValue, err := instance.GetProperty("AdjustIncrement")
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

// SetAdjustReason sets the value of AdjustReason for the instance
func (instance *ReadyThread) SetPropertyAdjustReason(value int8) (err error) {
	return instance.SetProperty("AdjustReason", (value))
}

// GetAdjustReason gets the value of AdjustReason for the instance
func (instance *ReadyThread) GetPropertyAdjustReason() (value int8, err error) {
	retValue, err := instance.GetProperty("AdjustReason")
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

// SetFlag sets the value of Flag for the instance
func (instance *ReadyThread) SetPropertyFlag(value int8) (err error) {
	return instance.SetProperty("Flag", (value))
}

// GetFlag gets the value of Flag for the instance
func (instance *ReadyThread) GetPropertyFlag() (value int8, err error) {
	retValue, err := instance.GetProperty("Flag")
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

// SetReserved sets the value of Reserved for the instance
func (instance *ReadyThread) SetPropertyReserved(value int8) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *ReadyThread) GetPropertyReserved() (value int8, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetTThreadId sets the value of TThreadId for the instance
func (instance *ReadyThread) SetPropertyTThreadId(value uint32) (err error) {
	return instance.SetProperty("TThreadId", (value))
}

// GetTThreadId gets the value of TThreadId for the instance
func (instance *ReadyThread) GetPropertyTThreadId() (value uint32, err error) {
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
