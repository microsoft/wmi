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

// WSP_ReplicationGroupStorageModificationEvent struct
type WSP_ReplicationGroupStorageModificationEvent struct { 
	*MSFT_StorageModificationEvent
}

	func NewWSP_ReplicationGroupStorageModificationEventEx1(instance *cim.WmiInstance) (newInstance *WSP_ReplicationGroupStorageModificationEvent, err error) {tmp, err := NewMSFT_StorageModificationEventEx1(instance)
		
	if err != nil { return }
	newInstance = &WSP_ReplicationGroupStorageModificationEvent {
	MSFT_StorageModificationEvent: tmp,
	}
	return
	}
	

	func NewWSP_ReplicationGroupStorageModificationEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WSP_ReplicationGroupStorageModificationEvent, err error) {tmp, err := NewMSFT_StorageModificationEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WSP_ReplicationGroupStorageModificationEvent {
	MSFT_StorageModificationEvent: tmp,
	}
	return
	}
	

