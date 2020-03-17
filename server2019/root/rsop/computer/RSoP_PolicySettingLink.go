// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSoP_PolicySettingLink struct
type RSoP_PolicySettingLink struct {
	cim.WmiInstance

	//
	setting RSOP_PolicySetting

	//
	status RSoP_PolicySettingStatus
}

// Setsetting sets the value of setting for the instance
func (instance *RSoP_PolicySettingLink) SetPropertysetting(value RSOP_PolicySetting) (err error) {
	return instance.SetProperty("setting", value)
}

// Getsetting gets the value of setting for the instance
func (instance *RSoP_PolicySettingLink) GetPropertysetting() (value RSOP_PolicySetting, err error) {
	retValue, err := instance.GetProperty("setting")
	if err != nil {
		return
	}
	value, ok := retValue.(RSOP_PolicySetting)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setstatus sets the value of status for the instance
func (instance *RSoP_PolicySettingLink) SetPropertystatus(value RSoP_PolicySettingStatus) (err error) {
	return instance.SetProperty("status", value)
}

// Getstatus gets the value of status for the instance
func (instance *RSoP_PolicySettingLink) GetPropertystatus() (value RSoP_PolicySettingStatus, err error) {
	retValue, err := instance.GetProperty("status")
	if err != nil {
		return
	}
	value, ok := retValue.(RSoP_PolicySettingStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}
