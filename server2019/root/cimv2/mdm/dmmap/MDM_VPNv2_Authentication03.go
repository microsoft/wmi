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

// MDM_VPNv2_Authentication03 struct
type MDM_VPNv2_Authentication03 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	MachineMethod string

	//
	ParentID string

	//
	UserMethod string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_VPNv2_Authentication03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_Authentication03) GetPropertyInstanceID() (value string, err error) {
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

// SetMachineMethod sets the value of MachineMethod for the instance
func (instance *MDM_VPNv2_Authentication03) SetPropertyMachineMethod(value string) (err error) {
	return instance.SetProperty("MachineMethod", value)
}

// GetMachineMethod gets the value of MachineMethod for the instance
func (instance *MDM_VPNv2_Authentication03) GetPropertyMachineMethod() (value string, err error) {
	retValue, err := instance.GetProperty("MachineMethod")
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
func (instance *MDM_VPNv2_Authentication03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_Authentication03) GetPropertyParentID() (value string, err error) {
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

// SetUserMethod sets the value of UserMethod for the instance
func (instance *MDM_VPNv2_Authentication03) SetPropertyUserMethod(value string) (err error) {
	return instance.SetProperty("UserMethod", value)
}

// GetUserMethod gets the value of UserMethod for the instance
func (instance *MDM_VPNv2_Authentication03) GetPropertyUserMethod() (value string, err error) {
	retValue, err := instance.GetProperty("UserMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
