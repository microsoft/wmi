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

// Msvm_ResourcePool struct
type Msvm_ResourcePool struct {
	CIM_ResourcePool
}

func (instance *Msvm_ResourcePool) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_ResourcePool) GetRelatedResourcePoolSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePoolSettingData")
}

func (instance *Msvm_ResourcePool) GetRelatedResourcePoolConfigurationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePoolConfigurationService")
}

func (instance *Msvm_ResourcePool) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
