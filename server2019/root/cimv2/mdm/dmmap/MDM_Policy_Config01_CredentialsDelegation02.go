// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MDM_Policy_Config01_CredentialsDelegation02 struct
type MDM_Policy_Config01_CredentialsDelegation02 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	RemoteHostAllowsDelegationOfNonExportableCredentials string
}

func NewMDM_Policy_Config01_CredentialsDelegation02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_CredentialsDelegation02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_CredentialsDelegation02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_CredentialsDelegation02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_CredentialsDelegation02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_CredentialsDelegation02{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_CredentialsDelegation02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_CredentialsDelegation02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_CredentialsDelegation02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_CredentialsDelegation02) GetPropertyParentID() (value string, err error) {
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

// SetRemoteHostAllowsDelegationOfNonExportableCredentials sets the value of RemoteHostAllowsDelegationOfNonExportableCredentials for the instance
func (instance *MDM_Policy_Config01_CredentialsDelegation02) SetPropertyRemoteHostAllowsDelegationOfNonExportableCredentials(value string) (err error) {
	return instance.SetProperty("RemoteHostAllowsDelegationOfNonExportableCredentials", (value))
}

// GetRemoteHostAllowsDelegationOfNonExportableCredentials gets the value of RemoteHostAllowsDelegationOfNonExportableCredentials for the instance
func (instance *MDM_Policy_Config01_CredentialsDelegation02) GetPropertyRemoteHostAllowsDelegationOfNonExportableCredentials() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteHostAllowsDelegationOfNonExportableCredentials")
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
