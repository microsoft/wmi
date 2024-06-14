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

// SPACES_VirtualDiskToStorageTier struct
type SPACES_VirtualDiskToStorageTier struct {
	*MSFT_VirtualDiskToStorageTier
}

func NewSPACES_VirtualDiskToStorageTierEx1(instance *cim.WmiInstance) (newInstance *SPACES_VirtualDiskToStorageTier, err error) {
	tmp, err := NewMSFT_VirtualDiskToStorageTierEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_VirtualDiskToStorageTier{
		MSFT_VirtualDiskToStorageTier: tmp,
	}
	return
}

func NewSPACES_VirtualDiskToStorageTierEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_VirtualDiskToStorageTier, err error) {
	tmp, err := NewMSFT_VirtualDiskToStorageTierEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_VirtualDiskToStorageTier{
		MSFT_VirtualDiskToStorageTier: tmp,
	}
	return
}
