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

// Msvm_TimeSyncComponentSettingData struct
type Msvm_TimeSyncComponentSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	EnabledState uint16
}

func NewMsvm_TimeSyncComponentSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_TimeSyncComponentSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_TimeSyncComponentSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_TimeSyncComponentSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_TimeSyncComponentSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_TimeSyncComponentSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_TimeSyncComponentSettingData) SetPropertyEnabledState(value uint16) (err error) {
	return instance.SetProperty("EnabledState", value)
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_TimeSyncComponentSettingData) GetPropertyEnabledState() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_TimeSyncComponentSettingData) GetRelatedVirtualSystemSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemSettingData")
}

func (instance *Msvm_TimeSyncComponentSettingData) GetRelatedTimeSyncComponent() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_TimeSyncComponent")
}
