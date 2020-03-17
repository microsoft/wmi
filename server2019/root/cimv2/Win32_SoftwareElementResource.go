// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_SoftwareElementResource struct
type Win32_SoftwareElementResource struct {
	Win32_ManagedSystemElementResource

	//
	Element Win32_SoftwareElement

	//
	Setting Win32_MSIResource
}

// SetElement sets the value of Element for the instance
func (instance *Win32_SoftwareElementResource) SetPropertyElement(value Win32_SoftwareElement) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *Win32_SoftwareElementResource) GetPropertyElement() (value Win32_SoftwareElement, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_SoftwareElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetting sets the value of Setting for the instance
func (instance *Win32_SoftwareElementResource) SetPropertySetting(value Win32_MSIResource) (err error) {
	return instance.SetProperty("Setting", value)
}

// GetSetting gets the value of Setting for the instance
func (instance *Win32_SoftwareElementResource) GetPropertySetting() (value Win32_MSIResource, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_MSIResource)
	if !ok {
		// TODO: Set an error
	}
	return
}
