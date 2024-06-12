// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.Cluster.Scaleout
//////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ClusterSetClusterToMember struct
type MSFT_ClusterSetClusterToMember struct {
	*CIM_Component
}

func NewMSFT_ClusterSetClusterToMemberEx1(instance *cim.WmiInstance) (newInstance *MSFT_ClusterSetClusterToMember, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetClusterToMember{
		CIM_Component: tmp,
	}
	return
}

func NewMSFT_ClusterSetClusterToMemberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ClusterSetClusterToMember, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetClusterToMember{
		CIM_Component: tmp,
	}
	return
}
