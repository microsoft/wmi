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

// Win32_PageFileSetting struct
type Win32_PageFileSetting struct {
	*CIM_Setting

	//
	InitialSize uint32

	//
	MaximumSize uint32

	//
	Name string
}

func NewWin32_PageFileSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_PageFileSetting, err error) {
	tmp, err := NewCIM_SettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PageFileSetting{
		CIM_Setting: tmp,
	}
	return
}

func NewWin32_PageFileSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PageFileSetting, err error) {
	tmp, err := NewCIM_SettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PageFileSetting{
		CIM_Setting: tmp,
	}
	return
}

// SetInitialSize sets the value of InitialSize for the instance
func (instance *Win32_PageFileSetting) SetPropertyInitialSize(value uint32) (err error) {
	return instance.SetProperty("InitialSize", value)
}

// GetInitialSize gets the value of InitialSize for the instance
func (instance *Win32_PageFileSetting) GetPropertyInitialSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("InitialSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumSize sets the value of MaximumSize for the instance
func (instance *Win32_PageFileSetting) SetPropertyMaximumSize(value uint32) (err error) {
	return instance.SetProperty("MaximumSize", value)
}

// GetMaximumSize gets the value of MaximumSize for the instance
func (instance *Win32_PageFileSetting) GetPropertyMaximumSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *Win32_PageFileSetting) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Win32_PageFileSetting) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
