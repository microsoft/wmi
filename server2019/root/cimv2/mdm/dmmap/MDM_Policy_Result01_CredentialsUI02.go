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

// MDM_Policy_Result01_CredentialsUI02 struct
type MDM_Policy_Result01_CredentialsUI02 struct {
	cim.WmiInstance

	//
	DisablePasswordReveal string

	//
	EnumerateAdministrators string

	//
	InstanceID string

	//
	ParentID string
}

// SetDisablePasswordReveal sets the value of DisablePasswordReveal for the instance
func (instance *MDM_Policy_Result01_CredentialsUI02) SetPropertyDisablePasswordReveal(value string) (err error) {
	return instance.SetProperty("DisablePasswordReveal", value)
}

// GetDisablePasswordReveal gets the value of DisablePasswordReveal for the instance
func (instance *MDM_Policy_Result01_CredentialsUI02) GetPropertyDisablePasswordReveal() (value string, err error) {
	retValue, err := instance.GetProperty("DisablePasswordReveal")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnumerateAdministrators sets the value of EnumerateAdministrators for the instance
func (instance *MDM_Policy_Result01_CredentialsUI02) SetPropertyEnumerateAdministrators(value string) (err error) {
	return instance.SetProperty("EnumerateAdministrators", value)
}

// GetEnumerateAdministrators gets the value of EnumerateAdministrators for the instance
func (instance *MDM_Policy_Result01_CredentialsUI02) GetPropertyEnumerateAdministrators() (value string, err error) {
	retValue, err := instance.GetProperty("EnumerateAdministrators")
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
func (instance *MDM_Policy_Result01_CredentialsUI02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_CredentialsUI02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_CredentialsUI02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_CredentialsUI02) GetPropertyParentID() (value string, err error) {
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
