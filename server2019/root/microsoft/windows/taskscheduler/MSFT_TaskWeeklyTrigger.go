// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

// MSFT_TaskWeeklyTrigger struct
type MSFT_TaskWeeklyTrigger struct {
	MSFT_TaskTrigger

	//
	DaysOfWeek uint16

	//
	RandomDelay string

	//
	WeeksInterval uint16
}

// SetDaysOfWeek sets the value of DaysOfWeek for the instance
func (instance *MSFT_TaskWeeklyTrigger) SetPropertyDaysOfWeek(value uint16) (err error) {
	return instance.SetProperty("DaysOfWeek", value)
}

// GetDaysOfWeek gets the value of DaysOfWeek for the instance
func (instance *MSFT_TaskWeeklyTrigger) GetPropertyDaysOfWeek() (value uint16, err error) {
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

// SetRandomDelay sets the value of RandomDelay for the instance
func (instance *MSFT_TaskWeeklyTrigger) SetPropertyRandomDelay(value string) (err error) {
	return instance.SetProperty("RandomDelay", value)
}

// GetRandomDelay gets the value of RandomDelay for the instance
func (instance *MSFT_TaskWeeklyTrigger) GetPropertyRandomDelay() (value string, err error) {
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

// SetWeeksInterval sets the value of WeeksInterval for the instance
func (instance *MSFT_TaskWeeklyTrigger) SetPropertyWeeksInterval(value uint16) (err error) {
	return instance.SetProperty("WeeksInterval", value)
}

// GetWeeksInterval gets the value of WeeksInterval for the instance
func (instance *MSFT_TaskWeeklyTrigger) GetPropertyWeeksInterval() (value uint16, err error) {
	retValue, err := instance.GetProperty("WeeksInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
