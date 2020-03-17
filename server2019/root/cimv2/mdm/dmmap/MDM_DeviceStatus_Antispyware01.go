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

// MDM_DeviceStatus_Antispyware01 struct
type MDM_DeviceStatus_Antispyware01 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	SignatureStatus int32

	//
	Status int32
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_Antispyware01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_Antispyware01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_DeviceStatus_Antispyware01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_Antispyware01) GetPropertyParentID() (value string, err error) {
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

// SetSignatureStatus sets the value of SignatureStatus for the instance
func (instance *MDM_DeviceStatus_Antispyware01) SetPropertySignatureStatus(value int32) (err error) {
	return instance.SetProperty("SignatureStatus", value)
}

// GetSignatureStatus gets the value of SignatureStatus for the instance
func (instance *MDM_DeviceStatus_Antispyware01) GetPropertySignatureStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("SignatureStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MDM_DeviceStatus_Antispyware01) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_DeviceStatus_Antispyware01) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
