// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Result01_RemoteDesktopServices02 struct
type MDM_Policy_Result01_RemoteDesktopServices02 struct {
	*cim.WmiInstance

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

func NewMDM_Policy_Result01_RemoteDesktopServices02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_RemoteDesktopServices02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_RemoteDesktopServices02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_RemoteDesktopServices02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_RemoteDesktopServices02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_RemoteDesktopServices02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowUsersToConnectRemotely sets the value of AllowUsersToConnectRemotely for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyAllowUsersToConnectRemotely(value string) (err error) {
	return instance.SetProperty("AllowUsersToConnectRemotely", (value))
}

// GetAllowUsersToConnectRemotely gets the value of AllowUsersToConnectRemotely for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyAllowUsersToConnectRemotely() (value string, err error) {
	retValue, err := instance.GetProperty("AllowUsersToConnectRemotely")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetClientConnectionEncryptionLevel sets the value of ClientConnectionEncryptionLevel for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyClientConnectionEncryptionLevel(value string) (err error) {
	return instance.SetProperty("ClientConnectionEncryptionLevel", (value))
}

// GetClientConnectionEncryptionLevel gets the value of ClientConnectionEncryptionLevel for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyClientConnectionEncryptionLevel() (value string, err error) {
	retValue, err := instance.GetProperty("ClientConnectionEncryptionLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDoNotAllowDriveRedirection sets the value of DoNotAllowDriveRedirection for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyDoNotAllowDriveRedirection(value string) (err error) {
	return instance.SetProperty("DoNotAllowDriveRedirection", (value))
}

// GetDoNotAllowDriveRedirection gets the value of DoNotAllowDriveRedirection for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyDoNotAllowDriveRedirection() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotAllowDriveRedirection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDoNotAllowPasswordSaving sets the value of DoNotAllowPasswordSaving for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyDoNotAllowPasswordSaving(value string) (err error) {
	return instance.SetProperty("DoNotAllowPasswordSaving", (value))
}

// GetDoNotAllowPasswordSaving gets the value of DoNotAllowPasswordSaving for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyDoNotAllowPasswordSaving() (value string, err error) {
	retValue, err := instance.GetProperty("DoNotAllowPasswordSaving")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPromptForPasswordUponConnection sets the value of PromptForPasswordUponConnection for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyPromptForPasswordUponConnection(value string) (err error) {
	return instance.SetProperty("PromptForPasswordUponConnection", (value))
}

// GetPromptForPasswordUponConnection gets the value of PromptForPasswordUponConnection for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyPromptForPasswordUponConnection() (value string, err error) {
	retValue, err := instance.GetProperty("PromptForPasswordUponConnection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetRequireSecureRPCCommunication sets the value of RequireSecureRPCCommunication for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) SetPropertyRequireSecureRPCCommunication(value string) (err error) {
	return instance.SetProperty("RequireSecureRPCCommunication", (value))
}

// GetRequireSecureRPCCommunication gets the value of RequireSecureRPCCommunication for the instance
func (instance *MDM_Policy_Result01_RemoteDesktopServices02) GetPropertyRequireSecureRPCCommunication() (value string, err error) {
	retValue, err := instance.GetProperty("RequireSecureRPCCommunication")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
