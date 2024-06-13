// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_MemberOfCollection struct
type Msvm_MemberOfCollection struct {
	*CIM_MemberOfCollection
}

func NewMsvm_MemberOfCollectionEx1(instance *cim.WmiInstance) (newInstance *Msvm_MemberOfCollection, err error) {
	tmp, err := NewCIM_MemberOfCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_MemberOfCollection{
		CIM_MemberOfCollection: tmp,
	}
	return
}

func NewMsvm_MemberOfCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_MemberOfCollection, err error) {
	tmp, err := NewCIM_MemberOfCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_MemberOfCollection{
		CIM_MemberOfCollection: tmp,
	}
	return
}
