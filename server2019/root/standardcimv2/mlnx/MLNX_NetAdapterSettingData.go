// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// MLNX_NetAdapterSettingData struct
type MLNX_NetAdapterSettingData struct {
	CIM_SettingData

	//
	InterfaceDescription string

	//
	Name string

	//
	Source NetAdapterSettingData_Source

	//
	SystemName string
}

// SetInterfaceDescription sets the value of InterfaceDescription for the instance
func (instance *MLNX_NetAdapterSettingData) SetPropertyInterfaceDescription(value string) (err error) {
	return instance.SetProperty("InterfaceDescription", value)
}

// GetInterfaceDescription gets the value of InterfaceDescription for the instance
func (instance *MLNX_NetAdapterSettingData) GetPropertyInterfaceDescription() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MLNX_NetAdapterSettingData) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MLNX_NetAdapterSettingData) GetPropertyName() (value string, err error) {
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
func (instance *MLNX_NetAdapterSettingData) SetPropertySource(value NetAdapterSettingData_Source) (err error) {
	return instance.SetProperty("Source", value)
}

// GetSource gets the value of Source for the instance
func (instance *MLNX_NetAdapterSettingData) GetPropertySource() (value NetAdapterSettingData_Source, err error) {
	retValue, err := instance.GetProperty("Source")
	if err != nil {
		return
	}
	value, ok := retValue.(NetAdapterSettingData_Source)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *MLNX_NetAdapterSettingData) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *MLNX_NetAdapterSettingData) GetPropertySystemName() (value string, err error) {
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
