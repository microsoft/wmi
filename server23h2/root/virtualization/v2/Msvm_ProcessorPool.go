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

// Msvm_ProcessorPool struct
type Msvm_ProcessorPool struct {
	*CIM_ResourcePool
}

func NewMsvm_ProcessorPoolEx1(instance *cim.WmiInstance) (newInstance *Msvm_ProcessorPool, err error) {
	tmp, err := NewCIM_ResourcePoolEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ProcessorPool{
		CIM_ResourcePool: tmp,
	}
	return
}

func NewMsvm_ProcessorPoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ProcessorPool, err error) {
	tmp, err := NewCIM_ResourcePoolEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ProcessorPool{
		CIM_ResourcePool: tmp,
	}
	return
}

//

// <param name="ProcessorCount" type="uint16 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ProcessorPool) CalculatePossibleReserve( /* IN */ ProcessorCount uint16) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CalculatePossibleReserve", ProcessorCount)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

func (instance *Msvm_ProcessorPool) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_ProcessorPool) GetRelatedResourcePoolConfigurationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePoolConfigurationService")
}

func (instance *Msvm_ProcessorPool) GetRelatedResourcePoolSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePoolSettingData")
}

func (instance *Msvm_ProcessorPool) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}

func (instance *Msvm_ProcessorPool) GetRelatedProcessor() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_Processor")
}
