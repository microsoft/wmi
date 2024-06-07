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

// Msvm_ElementAllocatedFromPool struct
type Msvm_ElementAllocatedFromPool struct { 
	*CIM_ElementAllocatedFromPool
}

	func NewMsvm_ElementAllocatedFromPoolEx1(instance *cim.WmiInstance) (newInstance *Msvm_ElementAllocatedFromPool, err error) {tmp, err := NewCIM_ElementAllocatedFromPoolEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_ElementAllocatedFromPool {
	CIM_ElementAllocatedFromPool: tmp,
	}
	return
	}
	

	func NewMsvm_ElementAllocatedFromPoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_ElementAllocatedFromPool, err error) {tmp, err := NewCIM_ElementAllocatedFromPoolEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_ElementAllocatedFromPool {
	CIM_ElementAllocatedFromPool: tmp,
	}
	return
	}
	

