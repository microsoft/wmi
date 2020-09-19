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

// WSP_StorageHealthStatusChangeEvent struct
type WSP_StorageHealthStatusChangeEvent struct {
	*MSFT_StorageHealthStatusChangeEvent
}

func NewWSP_StorageHealthStatusChangeEventEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageHealthStatusChangeEvent, err error) {
	tmp, err := NewMSFT_StorageHealthStatusChangeEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageHealthStatusChangeEvent{
		MSFT_StorageHealthStatusChangeEvent: tmp,
	}
	return
}

func NewWSP_StorageHealthStatusChangeEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageHealthStatusChangeEvent, err error) {
	tmp, err := NewMSFT_StorageHealthStatusChangeEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageHealthStatusChangeEvent{
		MSFT_StorageHealthStatusChangeEvent: tmp,
	}
	return
}
