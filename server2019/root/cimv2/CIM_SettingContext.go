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

// CIM_SettingContext struct
type CIM_SettingContext struct {
	cim.WmiInstance

	//
	Context CIM_Configuration

	//
	Setting CIM_Setting
}

// SetContext sets the value of Context for the instance
func (instance *CIM_SettingContext) SetPropertyContext(value CIM_Configuration) (err error) {
	return instance.SetProperty("Context", value)
}

// GetContext gets the value of Context for the instance
func (instance *CIM_SettingContext) GetPropertyContext() (value CIM_Configuration, err error) {
	retValue, err := instance.GetProperty("Context")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Configuration)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetting sets the value of Setting for the instance
func (instance *CIM_SettingContext) SetPropertySetting(value CIM_Setting) (err error) {
	return instance.SetProperty("Setting", value)
}

// GetSetting gets the value of Setting for the instance
func (instance *CIM_SettingContext) GetPropertySetting() (value CIM_Setting, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Setting)
	if !ok {
		// TODO: Set an error
	}
	return
}
