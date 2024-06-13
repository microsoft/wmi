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

// TP_V2_TimerSetNtTimer struct
type TP_V2_TimerSetNtTimer struct {
	*ThreadPoolTrace_V2

	//
	DueTime uint64

	//
	SubQueue uint32

	//
	TolerableDelay uint32
}

func NewTP_V2_TimerSetNtTimerEx1(instance *cim.WmiInstance) (newInstance *TP_V2_TimerSetNtTimer, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &TP_V2_TimerSetNtTimer{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

func NewTP_V2_TimerSetNtTimerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TP_V2_TimerSetNtTimer, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TP_V2_TimerSetNtTimer{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

// SetDueTime sets the value of DueTime for the instance
func (instance *TP_V2_TimerSetNtTimer) SetPropertyDueTime(value uint64) (err error) {
	return instance.SetProperty("DueTime", (value))
}

// GetDueTime gets the value of DueTime for the instance
func (instance *TP_V2_TimerSetNtTimer) GetPropertyDueTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("DueTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSubQueue sets the value of SubQueue for the instance
func (instance *TP_V2_TimerSetNtTimer) SetPropertySubQueue(value uint32) (err error) {
	return instance.SetProperty("SubQueue", (value))
}

// GetSubQueue gets the value of SubQueue for the instance
func (instance *TP_V2_TimerSetNtTimer) GetPropertySubQueue() (value uint32, err error) {
	retValue, err := instance.GetProperty("SubQueue")
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

// SetTolerableDelay sets the value of TolerableDelay for the instance
func (instance *TP_V2_TimerSetNtTimer) SetPropertyTolerableDelay(value uint32) (err error) {
	return instance.SetProperty("TolerableDelay", (value))
}

// GetTolerableDelay gets the value of TolerableDelay for the instance
func (instance *TP_V2_TimerSetNtTimer) GetPropertyTolerableDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("TolerableDelay")
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
