// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ProcessStopTrace struct
type Win32_ProcessStopTrace struct {
	Win32_ProcessTrace

	//
	ExitStatus uint32
}

// SetExitStatus sets the value of ExitStatus for the instance
func (instance *Win32_ProcessStopTrace) SetPropertyExitStatus(value uint32) (err error) {
	return instance.SetProperty("ExitStatus", value)
}

// GetExitStatus gets the value of ExitStatus for the instance
func (instance *Win32_ProcessStopTrace) GetPropertyExitStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExitStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
