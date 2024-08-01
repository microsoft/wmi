// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Dns
//////////////////////////////////////////////
package dns

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_DnsClientNrptGlobal struct
type PS_DnsClientNrptGlobal struct {
	*cim.WmiInstance
}

func NewPS_DnsClientNrptGlobalEx1(instance *cim.WmiInstance) (newInstance *PS_DnsClientNrptGlobal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_DnsClientNrptGlobal{
		WmiInstance: tmp,
	}
	return
}

func NewPS_DnsClientNrptGlobalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_DnsClientNrptGlobal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_DnsClientNrptGlobal{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="EnableDAForAllNetworks" type="string "></param>
// <param name="GpoName" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="QueryPolicy" type="string "></param>
// <param name="SecureNameQueryFallback" type="string "></param>
// <param name="Server" type="string "></param>

// <param name="cmdletOutput" type="DnsClientNrptGlobal "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DnsClientNrptGlobal) Set( /* IN */ EnableDAForAllNetworks string,
	/* IN */ GpoName string,
	/* IN */ SecureNameQueryFallback string,
	/* IN */ QueryPolicy string,
	/* IN */ Server string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput DnsClientNrptGlobal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", EnableDAForAllNetworks, GpoName, SecureNameQueryFallback, QueryPolicy, Server, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="GpoName" type="string "></param>
// <param name="Server" type="string "></param>

// <param name="cmdletOutput" type="DnsClientNrptGlobal "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DnsClientNrptGlobal) Get( /* IN */ Server string,
	/* IN */ GpoName string,
	/* OUT */ cmdletOutput DnsClientNrptGlobal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Server, GpoName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
