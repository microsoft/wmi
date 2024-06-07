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

// SPACES_VirtualDiskToPhysicalDisk struct
type SPACES_VirtualDiskToPhysicalDisk struct { 
	*MSFT_VirtualDiskToPhysicalDisk
}

	func NewSPACES_VirtualDiskToPhysicalDiskEx1(instance *cim.WmiInstance) (newInstance *SPACES_VirtualDiskToPhysicalDisk, err error) {tmp, err := NewMSFT_VirtualDiskToPhysicalDiskEx1(instance)
		
	if err != nil { return }
	newInstance = &SPACES_VirtualDiskToPhysicalDisk {
	MSFT_VirtualDiskToPhysicalDisk: tmp,
	}
	return
	}
	

	func NewSPACES_VirtualDiskToPhysicalDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SPACES_VirtualDiskToPhysicalDisk, err error) {tmp, err := NewMSFT_VirtualDiskToPhysicalDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SPACES_VirtualDiskToPhysicalDisk {
	MSFT_VirtualDiskToPhysicalDisk: tmp,
	}
	return
	}
	

