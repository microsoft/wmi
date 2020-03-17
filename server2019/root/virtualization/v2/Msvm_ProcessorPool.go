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

// Msvm_ProcessorPool struct
type Msvm_ProcessorPool struct {
	CIM_ResourcePool
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

func (instance *Msvm_ProcessorPool) GetRelatedResourcePoolSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePoolSettingData")
}

func (instance *Msvm_ProcessorPool) GetRelatedResourcePoolConfigurationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePoolConfigurationService")
}

func (instance *Msvm_ProcessorPool) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}

func (instance *Msvm_ProcessorPool) GetRelatedProcessor() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_Processor")
}
