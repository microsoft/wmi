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

// CIM_CompatibleProduct struct
type CIM_CompatibleProduct struct {
	cim.WmiInstance

	//
	CompatibilityDescription string

	//
	CompatibleProduct CIM_Product

	//
	Product CIM_Product
}

// SetCompatibilityDescription sets the value of CompatibilityDescription for the instance
func (instance *CIM_CompatibleProduct) SetPropertyCompatibilityDescription(value string) (err error) {
	return instance.SetProperty("CompatibilityDescription", value)
}

// GetCompatibilityDescription gets the value of CompatibilityDescription for the instance
func (instance *CIM_CompatibleProduct) GetPropertyCompatibilityDescription() (value string, err error) {
	retValue, err := instance.GetProperty("CompatibilityDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompatibleProduct sets the value of CompatibleProduct for the instance
func (instance *CIM_CompatibleProduct) SetPropertyCompatibleProduct(value CIM_Product) (err error) {
	return instance.SetProperty("CompatibleProduct", value)
}

// GetCompatibleProduct gets the value of CompatibleProduct for the instance
func (instance *CIM_CompatibleProduct) GetPropertyCompatibleProduct() (value CIM_Product, err error) {
	retValue, err := instance.GetProperty("CompatibleProduct")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Product)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProduct sets the value of Product for the instance
func (instance *CIM_CompatibleProduct) SetPropertyProduct(value CIM_Product) (err error) {
	return instance.SetProperty("Product", value)
}

// GetProduct gets the value of Product for the instance
func (instance *CIM_CompatibleProduct) GetPropertyProduct() (value CIM_Product, err error) {
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
