// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Storage.Providers_v2
//
// ////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_StorageAlertEvent struct
type WSP_StorageAlertEvent struct {
	*MSFT_StorageAlertEvent
}

func NewWSP_StorageAlertEventEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageAlertEvent, err error) {
	tmp, err := NewMSFT_StorageAlertEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageAlertEvent{
		MSFT_StorageAlertEvent: tmp,
	}
	return
}

func NewWSP_StorageAlertEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageAlertEvent, err error) {
	tmp, err := NewMSFT_StorageAlertEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageAlertEvent{
		MSFT_StorageAlertEvent: tmp,
	}
	return
}
