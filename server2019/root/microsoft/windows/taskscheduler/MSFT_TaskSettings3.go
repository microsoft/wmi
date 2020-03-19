// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_TaskSettings3 struct
type MSFT_TaskSettings3 struct {
	*MSFT_TaskSettings2

	//
	MaintenanceSettings MSFT_TaskMaintenanceSettings

	//
	Volatile bool
}

func NewMSFT_TaskSettings3Ex1(instance *cim.WmiInstance) (newInstance *MSFT_TaskSettings3, err error) {
	tmp, err := NewMSFT_TaskSettings2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskSettings3{
		MSFT_TaskSettings2: tmp,
	}
	return
}

func NewMSFT_TaskSettings3Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskSettings3, err error) {
	tmp, err := NewMSFT_TaskSettings2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskSettings3{
		MSFT_TaskSettings2: tmp,
	}
	return
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
