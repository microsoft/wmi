// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NSCBC80CE2_C2E8_4CBC_8C8C_1C374E56CEFC.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrNetShareSetting struct
type RSOP_PolmkrNetShareSetting struct {
	*RSOP_PolmkrServerSetting
}

func NewRSOP_PolmkrNetShareSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrNetShareSetting, err error) {
	tmp, err := NewRSOP_PolmkrServerSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrNetShareSetting{
		RSOP_PolmkrServerSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrNetShareSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrNetShareSetting, err error) {
	tmp, err := NewRSOP_PolmkrServerSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrNetShareSetting{
		RSOP_PolmkrServerSetting: tmp,
	}
	return
}
