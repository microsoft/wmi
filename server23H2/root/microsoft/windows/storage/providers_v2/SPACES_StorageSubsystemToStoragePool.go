// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_StorageSubsystemToStoragePool struct
type SPACES_StorageSubsystemToStoragePool struct {
	*MSFT_StorageSubSystemToStoragePool
}

func NewSPACES_StorageSubsystemToStoragePoolEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageSubsystemToStoragePool, err error) {
	tmp, err := NewMSFT_StorageSubSystemToStoragePoolEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystemToStoragePool{
		MSFT_StorageSubSystemToStoragePool: tmp,
	}
	return
}

func NewSPACES_StorageSubsystemToStoragePoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageSubsystemToStoragePool, err error) {
	tmp, err := NewMSFT_StorageSubSystemToStoragePoolEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystemToStoragePool{
		MSFT_StorageSubSystemToStoragePool: tmp,
	}
	return
}
