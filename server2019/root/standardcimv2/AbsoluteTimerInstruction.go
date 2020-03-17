// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// __AbsoluteTimerInstruction struct
type __AbsoluteTimerInstruction struct {
	__TimerInstruction

	//
	EventDateTime string
}

// SetEventDateTime sets the value of EventDateTime for the instance
func (instance *__AbsoluteTimerInstruction) SetPropertyEventDateTime(value string) (err error) {
	return instance.SetProperty("EventDateTime", value)
}

// GetEventDateTime gets the value of EventDateTime for the instance
func (instance *__AbsoluteTimerInstruction) GetPropertyEventDateTime() (value string, err error) {
	retValue, err := instance.GetProperty("EventDateTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
