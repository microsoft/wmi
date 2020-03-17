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

// MDM_ApplicationControl struct
type MDM_ApplicationControl struct {
	cim.WmiInstance

	//
	DeviceID string

	//
	InstanceID string

	//
	ParentID string

	//
	TenantID string
}

// SetDeviceID sets the value of DeviceID for the instance
func (instance *MDM_ApplicationControl) SetPropertyDeviceID(value string) (err error) {
	return instance.SetProperty("DeviceID", value)
}

// GetDeviceID gets the value of DeviceID for the instance
func (instance *MDM_ApplicationControl) GetPropertyDeviceID() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceID")
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
func (instance *MDM_ApplicationControl) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_ApplicationControl) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_ApplicationControl) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_ApplicationControl) GetPropertyParentID() (value string, err error) {
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

// SetTenantID sets the value of TenantID for the instance
func (instance *MDM_ApplicationControl) SetPropertyTenantID(value string) (err error) {
	return instance.SetProperty("TenantID", value)
}

// GetTenantID gets the value of TenantID for the instance
func (instance *MDM_ApplicationControl) GetPropertyTenantID() (value string, err error) {
	retValue, err := instance.GetProperty("TenantID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
