// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetConSecRuleFilterByAddress struct
type MSFT_NetConSecRuleFilterByAddress struct {
	*MSFT_NetConSecRuleFilters
}

func NewMSFT_NetConSecRuleFilterByAddressEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetConSecRuleFilterByAddress, err error) {
	tmp, err := NewMSFT_NetConSecRuleFiltersEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleFilterByAddress{
		MSFT_NetConSecRuleFilters: tmp,
	}
	return
}

func NewMSFT_NetConSecRuleFilterByAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetConSecRuleFilterByAddress, err error) {
	tmp, err := NewMSFT_NetConSecRuleFiltersEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleFilterByAddress{
		MSFT_NetConSecRuleFilters: tmp,
	}
	return
}
