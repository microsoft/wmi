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

// Win32_InstalledProgramFramework struct
type Win32_InstalledProgramFramework struct {
	cim.WmiInstance

	//
	FrameworkName string

	//
	FrameworkPublisher string

	//
	FrameworkVersion string

	//
	FrameworkVersionActual string

	//
	IsPrivate bool

	//
	ProgramId string
}

// SetFrameworkName sets the value of FrameworkName for the instance
func (instance *Win32_InstalledProgramFramework) SetPropertyFrameworkName(value string) (err error) {
	return instance.SetProperty("FrameworkName", value)
}

// GetFrameworkName gets the value of FrameworkName for the instance
func (instance *Win32_InstalledProgramFramework) GetPropertyFrameworkName() (value string, err error) {
	retValue, err := instance.GetProperty("FrameworkName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFrameworkPublisher sets the value of FrameworkPublisher for the instance
func (instance *Win32_InstalledProgramFramework) SetPropertyFrameworkPublisher(value string) (err error) {
	return instance.SetProperty("FrameworkPublisher", value)
}

// GetFrameworkPublisher gets the value of FrameworkPublisher for the instance
func (instance *Win32_InstalledProgramFramework) GetPropertyFrameworkPublisher() (value string, err error) {
	retValue, err := instance.GetProperty("FrameworkPublisher")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFrameworkVersion sets the value of FrameworkVersion for the instance
func (instance *Win32_InstalledProgramFramework) SetPropertyFrameworkVersion(value string) (err error) {
	return instance.SetProperty("FrameworkVersion", value)
}

// GetFrameworkVersion gets the value of FrameworkVersion for the instance
func (instance *Win32_InstalledProgramFramework) GetPropertyFrameworkVersion() (value string, err error) {
	retValue, err := instance.GetProperty("FrameworkVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFrameworkVersionActual sets the value of FrameworkVersionActual for the instance
func (instance *Win32_InstalledProgramFramework) SetPropertyFrameworkVersionActual(value string) (err error) {
	return instance.SetProperty("FrameworkVersionActual", value)
}

// GetFrameworkVersionActual gets the value of FrameworkVersionActual for the instance
func (instance *Win32_InstalledProgramFramework) GetPropertyFrameworkVersionActual() (value string, err error) {
	retValue, err := instance.GetProperty("FrameworkVersionActual")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsPrivate sets the value of IsPrivate for the instance
func (instance *Win32_InstalledProgramFramework) SetPropertyIsPrivate(value bool) (err error) {
	return instance.SetProperty("IsPrivate", value)
}

// GetIsPrivate gets the value of IsPrivate for the instance
func (instance *Win32_InstalledProgramFramework) GetPropertyIsPrivate() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPrivate")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProgramId sets the value of ProgramId for the instance
func (instance *Win32_InstalledProgramFramework) SetPropertyProgramId(value string) (err error) {
	return instance.SetProperty("ProgramId", value)
}

// GetProgramId gets the value of ProgramId for the instance
func (instance *Win32_InstalledProgramFramework) GetPropertyProgramId() (value string, err error) {
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
