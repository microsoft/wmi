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

// SPACES_StorageNodeToPhysicalDisk struct
type SPACES_StorageNodeToPhysicalDisk struct {
	*MSFT_StorageNodeToPhysicalDisk
}

func NewSPACES_StorageNodeToPhysicalDiskEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageNodeToPhysicalDisk, err error) {
	tmp, err := NewMSFT_StorageNodeToPhysicalDiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageNodeToPhysicalDisk{
		MSFT_StorageNodeToPhysicalDisk: tmp,
	}
	return
}

func NewSPACES_StorageNodeToPhysicalDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageNodeToPhysicalDisk, err error) {
	tmp, err := NewMSFT_StorageNodeToPhysicalDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageNodeToPhysicalDisk{
		MSFT_StorageNodeToPhysicalDisk: tmp,
	}
	return
}
