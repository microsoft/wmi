// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_RegistryValue struct
type RSOP_RegistryValue struct {
	RSOP_SecuritySettings

	//
	Data string

	//
	Path string

	//
	Type RegistryValue_Type
}

// SetData sets the value of Data for the instance
func (instance *RSOP_RegistryValue) SetPropertyData(value string) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *RSOP_RegistryValue) GetPropertyData() (value string, err error) {
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
func (instance *RSOP_RegistryValue) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *RSOP_RegistryValue) GetPropertyPath() (value string, err error) {
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
func (instance *RSOP_RegistryValue) SetPropertyType(value RegistryValue_Type) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *RSOP_RegistryValue) GetPropertyType() (value RegistryValue_Type, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(RegistryValue_Type)
	if !ok {
		// TODO: Set an error
	}
	return
}
