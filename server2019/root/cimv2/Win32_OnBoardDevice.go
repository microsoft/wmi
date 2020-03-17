// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_OnBoardDevice struct
type Win32_OnBoardDevice struct {
	CIM_PhysicalComponent

	//
	DeviceType uint16

	//
	Enabled bool
}

// SetDeviceType sets the value of DeviceType for the instance
func (instance *Win32_OnBoardDevice) SetPropertyDeviceType(value uint16) (err error) {
	return instance.SetProperty("DeviceType", value)
}

// GetDeviceType gets the value of DeviceType for the instance
func (instance *Win32_OnBoardDevice) GetPropertyDeviceType() (value uint16, err error) {
	retValue, err := instance.GetProperty("DeviceType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *Win32_OnBoardDevice) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *Win32_OnBoardDevice) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
