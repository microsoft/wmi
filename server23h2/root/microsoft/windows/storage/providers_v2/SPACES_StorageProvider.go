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

// SPACES_StorageProvider struct
type SPACES_StorageProvider struct {
	*MSFT_StorageProvider
}

func NewSPACES_StorageProviderEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageProvider, err error) {
	tmp, err := NewMSFT_StorageProviderEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageProvider{
		MSFT_StorageProvider: tmp,
	}
	return
}

func NewSPACES_StorageProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageProvider, err error) {
	tmp, err := NewMSFT_StorageProviderEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageProvider{
		MSFT_StorageProvider: tmp,
	}
	return
}
