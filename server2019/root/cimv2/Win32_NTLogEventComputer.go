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

// Win32_NTLogEventComputer struct
type Win32_NTLogEventComputer struct {
	cim.WmiInstance

	//
	Computer Win32_ComputerSystem

	//
	Record Win32_NTLogEvent
}

// SetComputer sets the value of Computer for the instance
func (instance *Win32_NTLogEventComputer) SetPropertyComputer(value Win32_ComputerSystem) (err error) {
	return instance.SetProperty("Computer", value)
}

// GetComputer gets the value of Computer for the instance
func (instance *Win32_NTLogEventComputer) GetPropertyComputer() (value Win32_ComputerSystem, err error) {
	retValue, err := instance.GetProperty("Computer")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_ComputerSystem)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecord sets the value of Record for the instance
func (instance *Win32_NTLogEventComputer) SetPropertyRecord(value Win32_NTLogEvent) (err error) {
	return instance.SetProperty("Record", value)
}

// GetRecord gets the value of Record for the instance
func (instance *Win32_NTLogEventComputer) GetPropertyRecord() (value Win32_NTLogEvent, err error) {
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
