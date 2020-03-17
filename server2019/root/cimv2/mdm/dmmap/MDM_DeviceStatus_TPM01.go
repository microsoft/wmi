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

// MDM_DeviceStatus_TPM01 struct
type MDM_DeviceStatus_TPM01 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	SpecificationVersion string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_TPM01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceStatus_TPM01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_DeviceStatus_TPM01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceStatus_TPM01) GetPropertyParentID() (value string, err error) {
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

// SetSpecificationVersion sets the value of SpecificationVersion for the instance
func (instance *MDM_DeviceStatus_TPM01) SetPropertySpecificationVersion(value string) (err error) {
	return instance.SetProperty("SpecificationVersion", value)
}

// GetSpecificationVersion gets the value of SpecificationVersion for the instance
func (instance *MDM_DeviceStatus_TPM01) GetPropertySpecificationVersion() (value string, err error) {
	retValue, err := instance.GetProperty("SpecificationVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
