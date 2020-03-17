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

// MSFT_TaskDynamicInfo struct
type MSFT_TaskDynamicInfo struct {
	cim.WmiInstance

	//
	LastRunTime string

	//
	LastTaskResult uint32

	//
	NextRunTime string

	//
	NumberOfMissedRuns uint32

	//
	TaskName string

	//
	TaskPath string
}

// SetLastRunTime sets the value of LastRunTime for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyLastRunTime(value string) (err error) {
	return instance.SetProperty("LastRunTime", value)
}

// GetLastRunTime gets the value of LastRunTime for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyLastRunTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastRunTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastTaskResult sets the value of LastTaskResult for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyLastTaskResult(value uint32) (err error) {
	return instance.SetProperty("LastTaskResult", value)
}

// GetLastTaskResult gets the value of LastTaskResult for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyLastTaskResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastTaskResult")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNextRunTime sets the value of NextRunTime for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyNextRunTime(value string) (err error) {
	return instance.SetProperty("NextRunTime", value)
}

// GetNextRunTime gets the value of NextRunTime for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyNextRunTime() (value string, err error) {
	retValue, err := instance.GetProperty("NextRunTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfMissedRuns sets the value of NumberOfMissedRuns for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyNumberOfMissedRuns(value uint32) (err error) {
	return instance.SetProperty("NumberOfMissedRuns", value)
}

// GetNumberOfMissedRuns gets the value of NumberOfMissedRuns for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyNumberOfMissedRuns() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfMissedRuns")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTaskName sets the value of TaskName for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyTaskName(value string) (err error) {
	return instance.SetProperty("TaskName", value)
}

// GetTaskName gets the value of TaskName for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyTaskName() (value string, err error) {
	retValue, err := instance.GetProperty("TaskName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTaskPath sets the value of TaskPath for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyTaskPath(value string) (err error) {
	return instance.SetProperty("TaskPath", value)
}

// GetTaskPath gets the value of TaskPath for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyTaskPath() (value string, err error) {
	retValue, err := instance.GetProperty("TaskPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
