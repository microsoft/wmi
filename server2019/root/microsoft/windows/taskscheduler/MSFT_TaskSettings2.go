// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

// MSFT_TaskSettings2 struct
type MSFT_TaskSettings2 struct {
	MSFT_TaskSettings

	//
	DisallowStartOnRemoteAppSession bool

	//
	UseUnifiedSchedulingEngine bool
}

// SetDisallowStartOnRemoteAppSession sets the value of DisallowStartOnRemoteAppSession for the instance
func (instance *MSFT_TaskSettings2) SetPropertyDisallowStartOnRemoteAppSession(value bool) (err error) {
	return instance.SetProperty("DisallowStartOnRemoteAppSession", value)
}

// GetDisallowStartOnRemoteAppSession gets the value of DisallowStartOnRemoteAppSession for the instance
func (instance *MSFT_TaskSettings2) GetPropertyDisallowStartOnRemoteAppSession() (value bool, err error) {
	retValue, err := instance.GetProperty("DisallowStartOnRemoteAppSession")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseUnifiedSchedulingEngine sets the value of UseUnifiedSchedulingEngine for the instance
func (instance *MSFT_TaskSettings2) SetPropertyUseUnifiedSchedulingEngine(value bool) (err error) {
	return instance.SetProperty("UseUnifiedSchedulingEngine", value)
}

// GetUseUnifiedSchedulingEngine gets the value of UseUnifiedSchedulingEngine for the instance
func (instance *MSFT_TaskSettings2) GetPropertyUseUnifiedSchedulingEngine() (value bool, err error) {
	retValue, err := instance.GetProperty("UseUnifiedSchedulingEngine")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
