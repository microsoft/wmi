// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnConnectionProxy struct
type VpnConnectionProxy struct {
	*cim.WmiInstance

	//
	AutoConfigurationScript string

	//
	AutoDetect bool

	//
	BypassProxyForLocal bool

	//
	ConnectionName string

	//
	ExceptionPrefix []string

	//
	ProxyServer string
}

func NewVpnConnectionProxyEx1(instance *cim.WmiInstance) (newInstance *VpnConnectionProxy, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnConnectionProxy{
		WmiInstance: tmp,
	}
	return
}

func NewVpnConnectionProxyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnConnectionProxy, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnConnectionProxy{
		WmiInstance: tmp,
	}
	return
}

// SetAutoConfigurationScript sets the value of AutoConfigurationScript for the instance
func (instance *VpnConnectionProxy) SetPropertyAutoConfigurationScript(value string) (err error) {
	return instance.SetProperty("AutoConfigurationScript", (value))
}

// GetAutoConfigurationScript gets the value of AutoConfigurationScript for the instance
func (instance *VpnConnectionProxy) GetPropertyAutoConfigurationScript() (value string, err error) {
	retValue, err := instance.GetProperty("AutoConfigurationScript")
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

// SetAutoDetect sets the value of AutoDetect for the instance
func (instance *VpnConnectionProxy) SetPropertyAutoDetect(value bool) (err error) {
	return instance.SetProperty("AutoDetect", (value))
}

// GetAutoDetect gets the value of AutoDetect for the instance
func (instance *VpnConnectionProxy) GetPropertyAutoDetect() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoDetect")
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

// SetBypassProxyForLocal sets the value of BypassProxyForLocal for the instance
func (instance *VpnConnectionProxy) SetPropertyBypassProxyForLocal(value bool) (err error) {
	return instance.SetProperty("BypassProxyForLocal", (value))
}

// GetBypassProxyForLocal gets the value of BypassProxyForLocal for the instance
func (instance *VpnConnectionProxy) GetPropertyBypassProxyForLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("BypassProxyForLocal")
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

// SetConnectionName sets the value of ConnectionName for the instance
func (instance *VpnConnectionProxy) SetPropertyConnectionName(value string) (err error) {
	return instance.SetProperty("ConnectionName", (value))
}

// GetConnectionName gets the value of ConnectionName for the instance
func (instance *VpnConnectionProxy) GetPropertyConnectionName() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionName")
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

// SetExceptionPrefix sets the value of ExceptionPrefix for the instance
func (instance *VpnConnectionProxy) SetPropertyExceptionPrefix(value []string) (err error) {
	return instance.SetProperty("ExceptionPrefix", (value))
}

// GetExceptionPrefix gets the value of ExceptionPrefix for the instance
func (instance *VpnConnectionProxy) GetPropertyExceptionPrefix() (value []string, err error) {
	retValue, err := instance.GetProperty("ExceptionPrefix")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetProxyServer sets the value of ProxyServer for the instance
func (instance *VpnConnectionProxy) SetPropertyProxyServer(value string) (err error) {
	return instance.SetProperty("ProxyServer", (value))
}

// GetProxyServer gets the value of ProxyServer for the instance
func (instance *VpnConnectionProxy) GetPropertyProxyServer() (value string, err error) {
	retValue, err := instance.GetProperty("ProxyServer")
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
