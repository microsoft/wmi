// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// VpnConnectionTriggerTrustedNetwork struct
type VpnConnectionTriggerTrustedNetwork struct {
	cim.WmiInstance

	//
	ConnectionName string

	//
	DnsSuffix []string
}

// SetConnectionName sets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerTrustedNetwork) SetPropertyConnectionName(value string) (err error) {
	return instance.SetProperty("ConnectionName", value)
}

// GetConnectionName gets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerTrustedNetwork) GetPropertyConnectionName() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSuffix sets the value of DnsSuffix for the instance
func (instance *VpnConnectionTriggerTrustedNetwork) SetPropertyDnsSuffix(value []string) (err error) {
	return instance.SetProperty("DnsSuffix", value)
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *VpnConnectionTriggerTrustedNetwork) GetPropertyDnsSuffix() (value []string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
