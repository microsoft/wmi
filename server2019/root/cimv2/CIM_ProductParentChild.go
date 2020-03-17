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

// CIM_ProductParentChild struct
type CIM_ProductParentChild struct {
	cim.WmiInstance

	//
	Child CIM_Product

	//
	Parent CIM_Product
}

// SetChild sets the value of Child for the instance
func (instance *CIM_ProductParentChild) SetPropertyChild(value CIM_Product) (err error) {
	return instance.SetProperty("Child", value)
}

// GetChild gets the value of Child for the instance
func (instance *CIM_ProductParentChild) GetPropertyChild() (value CIM_Product, err error) {
	retValue, err := instance.GetProperty("Child")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Product)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParent sets the value of Parent for the instance
func (instance *CIM_ProductParentChild) SetPropertyParent(value CIM_Product) (err error) {
	return instance.SetProperty("Parent", value)
}

// GetParent gets the value of Parent for the instance
func (instance *CIM_ProductParentChild) GetPropertyParent() (value CIM_Product, err error) {
	retValue, err := instance.GetProperty("Parent")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Product)
	if !ok {
		// TODO: Set an error
	}
	return
}
