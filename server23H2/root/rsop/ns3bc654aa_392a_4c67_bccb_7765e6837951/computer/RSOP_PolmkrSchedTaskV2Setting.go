// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS3BC654AA_392A_4C67_BCCB_7765E6837951.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrSchedTaskV2Setting struct
type RSOP_PolmkrSchedTaskV2Setting struct {
	*RSOP_PolmkrTaskSetting
}

func NewRSOP_PolmkrSchedTaskV2SettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrSchedTaskV2Setting, err error) {
	tmp, err := NewRSOP_PolmkrTaskSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrSchedTaskV2Setting{
		RSOP_PolmkrTaskSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrSchedTaskV2SettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrSchedTaskV2Setting, err error) {
	tmp, err := NewRSOP_PolmkrTaskSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrSchedTaskV2Setting{
		RSOP_PolmkrTaskSetting: tmp,
	}
	return
}
