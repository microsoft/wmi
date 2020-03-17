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

// CIM_ProductFRU struct
type CIM_ProductFRU struct {
	cim.WmiInstance

	//
	FRU CIM_FRU

	//
	Product CIM_Product
}

// SetFRU sets the value of FRU for the instance
func (instance *CIM_ProductFRU) SetPropertyFRU(value CIM_FRU) (err error) {
	return instance.SetProperty("FRU", value)
}

// GetFRU gets the value of FRU for the instance
func (instance *CIM_ProductFRU) GetPropertyFRU() (value CIM_FRU, err error) {
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

// SetProduct sets the value of Product for the instance
func (instance *CIM_ProductFRU) SetPropertyProduct(value CIM_Product) (err error) {
	return instance.SetProperty("Product", value)
}

// GetProduct gets the value of Product for the instance
func (instance *CIM_ProductFRU) GetPropertyProduct() (value CIM_Product, err error) {
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
