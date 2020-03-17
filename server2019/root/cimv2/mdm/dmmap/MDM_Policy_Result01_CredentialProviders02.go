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

// MDM_Policy_Result01_CredentialProviders02 struct
type MDM_Policy_Result01_CredentialProviders02 struct {
	cim.WmiInstance

	//
	AllowPINLogon string

	//
	BlockPicturePassword string

	//
	DisableAutomaticReDeploymentCredentials int32

	//
	InstanceID string

	//
	ParentID string
}

// SetAllowPINLogon sets the value of AllowPINLogon for the instance
func (instance *MDM_Policy_Result01_CredentialProviders02) SetPropertyAllowPINLogon(value string) (err error) {
	return instance.SetProperty("AllowPINLogon", value)
}

// GetAllowPINLogon gets the value of AllowPINLogon for the instance
func (instance *MDM_Policy_Result01_CredentialProviders02) GetPropertyAllowPINLogon() (value string, err error) {
	retValue, err := instance.GetProperty("AllowPINLogon")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBlockPicturePassword sets the value of BlockPicturePassword for the instance
func (instance *MDM_Policy_Result01_CredentialProviders02) SetPropertyBlockPicturePassword(value string) (err error) {
	return instance.SetProperty("BlockPicturePassword", value)
}

// GetBlockPicturePassword gets the value of BlockPicturePassword for the instance
func (instance *MDM_Policy_Result01_CredentialProviders02) GetPropertyBlockPicturePassword() (value string, err error) {
	retValue, err := instance.GetProperty("BlockPicturePassword")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableAutomaticReDeploymentCredentials sets the value of DisableAutomaticReDeploymentCredentials for the instance
func (instance *MDM_Policy_Result01_CredentialProviders02) SetPropertyDisableAutomaticReDeploymentCredentials(value int32) (err error) {
	return instance.SetProperty("DisableAutomaticReDeploymentCredentials", value)
}

// GetDisableAutomaticReDeploymentCredentials gets the value of DisableAutomaticReDeploymentCredentials for the instance
func (instance *MDM_Policy_Result01_CredentialProviders02) GetPropertyDisableAutomaticReDeploymentCredentials() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableAutomaticReDeploymentCredentials")
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
func (instance *MDM_Policy_Result01_CredentialProviders02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_CredentialProviders02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_CredentialProviders02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_CredentialProviders02) GetPropertyParentID() (value string, err error) {
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
