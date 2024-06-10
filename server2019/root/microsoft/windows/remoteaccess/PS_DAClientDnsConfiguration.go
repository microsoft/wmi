// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_DAClientDnsConfiguration struct
type PS_DAClientDnsConfiguration struct {
	*cim.WmiInstance
}

func NewPS_DAClientDnsConfigurationEx1(instance *cim.WmiInstance) (newInstance *PS_DAClientDnsConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_DAClientDnsConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewPS_DAClientDnsConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_DAClientDnsConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_DAClientDnsConfiguration{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="DnsIPAddress" type="string []"></param>
// <param name="DnsSuffix" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="ProxyServer" type="string "></param>

// <param name="cmdletOutput" type="DnsClientNrptRule "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAClientDnsConfiguration) Add( /* IN */ DnsIPAddress []string,
	/* IN */ DnsSuffix string,
	/* IN */ ProxyServer string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput DnsClientNrptRule) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", DnsIPAddress, DnsSuffix, ProxyServer, ComputerName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="DnsSuffix" type="string []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DnsClientNrptRule []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAClientDnsConfiguration) Remove( /* IN */ DnsSuffix []string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput []DnsClientNrptRule) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", DnsSuffix, ComputerName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="DnsIPAddress" type="string []"></param>
// <param name="DnsSuffix" type="string "></param>
// <param name="Local" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="ProxyServer" type="string "></param>

// <param name="cmdletOutput" type="DAClientDnsConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAClientDnsConfiguration) Set( /* IN */ DnsSuffix string,
	/* IN */ DnsIPAddress []string,
	/* IN */ ProxyServer string,
	/* IN */ Local string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput DAClientDnsConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", DnsSuffix, DnsIPAddress, ProxyServer, Local, ComputerName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>

// <param name="cmdletOutput" type="DAClientDnsConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAClientDnsConfiguration) Get( /* IN */ ComputerName string,
	/* OUT */ cmdletOutput DAClientDnsConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", ComputerName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
