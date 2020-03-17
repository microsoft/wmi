// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_MountPoint struct
type Win32_MountPoint struct {
	cim.WmiInstance

	//
	Directory Win32_Directory

	//
	Volume Win32_Volume
}

// SetDirectory sets the value of Directory for the instance
func (instance *Win32_MountPoint) SetPropertyDirectory(value Win32_Directory) (err error) {
	return instance.SetProperty("Directory", value)
}

// GetDirectory gets the value of Directory for the instance
func (instance *Win32_MountPoint) GetPropertyDirectory() (value Win32_Directory, err error) {
	retValue, err := instance.GetProperty("Directory")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_Directory)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolume sets the value of Volume for the instance
func (instance *Win32_MountPoint) SetPropertyVolume(value Win32_Volume) (err error) {
	return instance.SetProperty("Volume", value)
}

// GetVolume gets the value of Volume for the instance
func (instance *Win32_MountPoint) GetPropertyVolume() (value Win32_Volume, err error) {
	retValue, err := instance.GetProperty("Volume")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_Volume)
	if !ok {
		// TODO: Set an error
	}
	return
}
