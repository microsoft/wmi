// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

// CIM_SettingsDefineCapabilities struct
type CIM_SettingsDefineCapabilities struct {
	CIM_Component

	//
	PropertyPolicy uint16

	//
	ValueRange uint16

	//
	ValueRole uint16
}

// SetPropertyPolicy sets the value of PropertyPolicy for the instance
func (instance *CIM_SettingsDefineCapabilities) SetPropertyPropertyPolicy(value uint16) (err error) {
	return instance.SetProperty("PropertyPolicy", value)
}

// GetPropertyPolicy gets the value of PropertyPolicy for the instance
func (instance *CIM_SettingsDefineCapabilities) GetPropertyPropertyPolicy() (value uint16, err error) {
	retValue, err := instance.GetProperty("PropertyPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValueRange sets the value of ValueRange for the instance
func (instance *CIM_SettingsDefineCapabilities) SetPropertyValueRange(value uint16) (err error) {
	return instance.SetProperty("ValueRange", value)
}

// GetValueRange gets the value of ValueRange for the instance
func (instance *CIM_SettingsDefineCapabilities) GetPropertyValueRange() (value uint16, err error) {
	retValue, err := instance.GetProperty("ValueRange")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValueRole sets the value of ValueRole for the instance
func (instance *CIM_SettingsDefineCapabilities) SetPropertyValueRole(value uint16) (err error) {
	return instance.SetProperty("ValueRole", value)
}

// GetValueRole gets the value of ValueRole for the instance
func (instance *CIM_SettingsDefineCapabilities) GetPropertyValueRole() (value uint16, err error) {
	retValue, err := instance.GetProperty("ValueRole")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
