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

// MDM_EnterpriseModernAppManagement_AppSettingPolicy04 struct
type MDM_EnterpriseModernAppManagement_AppSettingPolicy04 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	IsVariableLeaf bool

	//
	ParentID string

	//
	Value string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppSettingPolicy04) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppSettingPolicy04) GetPropertyInstanceID() (value string, err error) {
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

// SetIsVariableLeaf sets the value of IsVariableLeaf for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppSettingPolicy04) SetPropertyIsVariableLeaf(value bool) (err error) {
	return instance.SetProperty("IsVariableLeaf", value)
}

// GetIsVariableLeaf gets the value of IsVariableLeaf for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppSettingPolicy04) GetPropertyIsVariableLeaf() (value bool, err error) {
	retValue, err := instance.GetProperty("IsVariableLeaf")
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
func (instance *MDM_EnterpriseModernAppManagement_AppSettingPolicy04) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppSettingPolicy04) GetPropertyParentID() (value string, err error) {
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

// SetValue sets the value of Value for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppSettingPolicy04) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppSettingPolicy04) GetPropertyValue() (value string, err error) {
	retValue, err := instance.GetProperty("Value")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
