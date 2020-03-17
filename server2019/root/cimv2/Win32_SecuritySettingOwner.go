// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_SecuritySettingOwner struct
type Win32_SecuritySettingOwner struct {
	cim.WmiInstance

	//
	Owner Win32_SID

	//
	SecuritySetting Win32_SecuritySetting
}

// SetOwner sets the value of Owner for the instance
func (instance *Win32_SecuritySettingOwner) SetPropertyOwner(value Win32_SID) (err error) {
	return instance.SetProperty("Owner", value)
}

// GetOwner gets the value of Owner for the instance
func (instance *Win32_SecuritySettingOwner) GetPropertyOwner() (value Win32_SID, err error) {
	retValue, err := instance.GetProperty("Owner")
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
func (instance *Win32_SecuritySettingOwner) SetPropertySecuritySetting(value Win32_SecuritySetting) (err error) {
	return instance.SetProperty("SecuritySetting", value)
}

// GetSecuritySetting gets the value of SecuritySetting for the instance
func (instance *Win32_SecuritySettingOwner) GetPropertySecuritySetting() (value Win32_SecuritySetting, err error) {
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
