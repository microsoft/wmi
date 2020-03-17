// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_VolumeChangeEvent struct
type Win32_VolumeChangeEvent struct {
	Win32_DeviceChangeEvent

	//
	DriveName string
}

// SetDriveName sets the value of DriveName for the instance
func (instance *Win32_VolumeChangeEvent) SetPropertyDriveName(value string) (err error) {
	return instance.SetProperty("DriveName", value)
}

// GetDriveName gets the value of DriveName for the instance
func (instance *Win32_VolumeChangeEvent) GetPropertyDriveName() (value string, err error) {
	retValue, err := instance.GetProperty("DriveName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
