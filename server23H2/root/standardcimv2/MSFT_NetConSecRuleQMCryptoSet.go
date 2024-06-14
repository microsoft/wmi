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

// MSFT_NetConSecRuleQMCryptoSet struct
type MSFT_NetConSecRuleQMCryptoSet struct {
	*MSFT_NetSARuleQMCrypto
}

func NewMSFT_NetConSecRuleQMCryptoSetEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetConSecRuleQMCryptoSet, err error) {
	tmp, err := NewMSFT_NetSARuleQMCryptoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleQMCryptoSet{
		MSFT_NetSARuleQMCrypto: tmp,
	}
	return
}

func NewMSFT_NetConSecRuleQMCryptoSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetConSecRuleQMCryptoSet, err error) {
	tmp, err := NewMSFT_NetSARuleQMCryptoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleQMCryptoSet{
		MSFT_NetSARuleQMCrypto: tmp,
	}
	return
}
