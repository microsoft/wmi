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

// WSP_StorageArrivalEvent struct
type WSP_StorageArrivalEvent struct {
	*MSFT_StorageArrivalEvent
}

func NewWSP_StorageArrivalEventEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageArrivalEvent, err error) {
	tmp, err := NewMSFT_StorageArrivalEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageArrivalEvent{
		MSFT_StorageArrivalEvent: tmp,
	}
	return
}

func NewWSP_StorageArrivalEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageArrivalEvent, err error) {
	tmp, err := NewMSFT_StorageArrivalEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageArrivalEvent{
		MSFT_StorageArrivalEvent: tmp,
	}
	return
}
