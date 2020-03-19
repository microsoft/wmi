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

// Msvm_GuestServiceInterfaceComponentSettingData struct
type Msvm_GuestServiceInterfaceComponentSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	DefaultEnabledStatePolicy GuestServiceInterfaceComponentSettingData_DefaultEnabledStatePolicy

	//
	EnabledState GuestServiceInterfaceComponentSettingData_EnabledState
}

func NewMsvm_GuestServiceInterfaceComponentSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_GuestServiceInterfaceComponentSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestServiceInterfaceComponentSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_GuestServiceInterfaceComponentSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_GuestServiceInterfaceComponentSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestServiceInterfaceComponentSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetDefaultEnabledStatePolicy sets the value of DefaultEnabledStatePolicy for the instance
func (instance *Msvm_GuestServiceInterfaceComponentSettingData) SetPropertyDefaultEnabledStatePolicy(value GuestServiceInterfaceComponentSettingData_DefaultEnabledStatePolicy) (err error) {
	return instance.SetProperty("DefaultEnabledStatePolicy", value)
}

// GetDefaultEnabledStatePolicy gets the value of DefaultEnabledStatePolicy for the instance
func (instance *Msvm_GuestServiceInterfaceComponentSettingData) GetPropertyDefaultEnabledStatePolicy() (value GuestServiceInterfaceComponentSettingData_DefaultEnabledStatePolicy, err error) {
	retValue, err := instance.GetProperty("DefaultEnabledStatePolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(GuestServiceInterfaceComponentSettingData_DefaultEnabledStatePolicy)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_GuestServiceInterfaceComponentSettingData) SetPropertyEnabledState(value GuestServiceInterfaceComponentSettingData_EnabledState) (err error) {
	return instance.SetProperty("EnabledState", value)
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_GuestServiceInterfaceComponentSettingData) GetPropertyEnabledState() (value GuestServiceInterfaceComponentSettingData_EnabledState, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(GuestServiceInterfaceComponentSettingData_EnabledState)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_GuestServiceInterfaceComponentSettingData) GetRelatedVirtualSystemSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemSettingData")
}

func (instance *Msvm_GuestServiceInterfaceComponentSettingData) GetRelatedGuestServiceInterfaceComponent() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_GuestServiceInterfaceComponent")
}
