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

// MDM_DeviceStatus struct
type MDM_DeviceStatus struct {
	cim.WmiInstance

	//
	DomainName string

	//
	InstanceID string

	//
	ParentID string

	//
	SecureBootState int32
}

// SetDomainName sets the value of DomainName for the instance
func (instance *MDM_DeviceStatus) SetPropertyDomainName(value string) (err error) {
	return instance.SetProperty("DomainName", value)
}

// GetDomainName gets the value of DomainName for the instance
func (instance *MDM_DeviceStatus) GetPropertyDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("DomainName")
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
func (instance *MDM_DeviceStatus) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_DeviceStatus) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceStatus) GetPropertyParentID() (value string, err error) {
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

// SetSecureBootState sets the value of SecureBootState for the instance
func (instance *MDM_DeviceStatus) SetPropertySecureBootState(value int32) (err error) {
	return instance.SetProperty("SecureBootState", value)
}

// GetSecureBootState gets the value of SecureBootState for the instance
func (instance *MDM_DeviceStatus) GetPropertySecureBootState() (value int32, err error) {
	retValue, err := instance.GetProperty("SecureBootState")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
