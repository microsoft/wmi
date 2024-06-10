// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NSE415D675_0F49_4CFF_A478_A2849757FFEE.User
//
// ////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrDriveSetting struct
type RSOP_PolmkrDriveSetting struct {
	*RSOP_PolmkrProSetting
}

func NewRSOP_PolmkrDriveSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrDriveSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrDriveSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrDriveSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrDriveSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrDriveSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}
