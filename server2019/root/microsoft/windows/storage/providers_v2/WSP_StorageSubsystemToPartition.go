// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_StorageSubsystemToPartition struct
type WSP_StorageSubsystemToPartition struct {
	*MSFT_StorageSubSystemToPartition
}

func NewWSP_StorageSubsystemToPartitionEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageSubsystemToPartition, err error) {
	tmp, err := NewMSFT_StorageSubSystemToPartitionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToPartition{
		MSFT_StorageSubSystemToPartition: tmp,
	}
	return
}

func NewWSP_StorageSubsystemToPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageSubsystemToPartition, err error) {
	tmp, err := NewMSFT_StorageSubSystemToPartitionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToPartition{
		MSFT_StorageSubSystemToPartition: tmp,
	}
	return
}
