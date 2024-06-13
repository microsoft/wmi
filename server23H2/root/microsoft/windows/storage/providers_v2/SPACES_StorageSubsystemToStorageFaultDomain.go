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

// SPACES_StorageSubsystemToStorageFaultDomain struct
type SPACES_StorageSubsystemToStorageFaultDomain struct {
	*MSFT_StorageSubSystemToStorageFaultDomain
}

func NewSPACES_StorageSubsystemToStorageFaultDomainEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageSubsystemToStorageFaultDomain, err error) {
	tmp, err := NewMSFT_StorageSubSystemToStorageFaultDomainEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystemToStorageFaultDomain{
		MSFT_StorageSubSystemToStorageFaultDomain: tmp,
	}
	return
}

func NewSPACES_StorageSubsystemToStorageFaultDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageSubsystemToStorageFaultDomain, err error) {
	tmp, err := NewMSFT_StorageSubSystemToStorageFaultDomainEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageSubsystemToStorageFaultDomain{
		MSFT_StorageSubSystemToStorageFaultDomain: tmp,
	}
	return
}
