// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// UmsContextSwitch struct
type UmsContextSwitch struct {
	*UmsEvent

	//
	KernelYieldCount uint32

	//
	MixedYieldCount uint32

	//
	ScheduledThreadId uint32

	//
	SwitchCount uint32

	//
	YieldCount uint32
}

func NewUmsContextSwitchEx1(instance *cim.WmiInstance) (newInstance *UmsContextSwitch, err error) {
	tmp, err := NewUmsEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &UmsContextSwitch{
		UmsEvent: tmp,
	}
	return
}

func NewUmsContextSwitchEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *UmsContextSwitch, err error) {
	tmp, err := NewUmsEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &UmsContextSwitch{
		UmsEvent: tmp,
	}
	return
}

// SetKernelYieldCount sets the value of KernelYieldCount for the instance
func (instance *UmsContextSwitch) SetPropertyKernelYieldCount(value uint32) (err error) {
	return instance.SetProperty("KernelYieldCount", (value))
}

// GetKernelYieldCount gets the value of KernelYieldCount for the instance
func (instance *UmsContextSwitch) GetPropertyKernelYieldCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("KernelYieldCount")
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

// SetMixedYieldCount sets the value of MixedYieldCount for the instance
func (instance *UmsContextSwitch) SetPropertyMixedYieldCount(value uint32) (err error) {
	return instance.SetProperty("MixedYieldCount", (value))
}

// GetMixedYieldCount gets the value of MixedYieldCount for the instance
func (instance *UmsContextSwitch) GetPropertyMixedYieldCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MixedYieldCount")
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

// SetScheduledThreadId sets the value of ScheduledThreadId for the instance
func (instance *UmsContextSwitch) SetPropertyScheduledThreadId(value uint32) (err error) {
	return instance.SetProperty("ScheduledThreadId", (value))
}

// GetScheduledThreadId gets the value of ScheduledThreadId for the instance
func (instance *UmsContextSwitch) GetPropertyScheduledThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScheduledThreadId")
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

// SetSwitchCount sets the value of SwitchCount for the instance
func (instance *UmsContextSwitch) SetPropertySwitchCount(value uint32) (err error) {
	return instance.SetProperty("SwitchCount", (value))
}

// GetSwitchCount gets the value of SwitchCount for the instance
func (instance *UmsContextSwitch) GetPropertySwitchCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SwitchCount")
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

// SetYieldCount sets the value of YieldCount for the instance
func (instance *UmsContextSwitch) SetPropertyYieldCount(value uint32) (err error) {
	return instance.SetProperty("YieldCount", (value))
}

// GetYieldCount gets the value of YieldCount for the instance
func (instance *UmsContextSwitch) GetPropertyYieldCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("YieldCount")
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
