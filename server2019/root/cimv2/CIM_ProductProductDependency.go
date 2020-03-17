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

// CIM_ProductProductDependency struct
type CIM_ProductProductDependency struct {
	cim.WmiInstance

	//
	DependentProduct CIM_Product

	//
	RequiredProduct CIM_Product

	//
	TypeOfDependency uint16
}

// SetDependentProduct sets the value of DependentProduct for the instance
func (instance *CIM_ProductProductDependency) SetPropertyDependentProduct(value CIM_Product) (err error) {
	return instance.SetProperty("DependentProduct", value)
}

// GetDependentProduct gets the value of DependentProduct for the instance
func (instance *CIM_ProductProductDependency) GetPropertyDependentProduct() (value CIM_Product, err error) {
	retValue, err := instance.GetProperty("DependentProduct")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Product)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequiredProduct sets the value of RequiredProduct for the instance
func (instance *CIM_ProductProductDependency) SetPropertyRequiredProduct(value CIM_Product) (err error) {
	return instance.SetProperty("RequiredProduct", value)
}

// GetRequiredProduct gets the value of RequiredProduct for the instance
func (instance *CIM_ProductProductDependency) GetPropertyRequiredProduct() (value CIM_Product, err error) {
	retValue, err := instance.GetProperty("RequiredProduct")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Product)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTypeOfDependency sets the value of TypeOfDependency for the instance
func (instance *CIM_ProductProductDependency) SetPropertyTypeOfDependency(value uint16) (err error) {
	return instance.SetProperty("TypeOfDependency", value)
}

// GetTypeOfDependency gets the value of TypeOfDependency for the instance
func (instance *CIM_ProductProductDependency) GetPropertyTypeOfDependency() (value uint16, err error) {
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
