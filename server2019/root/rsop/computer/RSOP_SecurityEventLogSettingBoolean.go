// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_SecurityEventLogSettingBoolean struct
type RSOP_SecurityEventLogSettingBoolean struct {
	RSOP_SecuritySettings

	//
	KeyName string

	//
	Setting bool

	//
	Type string
}

// SetKeyName sets the value of KeyName for the instance
func (instance *RSOP_SecurityEventLogSettingBoolean) SetPropertyKeyName(value string) (err error) {
	return instance.SetProperty("KeyName", value)
}

// GetKeyName gets the value of KeyName for the instance
func (instance *RSOP_SecurityEventLogSettingBoolean) GetPropertyKeyName() (value string, err error) {
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
func (instance *RSOP_SecurityEventLogSettingBoolean) SetPropertySetting(value bool) (err error) {
	return instance.SetProperty("Setting", value)
}

// GetSetting gets the value of Setting for the instance
func (instance *RSOP_SecurityEventLogSettingBoolean) GetPropertySetting() (value bool, err error) {
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

// SetType sets the value of Type for the instance
func (instance *RSOP_SecurityEventLogSettingBoolean) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *RSOP_SecurityEventLogSettingBoolean) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
