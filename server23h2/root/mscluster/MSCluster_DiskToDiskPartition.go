// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_DiskToDiskPartition struct
type MSCluster_DiskToDiskPartition struct {
	*CIM_Component
}

func NewMSCluster_DiskToDiskPartitionEx1(instance *cim.WmiInstance) (newInstance *MSCluster_DiskToDiskPartition, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_DiskToDiskPartition{
		CIM_Component: tmp,
	}
	return
}

func NewMSCluster_DiskToDiskPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_DiskToDiskPartition, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_DiskToDiskPartition{
		CIM_Component: tmp,
	}
	return
}
