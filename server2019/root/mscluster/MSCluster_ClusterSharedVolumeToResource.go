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

// MSCluster_ClusterSharedVolumeToResource struct
type MSCluster_ClusterSharedVolumeToResource struct {
	*CIM_Component
}

func NewMSCluster_ClusterSharedVolumeToResourceEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ClusterSharedVolumeToResource, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterSharedVolumeToResource{
		CIM_Component: tmp,
	}
	return
}

func NewMSCluster_ClusterSharedVolumeToResourceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ClusterSharedVolumeToResource, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterSharedVolumeToResource{
		CIM_Component: tmp,
	}
	return
}
