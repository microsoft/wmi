// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_ConfigSetting struct
type MDM_ConfigSetting struct {
	*cim.WmiInstance

	//
	SettingName string

	//
	SettingValue string
}

func NewMDM_ConfigSettingEx1(instance *cim.WmiInstance) (newInstance *MDM_ConfigSetting, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_ConfigSetting{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_ConfigSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_ConfigSetting, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_ConfigSetting{
		WmiInstance: tmp,
	}
	return
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
