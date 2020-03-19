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

// Win32_PerfRawData_LSM_UserInputDelayperSession struct
type Win32_PerfRawData_LSM_UserInputDelayperSession struct {
	*Win32_PerfRawData

	//
	MaxInputDelay uint64

	//
	MaxInputDelay_Base uint32
}

func NewWin32_PerfRawData_LSM_UserInputDelayperSessionEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_LSM_UserInputDelayperSession, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_LSM_UserInputDelayperSession{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_LSM_UserInputDelayperSessionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_LSM_UserInputDelayperSession, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_LSM_UserInputDelayperSession{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetMaxInputDelay sets the value of MaxInputDelay for the instance
func (instance *Win32_PerfRawData_LSM_UserInputDelayperSession) SetPropertyMaxInputDelay(value uint64) (err error) {
	return instance.SetProperty("MaxInputDelay", value)
}

// GetMaxInputDelay gets the value of MaxInputDelay for the instance
func (instance *Win32_PerfRawData_LSM_UserInputDelayperSession) GetPropertyMaxInputDelay() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxInputDelay")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxInputDelay_Base sets the value of MaxInputDelay_Base for the instance
func (instance *Win32_PerfRawData_LSM_UserInputDelayperSession) SetPropertyMaxInputDelay_Base(value uint32) (err error) {
	return instance.SetProperty("MaxInputDelay_Base", value)
}

// GetMaxInputDelay_Base gets the value of MaxInputDelay_Base for the instance
func (instance *Win32_PerfRawData_LSM_UserInputDelayperSession) GetPropertyMaxInputDelay_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInputDelay_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
