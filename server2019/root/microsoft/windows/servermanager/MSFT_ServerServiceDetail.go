// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerServiceDetail struct
type MSFT_ServerServiceDetail struct {
	cim.WmiInstance

	//
	DependentServices []string

	//
	Description string

	//
	DisplayName string

	//
	ExitCode uint64

	//
	IsDelayedAutoStart bool

	//
	IsTriggered bool

	//
	Name string

	//
	StartupType uint32

	//
	Status uint32

	//
	SupportedControlCodes uint32
}

// SetDependentServices sets the value of DependentServices for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyDependentServices(value []string) (err error) {
	return instance.SetProperty("DependentServices", value)
}

// GetDependentServices gets the value of DependentServices for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyDependentServices() (value []string, err error) {
	retValue, err := instance.GetProperty("DependentServices")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExitCode sets the value of ExitCode for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyExitCode(value uint64) (err error) {
	return instance.SetProperty("ExitCode", value)
}

// GetExitCode gets the value of ExitCode for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyExitCode() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExitCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsDelayedAutoStart sets the value of IsDelayedAutoStart for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyIsDelayedAutoStart(value bool) (err error) {
	return instance.SetProperty("IsDelayedAutoStart", value)
}

// GetIsDelayedAutoStart gets the value of IsDelayedAutoStart for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyIsDelayedAutoStart() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDelayedAutoStart")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsTriggered sets the value of IsTriggered for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyIsTriggered(value bool) (err error) {
	return instance.SetProperty("IsTriggered", value)
}

// GetIsTriggered gets the value of IsTriggered for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyIsTriggered() (value bool, err error) {
	retValue, err := instance.GetProperty("IsTriggered")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyName() (value string, err error) {
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

// SetStartupType sets the value of StartupType for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyStartupType(value uint32) (err error) {
	return instance.SetProperty("StartupType", value)
}

// GetStartupType gets the value of StartupType for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyStartupType() (value uint32, err error) {
	retValue, err := instance.GetProperty("StartupType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedControlCodes sets the value of SupportedControlCodes for the instance
func (instance *MSFT_ServerServiceDetail) SetPropertySupportedControlCodes(value uint32) (err error) {
	return instance.SetProperty("SupportedControlCodes", value)
}

// GetSupportedControlCodes gets the value of SupportedControlCodes for the instance
func (instance *MSFT_ServerServiceDetail) GetPropertySupportedControlCodes() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportedControlCodes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
