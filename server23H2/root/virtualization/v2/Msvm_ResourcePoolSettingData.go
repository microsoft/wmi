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

// Msvm_ResourcePoolSettingData struct
type Msvm_ResourcePoolSettingData struct {
	*Msvm_AbstractResourcePoolSettingData
}

func NewMsvm_ResourcePoolSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_ResourcePoolSettingData, err error) {
	tmp, err := NewMsvm_AbstractResourcePoolSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ResourcePoolSettingData{
		Msvm_AbstractResourcePoolSettingData: tmp,
	}
	return
}

func NewMsvm_ResourcePoolSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ResourcePoolSettingData, err error) {
	tmp, err := NewMsvm_AbstractResourcePoolSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ResourcePoolSettingData{
		Msvm_AbstractResourcePoolSettingData: tmp,
	}
	return
}

func (instance *Msvm_ResourcePoolSettingData) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}
