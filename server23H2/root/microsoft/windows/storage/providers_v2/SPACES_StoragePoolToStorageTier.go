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

// SPACES_StoragePoolToStorageTier struct
type SPACES_StoragePoolToStorageTier struct {
	*MSFT_StoragePoolToStorageTier
}

func NewSPACES_StoragePoolToStorageTierEx1(instance *cim.WmiInstance) (newInstance *SPACES_StoragePoolToStorageTier, err error) {
	tmp, err := NewMSFT_StoragePoolToStorageTierEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StoragePoolToStorageTier{
		MSFT_StoragePoolToStorageTier: tmp,
	}
	return
}

func NewSPACES_StoragePoolToStorageTierEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StoragePoolToStorageTier, err error) {
	tmp, err := NewMSFT_StoragePoolToStorageTierEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StoragePoolToStorageTier{
		MSFT_StoragePoolToStorageTier: tmp,
	}
	return
}
