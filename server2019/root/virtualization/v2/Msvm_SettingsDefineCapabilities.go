// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_SettingsDefineCapabilities struct
type Msvm_SettingsDefineCapabilities struct {
	CIM_SettingsDefineCapabilities

	//
	SupportStatement uint16
}

// SetSupportStatement sets the value of SupportStatement for the instance
func (instance *Msvm_SettingsDefineCapabilities) SetPropertySupportStatement(value uint16) (err error) {
	return instance.SetProperty("SupportStatement", value)
}

// GetSupportStatement gets the value of SupportStatement for the instance
func (instance *Msvm_SettingsDefineCapabilities) GetPropertySupportStatement() (value uint16, err error) {
	retValue, err := instance.GetProperty("SupportStatement")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
