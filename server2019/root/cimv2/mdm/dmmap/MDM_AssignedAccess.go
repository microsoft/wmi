// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_AssignedAccess struct
type MDM_AssignedAccess struct {
	*cim.WmiInstance

	//
	Configuration string

	//
	InstanceID string

	//
	KioskModeApp string

	//
	ParentID string

	//
	ShellLauncher string
}

func NewMDM_AssignedAccessEx1(instance *cim.WmiInstance) (newInstance *MDM_AssignedAccess, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_AssignedAccess{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_AssignedAccessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_AssignedAccess, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_AssignedAccess{
		WmiInstance: tmp,
	}
	return
}

// SetConfiguration sets the value of Configuration for the instance
func (instance *MDM_AssignedAccess) SetPropertyConfiguration(value string) (err error) {
	return instance.SetProperty("Configuration", value)
}

// GetConfiguration gets the value of Configuration for the instance
func (instance *MDM_AssignedAccess) GetPropertyConfiguration() (value string, err error) {
	retValue, err := instance.GetProperty("Configuration")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_AssignedAccess) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_AssignedAccess) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKioskModeApp sets the value of KioskModeApp for the instance
func (instance *MDM_AssignedAccess) SetPropertyKioskModeApp(value string) (err error) {
	return instance.SetProperty("KioskModeApp", value)
}

// GetKioskModeApp gets the value of KioskModeApp for the instance
func (instance *MDM_AssignedAccess) GetPropertyKioskModeApp() (value string, err error) {
	retValue, err := instance.GetProperty("KioskModeApp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_AssignedAccess) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_AssignedAccess) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShellLauncher sets the value of ShellLauncher for the instance
func (instance *MDM_AssignedAccess) SetPropertyShellLauncher(value string) (err error) {
	return instance.SetProperty("ShellLauncher", value)
}

// GetShellLauncher gets the value of ShellLauncher for the instance
func (instance *MDM_AssignedAccess) GetPropertyShellLauncher() (value string, err error) {
	retValue, err := instance.GetProperty("ShellLauncher")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
