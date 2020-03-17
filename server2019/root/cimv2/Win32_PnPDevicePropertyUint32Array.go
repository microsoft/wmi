// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PnPDevicePropertyUint32Array struct
type Win32_PnPDevicePropertyUint32Array struct {
	Win32_PnPDeviceProperty

	//
	Data []uint32
}

// SetData sets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyUint32Array) SetPropertyData(value []uint32) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyUint32Array) GetPropertyData() (value []uint32, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
