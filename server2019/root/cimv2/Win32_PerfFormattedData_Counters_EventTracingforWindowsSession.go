// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PerfFormattedData_Counters_EventTracingforWindowsSession struct
type Win32_PerfFormattedData_Counters_EventTracingforWindowsSession struct {
	*Win32_PerfFormattedData

	//
	BufferMemoryUsageNonPagedPool uint32

	//
	BufferMemoryUsagePagedPool uint32

	//
	EventsLoggedpersec uint64

	//
	EventsLost uint32

	//
	NumberofRealTimeConsumers uint32
}

func NewWin32_PerfFormattedData_Counters_EventTracingforWindowsSessionEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_EventTracingforWindowsSession{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_EventTracingforWindowsSessionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_EventTracingforWindowsSession{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetBufferMemoryUsageNonPagedPool sets the value of BufferMemoryUsageNonPagedPool for the instance
func (instance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession) SetPropertyBufferMemoryUsageNonPagedPool(value uint32) (err error) {
	return instance.SetProperty("BufferMemoryUsageNonPagedPool", value)
}

// GetBufferMemoryUsageNonPagedPool gets the value of BufferMemoryUsageNonPagedPool for the instance
func (instance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession) GetPropertyBufferMemoryUsageNonPagedPool() (value uint32, err error) {
	retValue, err := instance.GetProperty("BufferMemoryUsageNonPagedPool")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBufferMemoryUsagePagedPool sets the value of BufferMemoryUsagePagedPool for the instance
func (instance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession) SetPropertyBufferMemoryUsagePagedPool(value uint32) (err error) {
	return instance.SetProperty("BufferMemoryUsagePagedPool", value)
}

// GetBufferMemoryUsagePagedPool gets the value of BufferMemoryUsagePagedPool for the instance
func (instance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession) GetPropertyBufferMemoryUsagePagedPool() (value uint32, err error) {
	retValue, err := instance.GetProperty("BufferMemoryUsagePagedPool")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventsLoggedpersec sets the value of EventsLoggedpersec for the instance
func (instance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession) SetPropertyEventsLoggedpersec(value uint64) (err error) {
	return instance.SetProperty("EventsLoggedpersec", value)
}

// GetEventsLoggedpersec gets the value of EventsLoggedpersec for the instance
func (instance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession) GetPropertyEventsLoggedpersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EventsLoggedpersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventsLost sets the value of EventsLost for the instance
func (instance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession) SetPropertyEventsLost(value uint32) (err error) {
	return instance.SetProperty("EventsLost", value)
}

// GetEventsLost gets the value of EventsLost for the instance
func (instance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession) GetPropertyEventsLost() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventsLost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofRealTimeConsumers sets the value of NumberofRealTimeConsumers for the instance
func (instance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession) SetPropertyNumberofRealTimeConsumers(value uint32) (err error) {
	return instance.SetProperty("NumberofRealTimeConsumers", value)
}

// GetNumberofRealTimeConsumers gets the value of NumberofRealTimeConsumers for the instance
func (instance *Win32_PerfFormattedData_Counters_EventTracingforWindowsSession) GetPropertyNumberofRealTimeConsumers() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofRealTimeConsumers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
