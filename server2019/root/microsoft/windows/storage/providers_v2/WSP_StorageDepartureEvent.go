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

// WSP_StorageDepartureEvent struct
type WSP_StorageDepartureEvent struct {
	*MSFT_StorageDepartureEvent
}

func NewWSP_StorageDepartureEventEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageDepartureEvent, err error) {
	tmp, err := NewMSFT_StorageDepartureEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageDepartureEvent{
		MSFT_StorageDepartureEvent: tmp,
	}
	return
}

func NewWSP_StorageDepartureEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageDepartureEvent, err error) {
	tmp, err := NewMSFT_StorageDepartureEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageDepartureEvent{
		MSFT_StorageDepartureEvent: tmp,
	}
	return
}
