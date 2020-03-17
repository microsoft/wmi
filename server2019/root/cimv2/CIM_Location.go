// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_Location struct
type CIM_Location struct {
	cim.WmiInstance

	//
	Address string

	//
	Name string

	//
	PhysicalPosition string
}

// SetAddress sets the value of Address for the instance
func (instance *CIM_Location) SetPropertyAddress(value string) (err error) {
	return instance.SetProperty("Address", value)
}

// GetAddress gets the value of Address for the instance
func (instance *CIM_Location) GetPropertyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("Address")
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
func (instance *CIM_Location) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *CIM_Location) GetPropertyName() (value string, err error) {
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

// SetPhysicalPosition sets the value of PhysicalPosition for the instance
func (instance *CIM_Location) SetPropertyPhysicalPosition(value string) (err error) {
	return instance.SetProperty("PhysicalPosition", value)
}

// GetPhysicalPosition gets the value of PhysicalPosition for the instance
func (instance *CIM_Location) GetPropertyPhysicalPosition() (value string, err error) {
	retValue, err := instance.GetProperty("PhysicalPosition")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
