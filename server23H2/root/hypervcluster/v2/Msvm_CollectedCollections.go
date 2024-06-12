// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_CollectedCollections struct
type Msvm_CollectedCollections struct {
	*CIM_CollectedMSEs
}

func NewMsvm_CollectedCollectionsEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectedCollections, err error) {
	tmp, err := NewCIM_CollectedMSEsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectedCollections{
		CIM_CollectedMSEs: tmp,
	}
	return
}

func NewMsvm_CollectedCollectionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CollectedCollections, err error) {
	tmp, err := NewCIM_CollectedMSEsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectedCollections{
		CIM_CollectedMSEs: tmp,
	}
	return
}
