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

// MSFT_NetSARuleMMAuth struct
type MSFT_NetSARuleMMAuth struct { 
	*MSFT_NetSAActionInSARule
}

	func NewMSFT_NetSARuleMMAuthEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetSARuleMMAuth, err error) {tmp, err := NewMSFT_NetSAActionInSARuleEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetSARuleMMAuth {
	MSFT_NetSAActionInSARule: tmp,
	}
	return
	}
	

	func NewMSFT_NetSARuleMMAuthEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetSARuleMMAuth, err error) {tmp, err := NewMSFT_NetSAActionInSARuleEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetSARuleMMAuth {
	MSFT_NetSAActionInSARule: tmp,
	}
	return
	}
	

