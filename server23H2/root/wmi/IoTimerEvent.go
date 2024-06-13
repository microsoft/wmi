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

// IoTimerEvent struct
type IoTimerEvent struct {
	*PerfInfo_V2

	//
	DeviceObject uint32

	//
	TimerRoutine uint32
}

func NewIoTimerEventEx1(instance *cim.WmiInstance) (newInstance *IoTimerEvent, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &IoTimerEvent{
		PerfInfo_V2: tmp,
	}
	return
}

func NewIoTimerEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *IoTimerEvent, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &IoTimerEvent{
		PerfInfo_V2: tmp,
	}
	return
}

// SetDeviceObject sets the value of DeviceObject for the instance
func (instance *IoTimerEvent) SetPropertyDeviceObject(value uint32) (err error) {
	return instance.SetProperty("DeviceObject", (value))
}

// GetDeviceObject gets the value of DeviceObject for the instance
func (instance *IoTimerEvent) GetPropertyDeviceObject() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceObject")
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

// SetTimerRoutine sets the value of TimerRoutine for the instance
func (instance *IoTimerEvent) SetPropertyTimerRoutine(value uint32) (err error) {
	return instance.SetProperty("TimerRoutine", (value))
}

// GetTimerRoutine gets the value of TimerRoutine for the instance
func (instance *IoTimerEvent) GetPropertyTimerRoutine() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimerRoutine")
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
