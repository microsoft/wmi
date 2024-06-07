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

// Msvm_HostedDependency struct
type Msvm_HostedDependency struct { 
	*CIM_HostedDependency
}

	func NewMsvm_HostedDependencyEx1(instance *cim.WmiInstance) (newInstance *Msvm_HostedDependency, err error) {tmp, err := NewCIM_HostedDependencyEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_HostedDependency {
	CIM_HostedDependency: tmp,
	}
	return
	}
	

	func NewMsvm_HostedDependencyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_HostedDependency, err error) {tmp, err := NewCIM_HostedDependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_HostedDependency {
	CIM_HostedDependency: tmp,
	}
	return
	}
	

