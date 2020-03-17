// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerNetworkAdapter struct
type MSFT_ServerNetworkAdapter struct {
	cim.WmiInstance

	//
	Addresses []string

	//
	ConnectionStatus uint16

	//
	DHCPEnabled bool

	//
	Name string
}

// SetAddresses sets the value of Addresses for the instance
func (instance *MSFT_ServerNetworkAdapter) SetPropertyAddresses(value []string) (err error) {
	return instance.SetProperty("Addresses", value)
}

// GetAddresses gets the value of Addresses for the instance
func (instance *MSFT_ServerNetworkAdapter) GetPropertyAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("Addresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionStatus sets the value of ConnectionStatus for the instance
func (instance *MSFT_ServerNetworkAdapter) SetPropertyConnectionStatus(value uint16) (err error) {
	return instance.SetProperty("ConnectionStatus", value)
}

// GetConnectionStatus gets the value of ConnectionStatus for the instance
func (instance *MSFT_ServerNetworkAdapter) GetPropertyConnectionStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("ConnectionStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDHCPEnabled sets the value of DHCPEnabled for the instance
func (instance *MSFT_ServerNetworkAdapter) SetPropertyDHCPEnabled(value bool) (err error) {
	return instance.SetProperty("DHCPEnabled", value)
}

// GetDHCPEnabled gets the value of DHCPEnabled for the instance
func (instance *MSFT_ServerNetworkAdapter) GetPropertyDHCPEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DHCPEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_ServerNetworkAdapter) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServerNetworkAdapter) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
