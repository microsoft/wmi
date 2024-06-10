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

// WSP_StorageSubsystemToStorageHealth struct
type WSP_StorageSubsystemToStorageHealth struct {
	*MSFT_StorageSubSystemToStorageHealth
}

func NewWSP_StorageSubsystemToStorageHealthEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageSubsystemToStorageHealth, err error) {
	tmp, err := NewMSFT_StorageSubSystemToStorageHealthEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToStorageHealth{
		MSFT_StorageSubSystemToStorageHealth: tmp,
	}
	return
}

func NewWSP_StorageSubsystemToStorageHealthEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageSubsystemToStorageHealth, err error) {
	tmp, err := NewMSFT_StorageSubSystemToStorageHealthEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToStorageHealth{
		MSFT_StorageSubSystemToStorageHealth: tmp,
	}
	return
}
