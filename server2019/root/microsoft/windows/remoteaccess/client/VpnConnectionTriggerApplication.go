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

// VpnConnectionTriggerApplication struct
type VpnConnectionTriggerApplication struct {
	cim.WmiInstance

	//
	ApplicationID []string

	//
	ConnectionName string
}

// SetApplicationID sets the value of ApplicationID for the instance
func (instance *VpnConnectionTriggerApplication) SetPropertyApplicationID(value []string) (err error) {
	return instance.SetProperty("ApplicationID", value)
}

// GetApplicationID gets the value of ApplicationID for the instance
func (instance *VpnConnectionTriggerApplication) GetPropertyApplicationID() (value []string, err error) {
	retValue, err := instance.GetProperty("ApplicationID")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionName sets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerApplication) SetPropertyConnectionName(value string) (err error) {
	return instance.SetProperty("ConnectionName", value)
}

// GetConnectionName gets the value of ConnectionName for the instance
func (instance *VpnConnectionTriggerApplication) GetPropertyConnectionName() (value string, err error) {
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
