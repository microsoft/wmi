// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TaskWeeklyTrigger struct
type MSFT_TaskWeeklyTrigger struct {
	*MSFT_TaskTrigger

	//
	DaysOfWeek uint16

	//
	RandomDelay string

	//
	WeeksInterval uint16
}

func NewMSFT_TaskWeeklyTriggerEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskWeeklyTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskWeeklyTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

func NewMSFT_TaskWeeklyTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskWeeklyTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskWeeklyTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

// SetDaysOfWeek sets the value of DaysOfWeek for the instance
func (instance *MSFT_TaskWeeklyTrigger) SetPropertyDaysOfWeek(value uint16) (err error) {
	return instance.SetProperty("DaysOfWeek", (value))
}

// GetDaysOfWeek gets the value of DaysOfWeek for the instance
func (instance *MSFT_TaskWeeklyTrigger) GetPropertyDaysOfWeek() (value uint16, err error) {
	retValue, err := instance.GetProperty("DaysOfWeek")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetRandomDelay sets the value of RandomDelay for the instance
func (instance *MSFT_TaskWeeklyTrigger) SetPropertyRandomDelay(value string) (err error) {
	return instance.SetProperty("RandomDelay", (value))
}

// GetRandomDelay gets the value of RandomDelay for the instance
func (instance *MSFT_TaskWeeklyTrigger) GetPropertyRandomDelay() (value string, err error) {
	retValue, err := instance.GetProperty("RandomDelay")
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

// SetWeeksInterval sets the value of WeeksInterval for the instance
func (instance *MSFT_TaskWeeklyTrigger) SetPropertyWeeksInterval(value uint16) (err error) {
	return instance.SetProperty("WeeksInterval", (value))
}

// GetWeeksInterval gets the value of WeeksInterval for the instance
func (instance *MSFT_TaskWeeklyTrigger) GetPropertyWeeksInterval() (value uint16, err error) {
	retValue, err := instance.GetProperty("WeeksInterval")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
