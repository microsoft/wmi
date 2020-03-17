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

// MSFT_TaskAction struct
type MSFT_TaskAction struct {
	cim.WmiInstance

	//
	Id string
}

// SetId sets the value of Id for the instance
func (instance *MSFT_TaskAction) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_TaskAction) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
