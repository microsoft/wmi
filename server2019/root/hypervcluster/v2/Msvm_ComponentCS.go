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

// Msvm_ComponentCS struct
type Msvm_ComponentCS struct { 
	*CIM_ComponentCS
}

	func NewMsvm_ComponentCSEx1(instance *cim.WmiInstance) (newInstance *Msvm_ComponentCS, err error) {tmp, err := NewCIM_ComponentCSEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_ComponentCS {
	CIM_ComponentCS: tmp,
	}
	return
	}
	

	func NewMsvm_ComponentCSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_ComponentCS, err error) {tmp, err := NewCIM_ComponentCSEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_ComponentCS {
	CIM_ComponentCS: tmp,
	}
	return
	}
	

