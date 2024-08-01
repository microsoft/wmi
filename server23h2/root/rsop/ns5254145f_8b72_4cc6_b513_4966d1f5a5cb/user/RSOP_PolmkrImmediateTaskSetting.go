// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS5254145F_8B72_4CC6_B513_4966D1F5A5CB.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrImmediateTaskSetting struct
type RSOP_PolmkrImmediateTaskSetting struct {
	*RSOP_PolmkrTaskSetting
}

func NewRSOP_PolmkrImmediateTaskSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrImmediateTaskSetting, err error) {
	tmp, err := NewRSOP_PolmkrTaskSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrImmediateTaskSetting{
		RSOP_PolmkrTaskSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrImmediateTaskSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrImmediateTaskSetting, err error) {
	tmp, err := NewRSOP_PolmkrTaskSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrImmediateTaskSetting{
		RSOP_PolmkrTaskSetting: tmp,
	}
	return
}
