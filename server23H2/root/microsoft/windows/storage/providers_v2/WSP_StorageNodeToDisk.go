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

// WSP_StorageNodeToDisk struct
type WSP_StorageNodeToDisk struct {
	*MSFT_StorageNodeToDisk
}

func NewWSP_StorageNodeToDiskEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageNodeToDisk, err error) {
	tmp, err := NewMSFT_StorageNodeToDiskEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageNodeToDisk{
		MSFT_StorageNodeToDisk: tmp,
	}
	return
}

func NewWSP_StorageNodeToDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageNodeToDisk, err error) {
	tmp, err := NewMSFT_StorageNodeToDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageNodeToDisk{
		MSFT_StorageNodeToDisk: tmp,
	}
	return
}
