// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_DiskPartition struct
type MSCluster_DiskPartition struct {
	*MSCluster_ClusterDiskPartition
}

func NewMSCluster_DiskPartitionEx1(instance *cim.WmiInstance) (newInstance *MSCluster_DiskPartition, err error) {
	tmp, err := NewMSCluster_ClusterDiskPartitionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_DiskPartition{
		MSCluster_ClusterDiskPartition: tmp,
	}
	return
}

func NewMSCluster_DiskPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_DiskPartition, err error) {
	tmp, err := NewMSCluster_ClusterDiskPartitionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_DiskPartition{
		MSCluster_ClusterDiskPartition: tmp,
	}
	return
}
