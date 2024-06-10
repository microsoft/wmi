// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation struct {
	*Win32_PerfRawData

	//
	StoredProceduresInvokedPersec uint64

	//
	TaskLimitReached uint64

	//
	TaskLimitReachedPersec uint64

	//
	TasksAbortedPersec uint64

	//
	TasksRunning uint64

	//
	TasksStartedPersec uint64
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivationEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetStoredProceduresInvokedPersec sets the value of StoredProceduresInvokedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) SetPropertyStoredProceduresInvokedPersec(value uint64) (err error) {
	return instance.SetProperty("StoredProceduresInvokedPersec", (value))
}

// GetStoredProceduresInvokedPersec gets the value of StoredProceduresInvokedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) GetPropertyStoredProceduresInvokedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("StoredProceduresInvokedPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTaskLimitReached sets the value of TaskLimitReached for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) SetPropertyTaskLimitReached(value uint64) (err error) {
	return instance.SetProperty("TaskLimitReached", (value))
}

// GetTaskLimitReached gets the value of TaskLimitReached for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) GetPropertyTaskLimitReached() (value uint64, err error) {
	retValue, err := instance.GetProperty("TaskLimitReached")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTaskLimitReachedPersec sets the value of TaskLimitReachedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) SetPropertyTaskLimitReachedPersec(value uint64) (err error) {
	return instance.SetProperty("TaskLimitReachedPersec", (value))
}

// GetTaskLimitReachedPersec gets the value of TaskLimitReachedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) GetPropertyTaskLimitReachedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TaskLimitReachedPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTasksAbortedPersec sets the value of TasksAbortedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) SetPropertyTasksAbortedPersec(value uint64) (err error) {
	return instance.SetProperty("TasksAbortedPersec", (value))
}

// GetTasksAbortedPersec gets the value of TasksAbortedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) GetPropertyTasksAbortedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TasksAbortedPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTasksRunning sets the value of TasksRunning for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) SetPropertyTasksRunning(value uint64) (err error) {
	return instance.SetProperty("TasksRunning", (value))
}

// GetTasksRunning gets the value of TasksRunning for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) GetPropertyTasksRunning() (value uint64, err error) {
	retValue, err := instance.GetProperty("TasksRunning")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTasksStartedPersec sets the value of TasksStartedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) SetPropertyTasksStartedPersec(value uint64) (err error) {
	return instance.SetProperty("TasksStartedPersec", (value))
}

// GetTasksStartedPersec gets the value of TasksStartedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerActivation) GetPropertyTasksStartedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TasksStartedPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
