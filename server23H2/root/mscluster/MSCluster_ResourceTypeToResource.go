// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_ResourceTypeToResource struct
type MSCluster_ResourceTypeToResource struct {
	*CIM_Component
}

func NewMSCluster_ResourceTypeToResourceEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ResourceTypeToResource, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ResourceTypeToResource{
		CIM_Component: tmp,
	}
	return
}

func NewMSCluster_ResourceTypeToResourceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ResourceTypeToResource, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ResourceTypeToResource{
		CIM_Component: tmp,
	}
	return
}
