// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_BrowserSecurityZones struct
type MDM_BrowserSecurityZones struct {
	cim.WmiInstance

	//
	Exists bool

	//
	Namespace string

	//
	Zone uint32
}

// SetExists sets the value of Exists for the instance
func (instance *MDM_BrowserSecurityZones) SetPropertyExists(value bool) (err error) {
	return instance.SetProperty("Exists", value)
}

// GetExists gets the value of Exists for the instance
func (instance *MDM_BrowserSecurityZones) GetPropertyExists() (value bool, err error) {
	retValue, err := instance.GetProperty("Exists")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNamespace sets the value of Namespace for the instance
func (instance *MDM_BrowserSecurityZones) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", value)
}

// GetNamespace gets the value of Namespace for the instance
func (instance *MDM_BrowserSecurityZones) GetPropertyNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("Namespace")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetZone sets the value of Zone for the instance
func (instance *MDM_BrowserSecurityZones) SetPropertyZone(value uint32) (err error) {
	return instance.SetProperty("Zone", value)
}

// GetZone gets the value of Zone for the instance
func (instance *MDM_BrowserSecurityZones) GetPropertyZone() (value uint32, err error) {
	retValue, err := instance.GetProperty("Zone")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
