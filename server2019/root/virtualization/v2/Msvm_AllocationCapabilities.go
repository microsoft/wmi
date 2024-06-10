// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_AllocationCapabilities struct
type Msvm_AllocationCapabilities struct {
	*CIM_AllocationCapabilities
}

func NewMsvm_AllocationCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_AllocationCapabilities, err error) {
	tmp, err := NewCIM_AllocationCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_AllocationCapabilities{
		CIM_AllocationCapabilities: tmp,
	}
	return
}

func NewMsvm_AllocationCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_AllocationCapabilities, err error) {
	tmp, err := NewCIM_AllocationCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_AllocationCapabilities{
		CIM_AllocationCapabilities: tmp,
	}
	return
}

func (instance *Msvm_AllocationCapabilities) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}

func (instance *Msvm_AllocationCapabilities) GetRelatedResourceAllocationSettingData() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_ResourceAllocationSettingData")
}
