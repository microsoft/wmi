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

// SPACES_StorageNodeToStorageEnclosure struct
type SPACES_StorageNodeToStorageEnclosure struct {
	*MSFT_StorageNodeToStorageEnclosure
}

func NewSPACES_StorageNodeToStorageEnclosureEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageNodeToStorageEnclosure, err error) {
	tmp, err := NewMSFT_StorageNodeToStorageEnclosureEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageNodeToStorageEnclosure{
		MSFT_StorageNodeToStorageEnclosure: tmp,
	}
	return
}

func NewSPACES_StorageNodeToStorageEnclosureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageNodeToStorageEnclosure, err error) {
	tmp, err := NewMSFT_StorageNodeToStorageEnclosureEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageNodeToStorageEnclosure{
		MSFT_StorageNodeToStorageEnclosure: tmp,
	}
	return
}
