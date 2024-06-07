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

// Msvm_ManagementCollection struct
type Msvm_ManagementCollection struct { 
	*CIM_CollectionOfMSEs
}

	func NewMsvm_ManagementCollectionEx1(instance *cim.WmiInstance) (newInstance *Msvm_ManagementCollection, err error) {tmp, err := NewCIM_CollectionOfMSEsEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_ManagementCollection {
	CIM_CollectionOfMSEs: tmp,
	}
	return
	}
	

	func NewMsvm_ManagementCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_ManagementCollection, err error) {tmp, err := NewCIM_CollectionOfMSEsEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_ManagementCollection {
	CIM_CollectionOfMSEs: tmp,
	}
	return
	}
	

