// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

// MSFT_TaskTimeTrigger struct
type MSFT_TaskTimeTrigger struct {
	MSFT_TaskTrigger

	//
	RandomDelay string
}

// SetRandomDelay sets the value of RandomDelay for the instance
func (instance *MSFT_TaskTimeTrigger) SetPropertyRandomDelay(value string) (err error) {
	return instance.SetProperty("RandomDelay", value)
}

// GetRandomDelay gets the value of RandomDelay for the instance
func (instance *MSFT_TaskTimeTrigger) GetPropertyRandomDelay() (value string, err error) {
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
