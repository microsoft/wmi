// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	cimv2 "github.com/microsoft/wmi/server2019/root/cimv2"
)

// Win32_TSSessionDirectorySetting struct
type Win32_TSSessionDirectorySetting struct {
	CIM_ElementSetting

	//
	Element cimv2.Win32_TerminalService

	//
	Setting Win32_TSSessionDirectory
}

// SetElement sets the value of Element for the instance
func (instance *Win32_TSSessionDirectorySetting) SetPropertyElement(value cimv2.Win32_TerminalService) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *Win32_TSSessionDirectorySetting) GetPropertyElement() (value cimv2.Win32_TerminalService, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(cimv2.Win32_TerminalService)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetting sets the value of Setting for the instance
func (instance *Win32_TSSessionDirectorySetting) SetPropertySetting(value Win32_TSSessionDirectory) (err error) {
	return instance.SetProperty("Setting", value)
}

// GetSetting gets the value of Setting for the instance
func (instance *Win32_TSSessionDirectorySetting) GetPropertySetting() (value Win32_TSSessionDirectory, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_TSSessionDirectory)
	if !ok {
		// TODO: Set an error
	}
	return
}
