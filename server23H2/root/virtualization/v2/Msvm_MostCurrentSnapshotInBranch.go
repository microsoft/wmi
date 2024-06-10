// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_MostCurrentSnapshotInBranch struct
type Msvm_MostCurrentSnapshotInBranch struct {
	*CIM_MostCurrentSnapshotInBranch
}

func NewMsvm_MostCurrentSnapshotInBranchEx1(instance *cim.WmiInstance) (newInstance *Msvm_MostCurrentSnapshotInBranch, err error) {
	tmp, err := NewCIM_MostCurrentSnapshotInBranchEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_MostCurrentSnapshotInBranch{
		CIM_MostCurrentSnapshotInBranch: tmp,
	}
	return
}

func NewMsvm_MostCurrentSnapshotInBranchEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_MostCurrentSnapshotInBranch, err error) {
	tmp, err := NewCIM_MostCurrentSnapshotInBranchEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_MostCurrentSnapshotInBranch{
		CIM_MostCurrentSnapshotInBranch: tmp,
	}
	return
}
