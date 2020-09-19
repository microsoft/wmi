// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PowerSettingSubgroup struct
type Win32_PowerSettingSubgroup struct {
	*CIM_SettingData
}

func NewWin32_PowerSettingSubgroupEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSettingSubgroup, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingSubgroup{
		CIM_SettingData: tmp,
	}
	return
}

func NewWin32_PowerSettingSubgroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PowerSettingSubgroup, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingSubgroup{
		CIM_SettingData: tmp,
	}
	return
}
