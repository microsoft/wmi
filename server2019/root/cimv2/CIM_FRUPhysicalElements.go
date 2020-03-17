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

// CIM_FRUPhysicalElements struct
type CIM_FRUPhysicalElements struct {
	cim.WmiInstance

	//
	Component CIM_PhysicalElement

	//
	FRU CIM_FRU
}

// SetComponent sets the value of Component for the instance
func (instance *CIM_FRUPhysicalElements) SetPropertyComponent(value CIM_PhysicalElement) (err error) {
	return instance.SetProperty("Component", value)
}

// GetComponent gets the value of Component for the instance
func (instance *CIM_FRUPhysicalElements) GetPropertyComponent() (value CIM_PhysicalElement, err error) {
	retValue, err := instance.GetProperty("Component")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_PhysicalElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFRU sets the value of FRU for the instance
func (instance *CIM_FRUPhysicalElements) SetPropertyFRU(value CIM_FRU) (err error) {
	return instance.SetProperty("FRU", value)
}

// GetFRU gets the value of FRU for the instance
func (instance *CIM_FRUPhysicalElements) GetPropertyFRU() (value CIM_FRU, err error) {
	retValue, err := instance.GetProperty("FRU")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_FRU)
	if !ok {
		// TODO: Set an error
	}
	return
}
