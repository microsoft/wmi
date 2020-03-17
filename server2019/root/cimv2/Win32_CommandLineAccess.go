// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_CommandLineAccess struct
type Win32_CommandLineAccess struct {
	CIM_ServiceAccessPoint

	//
	CommandLine string
}

// SetCommandLine sets the value of CommandLine for the instance
func (instance *Win32_CommandLineAccess) SetPropertyCommandLine(value string) (err error) {
	return instance.SetProperty("CommandLine", value)
}

// GetCommandLine gets the value of CommandLine for the instance
func (instance *Win32_CommandLineAccess) GetPropertyCommandLine() (value string, err error) {
	retValue, err := instance.GetProperty("CommandLine")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
