// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.power
//
// ////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PowerSettingCapabilities struct
type Win32_PowerSettingCapabilities struct {
	*CIM_ElementCapabilities
}

func NewWin32_PowerSettingCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSettingCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}

func NewWin32_PowerSettingCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PowerSettingCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}
