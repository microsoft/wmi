// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetFirewallRuleFilterByService struct
type MSFT_NetFirewallRuleFilterByService struct {
	*MSFT_NetFirewallRuleFilters
}

func NewMSFT_NetFirewallRuleFilterByServiceEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetFirewallRuleFilterByService, err error) {
	tmp, err := NewMSFT_NetFirewallRuleFiltersEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallRuleFilterByService{
		MSFT_NetFirewallRuleFilters: tmp,
	}
	return
}

func NewMSFT_NetFirewallRuleFilterByServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetFirewallRuleFilterByService, err error) {
	tmp, err := NewMSFT_NetFirewallRuleFiltersEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallRuleFilterByService{
		MSFT_NetFirewallRuleFilters: tmp,
	}
	return
}
