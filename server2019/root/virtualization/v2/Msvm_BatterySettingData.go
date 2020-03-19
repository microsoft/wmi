// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_BatterySettingData struct
type Msvm_BatterySettingData struct {
	*CIM_ResourceAllocationSettingData
}

func NewMsvm_BatterySettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_BatterySettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_BatterySettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_BatterySettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_BatterySettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_BatterySettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func (instance *Msvm_BatterySettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
