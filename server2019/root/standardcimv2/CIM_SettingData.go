// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_SettingData struct
type CIM_SettingData struct {
	*CIM_ManagedElement
}

func NewCIM_SettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_SettingData, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SettingData{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_SettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SettingData, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SettingData{
		CIM_ManagedElement: tmp,
	}
	return
}
