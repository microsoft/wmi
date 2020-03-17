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

// Msvm_AllocationCapabilities struct
type Msvm_AllocationCapabilities struct {
	CIM_AllocationCapabilities
}

func (instance *Msvm_AllocationCapabilities) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}

func (instance *Msvm_AllocationCapabilities) GetRelatedResourceAllocationSettingData() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_ResourceAllocationSettingData")
}
