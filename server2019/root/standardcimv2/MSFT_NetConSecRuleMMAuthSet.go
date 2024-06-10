// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetConSecRuleMMAuthSet struct
type MSFT_NetConSecRuleMMAuthSet struct {
	*MSFT_NetSARuleMMAuth
}

func NewMSFT_NetConSecRuleMMAuthSetEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetConSecRuleMMAuthSet, err error) {
	tmp, err := NewMSFT_NetSARuleMMAuthEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleMMAuthSet{
		MSFT_NetSARuleMMAuth: tmp,
	}
	return
}

func NewMSFT_NetConSecRuleMMAuthSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetConSecRuleMMAuthSet, err error) {
	tmp, err := NewMSFT_NetSARuleMMAuthEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleMMAuthSet{
		MSFT_NetSARuleMMAuth: tmp,
	}
	return
}
