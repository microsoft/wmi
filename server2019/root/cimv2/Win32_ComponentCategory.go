// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ComponentCategory struct
type Win32_ComponentCategory struct {
	CIM_LogicalElement

	//
	CategoryId string
}

// SetCategoryId sets the value of CategoryId for the instance
func (instance *Win32_ComponentCategory) SetPropertyCategoryId(value string) (err error) {
	return instance.SetProperty("CategoryId", value)
}

// GetCategoryId gets the value of CategoryId for the instance
func (instance *Win32_ComponentCategory) GetPropertyCategoryId() (value string, err error) {
	retValue, err := instance.GetProperty("CategoryId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
