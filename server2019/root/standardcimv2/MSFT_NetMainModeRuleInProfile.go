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

// MSFT_NetMainModeRuleInProfile struct
type MSFT_NetMainModeRuleInProfile struct {
	*MSFT_NetRuleInProfile
}

func NewMSFT_NetMainModeRuleInProfileEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetMainModeRuleInProfile, err error) {
	tmp, err := NewMSFT_NetRuleInProfileEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetMainModeRuleInProfile{
		MSFT_NetRuleInProfile: tmp,
	}
	return
}

func NewMSFT_NetMainModeRuleInProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetMainModeRuleInProfile, err error) {
	tmp, err := NewMSFT_NetRuleInProfileEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetMainModeRuleInProfile{
		MSFT_NetRuleInProfile: tmp,
	}
	return
}
