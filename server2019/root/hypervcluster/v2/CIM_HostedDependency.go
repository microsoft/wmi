// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_HostedDependency struct
type CIM_HostedDependency struct {
	*CIM_Dependency
}

func NewCIM_HostedDependencyEx1(instance *cim.WmiInstance) (newInstance *CIM_HostedDependency, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_HostedDependency{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_HostedDependencyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_HostedDependency, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_HostedDependency{
		CIM_Dependency: tmp,
	}
	return
}
