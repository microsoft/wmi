// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS9606B916_B14A_4EB6_84F0_165F56E1DD34.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrServerSetting struct
type RSOP_PolmkrServerSetting struct {
	*RSOP_PolmkrSetting
}

func NewRSOP_PolmkrServerSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrServerSetting, err error) {
	tmp, err := NewRSOP_PolmkrSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrServerSetting{
		RSOP_PolmkrSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrServerSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrServerSetting, err error) {
	tmp, err := NewRSOP_PolmkrSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrServerSetting{
		RSOP_PolmkrSetting: tmp,
	}
	return
}