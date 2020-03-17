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

// Win32_OfflineFilesUserConfiguration struct
type Win32_OfflineFilesUserConfiguration struct {
	cim.WmiInstance

	//
	AssignedOfflineFiles []string

	//
	IsConfiguredByWMI bool

	//
	MakeAvailableOfflineButtonRemoved bool

	//
	WorkOfflineButtonRemoved bool
}

// SetAssignedOfflineFiles sets the value of AssignedOfflineFiles for the instance
func (instance *Win32_OfflineFilesUserConfiguration) SetPropertyAssignedOfflineFiles(value []string) (err error) {
	return instance.SetProperty("AssignedOfflineFiles", value)
}

// GetAssignedOfflineFiles gets the value of AssignedOfflineFiles for the instance
func (instance *Win32_OfflineFilesUserConfiguration) GetPropertyAssignedOfflineFiles() (value []string, err error) {
	retValue, err := instance.GetProperty("AssignedOfflineFiles")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsConfiguredByWMI sets the value of IsConfiguredByWMI for the instance
func (instance *Win32_OfflineFilesUserConfiguration) SetPropertyIsConfiguredByWMI(value bool) (err error) {
	return instance.SetProperty("IsConfiguredByWMI", value)
}

// GetIsConfiguredByWMI gets the value of IsConfiguredByWMI for the instance
func (instance *Win32_OfflineFilesUserConfiguration) GetPropertyIsConfiguredByWMI() (value bool, err error) {
	retValue, err := instance.GetProperty("IsConfiguredByWMI")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMakeAvailableOfflineButtonRemoved sets the value of MakeAvailableOfflineButtonRemoved for the instance
func (instance *Win32_OfflineFilesUserConfiguration) SetPropertyMakeAvailableOfflineButtonRemoved(value bool) (err error) {
	return instance.SetProperty("MakeAvailableOfflineButtonRemoved", value)
}

// GetMakeAvailableOfflineButtonRemoved gets the value of MakeAvailableOfflineButtonRemoved for the instance
func (instance *Win32_OfflineFilesUserConfiguration) GetPropertyMakeAvailableOfflineButtonRemoved() (value bool, err error) {
	retValue, err := instance.GetProperty("MakeAvailableOfflineButtonRemoved")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkOfflineButtonRemoved sets the value of WorkOfflineButtonRemoved for the instance
func (instance *Win32_OfflineFilesUserConfiguration) SetPropertyWorkOfflineButtonRemoved(value bool) (err error) {
	return instance.SetProperty("WorkOfflineButtonRemoved", value)
}

// GetWorkOfflineButtonRemoved gets the value of WorkOfflineButtonRemoved for the instance
func (instance *Win32_OfflineFilesUserConfiguration) GetPropertyWorkOfflineButtonRemoved() (value bool, err error) {
	retValue, err := instance.GetProperty("WorkOfflineButtonRemoved")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
