// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_Counters_StorportUnitQueue struct
type Win32_PerfFormattedData_Counters_StorportUnitQueue struct {
	*Win32_PerfFormattedData

	//
	OutstandingRequests uint32

	//
	QueuedRequests uint32
}

func NewWin32_PerfFormattedData_Counters_StorportUnitQueueEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_StorportUnitQueue, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_StorportUnitQueue{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_StorportUnitQueueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_StorportUnitQueue, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_StorportUnitQueue{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetOutstandingRequests sets the value of OutstandingRequests for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitQueue) SetPropertyOutstandingRequests(value uint32) (err error) {
	return instance.SetProperty("OutstandingRequests", (value))
}

// GetOutstandingRequests gets the value of OutstandingRequests for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitQueue) GetPropertyOutstandingRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutstandingRequests")
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

// SetQueuedRequests sets the value of QueuedRequests for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitQueue) SetPropertyQueuedRequests(value uint32) (err error) {
	return instance.SetProperty("QueuedRequests", (value))
}

// GetQueuedRequests gets the value of QueuedRequests for the instance
func (instance *Win32_PerfFormattedData_Counters_StorportUnitQueue) GetPropertyQueuedRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueuedRequests")
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
