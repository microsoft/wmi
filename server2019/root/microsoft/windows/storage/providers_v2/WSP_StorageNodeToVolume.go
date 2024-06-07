// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_StorageNodeToVolume struct
type WSP_StorageNodeToVolume struct { 
	*MSFT_StorageNodeToVolume
}

	func NewWSP_StorageNodeToVolumeEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageNodeToVolume, err error) {tmp, err := NewMSFT_StorageNodeToVolumeEx1(instance)
		
	if err != nil { return }
	newInstance = &WSP_StorageNodeToVolume {
	MSFT_StorageNodeToVolume: tmp,
	}
	return
	}
	

	func NewWSP_StorageNodeToVolumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WSP_StorageNodeToVolume, err error) {tmp, err := NewMSFT_StorageNodeToVolumeEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WSP_StorageNodeToVolume {
	MSFT_StorageNodeToVolume: tmp,
	}
	return
	}
	

