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

// MDM_WindowsAdvancedThreatProtection_DeviceTagging01 struct
type MDM_WindowsAdvancedThreatProtection_DeviceTagging01 struct {
	cim.WmiInstance

	//
	Criticality int32

	//
	Group string

	//
	IdMethod int32

	//
	InstanceID string

	//
	ParentID string
}

// SetCriticality sets the value of Criticality for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) SetPropertyCriticality(value int32) (err error) {
	return instance.SetProperty("Criticality", value)
}

// GetCriticality gets the value of Criticality for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) GetPropertyCriticality() (value int32, err error) {
	retValue, err := instance.GetProperty("Criticality")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGroup sets the value of Group for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) SetPropertyGroup(value string) (err error) {
	return instance.SetProperty("Group", value)
}

// GetGroup gets the value of Group for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) GetPropertyGroup() (value string, err error) {
	retValue, err := instance.GetProperty("Group")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdMethod sets the value of IdMethod for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) SetPropertyIdMethod(value int32) (err error) {
	return instance.SetProperty("IdMethod", value)
}

// GetIdMethod gets the value of IdMethod for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) GetPropertyIdMethod() (value int32, err error) {
	retValue, err := instance.GetProperty("IdMethod")
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
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection_DeviceTagging01) GetPropertyParentID() (value string, err error) {
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
