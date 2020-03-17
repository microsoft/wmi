// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_SettingsDefineState struct
type CIM_SettingsDefineState struct {
	cim.WmiInstance

	// The managed element.
	ManagedElement CIM_ManagedElement

	// The SettingData object that provides additional information about the current state and configuration of the ManagedElement.
	SettingData CIM_SettingData
}

// SetManagedElement sets the value of ManagedElement for the instance
func (instance *CIM_SettingsDefineState) SetPropertyManagedElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("ManagedElement", value)
}

// GetManagedElement gets the value of ManagedElement for the instance
func (instance *CIM_SettingsDefineState) GetPropertyManagedElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("ManagedElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettingData sets the value of SettingData for the instance
func (instance *CIM_SettingsDefineState) SetPropertySettingData(value CIM_SettingData) (err error) {
	return instance.SetProperty("SettingData", value)
}

// GetSettingData gets the value of SettingData for the instance
func (instance *CIM_SettingsDefineState) GetPropertySettingData() (value CIM_SettingData, err error) {
	retValue, err := instance.GetProperty("SettingData")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_SettingData)
	if !ok {
		// TODO: Set an error
	}
	return
}
