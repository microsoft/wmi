// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PowerSettingDefineCapabilities struct
type Win32_PowerSettingDefineCapabilities struct {
	*CIM_SettingsDefineCapabilities
}

func NewWin32_PowerSettingDefineCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSettingDefineCapabilities, err error) {
	tmp, err := NewCIM_SettingsDefineCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingDefineCapabilities{
		CIM_SettingsDefineCapabilities: tmp,
	}
	return
}

func NewWin32_PowerSettingDefineCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PowerSettingDefineCapabilities, err error) {
	tmp, err := NewCIM_SettingsDefineCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingDefineCapabilities{
		CIM_SettingsDefineCapabilities: tmp,
	}
	return
}
