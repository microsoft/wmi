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

// MDM_EnterpriseModernAppManagement_ReleaseManagement03_01 struct
type MDM_EnterpriseModernAppManagement_ReleaseManagement03_01 struct {
	cim.WmiInstance

	//
	ChannelId string

	//
	InstanceID string

	//
	ParentID string

	//
	ReleaseManagementId string
}

// SetChannelId sets the value of ChannelId for the instance
func (instance *MDM_EnterpriseModernAppManagement_ReleaseManagement03_01) SetPropertyChannelId(value string) (err error) {
	return instance.SetProperty("ChannelId", value)
}

// GetChannelId gets the value of ChannelId for the instance
func (instance *MDM_EnterpriseModernAppManagement_ReleaseManagement03_01) GetPropertyChannelId() (value string, err error) {
	retValue, err := instance.GetProperty("ChannelId")
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
func (instance *MDM_EnterpriseModernAppManagement_ReleaseManagement03_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_ReleaseManagement03_01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_EnterpriseModernAppManagement_ReleaseManagement03_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_ReleaseManagement03_01) GetPropertyParentID() (value string, err error) {
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

// SetReleaseManagementId sets the value of ReleaseManagementId for the instance
func (instance *MDM_EnterpriseModernAppManagement_ReleaseManagement03_01) SetPropertyReleaseManagementId(value string) (err error) {
	return instance.SetProperty("ReleaseManagementId", value)
}

// GetReleaseManagementId gets the value of ReleaseManagementId for the instance
func (instance *MDM_EnterpriseModernAppManagement_ReleaseManagement03_01) GetPropertyReleaseManagementId() (value string, err error) {
	retValue, err := instance.GetProperty("ReleaseManagementId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
