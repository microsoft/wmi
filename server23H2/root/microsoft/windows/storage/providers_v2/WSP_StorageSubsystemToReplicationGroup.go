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

// WSP_StorageSubsystemToReplicationGroup struct
type WSP_StorageSubsystemToReplicationGroup struct {
	*MSFT_StorageSubSystemToReplicationGroup
}

func NewWSP_StorageSubsystemToReplicationGroupEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageSubsystemToReplicationGroup, err error) {
	tmp, err := NewMSFT_StorageSubSystemToReplicationGroupEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToReplicationGroup{
		MSFT_StorageSubSystemToReplicationGroup: tmp,
	}
	return
}

func NewWSP_StorageSubsystemToReplicationGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageSubsystemToReplicationGroup, err error) {
	tmp, err := NewMSFT_StorageSubSystemToReplicationGroupEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToReplicationGroup{
		MSFT_StorageSubSystemToReplicationGroup: tmp,
	}
	return
}
