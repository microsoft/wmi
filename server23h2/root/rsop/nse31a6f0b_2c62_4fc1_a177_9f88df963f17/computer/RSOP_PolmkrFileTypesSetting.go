// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NSE31A6F0B_2C62_4FC1_A177_9F88DF963F17.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrFileTypesSetting struct
type RSOP_PolmkrFileTypesSetting struct {
	*RSOP_PolmkrFolderOptionSetting
}

func NewRSOP_PolmkrFileTypesSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrFileTypesSetting, err error) {
	tmp, err := NewRSOP_PolmkrFolderOptionSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrFileTypesSetting{
		RSOP_PolmkrFolderOptionSetting: tmp,
	}
	return
}

func NewRSOP_PolmkrFileTypesSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrFileTypesSetting, err error) {
	tmp, err := NewRSOP_PolmkrFolderOptionSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrFileTypesSetting{
		RSOP_PolmkrFolderOptionSetting: tmp,
	}
	return
}
