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

// MDM_Update_FailedUpdates01_01 struct
type MDM_Update_FailedUpdates01_01 struct {
	cim.WmiInstance

	//
	HResult int32

	//
	InstanceID string

	//
	ParentID string

	//
	State int32
}

// SetHResult sets the value of HResult for the instance
func (instance *MDM_Update_FailedUpdates01_01) SetPropertyHResult(value int32) (err error) {
	return instance.SetProperty("HResult", value)
}

// GetHResult gets the value of HResult for the instance
func (instance *MDM_Update_FailedUpdates01_01) GetPropertyHResult() (value int32, err error) {
	retValue, err := instance.GetProperty("HResult")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Update_FailedUpdates01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Update_FailedUpdates01_01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Update_FailedUpdates01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Update_FailedUpdates01_01) GetPropertyParentID() (value string, err error) {
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

// SetState sets the value of State for the instance
func (instance *MDM_Update_FailedUpdates01_01) SetPropertyState(value int32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MDM_Update_FailedUpdates01_01) GetPropertyState() (value int32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
