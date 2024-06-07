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

// Msvm_EthernetPortInfo struct
type Msvm_EthernetPortInfo struct { 
	*CIM_Dependency
}

	func NewMsvm_EthernetPortInfoEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetPortInfo, err error) {tmp, err := NewCIM_DependencyEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetPortInfo {
	CIM_Dependency: tmp,
	}
	return
	}
	

	func NewMsvm_EthernetPortInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_EthernetPortInfo, err error) {tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetPortInfo {
	CIM_Dependency: tmp,
	}
	return
	}
	

