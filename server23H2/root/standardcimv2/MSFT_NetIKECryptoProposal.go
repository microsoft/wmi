// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetIKECryptoProposal struct
type MSFT_NetIKECryptoProposal struct {
	*CIM_IKEProposal
}

func NewMSFT_NetIKECryptoProposalEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIKECryptoProposal, err error) {
	tmp, err := NewCIM_IKEProposalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKECryptoProposal{
		CIM_IKEProposal: tmp,
	}
	return
}

func NewMSFT_NetIKECryptoProposalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIKECryptoProposal, err error) {
	tmp, err := NewCIM_IKEProposalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKECryptoProposal{
		CIM_IKEProposal: tmp,
	}
	return
}
