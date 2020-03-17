// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_LSM_UserInputDelayperProcess struct
type Win32_PerfRawData_LSM_UserInputDelayperProcess struct {
	Win32_PerfRawData

	//
	MaxInputDelay uint64

	//
	MaxInputDelay_Base uint32
}

// SetMaxInputDelay sets the value of MaxInputDelay for the instance
func (instance *Win32_PerfRawData_LSM_UserInputDelayperProcess) SetPropertyMaxInputDelay(value uint64) (err error) {
	return instance.SetProperty("MaxInputDelay", value)
}

// GetMaxInputDelay gets the value of MaxInputDelay for the instance
func (instance *Win32_PerfRawData_LSM_UserInputDelayperProcess) GetPropertyMaxInputDelay() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxInputDelay")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxInputDelay_Base sets the value of MaxInputDelay_Base for the instance
func (instance *Win32_PerfRawData_LSM_UserInputDelayperProcess) SetPropertyMaxInputDelay_Base(value uint32) (err error) {
	return instance.SetProperty("MaxInputDelay_Base", value)
}

// GetMaxInputDelay_Base gets the value of MaxInputDelay_Base for the instance
func (instance *Win32_PerfRawData_LSM_UserInputDelayperProcess) GetPropertyMaxInputDelay_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInputDelay_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
