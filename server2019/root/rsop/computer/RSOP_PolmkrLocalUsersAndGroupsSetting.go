// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrLocalUsersAndGroupsSetting struct
type RSOP_PolmkrLocalUsersAndGroupsSetting struct {
	*RSOP_PolmkrProSetting
}

func NewRSOP_PolmkrLocalUsersAndGroupsSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrLocalUsersAndGroupsSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrLocalUsersAndGroupsSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrLocalUsersAndGroupsSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrLocalUsersAndGroupsSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrLocalUsersAndGroupsSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}
