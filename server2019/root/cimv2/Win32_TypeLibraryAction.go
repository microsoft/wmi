// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_TypeLibraryAction struct
type Win32_TypeLibraryAction struct {
	CIM_Action

	//
	Cost uint32

	//
	Language uint16

	//
	LibID string
}

// SetCost sets the value of Cost for the instance
func (instance *Win32_TypeLibraryAction) SetPropertyCost(value uint32) (err error) {
	return instance.SetProperty("Cost", value)
}

// GetCost gets the value of Cost for the instance
func (instance *Win32_TypeLibraryAction) GetPropertyCost() (value uint32, err error) {
	retValue, err := instance.GetProperty("Cost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLanguage sets the value of Language for the instance
func (instance *Win32_TypeLibraryAction) SetPropertyLanguage(value uint16) (err error) {
	return instance.SetProperty("Language", value)
}

// GetLanguage gets the value of Language for the instance
func (instance *Win32_TypeLibraryAction) GetPropertyLanguage() (value uint16, err error) {
	retValue, err := instance.GetProperty("Language")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLibID sets the value of LibID for the instance
func (instance *Win32_TypeLibraryAction) SetPropertyLibID(value string) (err error) {
	return instance.SetProperty("LibID", value)
}

// GetLibID gets the value of LibID for the instance
func (instance *Win32_TypeLibraryAction) GetPropertyLibID() (value string, err error) {
	retValue, err := instance.GetProperty("LibID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
