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

// Msvm_VirtualSystemMigrationCapabilities struct
type Msvm_VirtualSystemMigrationCapabilities struct {
	*CIM_VirtualSystemMigrationCapabilities
}

func NewMsvm_VirtualSystemMigrationCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemMigrationCapabilities, err error) {
	tmp, err := NewCIM_VirtualSystemMigrationCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationCapabilities{
		CIM_VirtualSystemMigrationCapabilities: tmp,
	}
	return
}

func NewMsvm_VirtualSystemMigrationCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemMigrationCapabilities, err error) {
	tmp, err := NewCIM_VirtualSystemMigrationCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemMigrationCapabilities{
		CIM_VirtualSystemMigrationCapabilities: tmp,
	}
	return
}

func (instance *Msvm_VirtualSystemMigrationCapabilities) GetRelatedVirtualSystemMigrationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemMigrationService")
}

func (instance *Msvm_VirtualSystemMigrationCapabilities) GetRelatedVirtualSystemMigrationSettingData() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_VirtualSystemMigrationSettingData")
}
