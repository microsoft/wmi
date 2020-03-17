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

// VpnCommonConfig struct
type VpnCommonConfig struct {
	cim.WmiInstance

	//
	ConnectionStatus string

	//
	DnsSuffix string

	//
	Guid string

	//
	IdleDisconnectSeconds uint32

	//
	IsAutoTriggerEnabled bool

	//
	Name string

	//
	ProfileType string

	//
	ProvisioningAuthority string

	//
	Proxy VpnConnectionProxy

	//
	RememberCredential bool

	//
	Routes []MSFT_NetRoute

	//
	ServerAddress string

	//
	ServerList []VpnServerAddress

	//
	SplitTunneling bool

	//
	VpnTrigger VpnConnectionTrigger
}

// SetConnectionStatus sets the value of ConnectionStatus for the instance
func (instance *VpnCommonConfig) SetPropertyConnectionStatus(value string) (err error) {
	return instance.SetProperty("ConnectionStatus", value)
}

// GetConnectionStatus gets the value of ConnectionStatus for the instance
func (instance *VpnCommonConfig) GetPropertyConnectionStatus() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionStatus")
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
func (instance *VpnCommonConfig) SetPropertyDnsSuffix(value string) (err error) {
	return instance.SetProperty("DnsSuffix", value)
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *VpnCommonConfig) GetPropertyDnsSuffix() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuid sets the value of Guid for the instance
func (instance *VpnCommonConfig) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", value)
}

// GetGuid gets the value of Guid for the instance
func (instance *VpnCommonConfig) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdleDisconnectSeconds sets the value of IdleDisconnectSeconds for the instance
func (instance *VpnCommonConfig) SetPropertyIdleDisconnectSeconds(value uint32) (err error) {
	return instance.SetProperty("IdleDisconnectSeconds", value)
}

// GetIdleDisconnectSeconds gets the value of IdleDisconnectSeconds for the instance
func (instance *VpnCommonConfig) GetPropertyIdleDisconnectSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("IdleDisconnectSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsAutoTriggerEnabled sets the value of IsAutoTriggerEnabled for the instance
func (instance *VpnCommonConfig) SetPropertyIsAutoTriggerEnabled(value bool) (err error) {
	return instance.SetProperty("IsAutoTriggerEnabled", value)
}

// GetIsAutoTriggerEnabled gets the value of IsAutoTriggerEnabled for the instance
func (instance *VpnCommonConfig) GetPropertyIsAutoTriggerEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsAutoTriggerEnabled")
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
func (instance *VpnCommonConfig) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *VpnCommonConfig) GetPropertyName() (value string, err error) {
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

// SetProfileType sets the value of ProfileType for the instance
func (instance *VpnCommonConfig) SetPropertyProfileType(value string) (err error) {
	return instance.SetProperty("ProfileType", value)
}

// GetProfileType gets the value of ProfileType for the instance
func (instance *VpnCommonConfig) GetPropertyProfileType() (value string, err error) {
	retValue, err := instance.GetProperty("ProfileType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProvisioningAuthority sets the value of ProvisioningAuthority for the instance
func (instance *VpnCommonConfig) SetPropertyProvisioningAuthority(value string) (err error) {
	return instance.SetProperty("ProvisioningAuthority", value)
}

// GetProvisioningAuthority gets the value of ProvisioningAuthority for the instance
func (instance *VpnCommonConfig) GetPropertyProvisioningAuthority() (value string, err error) {
	retValue, err := instance.GetProperty("ProvisioningAuthority")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProxy sets the value of Proxy for the instance
func (instance *VpnCommonConfig) SetPropertyProxy(value VpnConnectionProxy) (err error) {
	return instance.SetProperty("Proxy", value)
}

// GetProxy gets the value of Proxy for the instance
func (instance *VpnCommonConfig) GetPropertyProxy() (value VpnConnectionProxy, err error) {
	retValue, err := instance.GetProperty("Proxy")
	if err != nil {
		return
	}
	value, ok := retValue.(VpnConnectionProxy)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRememberCredential sets the value of RememberCredential for the instance
func (instance *VpnCommonConfig) SetPropertyRememberCredential(value bool) (err error) {
	return instance.SetProperty("RememberCredential", value)
}

// GetRememberCredential gets the value of RememberCredential for the instance
func (instance *VpnCommonConfig) GetPropertyRememberCredential() (value bool, err error) {
	retValue, err := instance.GetProperty("RememberCredential")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoutes sets the value of Routes for the instance
func (instance *VpnCommonConfig) SetPropertyRoutes(value []MSFT_NetRoute) (err error) {
	return instance.SetProperty("Routes", value)
}

// GetRoutes gets the value of Routes for the instance
func (instance *VpnCommonConfig) GetPropertyRoutes() (value []MSFT_NetRoute, err error) {
	retValue, err := instance.GetProperty("Routes")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_NetRoute)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerAddress sets the value of ServerAddress for the instance
func (instance *VpnCommonConfig) SetPropertyServerAddress(value string) (err error) {
	return instance.SetProperty("ServerAddress", value)
}

// GetServerAddress gets the value of ServerAddress for the instance
func (instance *VpnCommonConfig) GetPropertyServerAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ServerAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerList sets the value of ServerList for the instance
func (instance *VpnCommonConfig) SetPropertyServerList(value []VpnServerAddress) (err error) {
	return instance.SetProperty("ServerList", value)
}

// GetServerList gets the value of ServerList for the instance
func (instance *VpnCommonConfig) GetPropertyServerList() (value []VpnServerAddress, err error) {
	retValue, err := instance.GetProperty("ServerList")
	if err != nil {
		return
	}
	value, ok := retValue.([]VpnServerAddress)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSplitTunneling sets the value of SplitTunneling for the instance
func (instance *VpnCommonConfig) SetPropertySplitTunneling(value bool) (err error) {
	return instance.SetProperty("SplitTunneling", value)
}

// GetSplitTunneling gets the value of SplitTunneling for the instance
func (instance *VpnCommonConfig) GetPropertySplitTunneling() (value bool, err error) {
	retValue, err := instance.GetProperty("SplitTunneling")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVpnTrigger sets the value of VpnTrigger for the instance
func (instance *VpnCommonConfig) SetPropertyVpnTrigger(value VpnConnectionTrigger) (err error) {
	return instance.SetProperty("VpnTrigger", value)
}

// GetVpnTrigger gets the value of VpnTrigger for the instance
func (instance *VpnCommonConfig) GetPropertyVpnTrigger() (value VpnConnectionTrigger, err error) {
	retValue, err := instance.GetProperty("VpnTrigger")
	if err != nil {
		return
	}
	value, ok := retValue.(VpnConnectionTrigger)
	if !ok {
		// TODO: Set an error
	}
	return
}
