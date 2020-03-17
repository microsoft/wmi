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

// Win32_TerminalServiceToSetting struct
type Win32_TerminalServiceToSetting struct {
	CIM_ElementSetting

	//
	Element cimv2.Win32_TerminalService

	//
	Setting Win32_TerminalServiceSetting
}

// SetElement sets the value of Element for the instance
func (instance *Win32_TerminalServiceToSetting) SetPropertyElement(value cimv2.Win32_TerminalService) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *Win32_TerminalServiceToSetting) GetPropertyElement() (value cimv2.Win32_TerminalService, err error) {
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
func (instance *Win32_TerminalServiceToSetting) SetPropertySetting(value Win32_TerminalServiceSetting) (err error) {
	return instance.SetProperty("Setting", value)
}

// GetSetting gets the value of Setting for the instance
func (instance *Win32_TerminalServiceToSetting) GetPropertySetting() (value Win32_TerminalServiceSetting, err error) {
	retValue, err := instance.GetProperty("Setting")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_TerminalServiceSetting)
	if !ok {
		// TODO: Set an error
	}
	return
}
