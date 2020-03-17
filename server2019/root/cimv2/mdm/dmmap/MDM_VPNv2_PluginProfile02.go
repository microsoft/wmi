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

// MDM_VPNv2_PluginProfile02 struct
type MDM_VPNv2_PluginProfile02 struct {
	cim.WmiInstance

	//
	CustomConfiguration string

	//
	CustomStoreUrl string

	//
	InstanceID string

	//
	ParentID string

	//
	PluginPackageFamilyName string

	//
	ServerUrlList string
}

// SetCustomConfiguration sets the value of CustomConfiguration for the instance
func (instance *MDM_VPNv2_PluginProfile02) SetPropertyCustomConfiguration(value string) (err error) {
	return instance.SetProperty("CustomConfiguration", value)
}

// GetCustomConfiguration gets the value of CustomConfiguration for the instance
func (instance *MDM_VPNv2_PluginProfile02) GetPropertyCustomConfiguration() (value string, err error) {
	retValue, err := instance.GetProperty("CustomConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCustomStoreUrl sets the value of CustomStoreUrl for the instance
func (instance *MDM_VPNv2_PluginProfile02) SetPropertyCustomStoreUrl(value string) (err error) {
	return instance.SetProperty("CustomStoreUrl", value)
}

// GetCustomStoreUrl gets the value of CustomStoreUrl for the instance
func (instance *MDM_VPNv2_PluginProfile02) GetPropertyCustomStoreUrl() (value string, err error) {
	retValue, err := instance.GetProperty("CustomStoreUrl")
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
func (instance *MDM_VPNv2_PluginProfile02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_PluginProfile02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_VPNv2_PluginProfile02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_PluginProfile02) GetPropertyParentID() (value string, err error) {
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

// SetPluginPackageFamilyName sets the value of PluginPackageFamilyName for the instance
func (instance *MDM_VPNv2_PluginProfile02) SetPropertyPluginPackageFamilyName(value string) (err error) {
	return instance.SetProperty("PluginPackageFamilyName", value)
}

// GetPluginPackageFamilyName gets the value of PluginPackageFamilyName for the instance
func (instance *MDM_VPNv2_PluginProfile02) GetPropertyPluginPackageFamilyName() (value string, err error) {
	retValue, err := instance.GetProperty("PluginPackageFamilyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerUrlList sets the value of ServerUrlList for the instance
func (instance *MDM_VPNv2_PluginProfile02) SetPropertyServerUrlList(value string) (err error) {
	return instance.SetProperty("ServerUrlList", value)
}

// GetServerUrlList gets the value of ServerUrlList for the instance
func (instance *MDM_VPNv2_PluginProfile02) GetPropertyServerUrlList() (value string, err error) {
	retValue, err := instance.GetProperty("ServerUrlList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
