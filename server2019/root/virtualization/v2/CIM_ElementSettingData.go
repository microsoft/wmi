// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ElementSettingData struct
type CIM_ElementSettingData struct {
	cim.WmiInstance

	// An enumerated integer indicating that the referenced setting is currently being used in the operation of the element, or that this information is unknown.
	IsCurrent ElementSettingData_IsCurrent

	// An enumerated integer indicating that the referenced setting is a default setting for the element, or that this information is unknown.
	IsDefault ElementSettingData_IsDefault

	// An enumerated integer indicating whether or not the referenced setting is the next setting to be applied. For example, the application could take place on a re-initialization, reset, reconfiguration request. This could be a permanent setting, or a setting used only one time, as indicated by the flag. If it is a permanent setting then the setting is applied every time the managed element reinitializes, until this flag is manually reset. However, if it is single use, then the flag is automatically cleared after the settings are applied. Also note that if this flag is specified (i.e. set to value other than "Unknown"), then this takes precedence over any SettingData that may have been specified as Default. For example: If the managed element is a computer system, and the value of this flag is "Is Next", then the setting will be effective next time the system resets. And, unless this flag is changed, it will persist for subsequent system resets. However, if this flag is set to "Is Next For Single Use", then this setting will only be used once and the flag would be reset after that to "Is Not Next". So, in the above example, if the system reboots in a quick succession, the setting will not be used at the second reboot.
	IsNext ElementSettingData_IsNext

	// The managed element.
	ManagedElement CIM_ManagedElement

	// The SettingData object associated with the element.
	SettingData CIM_SettingData
}

// SetIsCurrent sets the value of IsCurrent for the instance
func (instance *CIM_ElementSettingData) SetPropertyIsCurrent(value ElementSettingData_IsCurrent) (err error) {
	return instance.SetProperty("IsCurrent", value)
}

// GetIsCurrent gets the value of IsCurrent for the instance
func (instance *CIM_ElementSettingData) GetPropertyIsCurrent() (value ElementSettingData_IsCurrent, err error) {
	retValue, err := instance.GetProperty("IsCurrent")
	if err != nil {
		return
	}
	value, ok := retValue.(ElementSettingData_IsCurrent)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsDefault sets the value of IsDefault for the instance
func (instance *CIM_ElementSettingData) SetPropertyIsDefault(value ElementSettingData_IsDefault) (err error) {
	return instance.SetProperty("IsDefault", value)
}

// GetIsDefault gets the value of IsDefault for the instance
func (instance *CIM_ElementSettingData) GetPropertyIsDefault() (value ElementSettingData_IsDefault, err error) {
	retValue, err := instance.GetProperty("IsDefault")
	if err != nil {
		return
	}
	value, ok := retValue.(ElementSettingData_IsDefault)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsNext sets the value of IsNext for the instance
func (instance *CIM_ElementSettingData) SetPropertyIsNext(value ElementSettingData_IsNext) (err error) {
	return instance.SetProperty("IsNext", value)
}

// GetIsNext gets the value of IsNext for the instance
func (instance *CIM_ElementSettingData) GetPropertyIsNext() (value ElementSettingData_IsNext, err error) {
	retValue, err := instance.GetProperty("IsNext")
	if err != nil {
		return
	}
	value, ok := retValue.(ElementSettingData_IsNext)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManagedElement sets the value of ManagedElement for the instance
func (instance *CIM_ElementSettingData) SetPropertyManagedElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("ManagedElement", value)
}

// GetManagedElement gets the value of ManagedElement for the instance
func (instance *CIM_ElementSettingData) GetPropertyManagedElement() (value CIM_ManagedElement, err error) {
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
func (instance *CIM_ElementSettingData) SetPropertySettingData(value CIM_SettingData) (err error) {
	return instance.SetProperty("SettingData", value)
}

// GetSettingData gets the value of SettingData for the instance
func (instance *CIM_ElementSettingData) GetPropertySettingData() (value CIM_SettingData, err error) {
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
