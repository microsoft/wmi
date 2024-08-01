// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NSE31A6F0B_2C62_4FC1_A177_9F88DF963F17.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrFolderOptionSettingVista struct
type RSOP_PolmkrFolderOptionSettingVista struct {
	*RSOP_PolmkrProSetting
}

func NewRSOP_PolmkrFolderOptionSettingVistaEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrFolderOptionSettingVista, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrFolderOptionSettingVista{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrFolderOptionSettingVistaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrFolderOptionSettingVista, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrFolderOptionSettingVista{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}
