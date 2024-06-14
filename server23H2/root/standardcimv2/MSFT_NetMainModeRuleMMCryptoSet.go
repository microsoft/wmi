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

// MSFT_NetMainModeRuleMMCryptoSet struct
type MSFT_NetMainModeRuleMMCryptoSet struct {
	*MSFT_NetSARuleMMCrypto
}

func NewMSFT_NetMainModeRuleMMCryptoSetEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetMainModeRuleMMCryptoSet, err error) {
	tmp, err := NewMSFT_NetSARuleMMCryptoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetMainModeRuleMMCryptoSet{
		MSFT_NetSARuleMMCrypto: tmp,
	}
	return
}

func NewMSFT_NetMainModeRuleMMCryptoSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetMainModeRuleMMCryptoSet, err error) {
	tmp, err := NewMSFT_NetSARuleMMCryptoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetMainModeRuleMMCryptoSet{
		MSFT_NetSARuleMMCrypto: tmp,
	}
	return
}
