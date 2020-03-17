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

// MDM_DeviceStatus_Compliance01 struct
type MDM_DeviceStatus_Compliance01 struct {
	cim.WmiInstance

	//
	EncryptionCompliance bool

	//
	InstanceID string

	//
	ParentID string
}

// SetEncryptionCompliance sets the value of EncryptionCompliance for the instance
func (instance *MDM_DeviceStatus_Compliance01) SetPropertyEncryptionCompliance(value bool) (err error) {
	return instance.SetProperty("EncryptionCompliance", value)
}

// GetEncryptionCompliance gets the value of EncryptionCompliance for the instance
func (instance *MDM_DeviceStatus_Compliance01) GetPropertyEncryptionCompliance() (value bool, err error) {
	retValue, err := instance.GetProperty("EncryptionCompliance")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_Compliance01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_Compliance01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_DeviceStatus_Compliance01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_Compliance01) GetPropertyParentID() (value string, err error) {
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
