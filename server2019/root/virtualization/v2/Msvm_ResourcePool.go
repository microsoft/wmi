// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ResourcePool struct
type Msvm_ResourcePool struct { 
	*CIM_ResourcePool
}

	func NewMsvm_ResourcePoolEx1(instance *cim.WmiInstance) (newInstance *Msvm_ResourcePool, err error) {tmp, err := NewCIM_ResourcePoolEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_ResourcePool {
	CIM_ResourcePool: tmp,
	}
	return
	}
	

	func NewMsvm_ResourcePoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_ResourcePool, err error) {tmp, err := NewCIM_ResourcePoolEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_ResourcePool {
	CIM_ResourcePool: tmp,
	}
	return
	}
	
func  (instance* Msvm_ResourcePool) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_ComputerSystem"); 
	}
	
func  (instance* Msvm_ResourcePool) GetRelatedResourcePoolSettingData() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_ResourcePoolSettingData"); 
	}
	
func  (instance* Msvm_ResourcePool) GetRelatedResourcePoolConfigurationService() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_ResourcePoolConfigurationService"); 
	}
	
func  (instance* Msvm_ResourcePool) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_AllocationCapabilities"); 
	}
	

