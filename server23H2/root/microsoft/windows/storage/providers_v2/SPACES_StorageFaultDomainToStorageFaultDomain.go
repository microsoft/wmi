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

// SPACES_StorageFaultDomainToStorageFaultDomain struct
type SPACES_StorageFaultDomainToStorageFaultDomain struct {
	*MSFT_StorageFaultDomainToStorageFaultDomain
}

func NewSPACES_StorageFaultDomainToStorageFaultDomainEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageFaultDomainToStorageFaultDomain, err error) {
	tmp, err := NewMSFT_StorageFaultDomainToStorageFaultDomainEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageFaultDomainToStorageFaultDomain{
		MSFT_StorageFaultDomainToStorageFaultDomain: tmp,
	}
	return
}

func NewSPACES_StorageFaultDomainToStorageFaultDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageFaultDomainToStorageFaultDomain, err error) {
	tmp, err := NewMSFT_StorageFaultDomainToStorageFaultDomainEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageFaultDomainToStorageFaultDomain{
		MSFT_StorageFaultDomainToStorageFaultDomain: tmp,
	}
	return
}
