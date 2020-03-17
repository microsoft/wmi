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

// VpnConnectionTrigger struct
type VpnConnectionTrigger struct {
	cim.WmiInstance

	//
	ApplicationID []string

	//
	ConnectionName string

	//
	dnsConfig []VpnConnectionTriggerDnsConfiguration

	//
	DnsSuffixSearchList []string

	//
	TrustedNetwork []string
}

// SetApplicationID sets the value of ApplicationID for the instance
func (instance *VpnConnectionTrigger) SetPropertyApplicationID(value []string) (err error) {
	return instance.SetProperty("ApplicationID", value)
}

// GetApplicationID gets the value of ApplicationID for the instance
func (instance *VpnConnectionTrigger) GetPropertyApplicationID() (value []string, err error) {
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
func (instance *VpnConnectionTrigger) SetPropertyConnectionName(value string) (err error) {
	return instance.SetProperty("ConnectionName", value)
}

// GetConnectionName gets the value of ConnectionName for the instance
func (instance *VpnConnectionTrigger) GetPropertyConnectionName() (value string, err error) {
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

// SetdnsConfig sets the value of dnsConfig for the instance
func (instance *VpnConnectionTrigger) SetPropertydnsConfig(value []VpnConnectionTriggerDnsConfiguration) (err error) {
	return instance.SetProperty("dnsConfig", value)
}

// GetdnsConfig gets the value of dnsConfig for the instance
func (instance *VpnConnectionTrigger) GetPropertydnsConfig() (value []VpnConnectionTriggerDnsConfiguration, err error) {
	retValue, err := instance.GetProperty("dnsConfig")
	if err != nil {
		return
	}
	value, ok := retValue.([]VpnConnectionTriggerDnsConfiguration)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSuffixSearchList sets the value of DnsSuffixSearchList for the instance
func (instance *VpnConnectionTrigger) SetPropertyDnsSuffixSearchList(value []string) (err error) {
	return instance.SetProperty("DnsSuffixSearchList", value)
}

// GetDnsSuffixSearchList gets the value of DnsSuffixSearchList for the instance
func (instance *VpnConnectionTrigger) GetPropertyDnsSuffixSearchList() (value []string, err error) {
	retValue, err := instance.GetProperty("DnsSuffixSearchList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedNetwork sets the value of TrustedNetwork for the instance
func (instance *VpnConnectionTrigger) SetPropertyTrustedNetwork(value []string) (err error) {
	return instance.SetProperty("TrustedNetwork", value)
}

// GetTrustedNetwork gets the value of TrustedNetwork for the instance
func (instance *VpnConnectionTrigger) GetPropertyTrustedNetwork() (value []string, err error) {
	retValue, err := instance.GetProperty("TrustedNetwork")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
