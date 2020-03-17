// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_SystemService struct
type RSOP_SystemService struct {
	RSOP_SecuritySettings

	//
	SDDLString string

	//
	Service string

	//
	StartupMode SystemService_StartupMode
}

// SetSDDLString sets the value of SDDLString for the instance
func (instance *RSOP_SystemService) SetPropertySDDLString(value string) (err error) {
	return instance.SetProperty("SDDLString", value)
}

// GetSDDLString gets the value of SDDLString for the instance
func (instance *RSOP_SystemService) GetPropertySDDLString() (value string, err error) {
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
func (instance *RSOP_SystemService) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *RSOP_SystemService) GetPropertyService() (value string, err error) {
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
func (instance *RSOP_SystemService) SetPropertyStartupMode(value SystemService_StartupMode) (err error) {
	return instance.SetProperty("StartupMode", value)
}

// GetStartupMode gets the value of StartupMode for the instance
func (instance *RSOP_SystemService) GetPropertyStartupMode() (value SystemService_StartupMode, err error) {
	retValue, err := instance.GetProperty("StartupMode")
	if err != nil {
		return
	}
	value, ok := retValue.(SystemService_StartupMode)
	if !ok {
		// TODO: Set an error
	}
	return
}
