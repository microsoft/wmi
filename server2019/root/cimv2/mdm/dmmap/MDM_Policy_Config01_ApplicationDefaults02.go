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

// MDM_Policy_Config01_ApplicationDefaults02 struct
type MDM_Policy_Config01_ApplicationDefaults02 struct {
	cim.WmiInstance

	//
	DefaultAssociationsConfiguration string

	//
	EnableAppUriHandlers int32

	//
	InstanceID string

	//
	ParentID string
}

// SetDefaultAssociationsConfiguration sets the value of DefaultAssociationsConfiguration for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) SetPropertyDefaultAssociationsConfiguration(value string) (err error) {
	return instance.SetProperty("DefaultAssociationsConfiguration", value)
}

// GetDefaultAssociationsConfiguration gets the value of DefaultAssociationsConfiguration for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) GetPropertyDefaultAssociationsConfiguration() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultAssociationsConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableAppUriHandlers sets the value of EnableAppUriHandlers for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) SetPropertyEnableAppUriHandlers(value int32) (err error) {
	return instance.SetProperty("EnableAppUriHandlers", value)
}

// GetEnableAppUriHandlers gets the value of EnableAppUriHandlers for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) GetPropertyEnableAppUriHandlers() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableAppUriHandlers")
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
func (instance *MDM_Policy_Config01_ApplicationDefaults02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_ApplicationDefaults02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) GetPropertyParentID() (value string, err error) {
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
