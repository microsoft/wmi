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

// CIM_FRUIncludesProduct struct
type CIM_FRUIncludesProduct struct {
	cim.WmiInstance

	//
	Component CIM_Product

	//
	FRU CIM_FRU
}

// SetComponent sets the value of Component for the instance
func (instance *CIM_FRUIncludesProduct) SetPropertyComponent(value CIM_Product) (err error) {
	return instance.SetProperty("Component", value)
}

// GetComponent gets the value of Component for the instance
func (instance *CIM_FRUIncludesProduct) GetPropertyComponent() (value CIM_Product, err error) {
	retValue, err := instance.GetProperty("Component")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Product)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFRU sets the value of FRU for the instance
func (instance *CIM_FRUIncludesProduct) SetPropertyFRU(value CIM_FRU) (err error) {
	return instance.SetProperty("FRU", value)
}

// GetFRU gets the value of FRU for the instance
func (instance *CIM_FRUIncludesProduct) GetPropertyFRU() (value CIM_FRU, err error) {
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
