// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ClassicCOMClass struct
type Win32_ClassicCOMClass struct {
	Win32_COMClass

	//
	ComponentId string
}

// SetComponentId sets the value of ComponentId for the instance
func (instance *Win32_ClassicCOMClass) SetPropertyComponentId(value string) (err error) {
	return instance.SetProperty("ComponentId", value)
}

// GetComponentId gets the value of ComponentId for the instance
func (instance *Win32_ClassicCOMClass) GetPropertyComponentId() (value string, err error) {
	retValue, err := instance.GetProperty("ComponentId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
