// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_GuestServiceInterfaceSettingDataComponent struct
type Msvm_GuestServiceInterfaceSettingDataComponent struct {
	*CIM_Component
}

func NewMsvm_GuestServiceInterfaceSettingDataComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_GuestServiceInterfaceSettingDataComponent, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestServiceInterfaceSettingDataComponent{
		CIM_Component: tmp,
	}
	return
}

func NewMsvm_GuestServiceInterfaceSettingDataComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_GuestServiceInterfaceSettingDataComponent, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestServiceInterfaceSettingDataComponent{
		CIM_Component: tmp,
	}
	return
}
