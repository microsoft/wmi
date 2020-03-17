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

// MDM_PassportForWork_DynamicLock01 struct
type MDM_PassportForWork_DynamicLock01 struct {
	cim.WmiInstance

	//
	DynamicLock bool

	//
	InstanceID string

	//
	ParentID string

	//
	Plugins string
}

// SetDynamicLock sets the value of DynamicLock for the instance
func (instance *MDM_PassportForWork_DynamicLock01) SetPropertyDynamicLock(value bool) (err error) {
	return instance.SetProperty("DynamicLock", value)
}

// GetDynamicLock gets the value of DynamicLock for the instance
func (instance *MDM_PassportForWork_DynamicLock01) GetPropertyDynamicLock() (value bool, err error) {
	retValue, err := instance.GetProperty("DynamicLock")
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
func (instance *MDM_PassportForWork_DynamicLock01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_PassportForWork_DynamicLock01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_PassportForWork_DynamicLock01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_PassportForWork_DynamicLock01) GetPropertyParentID() (value string, err error) {
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

// SetPlugins sets the value of Plugins for the instance
func (instance *MDM_PassportForWork_DynamicLock01) SetPropertyPlugins(value string) (err error) {
	return instance.SetProperty("Plugins", value)
}

// GetPlugins gets the value of Plugins for the instance
func (instance *MDM_PassportForWork_DynamicLock01) GetPropertyPlugins() (value string, err error) {
	retValue, err := instance.GetProperty("Plugins")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
