// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Bus struct
type Win32_Bus struct {
	CIM_LogicalDevice

	//
	BusNum uint32

	//
	BusType uint32
}

// SetBusNum sets the value of BusNum for the instance
func (instance *Win32_Bus) SetPropertyBusNum(value uint32) (err error) {
	return instance.SetProperty("BusNum", value)
}

// GetBusNum gets the value of BusNum for the instance
func (instance *Win32_Bus) GetPropertyBusNum() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusNum")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBusType sets the value of BusType for the instance
func (instance *Win32_Bus) SetPropertyBusType(value uint32) (err error) {
	return instance.SetProperty("BusType", value)
}

// GetBusType gets the value of BusType for the instance
func (instance *Win32_Bus) GetPropertyBusType() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
