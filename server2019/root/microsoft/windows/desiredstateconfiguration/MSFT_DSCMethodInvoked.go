// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DSCMethodInvoked struct
type MSFT_DSCMethodInvoked struct {
	cim.WmiInstance

	//
	ConfigurationData []uint8

	//
	ConfigurationNumber uint8

	//
	Flags uint32

	//
	Guid string

	//
	MethodNumber uint8

	//
	ModuleName string

	//
	ResourceName string

	//
	UserSid string
}

// SetConfigurationData sets the value of ConfigurationData for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyConfigurationData(value []uint8) (err error) {
	return instance.SetProperty("ConfigurationData", value)
}

// GetConfigurationData gets the value of ConfigurationData for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyConfigurationData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("ConfigurationData")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigurationNumber sets the value of ConfigurationNumber for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyConfigurationNumber(value uint8) (err error) {
	return instance.SetProperty("ConfigurationNumber", value)
}

// GetConfigurationNumber gets the value of ConfigurationNumber for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyConfigurationNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("ConfigurationNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuid sets the value of Guid for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", value)
}

// GetGuid gets the value of Guid for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMethodNumber sets the value of MethodNumber for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyMethodNumber(value uint8) (err error) {
	return instance.SetProperty("MethodNumber", value)
}

// GetMethodNumber gets the value of MethodNumber for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyMethodNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("MethodNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModuleName sets the value of ModuleName for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyModuleName(value string) (err error) {
	return instance.SetProperty("ModuleName", value)
}

// GetModuleName gets the value of ModuleName for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyModuleName() (value string, err error) {
	retValue, err := instance.GetProperty("ModuleName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceName sets the value of ResourceName for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyResourceName(value string) (err error) {
	return instance.SetProperty("ResourceName", value)
}

// GetResourceName gets the value of ResourceName for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyResourceName() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserSid sets the value of UserSid for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyUserSid(value string) (err error) {
	return instance.SetProperty("UserSid", value)
}

// GetUserSid gets the value of UserSid for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyUserSid() (value string, err error) {
	retValue, err := instance.GetProperty("UserSid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
