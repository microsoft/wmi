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

// Msvm_CollectedVirtualSystems struct
type Msvm_CollectedVirtualSystems struct { 
	*CIM_CollectedMSEs
}

	func NewMsvm_CollectedVirtualSystemsEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectedVirtualSystems, err error) {tmp, err := NewCIM_CollectedMSEsEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_CollectedVirtualSystems {
	CIM_CollectedMSEs: tmp,
	}
	return
	}
	

	func NewMsvm_CollectedVirtualSystemsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_CollectedVirtualSystems, err error) {tmp, err := NewCIM_CollectedMSEsEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_CollectedVirtualSystems {
	CIM_CollectedMSEs: tmp,
	}
	return
	}
	

