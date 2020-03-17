// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TerminalServiceSettingError struct
type Win32_TerminalServiceSettingError struct {
	__ExtendedStatus

	//
	TerminalServiceMode int32
}

// SetTerminalServiceMode sets the value of TerminalServiceMode for the instance
func (instance *Win32_TerminalServiceSettingError) SetPropertyTerminalServiceMode(value int32) (err error) {
	return instance.SetProperty("TerminalServiceMode", value)
}

// GetTerminalServiceMode gets the value of TerminalServiceMode for the instance
func (instance *Win32_TerminalServiceSettingError) GetPropertyTerminalServiceMode() (value int32, err error) {
	retValue, err := instance.GetProperty("TerminalServiceMode")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
