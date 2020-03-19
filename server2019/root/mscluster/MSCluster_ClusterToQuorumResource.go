// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_ClusterToQuorumResource struct
type MSCluster_ClusterToQuorumResource struct {
	*CIM_Component
}

func NewMSCluster_ClusterToQuorumResourceEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ClusterToQuorumResource, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterToQuorumResource{
		CIM_Component: tmp,
	}
	return
}

func NewMSCluster_ClusterToQuorumResourceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ClusterToQuorumResource, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterToQuorumResource{
		CIM_Component: tmp,
	}
	return
}
