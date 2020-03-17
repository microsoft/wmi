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

// MDM_Policy_User_Config01_Desktop02 struct
type MDM_Policy_User_Config01_Desktop02 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	PreventUserRedirectionOfProfileFolders string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_Desktop02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_Desktop02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_User_Config01_Desktop02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Config01_Desktop02) GetPropertyParentID() (value string, err error) {
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

// SetPreventUserRedirectionOfProfileFolders sets the value of PreventUserRedirectionOfProfileFolders for the instance
func (instance *MDM_Policy_User_Config01_Desktop02) SetPropertyPreventUserRedirectionOfProfileFolders(value string) (err error) {
	return instance.SetProperty("PreventUserRedirectionOfProfileFolders", value)
}

// GetPreventUserRedirectionOfProfileFolders gets the value of PreventUserRedirectionOfProfileFolders for the instance
func (instance *MDM_Policy_User_Config01_Desktop02) GetPropertyPreventUserRedirectionOfProfileFolders() (value string, err error) {
	retValue, err := instance.GetProperty("PreventUserRedirectionOfProfileFolders")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
