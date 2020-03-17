// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_TaskRepetitionPattern struct
type MSFT_TaskRepetitionPattern struct {
	cim.WmiInstance

	//
	Duration string

	//
	Interval string

	//
	StopAtDurationEnd bool
}

// SetDuration sets the value of Duration for the instance
func (instance *MSFT_TaskRepetitionPattern) SetPropertyDuration(value string) (err error) {
	return instance.SetProperty("Duration", value)
}

// GetDuration gets the value of Duration for the instance
func (instance *MSFT_TaskRepetitionPattern) GetPropertyDuration() (value string, err error) {
	retValue, err := instance.GetProperty("Duration")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterval sets the value of Interval for the instance
func (instance *MSFT_TaskRepetitionPattern) SetPropertyInterval(value string) (err error) {
	return instance.SetProperty("Interval", value)
}

// GetInterval gets the value of Interval for the instance
func (instance *MSFT_TaskRepetitionPattern) GetPropertyInterval() (value string, err error) {
	retValue, err := instance.GetProperty("Interval")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStopAtDurationEnd sets the value of StopAtDurationEnd for the instance
func (instance *MSFT_TaskRepetitionPattern) SetPropertyStopAtDurationEnd(value bool) (err error) {
	return instance.SetProperty("StopAtDurationEnd", value)
}

// GetStopAtDurationEnd gets the value of StopAtDurationEnd for the instance
func (instance *MSFT_TaskRepetitionPattern) GetPropertyStopAtDurationEnd() (value bool, err error) {
	retValue, err := instance.GetProperty("StopAtDurationEnd")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
