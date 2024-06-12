// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_SystemComponentSettingData struct
type Msvm_SystemComponentSettingData struct {
	*CIM_SettingData
}

func NewMsvm_SystemComponentSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_SystemComponentSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SystemComponentSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_SystemComponentSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SystemComponentSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SystemComponentSettingData{
		CIM_SettingData: tmp,
	}
	return
}
