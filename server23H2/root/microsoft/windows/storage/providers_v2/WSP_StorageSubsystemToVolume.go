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

// WSP_StorageSubsystemToVolume struct
type WSP_StorageSubsystemToVolume struct {
	*MSFT_StorageSubSystemToVolume
}

func NewWSP_StorageSubsystemToVolumeEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageSubsystemToVolume, err error) {
	tmp, err := NewMSFT_StorageSubSystemToVolumeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToVolume{
		MSFT_StorageSubSystemToVolume: tmp,
	}
	return
}

func NewWSP_StorageSubsystemToVolumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageSubsystemToVolume, err error) {
	tmp, err := NewMSFT_StorageSubSystemToVolumeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageSubsystemToVolume{
		MSFT_StorageSubSystemToVolume: tmp,
	}
	return
}
