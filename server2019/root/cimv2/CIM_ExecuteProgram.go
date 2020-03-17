// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_ExecuteProgram struct
type CIM_ExecuteProgram struct {
	CIM_Action

	//
	CommandLine string

	//
	ProgramPath string
}

// SetCommandLine sets the value of CommandLine for the instance
func (instance *CIM_ExecuteProgram) SetPropertyCommandLine(value string) (err error) {
	return instance.SetProperty("CommandLine", value)
}

// GetCommandLine gets the value of CommandLine for the instance
func (instance *CIM_ExecuteProgram) GetPropertyCommandLine() (value string, err error) {
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

// SetProgramPath sets the value of ProgramPath for the instance
func (instance *CIM_ExecuteProgram) SetPropertyProgramPath(value string) (err error) {
	return instance.SetProperty("ProgramPath", value)
}

// GetProgramPath gets the value of ProgramPath for the instance
func (instance *CIM_ExecuteProgram) GetPropertyProgramPath() (value string, err error) {
	retValue, err := instance.GetProperty("ProgramPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
