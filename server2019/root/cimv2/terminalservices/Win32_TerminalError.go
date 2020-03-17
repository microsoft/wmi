// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TerminalError struct
type Win32_TerminalError struct {
	__ExtendedStatus

	//
	TerminalName string
}

// SetTerminalName sets the value of TerminalName for the instance
func (instance *Win32_TerminalError) SetPropertyTerminalName(value string) (err error) {
	return instance.SetProperty("TerminalName", value)
}

// GetTerminalName gets the value of TerminalName for the instance
func (instance *Win32_TerminalError) GetPropertyTerminalName() (value string, err error) {
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
