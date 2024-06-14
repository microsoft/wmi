// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS46D0FB7F_C78E_4A38_AB7D_0E745BB9B155.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrEnvironmentSetting struct
type RSOP_PolmkrEnvironmentSetting struct {
	*RSOP_PolmkrProSetting
}

func NewRSOP_PolmkrEnvironmentSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrEnvironmentSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrEnvironmentSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrEnvironmentSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrEnvironmentSetting, err error) {
	tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrEnvironmentSetting{
		RSOP_PolmkrProSetting: tmp,
	}
	return
}
