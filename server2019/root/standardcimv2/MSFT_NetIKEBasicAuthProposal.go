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

// MSFT_NetIKEBasicAuthProposal struct
type MSFT_NetIKEBasicAuthProposal struct { 
	*MSFT_NetIKEAuthProposal
}

	func NewMSFT_NetIKEBasicAuthProposalEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIKEBasicAuthProposal, err error) {tmp, err := NewMSFT_NetIKEAuthProposalEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetIKEBasicAuthProposal {
	MSFT_NetIKEAuthProposal: tmp,
	}
	return
	}
	

	func NewMSFT_NetIKEBasicAuthProposalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetIKEBasicAuthProposal, err error) {tmp, err := NewMSFT_NetIKEAuthProposalEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetIKEBasicAuthProposal {
	MSFT_NetIKEAuthProposal: tmp,
	}
	return
	}
	

