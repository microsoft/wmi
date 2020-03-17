// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ThreadTrace struct
type Win32_ThreadTrace struct {
	Win32_SystemTrace

	//
	ProcessID uint32

	//
	ThreadID uint32
}

// SetProcessID sets the value of ProcessID for the instance
func (instance *Win32_ThreadTrace) SetPropertyProcessID(value uint32) (err error) {
	return instance.SetProperty("ProcessID", value)
}

// GetProcessID gets the value of ProcessID for the instance
func (instance *Win32_ThreadTrace) GetPropertyProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreadID sets the value of ThreadID for the instance
func (instance *Win32_ThreadTrace) SetPropertyThreadID(value uint32) (err error) {
	return instance.SetProperty("ThreadID", value)
}

// GetThreadID gets the value of ThreadID for the instance
func (instance *Win32_ThreadTrace) GetPropertyThreadID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
