// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_ConfigSetting struct
type MDM_ConfigSetting struct {
	cim.WmiInstance

	//
	SettingName string

	//
	SettingValue string
}

// SetSettingName sets the value of SettingName for the instance
func (instance *MDM_ConfigSetting) SetPropertySettingName(value string) (err error) {
	return instance.SetProperty("SettingName", value)
}

// GetSettingName gets the value of SettingName for the instance
func (instance *MDM_ConfigSetting) GetPropertySettingName() (value string, err error) {
	retValue, err := instance.GetProperty("SettingName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettingValue sets the value of SettingValue for the instance
func (instance *MDM_ConfigSetting) SetPropertySettingValue(value string) (err error) {
	return instance.SetProperty("SettingValue", value)
}

// GetSettingValue gets the value of SettingValue for the instance
func (instance *MDM_ConfigSetting) GetPropertySettingValue() (value string, err error) {
	retValue, err := instance.GetProperty("SettingValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
