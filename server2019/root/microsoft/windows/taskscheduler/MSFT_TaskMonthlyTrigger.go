// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

// MSFT_TaskMonthlyTrigger struct
type MSFT_TaskMonthlyTrigger struct {
	MSFT_TaskTrigger

	//
	DaysOfMonth uint16

	//
	MonthOfYear uint16

	//
	RandomDelay string

	//
	RunOnLastDayOfMonth bool
}

// SetDaysOfMonth sets the value of DaysOfMonth for the instance
func (instance *MSFT_TaskMonthlyTrigger) SetPropertyDaysOfMonth(value uint16) (err error) {
	return instance.SetProperty("DaysOfMonth", value)
}

// GetDaysOfMonth gets the value of DaysOfMonth for the instance
func (instance *MSFT_TaskMonthlyTrigger) GetPropertyDaysOfMonth() (value uint16, err error) {
	retValue, err := instance.GetProperty("DaysOfMonth")
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
func (instance *MSFT_TaskMonthlyTrigger) SetPropertyMonthOfYear(value uint16) (err error) {
	return instance.SetProperty("MonthOfYear", value)
}

// GetMonthOfYear gets the value of MonthOfYear for the instance
func (instance *MSFT_TaskMonthlyTrigger) GetPropertyMonthOfYear() (value uint16, err error) {
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
func (instance *MSFT_TaskMonthlyTrigger) SetPropertyRandomDelay(value string) (err error) {
	return instance.SetProperty("RandomDelay", value)
}

// GetRandomDelay gets the value of RandomDelay for the instance
func (instance *MSFT_TaskMonthlyTrigger) GetPropertyRandomDelay() (value string, err error) {
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

// SetRunOnLastDayOfMonth sets the value of RunOnLastDayOfMonth for the instance
func (instance *MSFT_TaskMonthlyTrigger) SetPropertyRunOnLastDayOfMonth(value bool) (err error) {
	return instance.SetProperty("RunOnLastDayOfMonth", value)
}

// GetRunOnLastDayOfMonth gets the value of RunOnLastDayOfMonth for the instance
func (instance *MSFT_TaskMonthlyTrigger) GetPropertyRunOnLastDayOfMonth() (value bool, err error) {
	retValue, err := instance.GetProperty("RunOnLastDayOfMonth")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
