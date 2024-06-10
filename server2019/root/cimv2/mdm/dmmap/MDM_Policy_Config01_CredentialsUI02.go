// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm.dmmap
//
// ////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Config01_CredentialsUI02 struct
type MDM_Policy_Config01_CredentialsUI02 struct {
	*cim.WmiInstance

	//
	DisablePasswordReveal string

	//
	EnumerateAdministrators string

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Config01_CredentialsUI02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_CredentialsUI02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_CredentialsUI02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_CredentialsUI02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_CredentialsUI02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_CredentialsUI02{
		WmiInstance: tmp,
	}
	return
}

// SetDisablePasswordReveal sets the value of DisablePasswordReveal for the instance
func (instance *MDM_Policy_Config01_CredentialsUI02) SetPropertyDisablePasswordReveal(value string) (err error) {
	return instance.SetProperty("DisablePasswordReveal", (value))
}

// GetDisablePasswordReveal gets the value of DisablePasswordReveal for the instance
func (instance *MDM_Policy_Config01_CredentialsUI02) GetPropertyDisablePasswordReveal() (value string, err error) {
	retValue, err := instance.GetProperty("DisablePasswordReveal")
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

// SetEnumerateAdministrators sets the value of EnumerateAdministrators for the instance
func (instance *MDM_Policy_Config01_CredentialsUI02) SetPropertyEnumerateAdministrators(value string) (err error) {
	return instance.SetProperty("EnumerateAdministrators", (value))
}

// GetEnumerateAdministrators gets the value of EnumerateAdministrators for the instance
func (instance *MDM_Policy_Config01_CredentialsUI02) GetPropertyEnumerateAdministrators() (value string, err error) {
	retValue, err := instance.GetProperty("EnumerateAdministrators")
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
func (instance *MDM_Policy_Config01_CredentialsUI02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_CredentialsUI02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_CredentialsUI02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_CredentialsUI02) GetPropertyParentID() (value string, err error) {
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
