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

// WSP_StoragePoolToVolume struct
type WSP_StoragePoolToVolume struct {
	*MSFT_StoragePoolToVolume
}

func NewWSP_StoragePoolToVolumeEx1(instance *cim.WmiInstance) (newInstance *WSP_StoragePoolToVolume, err error) {
	tmp, err := NewMSFT_StoragePoolToVolumeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StoragePoolToVolume{
		MSFT_StoragePoolToVolume: tmp,
	}
	return
}

func NewWSP_StoragePoolToVolumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StoragePoolToVolume, err error) {
	tmp, err := NewMSFT_StoragePoolToVolumeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StoragePoolToVolume{
		MSFT_StoragePoolToVolume: tmp,
	}
	return
}
