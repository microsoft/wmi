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

// MDM_Policy_Result01_RemoteDesktopServices02 struct
type MDM_Policy_Result01_RemoteDesktopServices02 struct {
	cim.WmiInstance

	//
	AllowUsersToConnectRemotely string

	//
	ClientConnectionEncryptionLevel string

	//
	DoNotAllowDriveRedirection string

	//
	DoNotAllowPasswordSaving string

	//
	InstanceID string

	//
	ParentID string

	//
	PromptForPasswordUponConnection string

	//
	RequireSecureRPCCommunication string
}

// SetAllowUsersToConnectRemotely sets the value of AllowUsersToConnectRemotely for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyAllowUsersToConnectRemotely(value string) (err error) {
	return instance.SetProperty("AllowUsersToConnectRemotely", value)
}

// GetAllowUsersToConnectRemotely gets the value of AllowUsersToConnectRemotely for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyAllowUsersToConnectRemotely() (value string, err error) {
	retValue, err := instance.GetProperty("AllowUsersToConnectRemotely")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClientConnectionEncryptionLevel sets the value of ClientConnectionEncryptionLevel for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyClientConnectionEncryptionLevel(value string) (err error) {
	return instance.SetProperty("ClientConnectionEncryptionLevel", value)
}

// GetClientConnectionEncryptionLevel gets the value of ClientConnectionEncryptionLevel for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyClientConnectionEncryptionLevel() (value string, err error) {
	retValue, err := instance.GetProperty("ClientConnectionEncryptionLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDoNotAllowDriveRedirection sets the value of DoNotAllowDriveRedirection for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyDoNotAllowDriveRedirection(value string) (err error) {
	return instance.SetProperty("DoNotAllowDriveRedirection", value)
}

// GetDoNotAllowDriveRedirection gets the value of DoNotAllowDriveRedirection for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyDoNotAllowDriveRedirection() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotAllowDriveRedirection")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDoNotAllowPasswordSaving sets the value of DoNotAllowPasswordSaving for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyDoNotAllowPasswordSaving(value string) (err error) {
	return instance.SetProperty("DoNotAllowPasswordSaving", value)
}

// GetDoNotAllowPasswordSaving gets the value of DoNotAllowPasswordSaving for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyDoNotAllowPasswordSaving() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotAllowPasswordSaving")
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
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyParentID() (value string, err error) {
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

// SetPromptForPasswordUponConnection sets the value of PromptForPasswordUponConnection for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyPromptForPasswordUponConnection(value string) (err error) {
	return instance.SetProperty("PromptForPasswordUponConnection", value)
}

// GetPromptForPasswordUponConnection gets the value of PromptForPasswordUponConnection for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyPromptForPasswordUponConnection() (value string, err error) {
	retValue, err := instance.GetProperty("PromptForPasswordUponConnection")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireSecureRPCCommunication sets the value of RequireSecureRPCCommunication for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyRequireSecureRPCCommunication(value string) (err error) {
	return instance.SetProperty("RequireSecureRPCCommunication", value)
}

// GetRequireSecureRPCCommunication gets the value of RequireSecureRPCCommunication for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyRequireSecureRPCCommunication() (value string, err error) {
	retValue, err := instance.GetProperty("RequireSecureRPCCommunication")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
