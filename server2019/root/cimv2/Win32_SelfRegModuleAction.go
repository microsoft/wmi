// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_SelfRegModuleAction struct
type Win32_SelfRegModuleAction struct {
	CIM_Action

	//
	Cost uint16

	//
	File string
}

// SetCost sets the value of Cost for the instance
func (instance *Win32_SelfRegModuleAction) SetPropertyCost(value uint16) (err error) {
	return instance.SetProperty("Cost", value)
}

// GetCost gets the value of Cost for the instance
func (instance *Win32_SelfRegModuleAction) GetPropertyCost() (value uint16, err error) {
	retValue, err := instance.GetProperty("Cost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFile sets the value of File for the instance
func (instance *Win32_SelfRegModuleAction) SetPropertyFile(value string) (err error) {
	return instance.SetProperty("File", value)
}

// GetFile gets the value of File for the instance
func (instance *Win32_SelfRegModuleAction) GetPropertyFile() (value string, err error) {
	retValue, err := instance.GetProperty("File")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
