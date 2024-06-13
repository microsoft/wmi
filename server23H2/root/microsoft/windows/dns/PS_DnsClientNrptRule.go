// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Dns
//////////////////////////////////////////////
package dns

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_DnsClientNrptRule struct
type PS_DnsClientNrptRule struct {
	*cim.WmiInstance
}

func NewPS_DnsClientNrptRuleEx1(instance *cim.WmiInstance) (newInstance *PS_DnsClientNrptRule, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_DnsClientNrptRule{
		WmiInstance: tmp,
	}
	return
}

func NewPS_DnsClientNrptRuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_DnsClientNrptRule, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_DnsClientNrptRule{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Comment" type="string "></param>
// <param name="DAEnable" type="bool "></param>
// <param name="DAIPsecEncryptionType" type="string "></param>
// <param name="DAIPsecRequired" type="bool "></param>
// <param name="DANameServers" type="string []"></param>
// <param name="DAProxyServerName" type="string "></param>
// <param name="DAProxyType" type="string "></param>
// <param name="DisplayName" type="string "></param>
// <param name="DnsSecEnable" type="bool "></param>
// <param name="DnsSecIPsecEncryptionType" type="string "></param>
// <param name="DnsSecIPsecRequired" type="bool "></param>
// <param name="DnsSecValidationRequired" type="bool "></param>
// <param name="GpoName" type="string "></param>
// <param name="IPsecTrustAuthority" type="string "></param>
// <param name="NameEncoding" type="string "></param>
// <param name="NameServers" type="string []"></param>
// <param name="Namespace" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="Server" type="string "></param>

// <param name="cmdletOutput" type="DnsClientNrptRule "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DnsClientNrptRule) Add( /* IN */ GpoName string,
	/* IN */ DANameServers []string,
	/* IN */ DAIPsecRequired bool,
	/* IN */ DAIPsecEncryptionType string,
	/* IN */ DAProxyServerName string,
	/* IN */ DnsSecEnable bool,
	/* IN */ PassThru bool,
	/* IN */ DAProxyType string,
	/* IN */ DnsSecValidationRequired bool,
	/* IN */ DAEnable bool,
	/* IN */ IPsecTrustAuthority string,
	/* IN */ Comment string,
	/* IN */ DisplayName string,
	/* IN */ DnsSecIPsecRequired bool,
	/* IN */ DnsSecIPsecEncryptionType string,
	/* IN */ NameServers []string,
	/* IN */ NameEncoding string,
	/* IN */ Namespace []string,
	/* IN */ Server string,
	/* OUT */ cmdletOutput DnsClientNrptRule) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", GpoName, DANameServers, DAIPsecRequired, DAIPsecEncryptionType, DAProxyServerName, DnsSecEnable, PassThru, DAProxyType, DnsSecValidationRequired, DAEnable, IPsecTrustAuthority, Comment, DisplayName, DnsSecIPsecRequired, DnsSecIPsecEncryptionType, NameServers, NameEncoding, Namespace, Server)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="GpoName" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Server" type="string "></param>

// <param name="cmdletOutput" type="DnsClientNrptRule "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DnsClientNrptRule) Remove( /* IN */ GpoName string,
	/* IN */ Name string,
	/* IN */ PassThru bool,
	/* IN */ Server string,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput DnsClientNrptRule) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", GpoName, Name, PassThru, Server, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="GpoName" type="string "></param>
// <param name="Name" type="string []"></param>
// <param name="Server" type="string "></param>

// <param name="cmdletOutput" type="DnsClientNrptRule []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DnsClientNrptRule) Get( /* IN */ GpoName string,
	/* IN */ Name []string,
	/* IN */ Server string,
	/* OUT */ cmdletOutput []DnsClientNrptRule) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", GpoName, Name, Server)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Comment" type="string "></param>
// <param name="DAEnable" type="bool "></param>
// <param name="DAIPsecEncryptionType" type="string "></param>
// <param name="DAIPsecRequired" type="bool "></param>
// <param name="DANameServers" type="string []"></param>
// <param name="DAProxyServerName" type="string "></param>
// <param name="DAProxyType" type="string "></param>
// <param name="DisplayName" type="string "></param>
// <param name="DnsSecEnable" type="bool "></param>
// <param name="DnsSecIPsecEncryptionType" type="string "></param>
// <param name="DnsSecIPsecRequired" type="bool "></param>
// <param name="DnsSecValidationRequired" type="bool "></param>
// <param name="GpoName" type="string "></param>
// <param name="IPsecTrustAuthority" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="NameEncoding" type="string "></param>
// <param name="NameServers" type="string []"></param>
// <param name="Namespace" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="Server" type="string "></param>

// <param name="cmdletOutput" type="DnsClientNrptRule "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DnsClientNrptRule) Set( /* IN */ DAEnable bool,
	/* IN */ DAIPsecEncryptionType string,
	/* IN */ DAIPsecRequired bool,
	/* IN */ DANameServers []string,
	/* IN */ DAProxyServerName string,
	/* IN */ DAProxyType string,
	/* IN */ DisplayName string,
	/* IN */ PassThru bool,
	/* IN */ IPsecTrustAuthority string,
	/* IN */ Name string,
	/* IN */ NameEncoding string,
	/* IN */ NameServers []string,
	/* IN */ Namespace []string,
	/* IN */ Server string,
	/* IN */ Comment string,
	/* IN */ DnsSecEnable bool,
	/* IN */ DnsSecIPsecEncryptionType string,
	/* IN */ DnsSecIPsecRequired bool,
	/* IN */ DnsSecValidationRequired bool,
	/* IN */ GpoName string,
	/* OUT */ cmdletOutput DnsClientNrptRule) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", DAEnable, DAIPsecEncryptionType, DAIPsecRequired, DANameServers, DAProxyServerName, DAProxyType, DisplayName, PassThru, IPsecTrustAuthority, Name, NameEncoding, NameServers, Namespace, Server, Comment, DnsSecEnable, DnsSecIPsecEncryptionType, DnsSecIPsecRequired, DnsSecValidationRequired, GpoName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
