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

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode struct {
	*Win32_PerfFormattedData

	//
	DatabaseNodeMemoryKB uint64

	//
	ForeignNodeMemoryKB uint64

	//
	FreeNodeMemoryKB uint64

	//
	StolenNodeMemoryKB uint64

	//
	TargetNodeMemoryKB uint64

	//
	TotalNodeMemoryKB uint64
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNodeEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetDatabaseNodeMemoryKB sets the value of DatabaseNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) SetPropertyDatabaseNodeMemoryKB(value uint64) (err error) {
	return instance.SetProperty("DatabaseNodeMemoryKB", (value))
}

// GetDatabaseNodeMemoryKB gets the value of DatabaseNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) GetPropertyDatabaseNodeMemoryKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("DatabaseNodeMemoryKB")
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

// SetForeignNodeMemoryKB sets the value of ForeignNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) SetPropertyForeignNodeMemoryKB(value uint64) (err error) {
	return instance.SetProperty("ForeignNodeMemoryKB", (value))
}

// GetForeignNodeMemoryKB gets the value of ForeignNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) GetPropertyForeignNodeMemoryKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("ForeignNodeMemoryKB")
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

// SetFreeNodeMemoryKB sets the value of FreeNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) SetPropertyFreeNodeMemoryKB(value uint64) (err error) {
	return instance.SetProperty("FreeNodeMemoryKB", (value))
}

// GetFreeNodeMemoryKB gets the value of FreeNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) GetPropertyFreeNodeMemoryKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("FreeNodeMemoryKB")
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

// SetStolenNodeMemoryKB sets the value of StolenNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) SetPropertyStolenNodeMemoryKB(value uint64) (err error) {
	return instance.SetProperty("StolenNodeMemoryKB", (value))
}

// GetStolenNodeMemoryKB gets the value of StolenNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) GetPropertyStolenNodeMemoryKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("StolenNodeMemoryKB")
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

// SetTargetNodeMemoryKB sets the value of TargetNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) SetPropertyTargetNodeMemoryKB(value uint64) (err error) {
	return instance.SetProperty("TargetNodeMemoryKB", (value))
}

// GetTargetNodeMemoryKB gets the value of TargetNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) GetPropertyTargetNodeMemoryKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("TargetNodeMemoryKB")
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

// SetTotalNodeMemoryKB sets the value of TotalNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) SetPropertyTotalNodeMemoryKB(value uint64) (err error) {
	return instance.SetProperty("TotalNodeMemoryKB", (value))
}

// GetTotalNodeMemoryKB gets the value of TotalNodeMemoryKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDMemoryNode) GetPropertyTotalNodeMemoryKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalNodeMemoryKB")
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
