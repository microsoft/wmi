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

// TP_V2_TimerCancelled struct
type TP_V2_TimerCancelled struct {
	*ThreadPoolTrace_V2

	//
	SubQueue uint32

	//
	Timer uint32
}

func NewTP_V2_TimerCancelledEx1(instance *cim.WmiInstance) (newInstance *TP_V2_TimerCancelled, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &TP_V2_TimerCancelled{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

func NewTP_V2_TimerCancelledEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TP_V2_TimerCancelled, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TP_V2_TimerCancelled{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

// SetSubQueue sets the value of SubQueue for the instance
func (instance *TP_V2_TimerCancelled) SetPropertySubQueue(value uint32) (err error) {
	return instance.SetProperty("SubQueue", (value))
}

// GetSubQueue gets the value of SubQueue for the instance
func (instance *TP_V2_TimerCancelled) GetPropertySubQueue() (value uint32, err error) {
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

// SetTimer sets the value of Timer for the instance
func (instance *TP_V2_TimerCancelled) SetPropertyTimer(value uint32) (err error) {
	return instance.SetProperty("Timer", (value))
}

// GetTimer gets the value of Timer for the instance
func (instance *TP_V2_TimerCancelled) GetPropertyTimer() (value uint32, err error) {
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
