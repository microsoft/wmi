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

// AntiStarvationBoost struct
type AntiStarvationBoost struct {
	*Thread_V2

	//
	Priority uint8

	//
	ProcessorIndex uint16

	//
	Reserved uint8

	//
	ThreadId uint32
}

func NewAntiStarvationBoostEx1(instance *cim.WmiInstance) (newInstance *AntiStarvationBoost, err error) {
	tmp, err := NewThread_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &AntiStarvationBoost{
		Thread_V2: tmp,
	}
	return
}

func NewAntiStarvationBoostEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AntiStarvationBoost, err error) {
	tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AntiStarvationBoost{
		Thread_V2: tmp,
	}
	return
}

// SetPriority sets the value of Priority for the instance
func (instance *AntiStarvationBoost) SetPropertyPriority(value uint8) (err error) {
	return instance.SetProperty("Priority", (value))
}

// GetPriority gets the value of Priority for the instance
func (instance *AntiStarvationBoost) GetPropertyPriority() (value uint8, err error) {
	retValue, err := instance.GetProperty("Priority")
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

// SetProcessorIndex sets the value of ProcessorIndex for the instance
func (instance *AntiStarvationBoost) SetPropertyProcessorIndex(value uint16) (err error) {
	return instance.SetProperty("ProcessorIndex", (value))
}

// GetProcessorIndex gets the value of ProcessorIndex for the instance
func (instance *AntiStarvationBoost) GetPropertyProcessorIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorIndex")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetReserved sets the value of Reserved for the instance
func (instance *AntiStarvationBoost) SetPropertyReserved(value uint8) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *AntiStarvationBoost) GetPropertyReserved() (value uint8, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetThreadId sets the value of ThreadId for the instance
func (instance *AntiStarvationBoost) SetPropertyThreadId(value uint32) (err error) {
	return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *AntiStarvationBoost) GetPropertyThreadId() (value uint32, err error) {
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
