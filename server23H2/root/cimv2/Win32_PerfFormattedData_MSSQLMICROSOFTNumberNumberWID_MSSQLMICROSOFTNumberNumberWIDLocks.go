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

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks struct {
	*Win32_PerfFormattedData

	//
	AverageWaitTimems uint64

	//
	LockRequestsPersec uint64

	//
	LockTimeoutsPersec uint64

	//
	LockTimeoutstimeout0Persec uint64

	//
	LockWaitsPersec uint64

	//
	LockWaitTimems uint64

	//
	NumberofDeadlocksPersec uint64
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocksEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocksEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetAverageWaitTimems sets the value of AverageWaitTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) SetPropertyAverageWaitTimems(value uint64) (err error) {
	return instance.SetProperty("AverageWaitTimems", (value))
}

// GetAverageWaitTimems gets the value of AverageWaitTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) GetPropertyAverageWaitTimems() (value uint64, err error) {
	retValue, err := instance.GetProperty("AverageWaitTimems")
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

// SetLockRequestsPersec sets the value of LockRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) SetPropertyLockRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("LockRequestsPersec", (value))
}

// GetLockRequestsPersec gets the value of LockRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) GetPropertyLockRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LockRequestsPersec")
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

// SetLockTimeoutsPersec sets the value of LockTimeoutsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) SetPropertyLockTimeoutsPersec(value uint64) (err error) {
	return instance.SetProperty("LockTimeoutsPersec", (value))
}

// GetLockTimeoutsPersec gets the value of LockTimeoutsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) GetPropertyLockTimeoutsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LockTimeoutsPersec")
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

// SetLockTimeoutstimeout0Persec sets the value of LockTimeoutstimeout0Persec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) SetPropertyLockTimeoutstimeout0Persec(value uint64) (err error) {
	return instance.SetProperty("LockTimeoutstimeout0Persec", (value))
}

// GetLockTimeoutstimeout0Persec gets the value of LockTimeoutstimeout0Persec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) GetPropertyLockTimeoutstimeout0Persec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LockTimeoutstimeout0Persec")
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

// SetLockWaitsPersec sets the value of LockWaitsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) SetPropertyLockWaitsPersec(value uint64) (err error) {
	return instance.SetProperty("LockWaitsPersec", (value))
}

// GetLockWaitsPersec gets the value of LockWaitsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) GetPropertyLockWaitsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LockWaitsPersec")
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

// SetLockWaitTimems sets the value of LockWaitTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) SetPropertyLockWaitTimems(value uint64) (err error) {
	return instance.SetProperty("LockWaitTimems", (value))
}

// GetLockWaitTimems gets the value of LockWaitTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) GetPropertyLockWaitTimems() (value uint64, err error) {
	retValue, err := instance.GetProperty("LockWaitTimems")
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

// SetNumberofDeadlocksPersec sets the value of NumberofDeadlocksPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) SetPropertyNumberofDeadlocksPersec(value uint64) (err error) {
	return instance.SetProperty("NumberofDeadlocksPersec", (value))
}

// GetNumberofDeadlocksPersec gets the value of NumberofDeadlocksPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLocks) GetPropertyNumberofDeadlocksPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NumberofDeadlocksPersec")
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
