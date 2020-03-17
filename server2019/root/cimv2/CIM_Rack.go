// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_Rack struct
type CIM_Rack struct {
	CIM_PhysicalFrame

	//
	CountryDesignation string

	//
	TypeOfRack uint16
}

// SetCountryDesignation sets the value of CountryDesignation for the instance
func (instance *CIM_Rack) SetPropertyCountryDesignation(value string) (err error) {
	return instance.SetProperty("CountryDesignation", value)
}

// GetCountryDesignation gets the value of CountryDesignation for the instance
func (instance *CIM_Rack) GetPropertyCountryDesignation() (value string, err error) {
	retValue, err := instance.GetProperty("CountryDesignation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTypeOfRack sets the value of TypeOfRack for the instance
func (instance *CIM_Rack) SetPropertyTypeOfRack(value uint16) (err error) {
	return instance.SetProperty("TypeOfRack", value)
}

// GetTypeOfRack gets the value of TypeOfRack for the instance
func (instance *CIM_Rack) GetPropertyTypeOfRack() (value uint16, err error) {
	retValue, err := instance.GetProperty("TypeOfRack")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
