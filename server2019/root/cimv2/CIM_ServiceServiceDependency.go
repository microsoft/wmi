// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_ServiceServiceDependency struct
type CIM_ServiceServiceDependency struct {
	CIM_Dependency

	//
	TypeOfDependency uint16
}

// SetTypeOfDependency sets the value of TypeOfDependency for the instance
func (instance *CIM_ServiceServiceDependency) SetPropertyTypeOfDependency(value uint16) (err error) {
	return instance.SetProperty("TypeOfDependency", value)
}

// GetTypeOfDependency gets the value of TypeOfDependency for the instance
func (instance *CIM_ServiceServiceDependency) GetPropertyTypeOfDependency() (value uint16, err error) {
	retValue, err := instance.GetProperty("TypeOfDependency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
