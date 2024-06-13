// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_PartitionToReplicaPeer struct
type WSP_PartitionToReplicaPeer struct {
	*MSFT_PartitionToReplicaPeer
}

func NewWSP_PartitionToReplicaPeerEx1(instance *cim.WmiInstance) (newInstance *WSP_PartitionToReplicaPeer, err error) {
	tmp, err := NewMSFT_PartitionToReplicaPeerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_PartitionToReplicaPeer{
		MSFT_PartitionToReplicaPeer: tmp,
	}
	return
}

func NewWSP_PartitionToReplicaPeerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_PartitionToReplicaPeer, err error) {
	tmp, err := NewMSFT_PartitionToReplicaPeerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_PartitionToReplicaPeer{
		MSFT_PartitionToReplicaPeer: tmp,
	}
	return
}
