// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ReplicaSystemDependency struct
type Msvm_ReplicaSystemDependency struct {
	*CIM_Dependency
}

func NewMsvm_ReplicaSystemDependencyEx1(instance *cim.WmiInstance) (newInstance *Msvm_ReplicaSystemDependency, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicaSystemDependency{
		CIM_Dependency: tmp,
	}
	return
}

func NewMsvm_ReplicaSystemDependencyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ReplicaSystemDependency, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicaSystemDependency{
		CIM_Dependency: tmp,
	}
	return
}
