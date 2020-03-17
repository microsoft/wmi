// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PnPDevicePropertySecurityDescriptor struct
type Win32_PnPDevicePropertySecurityDescriptor struct {
	Win32_PnPDeviceProperty

	//
	Data Win32_SecurityDescriptor
}

// SetData sets the value of Data for the instance
func (instance *Win32_PnPDevicePropertySecurityDescriptor) SetPropertyData(value Win32_SecurityDescriptor) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *Win32_PnPDevicePropertySecurityDescriptor) GetPropertyData() (value Win32_SecurityDescriptor, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_SecurityDescriptor)
	if !ok {
		// TODO: Set an error
	}
	return
}
