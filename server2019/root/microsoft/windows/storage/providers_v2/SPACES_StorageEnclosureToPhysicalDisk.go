// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_StorageEnclosureToPhysicalDisk struct
type SPACES_StorageEnclosureToPhysicalDisk struct {
	*MSFT_StorageEnclosureToPhysicalDisk
}

func NewSPACES_StorageEnclosureToPhysicalDiskEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageEnclosureToPhysicalDisk, err error) {
	tmp, err := NewMSFT_StorageEnclosureToPhysicalDiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageEnclosureToPhysicalDisk{
		MSFT_StorageEnclosureToPhysicalDisk: tmp,
	}
	return
}

func NewSPACES_StorageEnclosureToPhysicalDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageEnclosureToPhysicalDisk, err error) {
	tmp, err := NewMSFT_StorageEnclosureToPhysicalDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageEnclosureToPhysicalDisk{
		MSFT_StorageEnclosureToPhysicalDisk: tmp,
	}
	return
}
