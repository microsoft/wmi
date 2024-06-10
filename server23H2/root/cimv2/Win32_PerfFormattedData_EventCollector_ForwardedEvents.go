// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_EventCollector_ForwardedEvents struct
type Win32_PerfFormattedData_EventCollector_ForwardedEvents struct {
	*Win32_PerfFormattedData

	//
	EventsDroppedBatchs uint32

	//
	EventsReceivedBatches uint32

	//
	LostEvents uint64

	//
	ProcessedEvents uint64

	//
	TimestampOfLatestBatch uint64
}

func NewWin32_PerfFormattedData_EventCollector_ForwardedEventsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_EventCollector_ForwardedEvents, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_EventCollector_ForwardedEvents{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_EventCollector_ForwardedEventsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_EventCollector_ForwardedEvents, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_EventCollector_ForwardedEvents{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetEventsDroppedBatchs sets the value of EventsDroppedBatchs for the instance
func (instance *Win32_PerfFormattedData_EventCollector_ForwardedEvents) SetPropertyEventsDroppedBatchs(value uint32) (err error) {
	return instance.SetProperty("EventsDroppedBatchs", (value))
}

// GetEventsDroppedBatchs gets the value of EventsDroppedBatchs for the instance
func (instance *Win32_PerfFormattedData_EventCollector_ForwardedEvents) GetPropertyEventsDroppedBatchs() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventsDroppedBatchs")
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

// SetEventsReceivedBatches sets the value of EventsReceivedBatches for the instance
func (instance *Win32_PerfFormattedData_EventCollector_ForwardedEvents) SetPropertyEventsReceivedBatches(value uint32) (err error) {
	return instance.SetProperty("EventsReceivedBatches", (value))
}

// GetEventsReceivedBatches gets the value of EventsReceivedBatches for the instance
func (instance *Win32_PerfFormattedData_EventCollector_ForwardedEvents) GetPropertyEventsReceivedBatches() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventsReceivedBatches")
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

// SetLostEvents sets the value of LostEvents for the instance
func (instance *Win32_PerfFormattedData_EventCollector_ForwardedEvents) SetPropertyLostEvents(value uint64) (err error) {
	return instance.SetProperty("LostEvents", (value))
}

// GetLostEvents gets the value of LostEvents for the instance
func (instance *Win32_PerfFormattedData_EventCollector_ForwardedEvents) GetPropertyLostEvents() (value uint64, err error) {
	retValue, err := instance.GetProperty("LostEvents")
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

// SetProcessedEvents sets the value of ProcessedEvents for the instance
func (instance *Win32_PerfFormattedData_EventCollector_ForwardedEvents) SetPropertyProcessedEvents(value uint64) (err error) {
	return instance.SetProperty("ProcessedEvents", (value))
}

// GetProcessedEvents gets the value of ProcessedEvents for the instance
func (instance *Win32_PerfFormattedData_EventCollector_ForwardedEvents) GetPropertyProcessedEvents() (value uint64, err error) {
	retValue, err := instance.GetProperty("ProcessedEvents")
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

// SetTimestampOfLatestBatch sets the value of TimestampOfLatestBatch for the instance
func (instance *Win32_PerfFormattedData_EventCollector_ForwardedEvents) SetPropertyTimestampOfLatestBatch(value uint64) (err error) {
	return instance.SetProperty("TimestampOfLatestBatch", (value))
}

// GetTimestampOfLatestBatch gets the value of TimestampOfLatestBatch for the instance
func (instance *Win32_PerfFormattedData_EventCollector_ForwardedEvents) GetPropertyTimestampOfLatestBatch() (value uint64, err error) {
	retValue, err := instance.GetProperty("TimestampOfLatestBatch")
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
