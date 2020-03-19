// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSCluster_ClusterUpgradedEvent struct
type MSCluster_ClusterUpgradedEvent struct {
	*cim.WmiInstance
}

func NewMSCluster_ClusterUpgradedEventEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ClusterUpgradedEvent, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterUpgradedEvent{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_ClusterUpgradedEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ClusterUpgradedEvent, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ClusterUpgradedEvent{
		WmiInstance: tmp,
	}
	return
}
