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

// Win32_ComClassAutoEmulator struct
type Win32_ComClassAutoEmulator struct {
	cim.WmiInstance

	//
	NewVersion Win32_ClassicCOMClass

	//
	OldVersion Win32_ClassicCOMClass
}

// SetNewVersion sets the value of NewVersion for the instance
func (instance *Win32_ComClassAutoEmulator) SetPropertyNewVersion(value Win32_ClassicCOMClass) (err error) {
	return instance.SetProperty("NewVersion", value)
}

// GetNewVersion gets the value of NewVersion for the instance
func (instance *Win32_ComClassAutoEmulator) GetPropertyNewVersion() (value Win32_ClassicCOMClass, err error) {
	retValue, err := instance.GetProperty("NewVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_ClassicCOMClass)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOldVersion sets the value of OldVersion for the instance
func (instance *Win32_ComClassAutoEmulator) SetPropertyOldVersion(value Win32_ClassicCOMClass) (err error) {
	return instance.SetProperty("OldVersion", value)
}

// GetOldVersion gets the value of OldVersion for the instance
func (instance *Win32_ComClassAutoEmulator) GetPropertyOldVersion() (value Win32_ClassicCOMClass, err error) {
	retValue, err := instance.GetProperty("OldVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_ClassicCOMClass)
	if !ok {
		// TODO: Set an error
	}
	return
}
