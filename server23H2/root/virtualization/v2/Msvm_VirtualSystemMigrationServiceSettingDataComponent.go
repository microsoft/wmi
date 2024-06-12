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

// Msvm_VirtualSystemMigrationServiceSettingDataComponent struct
type Msvm_VirtualSystemMigrationServiceSettingDataComponent struct {
	*CIM_Component
}

func NewMsvm_VirtualSystemMigrationServiceSettingDataComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemMigrationServiceSettingDataComponent, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationServiceSettingDataComponent{
		CIM_Component: tmp,
	}
	return
}

func NewMsvm_VirtualSystemMigrationServiceSettingDataComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemMigrationServiceSettingDataComponent, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationServiceSettingDataComponent{
		CIM_Component: tmp,
	}
	return
}
