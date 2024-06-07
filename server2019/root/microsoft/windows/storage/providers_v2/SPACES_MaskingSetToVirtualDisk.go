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

// SPACES_MaskingSetToVirtualDisk struct
type SPACES_MaskingSetToVirtualDisk struct { 
	*MSFT_MaskingSetToVirtualDisk
}

	func NewSPACES_MaskingSetToVirtualDiskEx1(instance *cim.WmiInstance) (newInstance *SPACES_MaskingSetToVirtualDisk, err error) {tmp, err := NewMSFT_MaskingSetToVirtualDiskEx1(instance)
		
	if err != nil { return }
	newInstance = &SPACES_MaskingSetToVirtualDisk {
	MSFT_MaskingSetToVirtualDisk: tmp,
	}
	return
	}
	

	func NewSPACES_MaskingSetToVirtualDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SPACES_MaskingSetToVirtualDisk, err error) {tmp, err := NewMSFT_MaskingSetToVirtualDiskEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SPACES_MaskingSetToVirtualDisk {
	MSFT_MaskingSetToVirtualDisk: tmp,
	}
	return
	}
	

