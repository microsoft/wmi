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

// SPACES_StorageFaultEvent struct
type SPACES_StorageFaultEvent struct { 
	*MSFT_StorageFaultEvent
}

	func NewSPACES_StorageFaultEventEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageFaultEvent, err error) {tmp, err := NewMSFT_StorageFaultEventEx1(instance)
		
	if err != nil { return }
	newInstance = &SPACES_StorageFaultEvent {
	MSFT_StorageFaultEvent: tmp,
	}
	return
	}
	

	func NewSPACES_StorageFaultEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SPACES_StorageFaultEvent, err error) {tmp, err := NewMSFT_StorageFaultEventEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SPACES_StorageFaultEvent {
	MSFT_StorageFaultEvent: tmp,
	}
	return
	}
	

