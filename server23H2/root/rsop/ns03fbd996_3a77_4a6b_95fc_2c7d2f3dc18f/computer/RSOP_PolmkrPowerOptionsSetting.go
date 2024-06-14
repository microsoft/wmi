// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS03FBD996_3A77_4A6B_95FC_2C7D2F3DC18F.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrPowerOptionsSetting struct
type RSOP_PolmkrPowerOptionsSetting struct {
	*RSOP_PolmkrPowerSetting
}

func NewRSOP_PolmkrPowerOptionsSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrPowerOptionsSetting, err error) {
	tmp, err := NewRSOP_PolmkrPowerSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrPowerOptionsSetting{
		RSOP_PolmkrPowerSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrPowerOptionsSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrPowerOptionsSetting, err error) {
	tmp, err := NewRSOP_PolmkrPowerSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrPowerOptionsSetting{
		RSOP_PolmkrPowerSetting: tmp,
	}
	return
}
