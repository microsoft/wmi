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

// RSOP_PolmkrImmediateTaskV2Setting struct
type RSOP_PolmkrImmediateTaskV2Setting struct {
	*RSOP_PolmkrTaskSetting
}

func NewRSOP_PolmkrImmediateTaskV2SettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrImmediateTaskV2Setting, err error) {
	tmp, err := NewRSOP_PolmkrTaskSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrImmediateTaskV2Setting{
		RSOP_PolmkrTaskSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrImmediateTaskV2SettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrImmediateTaskV2Setting, err error) {
	tmp, err := NewRSOP_PolmkrTaskSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrImmediateTaskV2Setting{
		RSOP_PolmkrTaskSetting: tmp,
	}
	return
}