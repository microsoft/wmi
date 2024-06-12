// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetSARuleMMCrypto struct
type MSFT_NetSARuleMMCrypto struct {
	*MSFT_NetSAActionInSARule
}

func NewMSFT_NetSARuleMMCryptoEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetSARuleMMCrypto, err error) {
	tmp, err := NewMSFT_NetSAActionInSARuleEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSARuleMMCrypto{
		MSFT_NetSAActionInSARule: tmp,
	}
	return
}

func NewMSFT_NetSARuleMMCryptoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetSARuleMMCrypto, err error) {
	tmp, err := NewMSFT_NetSAActionInSARuleEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSARuleMMCrypto{
		MSFT_NetSAActionInSARule: tmp,
	}
	return
}
