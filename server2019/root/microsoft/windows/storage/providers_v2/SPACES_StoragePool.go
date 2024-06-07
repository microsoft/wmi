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

// SPACES_StoragePool struct
type SPACES_StoragePool struct { 
	*MSFT_StoragePool
}

	func NewSPACES_StoragePoolEx1(instance *cim.WmiInstance) (newInstance *SPACES_StoragePool, err error) {tmp, err := NewMSFT_StoragePoolEx1(instance)
		
	if err != nil { return }
	newInstance = &SPACES_StoragePool {
	MSFT_StoragePool: tmp,
	}
	return
	}
	

	func NewSPACES_StoragePoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SPACES_StoragePool, err error) {tmp, err := NewMSFT_StoragePoolEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SPACES_StoragePool {
	MSFT_StoragePool: tmp,
	}
	return
	}
	

