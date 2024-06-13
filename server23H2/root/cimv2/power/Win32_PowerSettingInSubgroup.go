// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PowerSettingInSubgroup struct
type Win32_PowerSettingInSubgroup struct {
	*CIM_ConcreteComponent
}

func NewWin32_PowerSettingInSubgroupEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSettingInSubgroup, err error) {
	tmp, err := NewCIM_ConcreteComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingInSubgroup{
		CIM_ConcreteComponent: tmp,
	}
	return
}

func NewWin32_PowerSettingInSubgroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PowerSettingInSubgroup, err error) {
	tmp, err := NewCIM_ConcreteComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingInSubgroup{
		CIM_ConcreteComponent: tmp,
	}
	return
}
