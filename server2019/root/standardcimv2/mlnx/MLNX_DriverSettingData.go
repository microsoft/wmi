// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// MLNX_DriverSettingData struct
type MLNX_DriverSettingData struct {
	CIM_SettingData

	//
	Name string

	//
	Source DriverSettingData_Source

	//
	SystemName string
}

// SetName sets the value of Name for the instance
func (instance *MLNX_DriverSettingData) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MLNX_DriverSettingData) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSource sets the value of Source for the instance
func (instance *MLNX_DriverSettingData) SetPropertySource(value DriverSettingData_Source) (err error) {
	return instance.SetProperty("Source", value)
}

// GetSource gets the value of Source for the instance
func (instance *MLNX_DriverSettingData) GetPropertySource() (value DriverSettingData_Source, err error) {
	retValue, err := instance.GetProperty("Source")
	if err != nil {
		return
	}
	value, ok := retValue.(DriverSettingData_Source)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *MLNX_DriverSettingData) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *MLNX_DriverSettingData) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
