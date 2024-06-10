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

// MSFT_NetMainModeRuleFilterByAddress struct
type MSFT_NetMainModeRuleFilterByAddress struct {
	*MSFT_NetMainModeRuleFilters
}

func NewMSFT_NetMainModeRuleFilterByAddressEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetMainModeRuleFilterByAddress, err error) {
	tmp, err := NewMSFT_NetMainModeRuleFiltersEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetMainModeRuleFilterByAddress{
		MSFT_NetMainModeRuleFilters: tmp,
	}
	return
}

func NewMSFT_NetMainModeRuleFilterByAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetMainModeRuleFilterByAddress, err error) {
	tmp, err := NewMSFT_NetMainModeRuleFiltersEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetMainModeRuleFilterByAddress{
		MSFT_NetMainModeRuleFilters: tmp,
	}
	return
}
