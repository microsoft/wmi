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

// Msvm_VirtualFcSwitchSettingData struct
type Msvm_VirtualFcSwitchSettingData struct {
	*CIM_VirtualSystemSettingData
}

func NewMsvm_VirtualFcSwitchSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualFcSwitchSettingData, err error) {
	tmp, err := NewCIM_VirtualSystemSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualFcSwitchSettingData{
		CIM_VirtualSystemSettingData: tmp,
	}
	return
}

func NewMsvm_VirtualFcSwitchSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualFcSwitchSettingData, err error) {
	tmp, err := NewCIM_VirtualSystemSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualFcSwitchSettingData{
		CIM_VirtualSystemSettingData: tmp,
	}
	return
}
