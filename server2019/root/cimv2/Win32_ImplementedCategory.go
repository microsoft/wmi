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

// Win32_ImplementedCategory struct
type Win32_ImplementedCategory struct {
	cim.WmiInstance

	//
	Category Win32_ComponentCategory

	//
	Component Win32_ClassicCOMClass
}

// SetCategory sets the value of Category for the instance
func (instance *Win32_ImplementedCategory) SetPropertyCategory(value Win32_ComponentCategory) (err error) {
	return instance.SetProperty("Category", value)
}

// GetCategory gets the value of Category for the instance
func (instance *Win32_ImplementedCategory) GetPropertyCategory() (value Win32_ComponentCategory, err error) {
	retValue, err := instance.GetProperty("Category")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_ComponentCategory)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComponent sets the value of Component for the instance
func (instance *Win32_ImplementedCategory) SetPropertyComponent(value Win32_ClassicCOMClass) (err error) {
	return instance.SetProperty("Component", value)
}

// GetComponent gets the value of Component for the instance
func (instance *Win32_ImplementedCategory) GetPropertyComponent() (value Win32_ClassicCOMClass, err error) {
	retValue, err := instance.GetProperty("Component")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_ClassicCOMClass)
	if !ok {
		// TODO: Set an error
	}
	return
}
