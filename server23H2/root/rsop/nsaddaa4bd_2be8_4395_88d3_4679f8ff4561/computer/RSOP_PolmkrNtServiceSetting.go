// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NSADDAA4BD_2BE8_4395_88D3_4679F8FF4561.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrNtServiceSetting struct
type RSOP_PolmkrNtServiceSetting struct {
	*RSOP_PolmkrProSetting
}

func NewRSOP_PolmkrNtServiceSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrNtServiceSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrNtServiceSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrNtServiceSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrNtServiceSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrNtServiceSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}