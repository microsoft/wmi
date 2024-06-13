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

// WSP_ReplicationGroupStorageDepartureEvent struct
type WSP_ReplicationGroupStorageDepartureEvent struct {
	*MSFT_StorageDepartureEvent
}

func NewWSP_ReplicationGroupStorageDepartureEventEx1(instance *cim.WmiInstance) (newInstance *WSP_ReplicationGroupStorageDepartureEvent, err error) {
	tmp, err := NewMSFT_StorageDepartureEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationGroupStorageDepartureEvent{
		MSFT_StorageDepartureEvent: tmp,
	}
	return
}

func NewWSP_ReplicationGroupStorageDepartureEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_ReplicationGroupStorageDepartureEvent, err error) {
	tmp, err := NewMSFT_StorageDepartureEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_ReplicationGroupStorageDepartureEvent{
		MSFT_StorageDepartureEvent: tmp,
	}
	return
}
