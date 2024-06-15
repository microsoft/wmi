// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualEthernetSwitchSettingDataComponent struct
type Msvm_VirtualEthernetSwitchSettingDataComponent struct {
	*CIM_Component
}

func NewMsvm_VirtualEthernetSwitchSettingDataComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualEthernetSwitchSettingDataComponent, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchSettingDataComponent{
		CIM_Component: tmp,
	}
	return
}

func NewMsvm_VirtualEthernetSwitchSettingDataComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualEthernetSwitchSettingDataComponent, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchSettingDataComponent{
		CIM_Component: tmp,
	}
	return
}
