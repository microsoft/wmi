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

// Win32_ProductResource struct
type Win32_ProductResource struct {
	cim.WmiInstance

	//
	Product Win32_Product

	//
	Resource Win32_MSIResource
}

// SetProduct sets the value of Product for the instance
func (instance *Win32_ProductResource) SetPropertyProduct(value Win32_Product) (err error) {
	return instance.SetProperty("Product", value)
}

// GetProduct gets the value of Product for the instance
func (instance *Win32_ProductResource) GetPropertyProduct() (value Win32_Product, err error) {
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

// SetResource sets the value of Resource for the instance
func (instance *Win32_ProductResource) SetPropertyResource(value Win32_MSIResource) (err error) {
	return instance.SetProperty("Resource", value)
}

// GetResource gets the value of Resource for the instance
func (instance *Win32_ProductResource) GetPropertyResource() (value Win32_MSIResource, err error) {
	retValue, err := instance.GetProperty("Resource")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_MSIResource)
	if !ok {
		// TODO: Set an error
	}
	return
}
