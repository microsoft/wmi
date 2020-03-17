// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_GlobalSaclSetting struct
type RSOP_GlobalSaclSetting struct {
	RSOP_PolicySetting

	//
	SettingType string

	//
	SettingValue string
}

// SetSettingType sets the value of SettingType for the instance
func (instance *RSOP_GlobalSaclSetting) SetPropertySettingType(value string) (err error) {
	return instance.SetProperty("SettingType", value)
}

// GetSettingType gets the value of SettingType for the instance
func (instance *RSOP_GlobalSaclSetting) GetPropertySettingType() (value string, err error) {
	retValue, err := instance.GetProperty("SettingType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettingValue sets the value of SettingValue for the instance
func (instance *RSOP_GlobalSaclSetting) SetPropertySettingValue(value string) (err error) {
	return instance.SetProperty("SettingValue", value)
}

// GetSettingValue gets the value of SettingValue for the instance
func (instance *RSOP_GlobalSaclSetting) GetPropertySettingValue() (value string, err error) {
	retValue, err := instance.GetProperty("SettingValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
