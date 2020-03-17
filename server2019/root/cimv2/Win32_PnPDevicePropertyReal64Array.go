// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PnPDevicePropertyReal64Array struct
type Win32_PnPDevicePropertyReal64Array struct {
	Win32_PnPDeviceProperty

	//
	Data []float64
}

// SetData sets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyReal64Array) SetPropertyData(value []float64) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyReal64Array) GetPropertyData() (value []float64, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.([]float64)
	if !ok {
		// TODO: Set an error
	}
	return
}
