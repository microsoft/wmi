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

// SPACES_StorageRack struct
type SPACES_StorageRack struct {
	*MSFT_StorageRack
}

func NewSPACES_StorageRackEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageRack, err error) {
	tmp, err := NewMSFT_StorageRackEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageRack{
		MSFT_StorageRack: tmp,
	}
	return
}

func NewSPACES_StorageRackEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageRack, err error) {
	tmp, err := NewMSFT_StorageRackEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageRack{
		MSFT_StorageRack: tmp,
	}
	return
}
