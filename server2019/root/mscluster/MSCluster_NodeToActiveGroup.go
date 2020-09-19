// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_NodeToActiveGroup struct
type MSCluster_NodeToActiveGroup struct {
	*CIM_Component
}

func NewMSCluster_NodeToActiveGroupEx1(instance *cim.WmiInstance) (newInstance *MSCluster_NodeToActiveGroup, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_NodeToActiveGroup{
		CIM_Component: tmp,
	}
	return
}

func NewMSCluster_NodeToActiveGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_NodeToActiveGroup, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_NodeToActiveGroup{
		CIM_Component: tmp,
	}
	return
}
