// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PnPDevicePropertySint64 struct
type Win32_PnPDevicePropertySint64 struct {
	Win32_PnPDeviceProperty

	//
	Data int64
}

// SetData sets the value of Data for the instance
func (instance *Win32_PnPDevicePropertySint64) SetPropertyData(value int64) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *Win32_PnPDevicePropertySint64) GetPropertyData() (value int64, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(int64)
	if !ok {
		// TODO: Set an error
	}
	return
}
