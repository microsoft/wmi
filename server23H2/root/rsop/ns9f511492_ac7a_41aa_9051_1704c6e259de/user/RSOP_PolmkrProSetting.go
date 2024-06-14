// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS9F511492_AC7A_41AA_9051_1704C6E259DE.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrProSetting struct
type RSOP_PolmkrProSetting struct {
	*RSOP_PolmkrSetting
}

func NewRSOP_PolmkrProSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrProSetting, err error) {
	tmp, err := NewRSOP_PolmkrSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrProSetting{
		RSOP_PolmkrSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrProSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrProSetting, err error) {
	tmp, err := NewRSOP_PolmkrSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrProSetting{
		RSOP_PolmkrSetting: tmp,
	}
	return
}