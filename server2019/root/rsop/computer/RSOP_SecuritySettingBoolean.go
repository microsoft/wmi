// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_SecuritySettingBoolean struct
type RSOP_SecuritySettingBoolean struct {
	RSOP_SecuritySettings

	//
	KeyName string

	//
	Setting bool
}

// SetKeyName sets the value of KeyName for the instance
func (instance *RSOP_SecuritySettingBoolean) SetPropertyKeyName(value string) (err error) {
	return instance.SetProperty("KeyName", value)
}

// GetKeyName gets the value of KeyName for the instance
func (instance *RSOP_SecuritySettingBoolean) GetPropertyKeyName() (value string, err error) {
	retValue, err := instance.GetProperty("KeyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetting sets the value of Setting for the instance
func (instance *RSOP_SecuritySettingBoolean) SetPropertySetting(value bool) (err error) {
	return instance.SetProperty("Setting", value)
}

// GetSetting gets the value of Setting for the instance
func (instance *RSOP_SecuritySettingBoolean) GetPropertySetting() (value bool, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
