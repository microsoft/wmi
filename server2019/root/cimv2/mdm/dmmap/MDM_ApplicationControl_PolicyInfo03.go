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

// MDM_ApplicationControl_PolicyInfo03 struct
type MDM_ApplicationControl_PolicyInfo03 struct {
	cim.WmiInstance

	//
	FriendlyName string

	//
	InstanceID string

	//
	IsAuthorized bool

	//
	IsDeployed bool

	//
	IsEffective bool

	//
	ParentID string

	//
	Status int32

	//
	Version string
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyInstanceID() (value string, err error) {
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

// SetIsAuthorized sets the value of IsAuthorized for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyIsAuthorized(value bool) (err error) {
	return instance.SetProperty("IsAuthorized", value)
}

// GetIsAuthorized gets the value of IsAuthorized for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyIsAuthorized() (value bool, err error) {
	retValue, err := instance.GetProperty("IsAuthorized")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsDeployed sets the value of IsDeployed for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyIsDeployed(value bool) (err error) {
	return instance.SetProperty("IsDeployed", value)
}

// GetIsDeployed gets the value of IsDeployed for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyIsDeployed() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDeployed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsEffective sets the value of IsEffective for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyIsEffective(value bool) (err error) {
	return instance.SetProperty("IsEffective", value)
}

// GetIsEffective gets the value of IsEffective for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyIsEffective() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEffective")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyParentID() (value string, err error) {
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

// SetStatus sets the value of Status for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyStatus() (value int32, err error) {
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

// SetVersion sets the value of Version for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
