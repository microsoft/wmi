// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_MotherboardDevice struct
type Win32_MotherboardDevice struct {
	CIM_LogicalDevice

	//
	PrimaryBusType string

	//
	RevisionNumber string

	//
	SecondaryBusType string
}

// SetPrimaryBusType sets the value of PrimaryBusType for the instance
func (instance *Win32_MotherboardDevice) SetPropertyPrimaryBusType(value string) (err error) {
	return instance.SetProperty("PrimaryBusType", value)
}

// GetPrimaryBusType gets the value of PrimaryBusType for the instance
func (instance *Win32_MotherboardDevice) GetPropertyPrimaryBusType() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryBusType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRevisionNumber sets the value of RevisionNumber for the instance
func (instance *Win32_MotherboardDevice) SetPropertyRevisionNumber(value string) (err error) {
	return instance.SetProperty("RevisionNumber", value)
}

// GetRevisionNumber gets the value of RevisionNumber for the instance
func (instance *Win32_MotherboardDevice) GetPropertyRevisionNumber() (value string, err error) {
	retValue, err := instance.GetProperty("RevisionNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecondaryBusType sets the value of SecondaryBusType for the instance
func (instance *Win32_MotherboardDevice) SetPropertySecondaryBusType(value string) (err error) {
	return instance.SetProperty("SecondaryBusType", value)
}

// GetSecondaryBusType gets the value of SecondaryBusType for the instance
func (instance *Win32_MotherboardDevice) GetPropertySecondaryBusType() (value string, err error) {
	retValue, err := instance.GetProperty("SecondaryBusType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
