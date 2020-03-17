// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_GuestCommunicationServiceSettingData struct
type Msvm_GuestCommunicationServiceSettingData struct {
	CIM_SettingData

	// EnabledStatePolicy is an integer enumeration that indicates the enabled, disabled or default state of the Msvm_GuestCommunicationServiceSettingData.Enabled (2) indicates that the communication service is set to the enabled state.
	///Disabled (3) indicates that the communication service is set to the disabled state.
	///Deferred (8) indicates that the communication service state depends on DefaultEnabledStatePolicy in Msvm_GuestInterfaceComponentSettingData.
	///
	EnabledStatePolicy GuestCommunicationServiceSettingData_EnabledStatePolicy

	//
	Name string
}

// SetEnabledStatePolicy sets the value of EnabledStatePolicy for the instance
func (instance *Msvm_GuestCommunicationServiceSettingData) SetPropertyEnabledStatePolicy(value GuestCommunicationServiceSettingData_EnabledStatePolicy) (err error) {
	return instance.SetProperty("EnabledStatePolicy", value)
}

// GetEnabledStatePolicy gets the value of EnabledStatePolicy for the instance
func (instance *Msvm_GuestCommunicationServiceSettingData) GetPropertyEnabledStatePolicy() (value GuestCommunicationServiceSettingData_EnabledStatePolicy, err error) {
	retValue, err := instance.GetProperty("EnabledStatePolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(GuestCommunicationServiceSettingData_EnabledStatePolicy)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *Msvm_GuestCommunicationServiceSettingData) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Msvm_GuestCommunicationServiceSettingData) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
