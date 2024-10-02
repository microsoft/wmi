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

// Msvm_ResourcePoolConfigurationService struct
type Msvm_ResourcePoolConfigurationService struct {
	*CIM_Service
}

func NewMsvm_ResourcePoolConfigurationServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_ResourcePoolConfigurationService, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ResourcePoolConfigurationService{
		CIM_Service: tmp,
	}
	return
}

func NewMsvm_ResourcePoolConfigurationServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ResourcePoolConfigurationService, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ResourcePoolConfigurationService{
		CIM_Service: tmp,
	}
	return
}

// Starts a job to create a child ResourcePool. The ResourcePool will be scoped to the same System as this Service. If 0 is returned, then the task completed successfully and the use of ConcreteJob was not required. If the task will take some time to complete, a ConcreteJob will be created and its reference returned in the output parameter Job. The resulting pool will be a child pool.

// <param name="AllocationSettings" type="string []">String containing one or more embedded instances of CIM_ResourceAllocationSettingData that is used to specify the pools allocation related settings. This array must contain either one element for each elemnt in the ParentPools array or exactly one element. If this array contains one element and ParentPools contains more than one element, the AlllocationSettings specifies a shared capacity allocation that can be satisfied by any of the parent pools. This is used to restrict the resources that can be allocated from the child to pool to a lower limit than the aggregate capacity provided by its parents. This option is not supported by all resource types. If a resource type does not support shared capacity allocation, this method shall return "Not Supported".</param>
// <param name="ParentPools" type="CIM_ResourcePool []">The Pool(s) from which to create the new Pool.</param>
// <param name="PoolSettings" type="string ">String containing an embedded instance of a Msvm_ResourcePoolSettingData instance that is used to specify the pools non-allocation related settings.</param>

// <param name="Job" type="CIM_ConcreteJob ">Reference to the job (may be null if job completed).</param>
// <param name="Pool" type="CIM_ResourcePool ">A reference to the resulting pool.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ResourcePoolConfigurationService) CreatePool( /* IN */ PoolSettings string,
	/* IN */ ParentPools []CIM_ResourcePool,
	/* IN */ AllocationSettings []string,
	/* OUT */ Pool CIM_ResourcePool,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CreatePool", Action, PercentComplete, Timeout, PoolSettings, ParentPools, AllocationSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Start a job to change parent pool resource settings for resources assigned to a child pool. If 0 is returned, the function completed successfully and no ConcreteJob instance was required. If 4096/0x1000 is returned, a ConcreteJob will be started to change the parent pool. The Job's reference will be returned in the output parameter Job.

// <param name="AllocationSettings" type="string []">Optional string containing a representation of a CIM_SettingData instance that is used to specify the settings for the Parent Pool.</param>
// <param name="ChildPool" type="CIM_ResourcePool ">Reference to the child pool.</param>
// <param name="ParentPools" type="CIM_ResourcePool []">Reference to the parent pool(s).</param>

// <param name="Job" type="CIM_ConcreteJob ">Reference to the job (may be null if job completed).</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ResourcePoolConfigurationService) ModifyPoolResources( /* IN */ ChildPool CIM_ResourcePool,
	/* IN */ ParentPools []CIM_ResourcePool,
	/* IN */ AllocationSettings []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyPoolResources", Action, PercentComplete, Timeout, ChildPool, ParentPools, AllocationSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Start a job to change the non-allocation related settings of a child. If 0 is returned, the function completed successfully and no ConcreteJob instance was required. If 4096/0x1000 is returned, a ConcreteJob will be started to change the settings. The Job's reference will be returned in the output parameter Job.

// <param name="ChildPool" type="CIM_ResourcePool ">Reference to the child pool.</param>
// <param name="PoolSettings" type="string ">String containing an embedded instance of a Msvm_ResourcePoolSettingData that is used to specify the settings for the Pool.</param>

// <param name="Job" type="CIM_ConcreteJob ">Reference to the job (may be null if job completed).</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ResourcePoolConfigurationService) ModifyPoolSettings( /* IN */ ChildPool CIM_ResourcePool,
	/* IN */ PoolSettings string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyPoolSettings", Action, PercentComplete, Timeout, ChildPool, PoolSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Start a job to delete a ResourcePool. No allocations may be outstanding or the delete will fail with "In Use." If the resource pool is a root resource pool, any host resources are returned back to the underlying system. If 0 is returned, the function completed successfully, and no ConcreteJob was required. If 4096/0x1000 is returned, a ConcreteJob will be started to delete the ResourcePool. A reference to the Job is returned in the Job parameter.

// <param name="Pool" type="CIM_ResourcePool ">Reference to the pool to delete.</param>

// <param name="Job" type="CIM_ConcreteJob ">Reference to the job (may be null if job completed).</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ResourcePoolConfigurationService) DeletePool( /* IN */ Pool CIM_ResourcePool,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DeletePool", Action, PercentComplete, Timeout, Pool)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_ResourcePoolConfigurationService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_ResourcePoolConfigurationService) GetRelatedResourcePool() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_ResourcePool")
}

func (instance *Msvm_ResourcePoolConfigurationService) GetRelatedSynth3dVideoPool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_Synth3dVideoPool")
}

func (instance *Msvm_ResourcePoolConfigurationService) GetRelatedProcessorPool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ProcessorPool")
}

func (instance *Msvm_ResourcePoolConfigurationService) GetRelatedResourcePoolConfigurationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePoolConfigurationCapabilities")
}
