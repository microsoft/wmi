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

// SPACES_StorageSubsystemToVirtualDisk struct
type SPACES_StorageSubsystemToVirtualDisk struct {
	*MSFT_StorageSubSystemToVirtualDisk
}

func NewSPACES_StorageSubsystemToVirtualDiskEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageSubsystemToVirtualDisk, err error) {
	tmp, err := NewMSFT_StorageSubSystemToVirtualDiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystemToVirtualDisk{
		MSFT_StorageSubSystemToVirtualDisk: tmp,
	}
	return
}

func NewSPACES_StorageSubsystemToVirtualDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageSubsystemToVirtualDisk, err error) {
	tmp, err := NewMSFT_StorageSubSystemToVirtualDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystemToVirtualDisk{
		MSFT_StorageSubSystemToVirtualDisk: tmp,
	}
	return
}
