// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_SnapshotOfVirtualSystem struct
type Msvm_SnapshotOfVirtualSystem struct {
	*CIM_SnapshotOfVirtualSystem
}

func NewMsvm_SnapshotOfVirtualSystemEx1(instance *cim.WmiInstance) (newInstance *Msvm_SnapshotOfVirtualSystem, err error) {
	tmp, err := NewCIM_SnapshotOfVirtualSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SnapshotOfVirtualSystem{
		CIM_SnapshotOfVirtualSystem: tmp,
	}
	return
}

func NewMsvm_SnapshotOfVirtualSystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SnapshotOfVirtualSystem, err error) {
	tmp, err := NewCIM_SnapshotOfVirtualSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SnapshotOfVirtualSystem{
		CIM_SnapshotOfVirtualSystem: tmp,
	}
	return
}
