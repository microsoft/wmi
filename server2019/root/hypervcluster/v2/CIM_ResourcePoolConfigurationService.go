// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ResourcePoolConfigurationService struct
type CIM_ResourcePoolConfigurationService struct { 
	*CIM_Service
}

	func NewCIM_ResourcePoolConfigurationServiceEx1(instance *cim.WmiInstance) (newInstance *CIM_ResourcePoolConfigurationService, err error) {tmp, err := NewCIM_ServiceEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_ResourcePoolConfigurationService {
	CIM_Service: tmp,
	}
	return
	}
	

	func NewCIM_ResourcePoolConfigurationServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_ResourcePoolConfigurationService, err error) {tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_ResourcePoolConfigurationService {
	CIM_Service: tmp,
	}
	return
	}
	

// Starts a job to create a root ResourcePool. The ResourcePool will be scoped to the same System as this Service. If 0 is returned, then the task completed successfully and the use of ConcreteJob was not required. If the task will take some time to complete, a ConcreteJob will be created and its reference returned in the output parameter Job. The resulting pool will be a root pool with no parent pool.

// <param name="ElementName" type="string ">A end user relevant name for the pool being created. If NULL, then a system supplied default name can be used. The value will be stored in the 'ElementName' property for the created pool.</param>
// <param name="HostResources" type="CIM_LogicalDevice []">Array of zero or more devices that are used to create the Pool or modify the source extents. All elements in the array must be of the same type.</param>
// <param name="ResourceType" type="string ">The type of resources the created poolwill manage. If HostResources contains elements, this property must mach their type.</param>

// <param name="Job" type="CIM_ConcreteJob ">Reference to the job (may be null if job completed).</param>
// <param name="Pool" type="CIM_ResourcePool ">On success, a reference to the resulting ResourcePool is returned. When a Job is returned, this may be NULL, in which case, the client must use the Job to find the resulting ResourcePool once the Job completes.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ResourcePoolConfigurationService) CreateResourcePool( /* IN */ ElementName string,
 /* IN */ HostResources []CIM_LogicalDevice,
 /* IN */ ResourceType string,
 /* OUT */ Pool CIM_ResourcePool,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("CreateResourcePool", Action, PercentComplete, Timeout , ElementName, HostResources, ResourceType)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// Start a job to create a sub-pool from a parent pool using the specified allocation settings If 0 is returned, the function completed successfully and no ConcreteJob instance was required. If 4096/0x1000 is returned, a ConcreteJob will be started to create the sub-pool. The Job's reference will be returned in the output parameter Job.

// <param name="ElementName" type="string ">A end user relevant name for the pool being created. If NULL, then a system supplied default name can be used. The value will be stored in the 'ElementName' property for the created element.</param>
// <param name="ParentPool" type="CIM_ResourcePool []">The Pool(s) from which to create the new Pool.</param>
// <param name="Settings" type="string []">String containing a representation of a CIM_SettingData instance that is used to specify the settings for the child Pool.</param>

// <param name="Job" type="CIM_ConcreteJob ">Reference to the job (may be null if job completed).</param>
// <param name="Pool" type="CIM_ResourcePool ">A reference to the resulting pool.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ResourcePoolConfigurationService) CreateChildResourcePool( /* IN */ ElementName string,
 /* IN */ Settings []string,
 /* IN */ ParentPool []CIM_ResourcePool,
 /* OUT */ Pool CIM_ResourcePool,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("CreateChildResourcePool", Action, PercentComplete, Timeout , ElementName, Settings, ParentPool)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// Start a job to delete a ResourcePool. No allocations may be outstanding or the delete will fail with "In Use." If the resource pool is a root resource pool, any host resources are returned back to the underlying system. If 0 is returned, the function completed successfully, and no ConcreteJob was required. If 4096/0x1000 is returned, a ConcreteJob will be started to delete the ResourcePool. A reference to the Job is returned in the Job parameter.

// <param name="Pool" type="CIM_ResourcePool ">Reference to the pool to delete.</param>

// <param name="Job" type="CIM_ConcreteJob ">Reference to the job (may be null if job completed).</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ResourcePoolConfigurationService) DeleteResourcePool( /* IN */ Pool CIM_ResourcePool,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("DeleteResourcePool", Action, PercentComplete, Timeout , Pool)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// Starts a job to add resources to a ResourcePool. If 0 is returned, then the task completed successfully and the use of ConcreteJob was not required. If the task will take some time to complete, a ConcreteJob will be created and its reference returned in the output parameter Job. The resulting pool will be a root pool with no parent pool.

// <param name="HostResources" type="CIM_LogicalDevice []">Array of CIM_LogicalDevice instances to add to the Pool.</param>
// <param name="Pool" type="CIM_ResourcePool ">The pool to add the resources to.</param>

// <param name="Job" type="CIM_ConcreteJob ">Reference to the job (may be null if job completed).</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ResourcePoolConfigurationService) AddResourcesToResourcePool( /* IN */ HostResources []CIM_LogicalDevice,
 /* IN */ Pool CIM_ResourcePool,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("AddResourcesToResourcePool", Action, PercentComplete, Timeout , HostResources, Pool)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// Starts a job to remove resources from a ResourcePool. If 0 is returned, then the task completed successfully and the use of ConcreteJob was not required. If the task will take some time to complete, a ConcreteJob will be created and its reference returned in the output parameter Job. The resulting pool will be a root pool with no parent pool.

// <param name="HostResources" type="CIM_LogicalDevice []">Array of CIM_LogicalDevice instances to remove from the Pool.</param>
// <param name="Pool" type="CIM_ResourcePool ">The pool to remove the resources from.</param>

// <param name="Job" type="CIM_ConcreteJob ">Reference to the job (may be null if job completed).</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ResourcePoolConfigurationService) RemoveResourcesFromResourcePool( /* IN */ HostResources []CIM_LogicalDevice,
 /* IN */ Pool CIM_ResourcePool,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("RemoveResourcesFromResourcePool", Action, PercentComplete, Timeout , HostResources, Pool)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// Start a job to change a parent pool using the specified allocation settings If 0 is returned, the function completed successfully and no ConcreteJob instance was required. If 4096/0x1000 is returned, a ConcreteJob will be started to change the parent pool. The Job's reference will be returned in the output parameter Job.

// <param name="ChildPool" type="CIM_ResourcePool ">Reference to the child pool.</param>
// <param name="ParentPool" type="CIM_ResourcePool []">Reference to the parent pool(s).</param>
// <param name="Settings" type="string []">Optional string containing a representation of a CIM_SettingData instance that is used to specify the settings for the Parent Pool.</param>

// <param name="Job" type="CIM_ConcreteJob ">Reference to the job (may be null if job completed).</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ResourcePoolConfigurationService) ChangeParentResourcePool( /* IN */ ChildPool CIM_ResourcePool,
 /* IN */ ParentPool []CIM_ResourcePool,
 /* IN */ Settings []string,
 /* OUT */ Job CIM_ConcreteJob,
/*Custom IN*/  Action cim.UserAction,
/*Custon IN*/  PercentComplete uint32,
/*Custon IN*/  Timeout uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodAsync("ChangeParentResourcePool", Action, PercentComplete, Timeout , ChildPool, ParentPool, Settings)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


