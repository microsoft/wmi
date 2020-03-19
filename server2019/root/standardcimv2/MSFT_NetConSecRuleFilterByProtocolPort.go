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

// MSFT_NetConSecRuleFilterByProtocolPort struct
type MSFT_NetConSecRuleFilterByProtocolPort struct {
	*MSFT_NetConSecRuleFilters
}

func NewMSFT_NetConSecRuleFilterByProtocolPortEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetConSecRuleFilterByProtocolPort, err error) {
	tmp, err := NewMSFT_NetConSecRuleFiltersEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleFilterByProtocolPort{
		MSFT_NetConSecRuleFilters: tmp,
	}
	return
}

func NewMSFT_NetConSecRuleFilterByProtocolPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetConSecRuleFilterByProtocolPort, err error) {
	tmp, err := NewMSFT_NetConSecRuleFiltersEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleFilterByProtocolPort{
		MSFT_NetConSecRuleFilters: tmp,
	}
	return
}
