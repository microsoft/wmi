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

// Msvm_CollectedReferencePoints struct
type Msvm_CollectedReferencePoints struct {
	*CIM_CollectedMSEs
}

func NewMsvm_CollectedReferencePointsEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectedReferencePoints, err error) {
	tmp, err := NewCIM_CollectedMSEsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectedReferencePoints{
		CIM_CollectedMSEs: tmp,
	}
	return
}

func NewMsvm_CollectedReferencePointsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CollectedReferencePoints, err error) {
	tmp, err := NewCIM_CollectedMSEsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectedReferencePoints{
		CIM_CollectedMSEs: tmp,
	}
	return
}
