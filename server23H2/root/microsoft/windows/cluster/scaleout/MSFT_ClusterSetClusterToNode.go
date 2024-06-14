// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Cluster.Scaleout
//////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ClusterSetClusterToNode struct
type MSFT_ClusterSetClusterToNode struct {
	*CIM_Component
}

func NewMSFT_ClusterSetClusterToNodeEx1(instance *cim.WmiInstance) (newInstance *MSFT_ClusterSetClusterToNode, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetClusterToNode{
		CIM_Component: tmp,
	}
	return
}

func NewMSFT_ClusterSetClusterToNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ClusterSetClusterToNode, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetClusterToNode{
		CIM_Component: tmp,
	}
	return
}
