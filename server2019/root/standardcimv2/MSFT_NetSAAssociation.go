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

// MSFT_NetSAAssociation struct
type MSFT_NetSAAssociation struct { 
	*CIM_Phase1SAUsedForPhase2
}

	func NewMSFT_NetSAAssociationEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetSAAssociation, err error) {tmp, err := NewCIM_Phase1SAUsedForPhase2Ex1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetSAAssociation {
	CIM_Phase1SAUsedForPhase2: tmp,
	}
	return
	}
	

	func NewMSFT_NetSAAssociationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetSAAssociation, err error) {tmp, err := NewCIM_Phase1SAUsedForPhase2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetSAAssociation {
	CIM_Phase1SAUsedForPhase2: tmp,
	}
	return
	}
	

