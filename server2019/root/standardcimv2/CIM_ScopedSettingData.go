// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ScopedSettingData struct
type CIM_ScopedSettingData struct {
	*CIM_SettingData
}

func NewCIM_ScopedSettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_ScopedSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ScopedSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewCIM_ScopedSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ScopedSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ScopedSettingData{
		CIM_SettingData: tmp,
	}
	return
}
