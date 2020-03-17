// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PnPDevicePropertySint16Array struct
type Win32_PnPDevicePropertySint16Array struct {
	Win32_PnPDeviceProperty

	//
	Data []int16
}

// SetData sets the value of Data for the instance
func (instance *Win32_PnPDevicePropertySint16Array) SetPropertyData(value []int16) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *Win32_PnPDevicePropertySint16Array) GetPropertyData() (value []int16, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.([]int16)
	if !ok {
		// TODO: Set an error
	}
	return
}
