// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PeriodicRestartSettings struct
type PeriodicRestartSettings struct {
	*EmbeddedObject

	//
	Memory uint32

	//
	PrivateMemory uint32

	//
	Requests uint32

	//
	Schedule []ScheduleElement

	//
	Time string
}

func NewPeriodicRestartSettingsEx1(instance *cim.WmiInstance) (newInstance *PeriodicRestartSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PeriodicRestartSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewPeriodicRestartSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PeriodicRestartSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PeriodicRestartSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetMemory sets the value of Memory for the instance
func (instance *PeriodicRestartSettings) SetPropertyMemory(value uint32) (err error) {
	return instance.SetProperty("Memory", (value))
}

// GetMemory gets the value of Memory for the instance
func (instance *PeriodicRestartSettings) GetPropertyMemory() (value uint32, err error) {
	retValue, err := instance.GetProperty("Memory")
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

// SetPrivateMemory sets the value of PrivateMemory for the instance
func (instance *PeriodicRestartSettings) SetPropertyPrivateMemory(value uint32) (err error) {
	return instance.SetProperty("PrivateMemory", (value))
}

// GetPrivateMemory gets the value of PrivateMemory for the instance
func (instance *PeriodicRestartSettings) GetPropertyPrivateMemory() (value uint32, err error) {
	retValue, err := instance.GetProperty("PrivateMemory")
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

// SetRequests sets the value of Requests for the instance
func (instance *PeriodicRestartSettings) SetPropertyRequests(value uint32) (err error) {
	return instance.SetProperty("Requests", (value))
}

// GetRequests gets the value of Requests for the instance
func (instance *PeriodicRestartSettings) GetPropertyRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("Requests")
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

// SetSchedule sets the value of Schedule for the instance
func (instance *PeriodicRestartSettings) SetPropertySchedule(value []ScheduleElement) (err error) {
	return instance.SetProperty("Schedule", (value))
}

// GetSchedule gets the value of Schedule for the instance
func (instance *PeriodicRestartSettings) GetPropertySchedule() (value []ScheduleElement, err error) {
	retValue, err := instance.GetProperty("Schedule")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ScheduleElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ScheduleElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ScheduleElement(valuetmp))
	}

	return
}

// SetTime sets the value of Time for the instance
func (instance *PeriodicRestartSettings) SetPropertyTime(value string) (err error) {
	return instance.SetProperty("Time", (value))
}

// GetTime gets the value of Time for the instance
func (instance *PeriodicRestartSettings) GetPropertyTime() (value string, err error) {
	retValue, err := instance.GetProperty("Time")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
