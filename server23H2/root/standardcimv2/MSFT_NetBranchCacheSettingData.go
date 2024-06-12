// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetBranchCacheSettingData struct
type MSFT_NetBranchCacheSettingData struct {
	*MSFT_NetSettingData
}

func NewMSFT_NetBranchCacheSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCacheSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetBranchCacheSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBranchCacheSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}
