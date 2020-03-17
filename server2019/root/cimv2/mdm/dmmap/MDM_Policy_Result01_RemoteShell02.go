// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Result01_RemoteShell02 struct
type MDM_Policy_Result01_RemoteShell02 struct {
	cim.WmiInstance

	//
	AllowRemoteShellAccess string

	//
	InstanceID string

	//
	MaxConcurrentUsers string

	//
	ParentID string

	//
	SpecifyIdleTimeout string

	//
	SpecifyMaxMemory string

	//
	SpecifyMaxProcesses string

	//
	SpecifyMaxRemoteShells string

	//
	SpecifyShellTimeout string
}

// SetAllowRemoteShellAccess sets the value of AllowRemoteShellAccess for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) SetPropertyAllowRemoteShellAccess(value string) (err error) {
	return instance.SetProperty("AllowRemoteShellAccess", value)
}

// GetAllowRemoteShellAccess gets the value of AllowRemoteShellAccess for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) GetPropertyAllowRemoteShellAccess() (value string, err error) {
	retValue, err := instance.GetProperty("AllowRemoteShellAccess")
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
func (instance *MDM_Policy_Result01_RemoteShell02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) GetPropertyInstanceID() (value string, err error) {
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

// SetMaxConcurrentUsers sets the value of MaxConcurrentUsers for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) SetPropertyMaxConcurrentUsers(value string) (err error) {
	return instance.SetProperty("MaxConcurrentUsers", value)
}

// GetMaxConcurrentUsers gets the value of MaxConcurrentUsers for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) GetPropertyMaxConcurrentUsers() (value string, err error) {
	retValue, err := instance.GetProperty("MaxConcurrentUsers")
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
func (instance *MDM_Policy_Result01_RemoteShell02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) GetPropertyParentID() (value string, err error) {
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

// SetSpecifyIdleTimeout sets the value of SpecifyIdleTimeout for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) SetPropertySpecifyIdleTimeout(value string) (err error) {
	return instance.SetProperty("SpecifyIdleTimeout", value)
}

// GetSpecifyIdleTimeout gets the value of SpecifyIdleTimeout for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) GetPropertySpecifyIdleTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyIdleTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecifyMaxMemory sets the value of SpecifyMaxMemory for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) SetPropertySpecifyMaxMemory(value string) (err error) {
	return instance.SetProperty("SpecifyMaxMemory", value)
}

// GetSpecifyMaxMemory gets the value of SpecifyMaxMemory for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) GetPropertySpecifyMaxMemory() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyMaxMemory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecifyMaxProcesses sets the value of SpecifyMaxProcesses for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) SetPropertySpecifyMaxProcesses(value string) (err error) {
	return instance.SetProperty("SpecifyMaxProcesses", value)
}

// GetSpecifyMaxProcesses gets the value of SpecifyMaxProcesses for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) GetPropertySpecifyMaxProcesses() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyMaxProcesses")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecifyMaxRemoteShells sets the value of SpecifyMaxRemoteShells for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) SetPropertySpecifyMaxRemoteShells(value string) (err error) {
	return instance.SetProperty("SpecifyMaxRemoteShells", value)
}

// GetSpecifyMaxRemoteShells gets the value of SpecifyMaxRemoteShells for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) GetPropertySpecifyMaxRemoteShells() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyMaxRemoteShells")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecifyShellTimeout sets the value of SpecifyShellTimeout for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) SetPropertySpecifyShellTimeout(value string) (err error) {
	return instance.SetProperty("SpecifyShellTimeout", value)
}

// GetSpecifyShellTimeout gets the value of SpecifyShellTimeout for the instance
func (instance *MDM_Policy_Result01_RemoteShell02) GetPropertySpecifyShellTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyShellTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
