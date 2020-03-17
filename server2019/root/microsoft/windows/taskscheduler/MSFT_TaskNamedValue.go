// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_TaskNamedValue struct
type MSFT_TaskNamedValue struct {
	cim.WmiInstance

	//
	Name string

	//
	Value string
}

// SetName sets the value of Name for the instance
func (instance *MSFT_TaskNamedValue) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_TaskNamedValue) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValue sets the value of Value for the instance
func (instance *MSFT_TaskNamedValue) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *MSFT_TaskNamedValue) GetPropertyValue() (value string, err error) {
	retValue, err := instance.GetProperty("Value")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
