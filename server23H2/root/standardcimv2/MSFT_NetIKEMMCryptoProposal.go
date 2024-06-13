// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetIKEMMCryptoProposal struct
type MSFT_NetIKEMMCryptoProposal struct {
	*MSFT_NetIKECryptoProposal
}

func NewMSFT_NetIKEMMCryptoProposalEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIKEMMCryptoProposal, err error) {
	tmp, err := NewMSFT_NetIKECryptoProposalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEMMCryptoProposal{
		MSFT_NetIKECryptoProposal: tmp,
	}
	return
}

func NewMSFT_NetIKEMMCryptoProposalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIKEMMCryptoProposal, err error) {
	tmp, err := NewMSFT_NetIKECryptoProposalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEMMCryptoProposal{
		MSFT_NetIKECryptoProposal: tmp,
	}
	return
}
