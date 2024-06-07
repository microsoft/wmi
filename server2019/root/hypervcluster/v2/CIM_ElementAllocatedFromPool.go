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

// CIM_ElementAllocatedFromPool struct
type CIM_ElementAllocatedFromPool struct { 
	*CIM_Dependency
}

	func NewCIM_ElementAllocatedFromPoolEx1(instance *cim.WmiInstance) (newInstance *CIM_ElementAllocatedFromPool, err error) {tmp, err := NewCIM_DependencyEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_ElementAllocatedFromPool {
	CIM_Dependency: tmp,
	}
	return
	}
	

	func NewCIM_ElementAllocatedFromPoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_ElementAllocatedFromPool, err error) {tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_ElementAllocatedFromPool {
	CIM_Dependency: tmp,
	}
	return
	}
	

