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

// MSFT_NetConSecRuleFilterByInterface struct
type MSFT_NetConSecRuleFilterByInterface struct {
	*MSFT_NetConSecRuleFilters
}

func NewMSFT_NetConSecRuleFilterByInterfaceEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetConSecRuleFilterByInterface, err error) {
	tmp, err := NewMSFT_NetConSecRuleFiltersEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleFilterByInterface{
		MSFT_NetConSecRuleFilters: tmp,
	}
	return
}

func NewMSFT_NetConSecRuleFilterByInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetConSecRuleFilterByInterface, err error) {
	tmp, err := NewMSFT_NetConSecRuleFiltersEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleFilterByInterface{
		MSFT_NetConSecRuleFilters: tmp,
	}
	return
}
