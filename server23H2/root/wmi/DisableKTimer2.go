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

// DisableKTimer2 struct
type DisableKTimer2 struct {
	*PerfInfo_V2

	//
	DisableCallback uint32

	//
	DisableContext uint32

	//
	Timer uint32

	//
	TimerFlags uint8
}

func NewDisableKTimer2Ex1(instance *cim.WmiInstance) (newInstance *DisableKTimer2, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &DisableKTimer2{
		PerfInfo_V2: tmp,
	}
	return
}

func NewDisableKTimer2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DisableKTimer2, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DisableKTimer2{
		PerfInfo_V2: tmp,
	}
	return
}

// SetDisableCallback sets the value of DisableCallback for the instance
func (instance *DisableKTimer2) SetPropertyDisableCallback(value uint32) (err error) {
	return instance.SetProperty("DisableCallback", (value))
}

// GetDisableCallback gets the value of DisableCallback for the instance
func (instance *DisableKTimer2) GetPropertyDisableCallback() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisableCallback")
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

// SetDisableContext sets the value of DisableContext for the instance
func (instance *DisableKTimer2) SetPropertyDisableContext(value uint32) (err error) {
	return instance.SetProperty("DisableContext", (value))
}

// GetDisableContext gets the value of DisableContext for the instance
func (instance *DisableKTimer2) GetPropertyDisableContext() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisableContext")
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

// SetTimer sets the value of Timer for the instance
func (instance *DisableKTimer2) SetPropertyTimer(value uint32) (err error) {
	return instance.SetProperty("Timer", (value))
}

// GetTimer gets the value of Timer for the instance
func (instance *DisableKTimer2) GetPropertyTimer() (value uint32, err error) {
	retValue, err := instance.GetProperty("Timer")
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

// SetTimerFlags sets the value of TimerFlags for the instance
func (instance *DisableKTimer2) SetPropertyTimerFlags(value uint8) (err error) {
	return instance.SetProperty("TimerFlags", (value))
}

// GetTimerFlags gets the value of TimerFlags for the instance
func (instance *DisableKTimer2) GetPropertyTimerFlags() (value uint8, err error) {
	retValue, err := instance.GetProperty("TimerFlags")
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
