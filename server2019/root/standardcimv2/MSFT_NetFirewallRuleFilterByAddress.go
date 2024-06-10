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

// MSFT_NetFirewallRuleFilterByAddress struct
type MSFT_NetFirewallRuleFilterByAddress struct {
	*MSFT_NetFirewallRuleFilters
}

func NewMSFT_NetFirewallRuleFilterByAddressEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetFirewallRuleFilterByAddress, err error) {
	tmp, err := NewMSFT_NetFirewallRuleFiltersEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallRuleFilterByAddress{
		MSFT_NetFirewallRuleFilters: tmp,
	}
	return
}

func NewMSFT_NetFirewallRuleFilterByAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetFirewallRuleFilterByAddress, err error) {
	tmp, err := NewMSFT_NetFirewallRuleFiltersEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallRuleFilterByAddress{
		MSFT_NetFirewallRuleFilters: tmp,
	}
	return
}
