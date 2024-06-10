// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.TaskScheduler
//
// ////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TaskDynamicInfo struct
type MSFT_TaskDynamicInfo struct {
	*cim.WmiInstance

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

func NewMSFT_TaskDynamicInfoEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskDynamicInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskDynamicInfo{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_TaskDynamicInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskDynamicInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskDynamicInfo{
		WmiInstance: tmp,
	}
	return
}

// SetLastRunTime sets the value of LastRunTime for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyLastRunTime(value string) (err error) {
	return instance.SetProperty("LastRunTime", (value))
}

// GetLastRunTime gets the value of LastRunTime for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyLastRunTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastRunTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLastTaskResult sets the value of LastTaskResult for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyLastTaskResult(value uint32) (err error) {
	return instance.SetProperty("LastTaskResult", (value))
}

// GetLastTaskResult gets the value of LastTaskResult for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyLastTaskResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastTaskResult")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetNextRunTime sets the value of NextRunTime for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyNextRunTime(value string) (err error) {
	return instance.SetProperty("NextRunTime", (value))
}

// GetNextRunTime gets the value of NextRunTime for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyNextRunTime() (value string, err error) {
	retValue, err := instance.GetProperty("NextRunTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetNumberOfMissedRuns sets the value of NumberOfMissedRuns for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyNumberOfMissedRuns(value uint32) (err error) {
	return instance.SetProperty("NumberOfMissedRuns", (value))
}

// GetNumberOfMissedRuns gets the value of NumberOfMissedRuns for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyNumberOfMissedRuns() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfMissedRuns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetTaskName sets the value of TaskName for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyTaskName(value string) (err error) {
	return instance.SetProperty("TaskName", (value))
}

// GetTaskName gets the value of TaskName for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyTaskName() (value string, err error) {
	retValue, err := instance.GetProperty("TaskName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetTaskPath sets the value of TaskPath for the instance
func (instance *MSFT_TaskDynamicInfo) SetPropertyTaskPath(value string) (err error) {
	return instance.SetProperty("TaskPath", (value))
}

// GetTaskPath gets the value of TaskPath for the instance
func (instance *MSFT_TaskDynamicInfo) GetPropertyTaskPath() (value string, err error) {
	retValue, err := instance.GetProperty("TaskPath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
