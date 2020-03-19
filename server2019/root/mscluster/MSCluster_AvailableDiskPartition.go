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

// MSCluster_AvailableDiskPartition struct
type MSCluster_AvailableDiskPartition struct {
	*MSCluster_ClusterDiskPartition
}

func NewMSCluster_AvailableDiskPartitionEx1(instance *cim.WmiInstance) (newInstance *MSCluster_AvailableDiskPartition, err error) {
	tmp, err := NewMSCluster_ClusterDiskPartitionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_AvailableDiskPartition{
		MSCluster_ClusterDiskPartition: tmp,
	}
	return
}

func NewMSCluster_AvailableDiskPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_AvailableDiskPartition, err error) {
	tmp, err := NewMSCluster_ClusterDiskPartitionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_AvailableDiskPartition{
		MSCluster_ClusterDiskPartition: tmp,
	}
	return
}
