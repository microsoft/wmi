// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetEventProvider struct
type MSFT_NetEventProvider struct { 
	*MSFT_NetEventProviderBase
}

	func NewMSFT_NetEventProviderEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventProvider, err error) {tmp, err := NewMSFT_NetEventProviderBaseEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetEventProvider {
	MSFT_NetEventProviderBase: tmp,
	}
	return
	}
	

	func NewMSFT_NetEventProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetEventProvider, err error) {tmp, err := NewMSFT_NetEventProviderBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetEventProvider {
	MSFT_NetEventProviderBase: tmp,
	}
	return
	}
	

