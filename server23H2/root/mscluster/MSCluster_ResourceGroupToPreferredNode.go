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

// MSCluster_ResourceGroupToPreferredNode struct
type MSCluster_ResourceGroupToPreferredNode struct {
	*CIM_Component
}

func NewMSCluster_ResourceGroupToPreferredNodeEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ResourceGroupToPreferredNode, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ResourceGroupToPreferredNode{
		CIM_Component: tmp,
	}
	return
}

func NewMSCluster_ResourceGroupToPreferredNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ResourceGroupToPreferredNode, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ResourceGroupToPreferredNode{
		CIM_Component: tmp,
	}
	return
}
