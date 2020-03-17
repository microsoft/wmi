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

// Win32_ProductCheck struct
type Win32_ProductCheck struct {
	cim.WmiInstance

	//
	Check CIM_Check

	//
	Product Win32_Product
}

// SetCheck sets the value of Check for the instance
func (instance *Win32_ProductCheck) SetPropertyCheck(value CIM_Check) (err error) {
	return instance.SetProperty("Check", value)
}

// GetCheck gets the value of Check for the instance
func (instance *Win32_ProductCheck) GetPropertyCheck() (value CIM_Check, err error) {
	retValue, err := instance.GetProperty("Check")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Check)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProduct sets the value of Product for the instance
func (instance *Win32_ProductCheck) SetPropertyProduct(value Win32_Product) (err error) {
	return instance.SetProperty("Product", value)
}

// GetProduct gets the value of Product for the instance
func (instance *Win32_ProductCheck) GetPropertyProduct() (value Win32_Product, err error) {
	retValue, err := instance.GetProperty("Product")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_Product)
	if !ok {
		// TODO: Set an error
	}
	return
}
