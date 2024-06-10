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

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationAgents struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationAgents struct {
	*Win32_PerfRawData

	//
	Running uint64
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationAgentsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationAgents, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationAgents{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationAgentsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationAgents, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationAgents{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetRunning sets the value of Running for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationAgents) SetPropertyRunning(value uint64) (err error) {
	return instance.SetProperty("Running", (value))
}

// GetRunning gets the value of Running for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationAgents) GetPropertyRunning() (value uint64, err error) {
	retValue, err := instance.GetProperty("Running")
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
