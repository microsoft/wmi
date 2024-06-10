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

// SPACES_StorageNodeToStoragePool struct
type SPACES_StorageNodeToStoragePool struct {
	*MSFT_StorageNodeToStoragePool
}

func NewSPACES_StorageNodeToStoragePoolEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageNodeToStoragePool, err error) {
	tmp, err := NewMSFT_StorageNodeToStoragePoolEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageNodeToStoragePool{
		MSFT_StorageNodeToStoragePool: tmp,
	}
	return
}

func NewSPACES_StorageNodeToStoragePoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageNodeToStoragePool, err error) {
	tmp, err := NewMSFT_StorageNodeToStoragePoolEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageNodeToStoragePool{
		MSFT_StorageNodeToStoragePool: tmp,
	}
	return
}
