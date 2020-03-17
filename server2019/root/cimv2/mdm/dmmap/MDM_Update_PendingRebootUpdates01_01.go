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

// MDM_Update_PendingRebootUpdates01_01 struct
type MDM_Update_PendingRebootUpdates01_01 struct {
	cim.WmiInstance

	//
	InstalledTime string

	//
	InstanceID string

	//
	ParentID string
}

// SetInstalledTime sets the value of InstalledTime for the instance
func (instance *MDM_Update_PendingRebootUpdates01_01) SetPropertyInstalledTime(value string) (err error) {
	return instance.SetProperty("InstalledTime", value)
}

// GetInstalledTime gets the value of InstalledTime for the instance
func (instance *MDM_Update_PendingRebootUpdates01_01) GetPropertyInstalledTime() (value string, err error) {
	retValue, err := instance.GetProperty("InstalledTime")
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
func (instance *MDM_Update_PendingRebootUpdates01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Update_PendingRebootUpdates01_01) GetPropertyInstanceID() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Update_PendingRebootUpdates01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Update_PendingRebootUpdates01_01) GetPropertyParentID() (value string, err error) {
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
