// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

// MSFT_TaskDailyTrigger struct
type MSFT_TaskDailyTrigger struct {
	MSFT_TaskTrigger

	//
	DaysInterval int16

	//
	RandomDelay string
}

// SetDaysInterval sets the value of DaysInterval for the instance
func (instance *MSFT_TaskDailyTrigger) SetPropertyDaysInterval(value int16) (err error) {
	return instance.SetProperty("DaysInterval", value)
}

// GetDaysInterval gets the value of DaysInterval for the instance
func (instance *MSFT_TaskDailyTrigger) GetPropertyDaysInterval() (value int16, err error) {
	retValue, err := instance.GetProperty("DaysInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(int16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRandomDelay sets the value of RandomDelay for the instance
func (instance *MSFT_TaskDailyTrigger) SetPropertyRandomDelay(value string) (err error) {
	return instance.SetProperty("RandomDelay", value)
}

// GetRandomDelay gets the value of RandomDelay for the instance
func (instance *MSFT_TaskDailyTrigger) GetPropertyRandomDelay() (value string, err error) {
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
