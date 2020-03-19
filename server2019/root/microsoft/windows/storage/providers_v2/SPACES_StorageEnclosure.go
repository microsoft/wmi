// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_StorageEnclosure struct
type SPACES_StorageEnclosure struct {
	*MSFT_StorageEnclosure
}

func NewSPACES_StorageEnclosureEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageEnclosure, err error) {
	tmp, err := NewMSFT_StorageEnclosureEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageEnclosure{
		MSFT_StorageEnclosure: tmp,
	}
	return
}

func NewSPACES_StorageEnclosureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageEnclosure, err error) {
	tmp, err := NewMSFT_StorageEnclosureEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageEnclosure{
		MSFT_StorageEnclosure: tmp,
	}
	return
}
