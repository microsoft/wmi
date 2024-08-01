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

// MSFT_NetFirewallRuleFilters struct
type MSFT_NetFirewallRuleFilters struct {
	*MSFT_NetPolicyRuleFilters
}

func NewMSFT_NetFirewallRuleFiltersEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetFirewallRuleFilters, err error) {
	tmp, err := NewMSFT_NetPolicyRuleFiltersEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallRuleFilters{
		MSFT_NetPolicyRuleFilters: tmp,
	}
	return
}

func NewMSFT_NetFirewallRuleFiltersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetFirewallRuleFilters, err error) {
	tmp, err := NewMSFT_NetPolicyRuleFiltersEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallRuleFilters{
		MSFT_NetPolicyRuleFilters: tmp,
	}
	return
}
