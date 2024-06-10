// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualSystemSettingDataComponent struct
type Msvm_VirtualSystemSettingDataComponent struct {
	*CIM_VirtualSystemSettingDataComponent
}

func NewMsvm_VirtualSystemSettingDataComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemSettingDataComponent, err error) {
	tmp, err := NewCIM_VirtualSystemSettingDataComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemSettingDataComponent{
		CIM_VirtualSystemSettingDataComponent: tmp,
	}
	return
}

func NewMsvm_VirtualSystemSettingDataComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemSettingDataComponent, err error) {
	tmp, err := NewCIM_VirtualSystemSettingDataComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemSettingDataComponent{
		CIM_VirtualSystemSettingDataComponent: tmp,
	}
	return
}
