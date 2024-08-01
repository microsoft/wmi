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

// MSFT_NetSwitchTeamMember struct
type MSFT_NetSwitchTeamMember struct {
	*MSFT_NetImPlatAdapter
}

func NewMSFT_NetSwitchTeamMemberEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetSwitchTeamMember, err error) {
	tmp, err := NewMSFT_NetImPlatAdapterEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSwitchTeamMember{
		MSFT_NetImPlatAdapter: tmp,
	}
	return
}

func NewMSFT_NetSwitchTeamMemberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetSwitchTeamMember, err error) {
	tmp, err := NewMSFT_NetImPlatAdapterEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSwitchTeamMember{
		MSFT_NetImPlatAdapter: tmp,
	}
	return
}
