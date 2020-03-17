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

// MDM_WindowsAdvancedThreatProtection struct
type MDM_WindowsAdvancedThreatProtection struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	Offboarding string

	//
	Onboarding string

	//
	ParentID string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection) GetPropertyInstanceID() (value string, err error) {
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

// SetOffboarding sets the value of Offboarding for the instance
func (instance *MDM_WindowsAdvancedThreatProtection) SetPropertyOffboarding(value string) (err error) {
	return instance.SetProperty("Offboarding", value)
}

// GetOffboarding gets the value of Offboarding for the instance
func (instance *MDM_WindowsAdvancedThreatProtection) GetPropertyOffboarding() (value string, err error) {
	retValue, err := instance.GetProperty("Offboarding")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOnboarding sets the value of Onboarding for the instance
func (instance *MDM_WindowsAdvancedThreatProtection) SetPropertyOnboarding(value string) (err error) {
	return instance.SetProperty("Onboarding", value)
}

// GetOnboarding gets the value of Onboarding for the instance
func (instance *MDM_WindowsAdvancedThreatProtection) GetPropertyOnboarding() (value string, err error) {
	retValue, err := instance.GetProperty("Onboarding")
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
func (instance *MDM_WindowsAdvancedThreatProtection) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_WindowsAdvancedThreatProtection) GetPropertyParentID() (value string, err error) {
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
