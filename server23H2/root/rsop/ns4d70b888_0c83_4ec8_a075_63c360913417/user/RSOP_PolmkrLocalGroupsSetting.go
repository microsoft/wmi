// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS4D70B888_0C83_4EC8_A075_63C360913417.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrLocalGroupsSetting struct
type RSOP_PolmkrLocalGroupsSetting struct {
	*RSOP_PolmkrLocalUsersAndGroupsSetting
}

func NewRSOP_PolmkrLocalGroupsSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrLocalGroupsSetting, err error) {
	tmp, err := NewRSOP_PolmkrLocalUsersAndGroupsSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrLocalGroupsSetting{
		RSOP_PolmkrLocalUsersAndGroupsSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrLocalGroupsSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrLocalGroupsSetting, err error) {
	tmp, err := NewRSOP_PolmkrLocalUsersAndGroupsSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrLocalGroupsSetting{
		RSOP_PolmkrLocalUsersAndGroupsSetting: tmp,
	}
	return
}