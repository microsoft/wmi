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

// MSFT_NetConSecRuleFilters struct
type MSFT_NetConSecRuleFilters struct {
	*MSFT_NetPolicyRuleFilters
}

func NewMSFT_NetConSecRuleFiltersEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetConSecRuleFilters, err error) {
	tmp, err := NewMSFT_NetPolicyRuleFiltersEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleFilters{
		MSFT_NetPolicyRuleFilters: tmp,
	}
	return
}

func NewMSFT_NetConSecRuleFiltersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetConSecRuleFilters, err error) {
	tmp, err := NewMSFT_NetPolicyRuleFiltersEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleFilters{
		MSFT_NetPolicyRuleFilters: tmp,
	}
	return
}
