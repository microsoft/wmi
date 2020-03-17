// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_SystemServiceBlocked struct
type RSOP_SystemServiceBlocked struct {
	RSOP_SecuritySettingsBlocked

	//
	SDDLString string

	//
	Service string

	//
	StartupMode SystemServiceBlocked_StartupMode
}

// SetSDDLString sets the value of SDDLString for the instance
func (instance *RSOP_SystemServiceBlocked) SetPropertySDDLString(value string) (err error) {
	return instance.SetProperty("SDDLString", value)
}

// GetSDDLString gets the value of SDDLString for the instance
func (instance *RSOP_SystemServiceBlocked) GetPropertySDDLString() (value string, err error) {
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

// SetService sets the value of Service for the instance
func (instance *RSOP_SystemServiceBlocked) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *RSOP_SystemServiceBlocked) GetPropertyService() (value string, err error) {
	retValue, err := instance.GetProperty("Service")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartupMode sets the value of StartupMode for the instance
func (instance *RSOP_SystemServiceBlocked) SetPropertyStartupMode(value SystemServiceBlocked_StartupMode) (err error) {
	return instance.SetProperty("StartupMode", value)
}

// GetStartupMode gets the value of StartupMode for the instance
func (instance *RSOP_SystemServiceBlocked) GetPropertyStartupMode() (value SystemServiceBlocked_StartupMode, err error) {
	retValue, err := instance.GetProperty("StartupMode")
	if err != nil {
		return
	}
	value, ok := retValue.(SystemServiceBlocked_StartupMode)
	if !ok {
		// TODO: Set an error
	}
	return
}
