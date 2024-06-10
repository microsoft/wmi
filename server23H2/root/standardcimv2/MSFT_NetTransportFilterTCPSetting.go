// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetTransportFilterTCPSetting struct
type MSFT_NetTransportFilterTCPSetting struct {
	*CIM_Dependency
}

func NewMSFT_NetTransportFilterTCPSettingEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetTransportFilterTCPSetting, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetTransportFilterTCPSetting{
		CIM_Dependency: tmp,
	}
	return
}

func NewMSFT_NetTransportFilterTCPSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetTransportFilterTCPSetting, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetTransportFilterTCPSetting{
		CIM_Dependency: tmp,
	}
	return
}
