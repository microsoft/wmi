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

// SPACES_StorageProviderToStorageSubsystem struct
type SPACES_StorageProviderToStorageSubsystem struct {
	*MSFT_StorageProviderToStorageSubSystem
}

func NewSPACES_StorageProviderToStorageSubsystemEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageProviderToStorageSubsystem, err error) {
	tmp, err := NewMSFT_StorageProviderToStorageSubSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageProviderToStorageSubsystem{
		MSFT_StorageProviderToStorageSubSystem: tmp,
	}
	return
}

func NewSPACES_StorageProviderToStorageSubsystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageProviderToStorageSubsystem, err error) {
	tmp, err := NewMSFT_StorageProviderToStorageSubSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageProviderToStorageSubsystem{
		MSFT_StorageProviderToStorageSubSystem: tmp,
	}
	return
}
