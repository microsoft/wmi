// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_TaskMonthlyDOWTrigger struct
type MSFT_TaskMonthlyDOWTrigger struct {
	*MSFT_TaskTrigger

	//
	DaysOfWeek uint16

	//
	MonthOfYear uint16

	//
	RandomDelay string

	//
	RunOnLastWeekOfMonth bool

	//
	WeeksOfMonth uint16
}

func NewMSFT_TaskMonthlyDOWTriggerEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskMonthlyDOWTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskMonthlyDOWTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

func NewMSFT_TaskMonthlyDOWTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskMonthlyDOWTrigger, err error) {
	tmp, err := NewMSFT_TaskTriggerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskMonthlyDOWTrigger{
		MSFT_TaskTrigger: tmp,
	}
	return
}

// SetDaysOfWeek sets the value of DaysOfWeek for the instance
func (instance *MSFT_TaskMonthlyDOWTrigger) SetPropertyDaysOfWeek(value uint16) (err error) {
	return instance.SetProperty("DaysOfWeek", value)
}

// GetDaysOfWeek gets the value of DaysOfWeek for the instance
func (instance *MSFT_TaskMonthlyDOWTrigger) GetPropertyDaysOfWeek() (value uint16, err error) {
	retValue, err := instance.GetProperty("DaysOfWeek")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonthOfYear sets the value of MonthOfYear for the instance
func (instance *MSFT_TaskMonthlyDOWTrigger) SetPropertyMonthOfYear(value uint16) (err error) {
	return instance.SetProperty("MonthOfYear", value)
}

// GetMonthOfYear gets the value of MonthOfYear for the instance
func (instance *MSFT_TaskMonthlyDOWTrigger) GetPropertyMonthOfYear() (value uint16, err error) {
	retValue, err := instance.GetProperty("MonthOfYear")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRandomDelay sets the value of RandomDelay for the instance
func (instance *MSFT_TaskMonthlyDOWTrigger) SetPropertyRandomDelay(value string) (err error) {
	return instance.SetProperty("RandomDelay", value)
}

// GetRandomDelay gets the value of RandomDelay for the instance
func (instance *MSFT_TaskMonthlyDOWTrigger) GetPropertyRandomDelay() (value string, err error) {
	retValue, err := instance.GetProperty("RandomDelay")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunOnLastWeekOfMonth sets the value of RunOnLastWeekOfMonth for the instance
func (instance *MSFT_TaskMonthlyDOWTrigger) SetPropertyRunOnLastWeekOfMonth(value bool) (err error) {
	return instance.SetProperty("RunOnLastWeekOfMonth", value)
}

// GetRunOnLastWeekOfMonth gets the value of RunOnLastWeekOfMonth for the instance
func (instance *MSFT_TaskMonthlyDOWTrigger) GetPropertyRunOnLastWeekOfMonth() (value bool, err error) {
	retValue, err := instance.GetProperty("RunOnLastWeekOfMonth")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWeeksOfMonth sets the value of WeeksOfMonth for the instance
func (instance *MSFT_TaskMonthlyDOWTrigger) SetPropertyWeeksOfMonth(value uint16) (err error) {
	return instance.SetProperty("WeeksOfMonth", value)
}

// GetWeeksOfMonth gets the value of WeeksOfMonth for the instance
func (instance *MSFT_TaskMonthlyDOWTrigger) GetPropertyWeeksOfMonth() (value uint16, err error) {
	retValue, err := instance.GetProperty("WeeksOfMonth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
