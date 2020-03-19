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

// WSP_FsStorageModificationEvent struct
type WSP_FsStorageModificationEvent struct {
	*MSFT_StorageModificationEvent
}

func NewWSP_FsStorageModificationEventEx1(instance *cim.WmiInstance) (newInstance *WSP_FsStorageModificationEvent, err error) {
	tmp, err := NewMSFT_StorageModificationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_FsStorageModificationEvent{
		MSFT_StorageModificationEvent: tmp,
	}
	return
}

func NewWSP_FsStorageModificationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_FsStorageModificationEvent, err error) {
	tmp, err := NewMSFT_StorageModificationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_FsStorageModificationEvent{
		MSFT_StorageModificationEvent: tmp,
	}
	return
}
