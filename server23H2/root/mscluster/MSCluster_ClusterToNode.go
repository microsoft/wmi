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

// MSCluster_ClusterToNode struct
type MSCluster_ClusterToNode struct {
	*CIM_ParticipatingCS
}

func NewMSCluster_ClusterToNodeEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ClusterToNode, err error) {
	tmp, err := NewCIM_ParticipatingCSEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterToNode{
		CIM_ParticipatingCS: tmp,
	}
	return
}

func NewMSCluster_ClusterToNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ClusterToNode, err error) {
	tmp, err := NewCIM_ParticipatingCSEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterToNode{
		CIM_ParticipatingCS: tmp,
	}
	return
}
