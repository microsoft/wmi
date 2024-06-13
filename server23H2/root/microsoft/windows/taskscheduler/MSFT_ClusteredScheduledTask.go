// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ClusteredScheduledTask struct
type MSFT_ClusteredScheduledTask struct {
	*cim.WmiInstance

	//
	ClusterName string

	//
	CurrentOwner string

	//
	Resource string

	//
	TaskDefinition MSFT_ScheduledTask

	//
	TaskName string

	//
	TaskType uint32
}

func NewMSFT_ClusteredScheduledTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_ClusteredScheduledTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusteredScheduledTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ClusteredScheduledTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ClusteredScheduledTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusteredScheduledTask{
		WmiInstance: tmp,
	}
	return
}

// SetClusterName sets the value of ClusterName for the instance
func (instance *MSFT_ClusteredScheduledTask) SetPropertyClusterName(value string) (err error) {
	return instance.SetProperty("ClusterName", (value))
}

// GetClusterName gets the value of ClusterName for the instance
func (instance *MSFT_ClusteredScheduledTask) GetPropertyClusterName() (value string, err error) {
	retValue, err := instance.GetProperty("ClusterName")
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

// SetCurrentOwner sets the value of CurrentOwner for the instance
func (instance *MSFT_ClusteredScheduledTask) SetPropertyCurrentOwner(value string) (err error) {
	return instance.SetProperty("CurrentOwner", (value))
}

// GetCurrentOwner gets the value of CurrentOwner for the instance
func (instance *MSFT_ClusteredScheduledTask) GetPropertyCurrentOwner() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentOwner")
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

// SetResource sets the value of Resource for the instance
func (instance *MSFT_ClusteredScheduledTask) SetPropertyResource(value string) (err error) {
	return instance.SetProperty("Resource", (value))
}

// GetResource gets the value of Resource for the instance
func (instance *MSFT_ClusteredScheduledTask) GetPropertyResource() (value string, err error) {
	retValue, err := instance.GetProperty("Resource")
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

// SetTaskDefinition sets the value of TaskDefinition for the instance
func (instance *MSFT_ClusteredScheduledTask) SetPropertyTaskDefinition(value MSFT_ScheduledTask) (err error) {
	return instance.SetProperty("TaskDefinition", (value))
}

// GetTaskDefinition gets the value of TaskDefinition for the instance
func (instance *MSFT_ClusteredScheduledTask) GetPropertyTaskDefinition() (value MSFT_ScheduledTask, err error) {
	retValue, err := instance.GetProperty("TaskDefinition")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_ScheduledTask)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_ScheduledTask is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_ScheduledTask(valuetmp)

	return
}

// SetTaskName sets the value of TaskName for the instance
func (instance *MSFT_ClusteredScheduledTask) SetPropertyTaskName(value string) (err error) {
	return instance.SetProperty("TaskName", (value))
}

// GetTaskName gets the value of TaskName for the instance
func (instance *MSFT_ClusteredScheduledTask) GetPropertyTaskName() (value string, err error) {
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

// SetTaskType sets the value of TaskType for the instance
func (instance *MSFT_ClusteredScheduledTask) SetPropertyTaskType(value uint32) (err error) {
	return instance.SetProperty("TaskType", (value))
}

// GetTaskType gets the value of TaskType for the instance
func (instance *MSFT_ClusteredScheduledTask) GetPropertyTaskType() (value uint32, err error) {
	retValue, err := instance.GetProperty("TaskType")
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
