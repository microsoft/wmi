// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TerminalSetting struct
type Win32_TerminalSetting struct {
	CIM_Setting

	//
	TerminalName string
}

// SetTerminalName sets the value of TerminalName for the instance
func (instance *Win32_TerminalSetting) SetPropertyTerminalName(value string) (err error) {
	return instance.SetProperty("TerminalName", value)
}

// GetTerminalName gets the value of TerminalName for the instance
func (instance *Win32_TerminalSetting) GetPropertyTerminalName() (value string, err error) {
	retValue, err := instance.GetProperty("TerminalName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
