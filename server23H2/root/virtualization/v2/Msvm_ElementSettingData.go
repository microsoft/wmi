// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ElementSettingData struct
type Msvm_ElementSettingData struct {
	*CIM_ElementSettingData
}

func NewMsvm_ElementSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_ElementSettingData, err error) {
	tmp, err := NewCIM_ElementSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ElementSettingData{
		CIM_ElementSettingData: tmp,
	}
	return
}

func NewMsvm_ElementSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ElementSettingData, err error) {
	tmp, err := NewCIM_ElementSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ElementSettingData{
		CIM_ElementSettingData: tmp,
	}
	return
}
