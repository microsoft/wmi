// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_BootConfiguration struct
type Win32_BootConfiguration struct {
	CIM_Setting

	//
	BootDirectory string

	//
	ConfigurationPath string

	//
	LastDrive string

	//
	Name string

	//
	ScratchDirectory string

	//
	TempDirectory string
}

// SetBootDirectory sets the value of BootDirectory for the instance
func (instance *Win32_BootConfiguration) SetPropertyBootDirectory(value string) (err error) {
	return instance.SetProperty("BootDirectory", value)
}

// GetBootDirectory gets the value of BootDirectory for the instance
func (instance *Win32_BootConfiguration) GetPropertyBootDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("BootDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigurationPath sets the value of ConfigurationPath for the instance
func (instance *Win32_BootConfiguration) SetPropertyConfigurationPath(value string) (err error) {
	return instance.SetProperty("ConfigurationPath", value)
}

// GetConfigurationPath gets the value of ConfigurationPath for the instance
func (instance *Win32_BootConfiguration) GetPropertyConfigurationPath() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigurationPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastDrive sets the value of LastDrive for the instance
func (instance *Win32_BootConfiguration) SetPropertyLastDrive(value string) (err error) {
	return instance.SetProperty("LastDrive", value)
}

// GetLastDrive gets the value of LastDrive for the instance
func (instance *Win32_BootConfiguration) GetPropertyLastDrive() (value string, err error) {
	retValue, err := instance.GetProperty("LastDrive")
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
func (instance *Win32_BootConfiguration) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Win32_BootConfiguration) GetPropertyName() (value string, err error) {
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

// SetScratchDirectory sets the value of ScratchDirectory for the instance
func (instance *Win32_BootConfiguration) SetPropertyScratchDirectory(value string) (err error) {
	return instance.SetProperty("ScratchDirectory", value)
}

// GetScratchDirectory gets the value of ScratchDirectory for the instance
func (instance *Win32_BootConfiguration) GetPropertyScratchDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("ScratchDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTempDirectory sets the value of TempDirectory for the instance
func (instance *Win32_BootConfiguration) SetPropertyTempDirectory(value string) (err error) {
	return instance.SetProperty("TempDirectory", value)
}

// GetTempDirectory gets the value of TempDirectory for the instance
func (instance *Win32_BootConfiguration) GetPropertyTempDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("TempDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
