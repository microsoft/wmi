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

// MDM_AppLocker_MSI03 struct
type MDM_AppLocker_MSI03 struct {
	cim.WmiInstance

	//
	EnforcementMode string

	//
	InstanceID string

	//
	ParentID string

	//
	Policy string
}

// SetEnforcementMode sets the value of EnforcementMode for the instance
func (instance *MDM_AppLocker_MSI03) SetPropertyEnforcementMode(value string) (err error) {
	return instance.SetProperty("EnforcementMode", value)
}

// GetEnforcementMode gets the value of EnforcementMode for the instance
func (instance *MDM_AppLocker_MSI03) GetPropertyEnforcementMode() (value string, err error) {
	retValue, err := instance.GetProperty("EnforcementMode")
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
func (instance *MDM_AppLocker_MSI03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_AppLocker_MSI03) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_AppLocker_MSI03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_AppLocker_MSI03) GetPropertyParentID() (value string, err error) {
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

// SetPolicy sets the value of Policy for the instance
func (instance *MDM_AppLocker_MSI03) SetPropertyPolicy(value string) (err error) {
	return instance.SetProperty("Policy", value)
}

// GetPolicy gets the value of Policy for the instance
func (instance *MDM_AppLocker_MSI03) GetPropertyPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("Policy")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
