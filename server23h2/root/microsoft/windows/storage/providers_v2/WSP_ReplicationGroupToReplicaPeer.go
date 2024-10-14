// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_ReplicationGroupToReplicaPeer struct
type WSP_ReplicationGroupToReplicaPeer struct {
	*MSFT_ReplicationGroupToReplicaPeer
}

func NewWSP_ReplicationGroupToReplicaPeerEx1(instance *cim.WmiInstance) (newInstance *WSP_ReplicationGroupToReplicaPeer, err error) {
	tmp, err := NewMSFT_ReplicationGroupToReplicaPeerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationGroupToReplicaPeer{
		MSFT_ReplicationGroupToReplicaPeer: tmp,
	}
	return
}

func NewWSP_ReplicationGroupToReplicaPeerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_ReplicationGroupToReplicaPeer, err error) {
	tmp, err := NewMSFT_ReplicationGroupToReplicaPeerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationGroupToReplicaPeer{
		MSFT_ReplicationGroupToReplicaPeer: tmp,
	}
	return
}
