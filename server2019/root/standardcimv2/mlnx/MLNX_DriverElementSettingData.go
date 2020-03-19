// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_DriverElementSettingData struct
type MLNX_DriverElementSettingData struct {
	*CIM_ElementSettingData
}

func NewMLNX_DriverElementSettingDataEx1(instance *cim.WmiInstance) (newInstance *MLNX_DriverElementSettingData, err error) {
	tmp, err := NewCIM_ElementSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverElementSettingData{
		CIM_ElementSettingData: tmp,
	}
	return
}

func NewMLNX_DriverElementSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DriverElementSettingData, err error) {
	tmp, err := NewCIM_ElementSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverElementSettingData{
		CIM_ElementSettingData: tmp,
	}
	return
}
