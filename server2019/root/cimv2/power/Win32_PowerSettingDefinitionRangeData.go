// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

// Win32_PowerSettingDefinitionRangeData struct
type Win32_PowerSettingDefinitionRangeData struct {
	CIM_SettingData

	//
	SettingValue uint32
}

// SetSettingValue sets the value of SettingValue for the instance
func (instance *Win32_PowerSettingDefinitionRangeData) SetPropertySettingValue(value uint32) (err error) {
	return instance.SetProperty("SettingValue", value)
}

// GetSettingValue gets the value of SettingValue for the instance
func (instance *Win32_PowerSettingDefinitionRangeData) GetPropertySettingValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("SettingValue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
