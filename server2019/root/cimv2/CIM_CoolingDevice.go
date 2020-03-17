// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_CoolingDevice struct
type CIM_CoolingDevice struct {
	CIM_LogicalDevice

	//
	ActiveCooling bool
}

// SetActiveCooling sets the value of ActiveCooling for the instance
func (instance *CIM_CoolingDevice) SetPropertyActiveCooling(value bool) (err error) {
	return instance.SetProperty("ActiveCooling", value)
}

// GetActiveCooling gets the value of ActiveCooling for the instance
func (instance *CIM_CoolingDevice) GetPropertyActiveCooling() (value bool, err error) {
	retValue, err := instance.GetProperty("ActiveCooling")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
