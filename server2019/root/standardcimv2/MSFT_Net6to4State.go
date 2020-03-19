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

// MSFT_Net6to4State struct
type MSFT_Net6to4State struct {
	*CIM_ElementSettingData
}

func NewMSFT_Net6to4StateEx1(instance *cim.WmiInstance) (newInstance *MSFT_Net6to4State, err error) {
	tmp, err := NewCIM_ElementSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_Net6to4State{
		CIM_ElementSettingData: tmp,
	}
	return
}

func NewMSFT_Net6to4StateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_Net6to4State, err error) {
	tmp, err := NewCIM_ElementSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_Net6to4State{
		CIM_ElementSettingData: tmp,
	}
	return
}
