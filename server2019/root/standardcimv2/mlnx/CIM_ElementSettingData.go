// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ElementSettingData struct
type CIM_ElementSettingData struct {
	cim.WmiInstance

	//
	IsCurrent ElementSettingData_IsCurrent

	//
	IsDefault ElementSettingData_IsDefault

	//
	IsNext ElementSettingData_IsNext

	//
	ManagedElement CIM_ManagedElement

	//
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
