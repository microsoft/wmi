// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_CollectionSettingData struct
type Msvm_CollectionSettingData struct {
	*CIM_SettingData
}

func NewMsvm_CollectionSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectionSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_CollectionSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CollectionSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionSettingData{
		CIM_SettingData: tmp,
	}
	return
}
