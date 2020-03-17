// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_GuestServiceInterfaceComponentSettingData struct
type Msvm_GuestServiceInterfaceComponentSettingData struct {
	CIM_ResourceAllocationSettingData

	//
	DefaultEnabledStatePolicy GuestServiceInterfaceComponentSettingData_DefaultEnabledStatePolicy

	//
	EnabledState GuestServiceInterfaceComponentSettingData_EnabledState
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
