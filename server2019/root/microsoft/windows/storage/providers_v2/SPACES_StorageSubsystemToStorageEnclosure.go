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

// SPACES_StorageSubsystemToStorageEnclosure struct
type SPACES_StorageSubsystemToStorageEnclosure struct {
	*MSFT_StorageSubSystemToStorageEnclosure
}

func NewSPACES_StorageSubsystemToStorageEnclosureEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageSubsystemToStorageEnclosure, err error) {
	tmp, err := NewMSFT_StorageSubSystemToStorageEnclosureEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystemToStorageEnclosure{
		MSFT_StorageSubSystemToStorageEnclosure: tmp,
	}
	return
}

func NewSPACES_StorageSubsystemToStorageEnclosureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageSubsystemToStorageEnclosure, err error) {
	tmp, err := NewMSFT_StorageSubSystemToStorageEnclosureEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystemToStorageEnclosure{
		MSFT_StorageSubSystemToStorageEnclosure: tmp,
	}
	return
}
