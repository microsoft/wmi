// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.MSCluster
//
// ////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_ResourceToPossibleOwner struct
type MSCluster_ResourceToPossibleOwner struct {
	*CIM_Component
}

func NewMSCluster_ResourceToPossibleOwnerEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ResourceToPossibleOwner, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ResourceToPossibleOwner{
		CIM_Component: tmp,
	}
	return
}

func NewMSCluster_ResourceToPossibleOwnerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ResourceToPossibleOwner, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ResourceToPossibleOwner{
		CIM_Component: tmp,
	}
	return
}
