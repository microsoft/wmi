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

// Msvm_EthernetPortFailoverSettingDataComponent struct
type Msvm_EthernetPortFailoverSettingDataComponent struct {
	*CIM_Component
}

func NewMsvm_EthernetPortFailoverSettingDataComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetPortFailoverSettingDataComponent, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetPortFailoverSettingDataComponent{
		CIM_Component: tmp,
	}
	return
}

func NewMsvm_EthernetPortFailoverSettingDataComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetPortFailoverSettingDataComponent, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetPortFailoverSettingDataComponent{
		CIM_Component: tmp,
	}
	return
}
