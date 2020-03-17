// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_SettingCheck struct
type CIM_SettingCheck struct {
	CIM_Check

	//
	CheckType uint16

	//
	EntryName string

	//
	EntryValue string

	//
	FileName string

	//
	SectionKey string
}

// SetCheckType sets the value of CheckType for the instance
func (instance *CIM_SettingCheck) SetPropertyCheckType(value uint16) (err error) {
	return instance.SetProperty("CheckType", value)
}

// GetCheckType gets the value of CheckType for the instance
func (instance *CIM_SettingCheck) GetPropertyCheckType() (value uint16, err error) {
	retValue, err := instance.GetProperty("CheckType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEntryName sets the value of EntryName for the instance
func (instance *CIM_SettingCheck) SetPropertyEntryName(value string) (err error) {
	return instance.SetProperty("EntryName", value)
}

// GetEntryName gets the value of EntryName for the instance
func (instance *CIM_SettingCheck) GetPropertyEntryName() (value string, err error) {
	retValue, err := instance.GetProperty("EntryName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEntryValue sets the value of EntryValue for the instance
func (instance *CIM_SettingCheck) SetPropertyEntryValue(value string) (err error) {
	return instance.SetProperty("EntryValue", value)
}

// GetEntryValue gets the value of EntryValue for the instance
func (instance *CIM_SettingCheck) GetPropertyEntryValue() (value string, err error) {
	retValue, err := instance.GetProperty("EntryValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileName sets the value of FileName for the instance
func (instance *CIM_SettingCheck) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", value)
}

// GetFileName gets the value of FileName for the instance
func (instance *CIM_SettingCheck) GetPropertyFileName() (value string, err error) {
	retValue, err := instance.GetProperty("FileName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSectionKey sets the value of SectionKey for the instance
func (instance *CIM_SettingCheck) SetPropertySectionKey(value string) (err error) {
	return instance.SetProperty("SectionKey", value)
}

// GetSectionKey gets the value of SectionKey for the instance
func (instance *CIM_SettingCheck) GetPropertySectionKey() (value string, err error) {
	retValue, err := instance.GetProperty("SectionKey")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
