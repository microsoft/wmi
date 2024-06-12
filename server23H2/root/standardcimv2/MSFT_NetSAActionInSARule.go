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

// MSFT_NetSAActionInSARule struct
type MSFT_NetSAActionInSARule struct {
	*CIM_PolicyActionInPolicyRule
}

func NewMSFT_NetSAActionInSARuleEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetSAActionInSARule, err error) {
	tmp, err := NewCIM_PolicyActionInPolicyRuleEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSAActionInSARule{
		CIM_PolicyActionInPolicyRule: tmp,
	}
	return
}

func NewMSFT_NetSAActionInSARuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetSAActionInSARule, err error) {
	tmp, err := NewCIM_PolicyActionInPolicyRuleEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSAActionInSARule{
		CIM_PolicyActionInPolicyRule: tmp,
	}
	return
}
