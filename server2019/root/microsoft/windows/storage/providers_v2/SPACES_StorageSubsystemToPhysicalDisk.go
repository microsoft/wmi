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

// SPACES_StorageSubsystemToPhysicalDisk struct
type SPACES_StorageSubsystemToPhysicalDisk struct { 
	*MSFT_StorageSubSystemToPhysicalDisk
}

	func NewSPACES_StorageSubsystemToPhysicalDiskEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageSubsystemToPhysicalDisk, err error) {tmp, err := NewMSFT_StorageSubSystemToPhysicalDiskEx1(instance)
		
	if err != nil { return }
	newInstance = &SPACES_StorageSubsystemToPhysicalDisk {
	MSFT_StorageSubSystemToPhysicalDisk: tmp,
	}
	return
	}
	

	func NewSPACES_StorageSubsystemToPhysicalDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SPACES_StorageSubsystemToPhysicalDisk, err error) {tmp, err := NewMSFT_StorageSubSystemToPhysicalDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SPACES_StorageSubsystemToPhysicalDisk {
	MSFT_StorageSubSystemToPhysicalDisk: tmp,
	}
	return
	}
	

