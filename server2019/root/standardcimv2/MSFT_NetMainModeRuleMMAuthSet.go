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

// MSFT_NetMainModeRuleMMAuthSet struct
type MSFT_NetMainModeRuleMMAuthSet struct { 
	*MSFT_NetSARuleMMAuth
}

	func NewMSFT_NetMainModeRuleMMAuthSetEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetMainModeRuleMMAuthSet, err error) {tmp, err := NewMSFT_NetSARuleMMAuthEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetMainModeRuleMMAuthSet {
	MSFT_NetSARuleMMAuth: tmp,
	}
	return
	}
	

	func NewMSFT_NetMainModeRuleMMAuthSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetMainModeRuleMMAuthSet, err error) {tmp, err := NewMSFT_NetSARuleMMAuthEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetMainModeRuleMMAuthSet {
	MSFT_NetSARuleMMAuth: tmp,
	}
	return
	}
	

