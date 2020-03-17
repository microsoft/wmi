// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

// __TimerInstruction struct
type __TimerInstruction struct {
	__EventGenerator

	//
	SkipIfPassed bool

	//
	TimerId string
}

// SetSkipIfPassed sets the value of SkipIfPassed for the instance
func (instance *__TimerInstruction) SetPropertySkipIfPassed(value bool) (err error) {
	return instance.SetProperty("SkipIfPassed", value)
}

// GetSkipIfPassed gets the value of SkipIfPassed for the instance
func (instance *__TimerInstruction) GetPropertySkipIfPassed() (value bool, err error) {
	retValue, err := instance.GetProperty("SkipIfPassed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimerId sets the value of TimerId for the instance
func (instance *__TimerInstruction) SetPropertyTimerId(value string) (err error) {
	return instance.SetProperty("TimerId", value)
}

// GetTimerId gets the value of TimerId for the instance
func (instance *__TimerInstruction) GetPropertyTimerId() (value string, err error) {
	retValue, err := instance.GetProperty("TimerId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
