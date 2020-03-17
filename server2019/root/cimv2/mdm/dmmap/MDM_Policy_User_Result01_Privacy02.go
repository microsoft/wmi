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

// MDM_Policy_User_Result01_Privacy02 struct
type MDM_Policy_User_Result01_Privacy02 struct {
	cim.WmiInstance

	//
	DisablePrivacyExperience int32

	//
	InstanceID string

	//
	ParentID string
}

// SetDisablePrivacyExperience sets the value of DisablePrivacyExperience for the instance
func (instance *MDM_Policy_User_Result01_Privacy02) SetPropertyDisablePrivacyExperience(value int32) (err error) {
	return instance.SetProperty("DisablePrivacyExperience", value)
}

// GetDisablePrivacyExperience gets the value of DisablePrivacyExperience for the instance
func (instance *MDM_Policy_User_Result01_Privacy02) GetPropertyDisablePrivacyExperience() (value int32, err error) {
	retValue, err := instance.GetProperty("DisablePrivacyExperience")
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
func (instance *MDM_Policy_User_Result01_Privacy02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_Privacy02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_User_Result01_Privacy02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Result01_Privacy02) GetPropertyParentID() (value string, err error) {
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
