// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ScheduledJob struct
type Win32_ScheduledJob struct {
	*CIM_Job

	//
	Command string

	//
	DaysOfMonth uint32

	//
	DaysOfWeek uint32

	//
	InteractWithDesktop bool

	//
	JobId uint32

	//
	RunRepeatedly bool
}

func NewWin32_ScheduledJobEx1(instance *cim.WmiInstance) (newInstance *Win32_ScheduledJob, err error) {
	tmp, err := NewCIM_JobEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_ScheduledJob{
		CIM_Job: tmp,
	}
	return
}

func NewWin32_ScheduledJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ScheduledJob, err error) {
	tmp, err := NewCIM_JobEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ScheduledJob{
		CIM_Job: tmp,
	}
	return
}

// SetCommand sets the value of Command for the instance
func (instance *Win32_ScheduledJob) SetPropertyCommand(value string) (err error) {
	return instance.SetProperty("Command", value)
}

// GetCommand gets the value of Command for the instance
func (instance *Win32_ScheduledJob) GetPropertyCommand() (value string, err error) {
	retValue, err := instance.GetProperty("Command")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaysOfMonth sets the value of DaysOfMonth for the instance
func (instance *Win32_ScheduledJob) SetPropertyDaysOfMonth(value uint32) (err error) {
	return instance.SetProperty("DaysOfMonth", value)
}

// GetDaysOfMonth gets the value of DaysOfMonth for the instance
func (instance *Win32_ScheduledJob) GetPropertyDaysOfMonth() (value uint32, err error) {
	retValue, err := instance.GetProperty("DaysOfMonth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaysOfWeek sets the value of DaysOfWeek for the instance
func (instance *Win32_ScheduledJob) SetPropertyDaysOfWeek(value uint32) (err error) {
	return instance.SetProperty("DaysOfWeek", value)
}

// GetDaysOfWeek gets the value of DaysOfWeek for the instance
func (instance *Win32_ScheduledJob) GetPropertyDaysOfWeek() (value uint32, err error) {
	retValue, err := instance.GetProperty("DaysOfWeek")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInteractWithDesktop sets the value of InteractWithDesktop for the instance
func (instance *Win32_ScheduledJob) SetPropertyInteractWithDesktop(value bool) (err error) {
	return instance.SetProperty("InteractWithDesktop", value)
}

// GetInteractWithDesktop gets the value of InteractWithDesktop for the instance
func (instance *Win32_ScheduledJob) GetPropertyInteractWithDesktop() (value bool, err error) {
	retValue, err := instance.GetProperty("InteractWithDesktop")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobId sets the value of JobId for the instance
func (instance *Win32_ScheduledJob) SetPropertyJobId(value uint32) (err error) {
	return instance.SetProperty("JobId", value)
}

// GetJobId gets the value of JobId for the instance
func (instance *Win32_ScheduledJob) GetPropertyJobId() (value uint32, err error) {
	retValue, err := instance.GetProperty("JobId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunRepeatedly sets the value of RunRepeatedly for the instance
func (instance *Win32_ScheduledJob) SetPropertyRunRepeatedly(value bool) (err error) {
	return instance.SetProperty("RunRepeatedly", value)
}

// GetRunRepeatedly gets the value of RunRepeatedly for the instance
func (instance *Win32_ScheduledJob) GetPropertyRunRepeatedly() (value bool, err error) {
	retValue, err := instance.GetProperty("RunRepeatedly")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Command" type="string "></param>
// <param name="DaysOfMonth" type="uint32 "></param>
// <param name="DaysOfWeek" type="uint32 "></param>
// <param name="InteractWithDesktop" type="bool "></param>
// <param name="RunRepeatedly" type="bool "></param>
// <param name="StartTime" type="string "></param>

// <param name="JobId" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_ScheduledJob) Create( /* IN */ Command string,
	/* IN */ StartTime string,
	/* OPTIONAL IN */ RunRepeatedly bool,
	/* OPTIONAL IN */ DaysOfWeek uint32,
	/* OPTIONAL IN */ DaysOfMonth uint32,
	/* OPTIONAL IN */ InteractWithDesktop bool,
	/* OUT */ JobId uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Create", Command, StartTime, RunRepeatedly, DaysOfWeek, DaysOfMonth, InteractWithDesktop)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_ScheduledJob) Delete() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Delete")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
