// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.AppBackgroundTask
//////////////////////////////////////////////
package appbackgroundtask

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_BackgroundTask struct
type MSFT_BackgroundTask struct {
	cim.WmiInstance

	//
	EntryPoint []string

	//
	PackageFullName string

	//
	PerfInfo []string

	//
	TaskID []string

	//
	TaskName []string
}

// SetEntryPoint sets the value of EntryPoint for the instance
func (instance *MSFT_BackgroundTask) SetPropertyEntryPoint(value []string) (err error) {
	return instance.SetProperty("EntryPoint", value)
}

// GetEntryPoint gets the value of EntryPoint for the instance
func (instance *MSFT_BackgroundTask) GetPropertyEntryPoint() (value []string, err error) {
	retValue, err := instance.GetProperty("EntryPoint")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageFullName sets the value of PackageFullName for the instance
func (instance *MSFT_BackgroundTask) SetPropertyPackageFullName(value string) (err error) {
	return instance.SetProperty("PackageFullName", value)
}

// GetPackageFullName gets the value of PackageFullName for the instance
func (instance *MSFT_BackgroundTask) GetPropertyPackageFullName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageFullName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPerfInfo sets the value of PerfInfo for the instance
func (instance *MSFT_BackgroundTask) SetPropertyPerfInfo(value []string) (err error) {
	return instance.SetProperty("PerfInfo", value)
}

// GetPerfInfo gets the value of PerfInfo for the instance
func (instance *MSFT_BackgroundTask) GetPropertyPerfInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("PerfInfo")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTaskID sets the value of TaskID for the instance
func (instance *MSFT_BackgroundTask) SetPropertyTaskID(value []string) (err error) {
	return instance.SetProperty("TaskID", value)
}

// GetTaskID gets the value of TaskID for the instance
func (instance *MSFT_BackgroundTask) GetPropertyTaskID() (value []string, err error) {
	retValue, err := instance.GetProperty("TaskID")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTaskName sets the value of TaskName for the instance
func (instance *MSFT_BackgroundTask) SetPropertyTaskName(value []string) (err error) {
	return instance.SetProperty("TaskName", value)
}

// GetTaskName gets the value of TaskName for the instance
func (instance *MSFT_BackgroundTask) GetPropertyTaskName() (value []string, err error) {
	retValue, err := instance.GetProperty("TaskName")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
