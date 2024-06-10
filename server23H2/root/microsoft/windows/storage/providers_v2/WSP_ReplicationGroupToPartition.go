// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Storage.Providers_v2
//
// ////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_ReplicationGroupToPartition struct
type WSP_ReplicationGroupToPartition struct {
	*MSFT_ReplicationGroupToPartition
}

func NewWSP_ReplicationGroupToPartitionEx1(instance *cim.WmiInstance) (newInstance *WSP_ReplicationGroupToPartition, err error) {
	tmp, err := NewMSFT_ReplicationGroupToPartitionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationGroupToPartition{
		MSFT_ReplicationGroupToPartition: tmp,
	}
	return
}

func NewWSP_ReplicationGroupToPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_ReplicationGroupToPartition, err error) {
	tmp, err := NewMSFT_ReplicationGroupToPartitionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationGroupToPartition{
		MSFT_ReplicationGroupToPartition: tmp,
	}
	return
}
