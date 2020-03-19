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

// WSP_StorageModificationEvent struct
type WSP_StorageModificationEvent struct {
	*MSFT_StorageModificationEvent
}

func NewWSP_StorageModificationEventEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageModificationEvent, err error) {
	tmp, err := NewMSFT_StorageModificationEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageModificationEvent{
		MSFT_StorageModificationEvent: tmp,
	}
	return
}

func NewWSP_StorageModificationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageModificationEvent, err error) {
	tmp, err := NewMSFT_StorageModificationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageModificationEvent{
		MSFT_StorageModificationEvent: tmp,
	}
	return
}
