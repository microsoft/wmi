// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Client
//
// ////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnCommonConfig struct
type VpnCommonConfig struct {
	*cim.WmiInstance

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

func NewVpnCommonConfigEx1(instance *cim.WmiInstance) (newInstance *VpnCommonConfig, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnCommonConfig{
		WmiInstance: tmp,
	}
	return
}

func NewVpnCommonConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnCommonConfig, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnCommonConfig{
		WmiInstance: tmp,
	}
	return
}

// SetConnectionStatus sets the value of ConnectionStatus for the instance
func (instance *VpnCommonConfig) SetPropertyConnectionStatus(value string) (err error) {
	return instance.SetProperty("ConnectionStatus", (value))
}

// GetConnectionStatus gets the value of ConnectionStatus for the instance
func (instance *VpnCommonConfig) GetPropertyConnectionStatus() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDnsSuffix sets the value of DnsSuffix for the instance
func (instance *VpnCommonConfig) SetPropertyDnsSuffix(value string) (err error) {
	return instance.SetProperty("DnsSuffix", (value))
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *VpnCommonConfig) GetPropertyDnsSuffix() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetGuid sets the value of Guid for the instance
func (instance *VpnCommonConfig) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *VpnCommonConfig) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetIdleDisconnectSeconds sets the value of IdleDisconnectSeconds for the instance
func (instance *VpnCommonConfig) SetPropertyIdleDisconnectSeconds(value uint32) (err error) {
	return instance.SetProperty("IdleDisconnectSeconds", (value))
}

// GetIdleDisconnectSeconds gets the value of IdleDisconnectSeconds for the instance
func (instance *VpnCommonConfig) GetPropertyIdleDisconnectSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("IdleDisconnectSeconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetIsAutoTriggerEnabled sets the value of IsAutoTriggerEnabled for the instance
func (instance *VpnCommonConfig) SetPropertyIsAutoTriggerEnabled(value bool) (err error) {
	return instance.SetProperty("IsAutoTriggerEnabled", (value))
}

// GetIsAutoTriggerEnabled gets the value of IsAutoTriggerEnabled for the instance
func (instance *VpnCommonConfig) GetPropertyIsAutoTriggerEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsAutoTriggerEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *VpnCommonConfig) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *VpnCommonConfig) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetProfileType sets the value of ProfileType for the instance
func (instance *VpnCommonConfig) SetPropertyProfileType(value string) (err error) {
	return instance.SetProperty("ProfileType", (value))
}

// GetProfileType gets the value of ProfileType for the instance
func (instance *VpnCommonConfig) GetPropertyProfileType() (value string, err error) {
	retValue, err := instance.GetProperty("ProfileType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetProvisioningAuthority sets the value of ProvisioningAuthority for the instance
func (instance *VpnCommonConfig) SetPropertyProvisioningAuthority(value string) (err error) {
	return instance.SetProperty("ProvisioningAuthority", (value))
}

// GetProvisioningAuthority gets the value of ProvisioningAuthority for the instance
func (instance *VpnCommonConfig) GetPropertyProvisioningAuthority() (value string, err error) {
	retValue, err := instance.GetProperty("ProvisioningAuthority")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetProxy sets the value of Proxy for the instance
func (instance *VpnCommonConfig) SetPropertyProxy(value VpnConnectionProxy) (err error) {
	return instance.SetProperty("Proxy", (value))
}

// GetProxy gets the value of Proxy for the instance
func (instance *VpnCommonConfig) GetPropertyProxy() (value VpnConnectionProxy, err error) {
	retValue, err := instance.GetProperty("Proxy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(VpnConnectionProxy)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " VpnConnectionProxy is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VpnConnectionProxy(valuetmp)

	return
}

// SetRememberCredential sets the value of RememberCredential for the instance
func (instance *VpnCommonConfig) SetPropertyRememberCredential(value bool) (err error) {
	return instance.SetProperty("RememberCredential", (value))
}

// GetRememberCredential gets the value of RememberCredential for the instance
func (instance *VpnCommonConfig) GetPropertyRememberCredential() (value bool, err error) {
	retValue, err := instance.GetProperty("RememberCredential")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetRoutes sets the value of Routes for the instance
func (instance *VpnCommonConfig) SetPropertyRoutes(value []MSFT_NetRoute) (err error) {
	return instance.SetProperty("Routes", (value))
}

// GetRoutes gets the value of Routes for the instance
func (instance *VpnCommonConfig) GetPropertyRoutes() (value []MSFT_NetRoute, err error) {
	retValue, err := instance.GetProperty("Routes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_NetRoute)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_NetRoute is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_NetRoute(valuetmp))
	}

	return
}

// SetServerAddress sets the value of ServerAddress for the instance
func (instance *VpnCommonConfig) SetPropertyServerAddress(value string) (err error) {
	return instance.SetProperty("ServerAddress", (value))
}

// GetServerAddress gets the value of ServerAddress for the instance
func (instance *VpnCommonConfig) GetPropertyServerAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ServerAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetServerList sets the value of ServerList for the instance
func (instance *VpnCommonConfig) SetPropertyServerList(value []VpnServerAddress) (err error) {
	return instance.SetProperty("ServerList", (value))
}

// GetServerList gets the value of ServerList for the instance
func (instance *VpnCommonConfig) GetPropertyServerList() (value []VpnServerAddress, err error) {
	retValue, err := instance.GetProperty("ServerList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(VpnServerAddress)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " VpnServerAddress is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, VpnServerAddress(valuetmp))
	}

	return
}

// SetSplitTunneling sets the value of SplitTunneling for the instance
func (instance *VpnCommonConfig) SetPropertySplitTunneling(value bool) (err error) {
	return instance.SetProperty("SplitTunneling", (value))
}

// GetSplitTunneling gets the value of SplitTunneling for the instance
func (instance *VpnCommonConfig) GetPropertySplitTunneling() (value bool, err error) {
	retValue, err := instance.GetProperty("SplitTunneling")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetVpnTrigger sets the value of VpnTrigger for the instance
func (instance *VpnCommonConfig) SetPropertyVpnTrigger(value VpnConnectionTrigger) (err error) {
	return instance.SetProperty("VpnTrigger", (value))
}

// GetVpnTrigger gets the value of VpnTrigger for the instance
func (instance *VpnCommonConfig) GetPropertyVpnTrigger() (value VpnConnectionTrigger, err error) {
	retValue, err := instance.GetProperty("VpnTrigger")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(VpnConnectionTrigger)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " VpnConnectionTrigger is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VpnConnectionTrigger(valuetmp)

	return
}
