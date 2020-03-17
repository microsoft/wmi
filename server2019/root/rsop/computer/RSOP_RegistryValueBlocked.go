// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_RegistryValueBlocked struct
type RSOP_RegistryValueBlocked struct {
	RSOP_SecuritySettingsBlocked

	//
	Data string

	//
	Path string

	//
	Type RegistryValueBlocked_Type
}

// SetData sets the value of Data for the instance
func (instance *RSOP_RegistryValueBlocked) SetPropertyData(value string) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *RSOP_RegistryValueBlocked) GetPropertyData() (value string, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *RSOP_RegistryValueBlocked) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *RSOP_RegistryValueBlocked) GetPropertyPath() (value string, err error) {
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

// SetType sets the value of Type for the instance
func (instance *RSOP_RegistryValueBlocked) SetPropertyType(value RegistryValueBlocked_Type) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *RSOP_RegistryValueBlocked) GetPropertyType() (value RegistryValueBlocked_Type, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(RegistryValueBlocked_Type)
	if !ok {
		// TODO: Set an error
	}
	return
}
