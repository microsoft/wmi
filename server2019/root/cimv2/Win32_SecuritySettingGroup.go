// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_SecuritySettingGroup struct
type Win32_SecuritySettingGroup struct {
	*cim.WmiInstance

	//
	Group Win32_SID

	//
	SecuritySetting Win32_SecuritySetting
}

func NewWin32_SecuritySettingGroupEx1(instance *cim.WmiInstance) (newInstance *Win32_SecuritySettingGroup, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_SecuritySettingGroup{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_SecuritySettingGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_SecuritySettingGroup, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_SecuritySettingGroup{
		WmiInstance: tmp,
	}
	return
}

// SetGroup sets the value of Group for the instance
func (instance *Win32_SecuritySettingGroup) SetPropertyGroup(value Win32_SID) (err error) {
	return instance.SetProperty("Group", value)
}

// GetGroup gets the value of Group for the instance
func (instance *Win32_SecuritySettingGroup) GetPropertyGroup() (value Win32_SID, err error) {
	retValue, err := instance.GetProperty("Group")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_SID)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecuritySetting sets the value of SecuritySetting for the instance
func (instance *Win32_SecuritySettingGroup) SetPropertySecuritySetting(value Win32_SecuritySetting) (err error) {
	return instance.SetProperty("SecuritySetting", value)
}

// GetSecuritySetting gets the value of SecuritySetting for the instance
func (instance *Win32_SecuritySettingGroup) GetPropertySecuritySetting() (value Win32_SecuritySetting, err error) {
	retValue, err := instance.GetProperty("SecuritySetting")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_SecuritySetting)
	if !ok {
		// TODO: Set an error
	}
	return
}
