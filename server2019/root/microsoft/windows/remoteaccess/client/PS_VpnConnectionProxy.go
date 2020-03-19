// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_VpnConnectionProxy struct
type PS_VpnConnectionProxy struct {
	*cim.WmiInstance

	//
	AutoConfigurationScript string

	//
	AutoDetect bool

	//
	BypassProxyForLocal bool

	//
	ExceptionPrefix []string

	//
	ProxyServer string
}

func NewPS_VpnConnectionProxyEx1(instance *cim.WmiInstance) (newInstance *PS_VpnConnectionProxy, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnConnectionProxy{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnConnectionProxyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnConnectionProxy, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnConnectionProxy{
		WmiInstance: tmp,
	}
	return
}

// SetAutoConfigurationScript sets the value of AutoConfigurationScript for the instance
func (instance *PS_VpnConnectionProxy) SetPropertyAutoConfigurationScript(value string) (err error) {
	return instance.SetProperty("AutoConfigurationScript", value)
}

// GetAutoConfigurationScript gets the value of AutoConfigurationScript for the instance
func (instance *PS_VpnConnectionProxy) GetPropertyAutoConfigurationScript() (value string, err error) {
	retValue, err := instance.GetProperty("AutoConfigurationScript")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoDetect sets the value of AutoDetect for the instance
func (instance *PS_VpnConnectionProxy) SetPropertyAutoDetect(value bool) (err error) {
	return instance.SetProperty("AutoDetect", value)
}

// GetAutoDetect gets the value of AutoDetect for the instance
func (instance *PS_VpnConnectionProxy) GetPropertyAutoDetect() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoDetect")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBypassProxyForLocal sets the value of BypassProxyForLocal for the instance
func (instance *PS_VpnConnectionProxy) SetPropertyBypassProxyForLocal(value bool) (err error) {
	return instance.SetProperty("BypassProxyForLocal", value)
}

// GetBypassProxyForLocal gets the value of BypassProxyForLocal for the instance
func (instance *PS_VpnConnectionProxy) GetPropertyBypassProxyForLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("BypassProxyForLocal")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExceptionPrefix sets the value of ExceptionPrefix for the instance
func (instance *PS_VpnConnectionProxy) SetPropertyExceptionPrefix(value []string) (err error) {
	return instance.SetProperty("ExceptionPrefix", value)
}

// GetExceptionPrefix gets the value of ExceptionPrefix for the instance
func (instance *PS_VpnConnectionProxy) GetPropertyExceptionPrefix() (value []string, err error) {
	retValue, err := instance.GetProperty("ExceptionPrefix")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProxyServer sets the value of ProxyServer for the instance
func (instance *PS_VpnConnectionProxy) SetPropertyProxyServer(value string) (err error) {
	return instance.SetProperty("ProxyServer", value)
}

// GetProxyServer gets the value of ProxyServer for the instance
func (instance *PS_VpnConnectionProxy) GetPropertyProxyServer() (value string, err error) {
	retValue, err := instance.GetProperty("ProxyServer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="AutoConfigurationScript" type="string "></param>
// <param name="AutoDetect" type="bool "></param>
// <param name="BypassProxyForLocal" type="bool "></param>
// <param name="ConnectionName" type="string "></param>
// <param name="ExceptionPrefix" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="ProxyServer" type="string "></param>

// <param name="cmdletOutput" type="VpnConnectionProxy "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionProxy) Set( /* IN */ AutoDetect bool,
	/* IN */ AutoConfigurationScript string,
	/* IN */ ProxyServer string,
	/* IN */ BypassProxyForLocal bool,
	/* IN */ ExceptionPrefix []string,
	/* IN */ ConnectionName string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput VpnConnectionProxy) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", AutoDetect, AutoConfigurationScript, ProxyServer, BypassProxyForLocal, ExceptionPrefix, ConnectionName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
