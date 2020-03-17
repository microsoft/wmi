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

// CIM_ProductPhysicalElements struct
type CIM_ProductPhysicalElements struct {
	cim.WmiInstance

	//
	Component CIM_PhysicalElement

	//
	Product CIM_Product
}

// SetComponent sets the value of Component for the instance
func (instance *CIM_ProductPhysicalElements) SetPropertyComponent(value CIM_PhysicalElement) (err error) {
	return instance.SetProperty("Component", value)
}

// GetComponent gets the value of Component for the instance
func (instance *CIM_ProductPhysicalElements) GetPropertyComponent() (value CIM_PhysicalElement, err error) {
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

// SetProduct sets the value of Product for the instance
func (instance *CIM_ProductPhysicalElements) SetPropertyProduct(value CIM_Product) (err error) {
	return instance.SetProperty("Product", value)
}

// GetProduct gets the value of Product for the instance
func (instance *CIM_ProductPhysicalElements) GetPropertyProduct() (value CIM_Product, err error) {
	retValue, err := instance.GetProperty("Product")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Product)
	if !ok {
		// TODO: Set an error
	}
	return
}
