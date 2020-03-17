// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Session struct
type Win32_Session struct {
	CIM_LogicalElement

	//
	StartTime string
}

// SetStartTime sets the value of StartTime for the instance
func (instance *Win32_Session) SetPropertyStartTime(value string) (err error) {
	return instance.SetProperty("StartTime", value)
}

// GetStartTime gets the value of StartTime for the instance
func (instance *Win32_Session) GetPropertyStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
