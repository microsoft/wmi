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

// MSFT_NetRuleInProfile struct
type MSFT_NetRuleInProfile struct {
	*CIM_PolicySetComponent
}

func NewMSFT_NetRuleInProfileEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetRuleInProfile, err error) {
	tmp, err := NewCIM_PolicySetComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetRuleInProfile{
		CIM_PolicySetComponent: tmp,
	}
	return
}

func NewMSFT_NetRuleInProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetRuleInProfile, err error) {
	tmp, err := NewCIM_PolicySetComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetRuleInProfile{
		CIM_PolicySetComponent: tmp,
	}
	return
}
