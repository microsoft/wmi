// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_Chip struct
type CIM_Chip struct {
	CIM_PhysicalComponent

	//
	FormFactor uint16
}

// SetFormFactor sets the value of FormFactor for the instance
func (instance *CIM_Chip) SetPropertyFormFactor(value uint16) (err error) {
	return instance.SetProperty("FormFactor", value)
}

// GetFormFactor gets the value of FormFactor for the instance
func (instance *CIM_Chip) GetPropertyFormFactor() (value uint16, err error) {
	retValue, err := instance.GetProperty("FormFactor")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
