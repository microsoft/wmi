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

// SPACES_StorageSubsystemToInitiatorId struct
type SPACES_StorageSubsystemToInitiatorId struct { 
	*MSFT_StorageSubSystemToInitiatorId
}

	func NewSPACES_StorageSubsystemToInitiatorIdEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageSubsystemToInitiatorId, err error) {tmp, err := NewMSFT_StorageSubSystemToInitiatorIdEx1(instance)
		
	if err != nil { return }
	newInstance = &SPACES_StorageSubsystemToInitiatorId {
	MSFT_StorageSubSystemToInitiatorId: tmp,
	}
	return
	}
	

	func NewSPACES_StorageSubsystemToInitiatorIdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SPACES_StorageSubsystemToInitiatorId, err error) {tmp, err := NewMSFT_StorageSubSystemToInitiatorIdEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SPACES_StorageSubsystemToInitiatorId {
	MSFT_StorageSubSystemToInitiatorId: tmp,
	}
	return
	}
	

