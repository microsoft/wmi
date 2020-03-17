// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PatchPackage struct
type Win32_PatchPackage struct {
	Win32_MSIResource

	//
	PatchID string

	//
	ProductCode string
}

// SetPatchID sets the value of PatchID for the instance
func (instance *Win32_PatchPackage) SetPropertyPatchID(value string) (err error) {
	return instance.SetProperty("PatchID", value)
}

// GetPatchID gets the value of PatchID for the instance
func (instance *Win32_PatchPackage) GetPropertyPatchID() (value string, err error) {
	retValue, err := instance.GetProperty("PatchID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProductCode sets the value of ProductCode for the instance
func (instance *Win32_PatchPackage) SetPropertyProductCode(value string) (err error) {
	return instance.SetProperty("ProductCode", value)
}

// GetProductCode gets the value of ProductCode for the instance
func (instance *Win32_PatchPackage) GetPropertyProductCode() (value string, err error) {
	retValue, err := instance.GetProperty("ProductCode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
