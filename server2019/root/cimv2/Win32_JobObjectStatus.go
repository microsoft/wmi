// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_JobObjectStatus struct
type Win32_JobObjectStatus struct {
	__ExtendedStatus

	//
	AdditionalDescription string

	//
	Win32ErrorCode uint32
}

// SetAdditionalDescription sets the value of AdditionalDescription for the instance
func (instance *Win32_JobObjectStatus) SetPropertyAdditionalDescription(value string) (err error) {
	return instance.SetProperty("AdditionalDescription", value)
}

// GetAdditionalDescription gets the value of AdditionalDescription for the instance
func (instance *Win32_JobObjectStatus) GetPropertyAdditionalDescription() (value string, err error) {
	retValue, err := instance.GetProperty("AdditionalDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWin32ErrorCode sets the value of Win32ErrorCode for the instance
func (instance *Win32_JobObjectStatus) SetPropertyWin32ErrorCode(value uint32) (err error) {
	return instance.SetProperty("Win32ErrorCode", value)
}

// GetWin32ErrorCode gets the value of Win32ErrorCode for the instance
func (instance *Win32_JobObjectStatus) GetPropertyWin32ErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("Win32ErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
