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

// MSFT_TaskMaintenanceSettings struct
type MSFT_TaskMaintenanceSettings struct {
	cim.WmiInstance

	//
	Deadline string

	//
	Exclusive bool

	//
	Period string
}

// SetDeadline sets the value of Deadline for the instance
func (instance *MSFT_TaskMaintenanceSettings) SetPropertyDeadline(value string) (err error) {
	return instance.SetProperty("Deadline", value)
}

// GetDeadline gets the value of Deadline for the instance
func (instance *MSFT_TaskMaintenanceSettings) GetPropertyDeadline() (value string, err error) {
	retValue, err := instance.GetProperty("Deadline")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExclusive sets the value of Exclusive for the instance
func (instance *MSFT_TaskMaintenanceSettings) SetPropertyExclusive(value bool) (err error) {
	return instance.SetProperty("Exclusive", value)
}

// GetExclusive gets the value of Exclusive for the instance
func (instance *MSFT_TaskMaintenanceSettings) GetPropertyExclusive() (value bool, err error) {
	retValue, err := instance.GetProperty("Exclusive")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeriod sets the value of Period for the instance
func (instance *MSFT_TaskMaintenanceSettings) SetPropertyPeriod(value string) (err error) {
	return instance.SetProperty("Period", value)
}

// GetPeriod gets the value of Period for the instance
func (instance *MSFT_TaskMaintenanceSettings) GetPropertyPeriod() (value string, err error) {
	retValue, err := instance.GetProperty("Period")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
