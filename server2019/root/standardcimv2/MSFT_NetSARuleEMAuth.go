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

// MSFT_NetSARuleEMAuth struct
type MSFT_NetSARuleEMAuth struct { 
	*MSFT_NetSAActionInSARule
}

	func NewMSFT_NetSARuleEMAuthEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetSARuleEMAuth, err error) {tmp, err := NewMSFT_NetSAActionInSARuleEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetSARuleEMAuth {
	MSFT_NetSAActionInSARule: tmp,
	}
	return
	}
	

	func NewMSFT_NetSARuleEMAuthEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetSARuleEMAuth, err error) {tmp, err := NewMSFT_NetSAActionInSARuleEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetSARuleEMAuth {
	MSFT_NetSAActionInSARule: tmp,
	}
	return
	}
	

