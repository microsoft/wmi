// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualSystemMigrationCapabilities struct
type Msvm_VirtualSystemMigrationCapabilities struct {
	CIM_VirtualSystemMigrationCapabilities
}

func (instance *Msvm_VirtualSystemMigrationCapabilities) GetRelatedVirtualSystemMigrationSettingData() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_VirtualSystemMigrationSettingData")
}

func (instance *Msvm_VirtualSystemMigrationCapabilities) GetRelatedVirtualSystemMigrationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemMigrationService")
}
