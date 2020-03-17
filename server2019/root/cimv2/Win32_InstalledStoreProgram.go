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

// Win32_InstalledStoreProgram struct
type Win32_InstalledStoreProgram struct {
	cim.WmiInstance

	//
	Architecture string

	//
	Language string

	//
	Name string

	//
	ProgramId string

	//
	Vendor string

	//
	Version string
}

// SetArchitecture sets the value of Architecture for the instance
func (instance *Win32_InstalledStoreProgram) SetPropertyArchitecture(value string) (err error) {
	return instance.SetProperty("Architecture", value)
}

// GetArchitecture gets the value of Architecture for the instance
func (instance *Win32_InstalledStoreProgram) GetPropertyArchitecture() (value string, err error) {
	retValue, err := instance.GetProperty("Architecture")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLanguage sets the value of Language for the instance
func (instance *Win32_InstalledStoreProgram) SetPropertyLanguage(value string) (err error) {
	return instance.SetProperty("Language", value)
}

// GetLanguage gets the value of Language for the instance
func (instance *Win32_InstalledStoreProgram) GetPropertyLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("Language")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *Win32_InstalledStoreProgram) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Win32_InstalledStoreProgram) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProgramId sets the value of ProgramId for the instance
func (instance *Win32_InstalledStoreProgram) SetPropertyProgramId(value string) (err error) {
	return instance.SetProperty("ProgramId", value)
}

// GetProgramId gets the value of ProgramId for the instance
func (instance *Win32_InstalledStoreProgram) GetPropertyProgramId() (value string, err error) {
	retValue, err := instance.GetProperty("ProgramId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVendor sets the value of Vendor for the instance
func (instance *Win32_InstalledStoreProgram) SetPropertyVendor(value string) (err error) {
	return instance.SetProperty("Vendor", value)
}

// GetVendor gets the value of Vendor for the instance
func (instance *Win32_InstalledStoreProgram) GetPropertyVendor() (value string, err error) {
	retValue, err := instance.GetProperty("Vendor")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *Win32_InstalledStoreProgram) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *Win32_InstalledStoreProgram) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
