// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

// MSFT_TaskSettings3 struct
type MSFT_TaskSettings3 struct {
	MSFT_TaskSettings2

	//
	MaintenanceSettings MSFT_TaskMaintenanceSettings

	//
	Volatile bool
}

// SetMaintenanceSettings sets the value of MaintenanceSettings for the instance
func (instance *MSFT_TaskSettings3) SetPropertyMaintenanceSettings(value MSFT_TaskMaintenanceSettings) (err error) {
	return instance.SetProperty("MaintenanceSettings", value)
}

// GetMaintenanceSettings gets the value of MaintenanceSettings for the instance
func (instance *MSFT_TaskSettings3) GetPropertyMaintenanceSettings() (value MSFT_TaskMaintenanceSettings, err error) {
	retValue, err := instance.GetProperty("MaintenanceSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_TaskMaintenanceSettings)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setvolatile sets the value of volatile for the instance
func (instance *MSFT_TaskSettings3) SetPropertyvolatile(value bool) (err error) {
	return instance.SetProperty("volatile", value)
}

// Getvolatile gets the value of volatile for the instance
func (instance *MSFT_TaskSettings3) GetPropertyvolatile() (value bool, err error) {
	retValue, err := instance.GetProperty("volatile")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
