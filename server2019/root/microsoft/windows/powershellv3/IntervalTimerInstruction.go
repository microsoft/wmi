// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Powershellv3
//////////////////////////////////////////////
package powershellv3

// __IntervalTimerInstruction struct
type __IntervalTimerInstruction struct {
	__TimerInstruction

	//
	IntervalBetweenEvents uint32
}

// SetIntervalBetweenEvents sets the value of IntervalBetweenEvents for the instance
func (instance *__IntervalTimerInstruction) SetPropertyIntervalBetweenEvents(value uint32) (err error) {
	return instance.SetProperty("IntervalBetweenEvents", value)
}

// GetIntervalBetweenEvents gets the value of IntervalBetweenEvents for the instance
func (instance *__IntervalTimerInstruction) GetPropertyIntervalBetweenEvents() (value uint32, err error) {
	retValue, err := instance.GetProperty("IntervalBetweenEvents")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
