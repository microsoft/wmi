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

// MDM_Update_InstallableUpdates01_01 struct
type MDM_Update_InstallableUpdates01_01 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	RevisionNumber int32

	//
	Type int32
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Update_InstallableUpdates01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Update_InstallableUpdates01_01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Update_InstallableUpdates01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Update_InstallableUpdates01_01) GetPropertyParentID() (value string, err error) {
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

// SetRevisionNumber sets the value of RevisionNumber for the instance
func (instance *MDM_Update_InstallableUpdates01_01) SetPropertyRevisionNumber(value int32) (err error) {
	return instance.SetProperty("RevisionNumber", value)
}

// GetRevisionNumber gets the value of RevisionNumber for the instance
func (instance *MDM_Update_InstallableUpdates01_01) GetPropertyRevisionNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("RevisionNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MDM_Update_InstallableUpdates01_01) SetPropertyType(value int32) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MDM_Update_InstallableUpdates01_01) GetPropertyType() (value int32, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
