// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_ScriptCmd struct
type RSOP_ScriptCmd struct {
	cim.WmiInstance

	//
	arguments string

	//
	executionTime string

	//
	script string
}

// Setarguments sets the value of arguments for the instance
func (instance *RSOP_ScriptCmd) SetPropertyarguments(value string) (err error) {
	return instance.SetProperty("arguments", value)
}

// Getarguments gets the value of arguments for the instance
func (instance *RSOP_ScriptCmd) GetPropertyarguments() (value string, err error) {
	retValue, err := instance.GetProperty("arguments")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetexecutionTime sets the value of executionTime for the instance
func (instance *RSOP_ScriptCmd) SetPropertyexecutionTime(value string) (err error) {
	return instance.SetProperty("executionTime", value)
}

// GetexecutionTime gets the value of executionTime for the instance
func (instance *RSOP_ScriptCmd) GetPropertyexecutionTime() (value string, err error) {
	retValue, err := instance.GetProperty("executionTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setscript sets the value of script for the instance
func (instance *RSOP_ScriptCmd) SetPropertyscript(value string) (err error) {
	return instance.SetProperty("script", value)
}

// Getscript gets the value of script for the instance
func (instance *RSOP_ScriptCmd) GetPropertyscript() (value string, err error) {
	retValue, err := instance.GetProperty("script")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
