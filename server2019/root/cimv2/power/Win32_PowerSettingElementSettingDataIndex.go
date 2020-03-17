// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

// Win32_PowerSettingElementSettingDataIndex struct
type Win32_PowerSettingElementSettingDataIndex struct {
	CIM_ElementSettingData

	//
	IsACSetting uint16
}

// SetIsACSetting sets the value of IsACSetting for the instance
func (instance *Win32_PowerSettingElementSettingDataIndex) SetPropertyIsACSetting(value uint16) (err error) {
	return instance.SetProperty("IsACSetting", value)
}

// GetIsACSetting gets the value of IsACSetting for the instance
func (instance *Win32_PowerSettingElementSettingDataIndex) GetPropertyIsACSetting() (value uint16, err error) {
	retValue, err := instance.GetProperty("IsACSetting")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
