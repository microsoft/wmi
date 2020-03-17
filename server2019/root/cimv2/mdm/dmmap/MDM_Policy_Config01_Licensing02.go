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

// MDM_Policy_Config01_Licensing02 struct
type MDM_Policy_Config01_Licensing02 struct {
	cim.WmiInstance

	//
	AllowWindowsEntitlementReactivation int32

	//
	DisallowKMSClientOnlineAVSValidation int32

	//
	InstanceID string

	//
	ParentID string
}

// SetAllowWindowsEntitlementReactivation sets the value of AllowWindowsEntitlementReactivation for the instance
func (instance *MDM_Policy_Config01_Licensing02) SetPropertyAllowWindowsEntitlementReactivation(value int32) (err error) {
	return instance.SetProperty("AllowWindowsEntitlementReactivation", value)
}

// GetAllowWindowsEntitlementReactivation gets the value of AllowWindowsEntitlementReactivation for the instance
func (instance *MDM_Policy_Config01_Licensing02) GetPropertyAllowWindowsEntitlementReactivation() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsEntitlementReactivation")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisallowKMSClientOnlineAVSValidation sets the value of DisallowKMSClientOnlineAVSValidation for the instance
func (instance *MDM_Policy_Config01_Licensing02) SetPropertyDisallowKMSClientOnlineAVSValidation(value int32) (err error) {
	return instance.SetProperty("DisallowKMSClientOnlineAVSValidation", value)
}

// GetDisallowKMSClientOnlineAVSValidation gets the value of DisallowKMSClientOnlineAVSValidation for the instance
func (instance *MDM_Policy_Config01_Licensing02) GetPropertyDisallowKMSClientOnlineAVSValidation() (value int32, err error) {
	retValue, err := instance.GetProperty("DisallowKMSClientOnlineAVSValidation")
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
func (instance *MDM_Policy_Config01_Licensing02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Licensing02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_Licensing02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Licensing02) GetPropertyParentID() (value string, err error) {
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
