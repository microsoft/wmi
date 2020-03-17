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

// MSFT_TaskPrincipal struct
type MSFT_TaskPrincipal struct {
	cim.WmiInstance

	//
	DisplayName string

	//
	GroupId string

	//
	Id string

	//
	LogonType TaskPrincipal_LogonType

	//
	RunLevel int32

	//
	UserId string
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGroupId sets the value of GroupId for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyGroupId(value string) (err error) {
	return instance.SetProperty("GroupId", value)
}

// GetGroupId gets the value of GroupId for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyGroupId() (value string, err error) {
	retValue, err := instance.GetProperty("GroupId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyId() (value string, err error) {
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

// SetLogonType sets the value of LogonType for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyLogonType(value TaskPrincipal_LogonType) (err error) {
	return instance.SetProperty("LogonType", value)
}

// GetLogonType gets the value of LogonType for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyLogonType() (value TaskPrincipal_LogonType, err error) {
	retValue, err := instance.GetProperty("LogonType")
	if err != nil {
		return
	}
	value, ok := retValue.(TaskPrincipal_LogonType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunLevel sets the value of RunLevel for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyRunLevel(value int32) (err error) {
	return instance.SetProperty("RunLevel", value)
}

// GetRunLevel gets the value of RunLevel for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyRunLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("RunLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserId sets the value of UserId for the instance
func (instance *MSFT_TaskPrincipal) SetPropertyUserId(value string) (err error) {
	return instance.SetProperty("UserId", value)
}

// GetUserId gets the value of UserId for the instance
func (instance *MSFT_TaskPrincipal) GetPropertyUserId() (value string, err error) {
	retValue, err := instance.GetProperty("UserId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
