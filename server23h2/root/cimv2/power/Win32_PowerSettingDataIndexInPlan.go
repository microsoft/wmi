// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PowerSettingDataIndexInPlan struct
type Win32_PowerSettingDataIndexInPlan struct {
	*CIM_ConcreteComponent
}

func NewWin32_PowerSettingDataIndexInPlanEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSettingDataIndexInPlan, err error) {
	tmp, err := NewCIM_ConcreteComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingDataIndexInPlan{
		CIM_ConcreteComponent: tmp,
	}
	return
}

func NewWin32_PowerSettingDataIndexInPlanEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PowerSettingDataIndexInPlan, err error) {
	tmp, err := NewCIM_ConcreteComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSettingDataIndexInPlan{
		CIM_ConcreteComponent: tmp,
	}
	return
}
