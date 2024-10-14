// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_SnapshotOfVirtualSystemCollection struct
type Msvm_SnapshotOfVirtualSystemCollection struct {
	*CIM_Dependency
}

func NewMsvm_SnapshotOfVirtualSystemCollectionEx1(instance *cim.WmiInstance) (newInstance *Msvm_SnapshotOfVirtualSystemCollection, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SnapshotOfVirtualSystemCollection{
		CIM_Dependency: tmp,
	}
	return
}

func NewMsvm_SnapshotOfVirtualSystemCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SnapshotOfVirtualSystemCollection, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SnapshotOfVirtualSystemCollection{
		CIM_Dependency: tmp,
	}
	return
}
