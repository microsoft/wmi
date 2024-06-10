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

// WSP_StorageHealth struct
type WSP_StorageHealth struct {
	*MSFT_StorageHealth
}

func NewWSP_StorageHealthEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageHealth, err error) {
	tmp, err := NewMSFT_StorageHealthEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageHealth{
		MSFT_StorageHealth: tmp,
	}
	return
}

func NewWSP_StorageHealthEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageHealth, err error) {
	tmp, err := NewMSFT_StorageHealthEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageHealth{
		MSFT_StorageHealth: tmp,
	}
	return
}
