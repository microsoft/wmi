// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

// Win32_PowerSettingDataIndex struct
type Win32_PowerSettingDataIndex struct {
	CIM_SettingData

	//
	SettingIndexValue uint32
}

// SetSettingIndexValue sets the value of SettingIndexValue for the instance
func (instance *Win32_PowerSettingDataIndex) SetPropertySettingIndexValue(value uint32) (err error) {
	return instance.SetProperty("SettingIndexValue", value)
}

// GetSettingIndexValue gets the value of SettingIndexValue for the instance
func (instance *Win32_PowerSettingDataIndex) GetPropertySettingIndexValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("SettingIndexValue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
