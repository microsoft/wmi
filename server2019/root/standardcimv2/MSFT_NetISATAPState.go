// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetISATAPState struct
type MSFT_NetISATAPState struct {
	*CIM_ElementSettingData
}

func NewMSFT_NetISATAPStateEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetISATAPState, err error) {
	tmp, err := NewCIM_ElementSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetISATAPState{
		CIM_ElementSettingData: tmp,
	}
	return
}

func NewMSFT_NetISATAPStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetISATAPState, err error) {
	tmp, err := NewCIM_ElementSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetISATAPState{
		CIM_ElementSettingData: tmp,
	}
	return
}
