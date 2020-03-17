// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_NTLogEventUser struct
type Win32_NTLogEventUser struct {
	cim.WmiInstance

	//
	Record Win32_NTLogEvent

	//
	User Win32_UserAccount
}

// SetRecord sets the value of Record for the instance
func (instance *Win32_NTLogEventUser) SetPropertyRecord(value Win32_NTLogEvent) (err error) {
	return instance.SetProperty("Record", value)
}

// GetRecord gets the value of Record for the instance
func (instance *Win32_NTLogEventUser) GetPropertyRecord() (value Win32_NTLogEvent, err error) {
	retValue, err := instance.GetProperty("Record")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_NTLogEvent)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUser sets the value of User for the instance
func (instance *Win32_NTLogEventUser) SetPropertyUser(value Win32_UserAccount) (err error) {
	return instance.SetProperty("User", value)
}

// GetUser gets the value of User for the instance
func (instance *Win32_NTLogEventUser) GetPropertyUser() (value Win32_UserAccount, err error) {
	retValue, err := instance.GetProperty("User")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_UserAccount)
	if !ok {
		// TODO: Set an error
	}
	return
}
