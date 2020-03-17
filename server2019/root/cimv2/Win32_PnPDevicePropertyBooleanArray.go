// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PnPDevicePropertyBooleanArray struct
type Win32_PnPDevicePropertyBooleanArray struct {
	Win32_PnPDeviceProperty

	//
	Data []bool
}

// SetData sets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyBooleanArray) SetPropertyData(value []bool) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyBooleanArray) GetPropertyData() (value []bool, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.([]bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
