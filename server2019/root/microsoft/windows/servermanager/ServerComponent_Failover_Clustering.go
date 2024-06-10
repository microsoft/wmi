// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.ServerManager
//
// ////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ServerComponent_Failover_Clustering struct
type ServerComponent_Failover_Clustering struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_Failover_ClusteringEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_Failover_Clustering, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_Failover_Clustering{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_Failover_ClusteringEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_Failover_Clustering, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_Failover_Clustering{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
