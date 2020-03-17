// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_RegistryKeyBlocked struct
type RSOP_RegistryKeyBlocked struct {
	RSOP_SecuritySettingsBlocked

	//
	Mode RegistryKeyBlocked_Mode

	//
	Path string

	//
	SDDLString string
}

// SetMode sets the value of Mode for the instance
func (instance *RSOP_RegistryKeyBlocked) SetPropertyMode(value RegistryKeyBlocked_Mode) (err error) {
	return instance.SetProperty("Mode", value)
}

// GetMode gets the value of Mode for the instance
func (instance *RSOP_RegistryKeyBlocked) GetPropertyMode() (value RegistryKeyBlocked_Mode, err error) {
	retValue, err := instance.GetProperty("Mode")
	if err != nil {
		return
	}
	value, ok := retValue.(RegistryKeyBlocked_Mode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *RSOP_RegistryKeyBlocked) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *RSOP_RegistryKeyBlocked) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSDDLString sets the value of SDDLString for the instance
func (instance *RSOP_RegistryKeyBlocked) SetPropertySDDLString(value string) (err error) {
	return instance.SetProperty("SDDLString", value)
}

// GetSDDLString gets the value of SDDLString for the instance
func (instance *RSOP_RegistryKeyBlocked) GetPropertySDDLString() (value string, err error) {
	retValue, err := instance.GetProperty("SDDLString")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
