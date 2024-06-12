// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_Counters_EventLogSubscriptions struct
type Win32_PerfRawData_Counters_EventLogSubscriptions struct {
	*Win32_PerfRawData

	//
	EventfilteroperationsPersec uint64

	//
	EventsPersec uint64
}

func NewWin32_PerfRawData_Counters_EventLogSubscriptionsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_EventLogSubscriptions, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_EventLogSubscriptions{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_EventLogSubscriptionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_EventLogSubscriptions, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_EventLogSubscriptions{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetEventfilteroperationsPersec sets the value of EventfilteroperationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_EventLogSubscriptions) SetPropertyEventfilteroperationsPersec(value uint64) (err error) {
	return instance.SetProperty("EventfilteroperationsPersec", (value))
}

// GetEventfilteroperationsPersec gets the value of EventfilteroperationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_EventLogSubscriptions) GetPropertyEventfilteroperationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EventfilteroperationsPersec")
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

// SetEventsPersec sets the value of EventsPersec for the instance
func (instance *Win32_PerfRawData_Counters_EventLogSubscriptions) SetPropertyEventsPersec(value uint64) (err error) {
	return instance.SetProperty("EventsPersec", (value))
}

// GetEventsPersec gets the value of EventsPersec for the instance
func (instance *Win32_PerfRawData_Counters_EventLogSubscriptions) GetPropertyEventsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EventsPersec")
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
