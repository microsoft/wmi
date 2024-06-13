// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_ResourceToDependentResource struct
type MSCluster_ResourceToDependentResource struct {
	*CIM_Dependency
}

func NewMSCluster_ResourceToDependentResourceEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ResourceToDependentResource, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ResourceToDependentResource{
		CIM_Dependency: tmp,
	}
	return
}

func NewMSCluster_ResourceToDependentResourceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ResourceToDependentResource, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ResourceToDependentResource{
		CIM_Dependency: tmp,
	}
	return
}
