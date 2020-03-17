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

// MSFT_TaskIdleSettings struct
type MSFT_TaskIdleSettings struct {
	cim.WmiInstance

	//
	IdleDuration string

	//
	RestartOnIdle bool

	//
	StopOnIdleEnd bool

	//
	WaitTimeout string
}

// SetIdleDuration sets the value of IdleDuration for the instance
func (instance *MSFT_TaskIdleSettings) SetPropertyIdleDuration(value string) (err error) {
	return instance.SetProperty("IdleDuration", value)
}

// GetIdleDuration gets the value of IdleDuration for the instance
func (instance *MSFT_TaskIdleSettings) GetPropertyIdleDuration() (value string, err error) {
	retValue, err := instance.GetProperty("IdleDuration")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestartOnIdle sets the value of RestartOnIdle for the instance
func (instance *MSFT_TaskIdleSettings) SetPropertyRestartOnIdle(value bool) (err error) {
	return instance.SetProperty("RestartOnIdle", value)
}

// GetRestartOnIdle gets the value of RestartOnIdle for the instance
func (instance *MSFT_TaskIdleSettings) GetPropertyRestartOnIdle() (value bool, err error) {
	retValue, err := instance.GetProperty("RestartOnIdle")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStopOnIdleEnd sets the value of StopOnIdleEnd for the instance
func (instance *MSFT_TaskIdleSettings) SetPropertyStopOnIdleEnd(value bool) (err error) {
	return instance.SetProperty("StopOnIdleEnd", value)
}

// GetStopOnIdleEnd gets the value of StopOnIdleEnd for the instance
func (instance *MSFT_TaskIdleSettings) GetPropertyStopOnIdleEnd() (value bool, err error) {
	retValue, err := instance.GetProperty("StopOnIdleEnd")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWaitTimeout sets the value of WaitTimeout for the instance
func (instance *MSFT_TaskIdleSettings) SetPropertyWaitTimeout(value string) (err error) {
	return instance.SetProperty("WaitTimeout", value)
}

// GetWaitTimeout gets the value of WaitTimeout for the instance
func (instance *MSFT_TaskIdleSettings) GetPropertyWaitTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("WaitTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
