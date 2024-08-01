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

// WSP_FileShareAccessControlEntry struct
type WSP_FileShareAccessControlEntry struct {
	*MSFT_FileShareAccessControlEntry
}

func NewWSP_FileShareAccessControlEntryEx1(instance *cim.WmiInstance) (newInstance *WSP_FileShareAccessControlEntry, err error) {
	tmp, err := NewMSFT_FileShareAccessControlEntryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_FileShareAccessControlEntry{
		MSFT_FileShareAccessControlEntry: tmp,
	}
	return
}

func NewWSP_FileShareAccessControlEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_FileShareAccessControlEntry, err error) {
	tmp, err := NewMSFT_FileShareAccessControlEntryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_FileShareAccessControlEntry{
		MSFT_FileShareAccessControlEntry: tmp,
	}
	return
}
