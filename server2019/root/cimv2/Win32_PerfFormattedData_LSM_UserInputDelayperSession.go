// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_LSM_UserInputDelayperSession struct
type Win32_PerfFormattedData_LSM_UserInputDelayperSession struct {
	Win32_PerfFormattedData

	//
	MaxInputDelay uint64
}

// SetMaxInputDelay sets the value of MaxInputDelay for the instance
func (instance *Win32_PerfFormattedData_LSM_UserInputDelayperSession) SetPropertyMaxInputDelay(value uint64) (err error) {
	return instance.SetProperty("MaxInputDelay", value)
}

// GetMaxInputDelay gets the value of MaxInputDelay for the instance
func (instance *Win32_PerfFormattedData_LSM_UserInputDelayperSession) GetPropertyMaxInputDelay() (value uint64, err error) {
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
