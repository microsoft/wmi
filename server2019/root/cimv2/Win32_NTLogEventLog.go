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

// Win32_NTLogEventLog struct
type Win32_NTLogEventLog struct {
	cim.WmiInstance

	//
	Log Win32_NTEventlogFile

	//
	Record Win32_NTLogEvent
}

// SetLog sets the value of Log for the instance
func (instance *Win32_NTLogEventLog) SetPropertyLog(value Win32_NTEventlogFile) (err error) {
	return instance.SetProperty("Log", value)
}

// GetLog gets the value of Log for the instance
func (instance *Win32_NTLogEventLog) GetPropertyLog() (value Win32_NTEventlogFile, err error) {
	retValue, err := instance.GetProperty("Log")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_NTEventlogFile)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecord sets the value of Record for the instance
func (instance *Win32_NTLogEventLog) SetPropertyRecord(value Win32_NTLogEvent) (err error) {
	return instance.SetProperty("Record", value)
}

// GetRecord gets the value of Record for the instance
func (instance *Win32_NTLogEventLog) GetPropertyRecord() (value Win32_NTLogEvent, err error) {
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
